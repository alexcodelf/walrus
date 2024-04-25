package k8s

import (
	"context"
	"fmt"

	wf "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type WorkflowReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = &WorkflowReconciler{}

func (r *WorkflowReconciler) Reconcile(ctx context.Context, req ctrlreconcile.Request) (ctrlreconcile.Result, error) {
	// Fetch the Workflow instance
	workflow := &wf.Workflow{}
	err := r.client.Get(ctx, req.NamespacedName, workflow)
	if err != nil {
		// Error reading the object - requeue the request.
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Ignore deleted workflows.
	if workflow.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	switch workflow.Status.Phase {
	case wf.WorkflowRunning:
		// TODO update run step status.

		for i := range workflow.Status.Nodes {
			node := workflow.Status.Nodes[i]

			fmt.Println(node.TemplateName)
		}

	case wf.WorkflowSucceeded:
		// TODO update run status.
	case wf.WorkflowFailed:
		// TODO update run status.
	}

	return ctrl.Result{}, err
}

func (r *WorkflowReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&wf.Workflow{}).
		Complete(r)
}
