package job

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	runbus "github.com/seal-io/walrus/pkg/bus/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	runstatus "github.com/seal-io/walrus/pkg/resourceruns/status"
	"github.com/seal-io/walrus/pkg/storage"
	"github.com/seal-io/walrus/utils/gopool"
)

type Reconciler struct {
	Logger         logr.Logger
	Kubeconfig     *rest.Config
	KubeClient     client.Client
	ModelClient    *model.Client
	StorageManager *storage.Manager
}

func (r Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	job := &batchv1.Job{}

	err := r.KubeClient.Get(ctx, req.NamespacedName, job)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	err = r.syncRunStatus(ctx, job)
	if err != nil && !model.IsNotFound(err) {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r Reconciler) Setup(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.Job{}).
		Complete(r)
}

// syncRunStatus sync the resource run status.
func (r Reconciler) syncRunStatus(ctx context.Context, job *batchv1.Job) (err error) {
	runID, ok := job.Labels[types.LabelWalrusResourceRunID]
	if !ok {
		// Not a deployer job.
		return nil
	}

	run, err := r.ModelClient.ResourceRuns().Get(ctx, object.ID(runID))
	if err != nil {
		return err
	}

	// If the run status is not running, then skip it.
	if !runstatus.IsStatusRunning(run) {
		return nil
	}

	if job.Status.Succeeded == 0 && job.Status.Failed == 0 {
		return nil
	}

	// Get job pods logs.
	record, err := r.getJobPodsLogs(ctx, job.Name)
	if err != nil {
		r.Logger.Error(err, "failed to get job pod logs", "resource-run", runID)
		record = err.Error()
	}

	update := r.ModelClient.ResourceRuns().UpdateOne(run)

	if job.Status.Succeeded > 0 {
		r.Logger.Info("succeed", "resource-run", runID)
		runstatus.SetStatusTrue(run, "")
	}

	if job.Status.Failed > 0 {
		r.Logger.Info("failed", "resource-run", runID)
		runstatus.SetStatusFalse(run, "please check the logs")
		// Clear component changes and summary when run failed.
		update.ClearComponentChanges().
			ClearComponentChangeSummary()
	}

	// Report to Resource run.
	if runstatus.IsStatusPlanCondition(run) {
		run.PlanRecord = record
	} else {
		run.Record = record
	}

	run.Status.SetSummary(status.WalkResourceRun(&run.Status))
	run.Duration = int(time.Since(*run.CreateTime).Seconds())

	run, err = update.
		SetStatus(run.Status).
		SetPlanRecord(run.PlanRecord).
		SetRecord(run.Record).
		SetDuration(run.Duration).
		Save(ctx)
	if err != nil {
		return err
	}

	// Clean plan files.
	r.cleanPlanFiles(run)

	return runbus.Notify(ctx, r.ModelClient, run)
}

// getJobPodsLogs returns the logs of all pods of a job.
func (r Reconciler) getJobPodsLogs(ctx context.Context, jobName string) (string, error) {
	clientSet, err := kubernetes.NewForConfig(r.Kubeconfig)
	if err != nil {
		return "", err
	}
	ls := "job-name=" + jobName

	pods, err := clientSet.CoreV1().Pods(types.WalrusSystemNamespace).
		List(ctx, metav1.ListOptions{LabelSelector: ls})
	if err != nil {
		return "", err
	}

	var logs string

	for _, pod := range pods.Items {
		var podLogs []byte

		podLogs, err = clientSet.CoreV1().Pods(types.WalrusSystemNamespace).
			GetLogs(pod.Name, &corev1.PodLogOptions{}).
			DoRaw(ctx)
		if err != nil {
			return "", err
		}
		logs += string(podLogs)
	}

	return logs, nil
}

func (r Reconciler) cleanPlanFiles(run *model.ResourceRun) {
	// When run status is planned, skip it.
	if runstatus.IsStatusPlanned(run) {
		return
	}

	gopool.Go(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		err := r.StorageManager.DeleteRunPlan(ctx, run)
		if err != nil {
			r.Logger.Error(err, "failed to delete run plan", "run", run.ID)
		}
	})
}
