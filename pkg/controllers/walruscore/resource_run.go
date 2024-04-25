package walruscore

import (
	"context"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/resourceruns/kubehelper"
	runoperation "github.com/seal-io/walrus/pkg/resourceruns/operation"
	runstatus "github.com/seal-io/walrus/pkg/resourceruns/status"
	"github.com/seal-io/walrus/pkg/systemkuberes"
)

// ResourceRunReconciler reconciles a v1.ResourceRun object.
type ResourceRunReconciler struct {
	client ctrlcli.Client

	wfManager *runoperation.ResourceRunWorkflowManager
}

var _ ctrlreconcile.Reconciler = (*ResourceRunReconciler)(nil)

func (r *ResourceRunReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = ctrllog.FromContext(ctx)
	run := new(walruscore.ResourceRun)

	// Fetch
	err := r.client.Get(ctx, req.NamespacedName, run)
	if err != nil {
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	if run.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	if apistatus.ResourceRunConditionRunning.IsTrue(run) ||
		apistatus.ResourceRunConditionRunning.IsFalse(run) ||
		apistatus.ResourceRunConditionRunning.IsUnknown(run) ||
		apistatus.ResourceRunConditionCanceled.IsTrue(run) {
		return ctrl.Result{}, nil
	}

	res := &walruscore.Resource{
		ObjectMeta: meta.ObjectMeta{
			Namespace: run.Namespace,
			Name:      run.Spec.ResourceName,
		},
	}
	err = r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(res), res)
	if err != nil {
		return ctrl.Result{}, err
	}

	template := &walruscore.Template{
		ObjectMeta: meta.ObjectMeta{
			Namespace: run.Spec.Template.Namespace,
			Name:      run.Spec.Template.Name,
		},
	}
	err = r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(template), template)
	if err != nil {
		return ctrl.Result{}, err
	}

	run.Status.ConfigSecretName = kubehelper.NormalizeResourceRunConfigSecretName(res.Name)
	run.Status.ComputedAttributes = res.Status.ComputedAttributes
	run.Status.TemplateFormat = template.Spec.TemplateFormat
	run.Status.ResourceRunTemplate = &walruscore.ResourceRunTemplateReference{
		Name:      systemkuberes.ResourceRunTemplateDefault,
		Namespace: res.Namespace,
	}
	run.Status.Steps, err = runoperation.ComputeResourceRunSteps(ctx, res, run)
	if err != nil {
		return ctrl.Result{}, err
	}

	apistatus.ResourceRunConditionRunning.Unknown(run, apistatus.ResourceRunConditionReasonRunning, "")

	run, err = runstatus.UpdateStatus(ctx, run)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Create workflows.
	err = r.wfManager.CreatePlanWorkflow(ctx, run)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ResourceRunReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()
	r.wfManager = runoperation.NewResourceRunWorkflowManager()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.ResourceRun{}).
		Complete(r)
}
