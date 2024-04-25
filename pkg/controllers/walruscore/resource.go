package walruscore

import (
	"context"
	"encoding/json"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	runoperation "github.com/seal-io/walrus/pkg/resourceruns/operation"
	"github.com/seal-io/walrus/pkg/resources"
)

// ResourceReconciler reconciles a v1.Resource object.
type ResourceReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ResourceReconciler)(nil)

func (r *ResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)
	res := new(walruscore.Resource)

	// Fetch
	err := r.client.Get(ctx, req.NamespacedName, res)
	if err != nil {
		logger.Error(err, "fetch resource")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Revoke if deleted.
	if res.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	// Revoke if resource in draft status.
	if apistatus.ResourceConditionUnDeployed.IsTrue(res) {
		return ctrl.Result{}, nil
	}

	// Revoke if resource in deploying status.
	if apistatus.ResourceConditionDeployed.IsUnknown(res) {
		return ctrl.Result{}, nil
	}

	// Check dependencies, if not ready, reconcile later.
	// TODO (alex): add dependencies check, if not ready, reconcile later.

	if apistatus.ResourceConditionDeployed.IsUnknown(res) {
		return ctrl.Result{}, nil
	}

	// Initialize status.
	apistatus.ResourceConditionDeployed.Unknown(res, apistatus.ResourceConditionReasonDeploying, "Deploying")

	computedAttributes, err := resources.GenComputedAttributes(ctx, res)
	if err != nil {
		logger.Error(err, "generate computed attributes")
		return ctrl.Result{}, nil
	}

	attrsRaw, err := json.Marshal(computedAttributes)
	if err != nil {
		logger.Error(err, "marshal computed attributes")
		return ctrl.Result{}, nil
	}

	res.Status.ComputedAttributes = runtime.RawExtension{
		Raw: attrsRaw,
	}

	hook, err := resources.GetOrCreateHook(ctx, r.client, res)
	if err != nil {
		logger.Error(err, "create resource hook")
		return ctrl.Result{}, nil
	}

	res.Status.ResourceHook = &walruscore.ResourceHookReference{
		Namespace: hook.Namespace,
		Name:      hook.Name,
	}

	res.Status.ConditionSummary = *apistatus.WalkResource(&res.Status.StatusDescriptor)
	err = r.client.Status().Update(ctx, res)
	if err != nil {
		logger.Error(err, "update resource status")
		return ctrl.Result{}, nil
	}

	// TODO (alex) fix resource run type, template resource run type.
	_, err = runoperation.CreateResourceRun(ctx, res, walruscore.ResourceRunTypeCreate)
	if err != nil {
		logger.Error(err, "create resource run")
		return ctrl.Result{}, nil
	}

	res.Status.ConditionSummary = *apistatus.WalkResource(&res.Status.StatusDescriptor)

	return ctrl.Result{}, nil
}

func (r *ResourceReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Resource{}).
		Complete(r)
}
