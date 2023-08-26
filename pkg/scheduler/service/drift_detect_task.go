package service

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	batchv1 "k8s.io/api/batch/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	toolswatch "k8s.io/client-go/tools/watch"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/deployer"
	"github.com/seal-io/walrus/pkg/deployer/terraform"
	deptypes "github.com/seal-io/walrus/pkg/deployer/types"
	pkgservice "github.com/seal-io/walrus/pkg/service"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/strs"
)

const (
	summaryStatusReady = "Ready"
)

type DriftDetectTask struct {
	logger       log.Logger
	modelClient  model.ClientSet
	kubeClient   *kubernetes.Clientset
	tlsCertified bool
	deployer     deptypes.Deployer
}

func NewServiceDriftDetectTask(
	logger log.Logger,
	mc model.ClientSet,
	kc *rest.Config,
	tlsCertified bool,
) (in *DriftDetectTask, err error) {
	// Create deployer.
	opts := deptypes.CreateOptions{
		Type:        terraform.DeployerType,
		ModelClient: mc,
		KubeConfig:  kc,
	}

	dp, err := deployer.Get(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	kubeClient, err := kubernetes.NewForConfig(kc)
	if err != nil {
		return nil, err
	}

	in = &DriftDetectTask{
		tlsCertified: tlsCertified,
		logger:       logger,
		modelClient:  mc,
		kubeClient:   kubeClient,
		deployer:     dp,
	}

	return
}

func (in *DriftDetectTask) Process(ctx context.Context, args ...any) error {
	if !settings.EnableDriftDetection.ShouldValueBool(ctx, in.modelClient) {
		// Disable drift detection.
		return nil
	}

	// NB(alex): group 10 services into one task group,
	// treat each service as a task unit to detect drift.
	query := in.modelClient.Services().Query().
		Where(
			func(s *sql.Selector) {
				s.Where(sqljson.ValueEQ(
					service.FieldStatus,
					summaryStatusReady,
					sqljson.Path("summaryStatus"),
				))
			},
			service.Or(
				service.DriftResultIsNil(),
				func(s *sql.Selector) {
					s.Where(sqljson.ValueLTE(
						service.FieldDriftResult,
						time.Now().Add(-time.Hour),
						sqljson.Path("time"),
					))
				},
			),
		)

	cnt, err := query.Clone().Count(ctx)
	if err != nil {
		return err
	}

	if cnt == 0 {
		return nil
	}

	const bks = 10

	bkc := cnt/bks + 1

	for i := 0; i < bkc; i++ {
		err := in.driftDetectServices(ctx, query, i*bks, bks)
		if err != nil {
			in.logger.Warnf("drift detect services failed: %v", err)
		}
	}

	return nil
}

func (in *DriftDetectTask) updateServiceStatus(
	ctx context.Context,
	entity *model.Service,
	s status.ConditionType,
) error {
	s.Reset(entity, "Detect service drift")

	return pkgservice.UpdateStatus(ctx, in.modelClient, entity)
}

func (in *DriftDetectTask) driftDetectServices(
	ctx context.Context,
	query *model.ServiceQuery,
	offset, limit int,
) error {
	jobGroupLabel := "drift-detect-" + strs.Hex(16)

	entities, err := query.Clone().
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return err
	}

	if len(entities) == 0 {
		return nil
	}

	for _, entity := range entities {
		if err = in.updateServiceStatus(ctx, entity, status.ServiceStatusDetected); err != nil {
			return err
		}

		opts := pkgservice.Options{
			TlsCertified: in.tlsCertified,
			Labels: map[string]string{
				terraform.K8sJobGroupAnno: jobGroupLabel,
			},
		}
		if err = pkgservice.Detect(ctx, in.modelClient, in.deployer, entity, opts); err != nil {
			return err
		}
	}

	// Wait for job complete.
	labelSelector := labels.Set(map[string]string{
		terraform.K8sJobGroupAnno: jobGroupLabel,
	})

	finished, err := waitJobCompleted(ctx, in.kubeClient, len(entities), metav1.ListOptions{
		ResourceVersion: "0",
		LabelSelector:   labelSelector.String(),
	})
	if err != nil && !kerrors.IsNotFound(err) {
		return err
	}

	if !finished {
		in.logger.Warnf("drift detect job is not completed")
	}

	return nil
}

// waitJobCompleted waits for the job completed.
func waitJobCompleted(
	ctx context.Context,
	kc *kubernetes.Clientset,
	jobNum int,
	opts metav1.ListOptions,
) (finished bool, err error) {
	watchFn := func(options metav1.ListOptions) (watcher watch.Interface, err error) {
		return kc.BatchV1().Jobs(types.WalrusSystemNamespace).Watch(ctx, opts)
	}

	listFn := func(options metav1.ListOptions) (result runtime.Object, err error) {
		return kc.BatchV1().Jobs(types.WalrusSystemNamespace).List(ctx, opts)
	}

	watcher, err := toolswatch.NewRetryWatcher("1", &cache.ListWatch{
		WatchFunc: watchFn,
		ListFunc:  listFn,
	})
	defer watcher.Stop()

	finishedJob := 0
	timer := time.NewTimer(time.Minute * 10)

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case event, ok := <-watcher.ResultChan():
		if !ok {
			return false, nil
		}

		item := event.Object.(*batchv1.Job)
		if item.Status.Succeeded == 1 || item.Status.Failed == 1 {
			finishedJob += 1
		}

		if finishedJob == jobNum {
			finished = true
			break
		}
	case <-timer.C:
		return false, errors.New("job is not completed in 10 minutes")
	}

	if finishedJob == jobNum {
		finished = true
	}

	return
}
