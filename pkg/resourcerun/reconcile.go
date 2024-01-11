package resourcerun

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/seal-io/walrus/pkg/bus/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RunReconciler struct {
	Logger      logr.Logger
	Kubeconfig  *rest.Config
	KubeClient  client.Client
	ModelClient *model.Client
}

func (r RunReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	job := &batchv1.Job{}

	err := r.KubeClient.Get(ctx, req.NamespacedName, job)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	err = r.syncResourceRunStatus(ctx, job)
	if err != nil && !model.IsNotFound(err) {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r RunReconciler) Setup(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.Job{}).
		Complete(r)
}

// syncResourceRunStatus sync the resource run status.
func (r RunReconciler) syncResourceRunStatus(ctx context.Context, job *batchv1.Job) (err error) {
	resRunID, ok := job.Labels[ResourceRunIDLabel]
	if !ok {
		// Not a deployer job.
		return nil
	}

	resourceRun, err := r.ModelClient.ResourceRuns().Get(ctx, object.ID(resRunID))
	if err != nil {
		return err
	}

	// If the resource run status is not running, then skip it.
	if !status.ResourceRunStatusReady.IsUnknown(resourceRun) {
		return nil
	}

	if job.Status.Succeeded == 0 && job.Status.Failed == 0 {
		return nil
	}

	status.ResourceRunStatusReady.True(resourceRun, "")

	// Get job pods logs.
	record, err := r.getJobPodsLogs(ctx, job.Name)
	if err != nil {
		r.Logger.Error(err, "failed to get job pod logs", "resource-run", resRunID)
		record = err.Error()
	}

	if job.Status.Succeeded > 0 {
		r.Logger.Info("succeed", "resource-run", resRunID)
	}

	if job.Status.Failed > 0 {
		r.Logger.Info("failed", "resource-run", resRunID)
		status.ResourceRunStatusReady.False(resourceRun, "")
	}

	// Report to resource run.
	resourceRun.Record = record
	resourceRun.Status.SetSummary(status.WalkResourceRun(&resourceRun.Status))
	resourceRun.Duration = int(time.Since(*resourceRun.CreateTime).Seconds())

	resourceRun, err = r.ModelClient.ResourceRuns().UpdateOne(resourceRun).
		SetStatus(resourceRun.Status).
		SetRecord(resourceRun.Record).
		SetDuration(resourceRun.Duration).
		Save(ctx)
	if err != nil {
		return err
	}

	return resourcerun.Notify(ctx, r.ModelClient, resourceRun)
}

// getJobPodsLogs returns the logs of all pods of a job.
func (r RunReconciler) getJobPodsLogs(ctx context.Context, jobName string) (string, error) {
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
