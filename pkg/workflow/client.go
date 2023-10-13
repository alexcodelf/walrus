package workflow

import (
	"context"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// WorkflowClient is the interface that defines the operations of workflow engine.
type WorkflowClient interface {
	// Submit submits a workflow to the workflow engine.
	Submit(context.Context, SubmitOptions) error
	// Delete deletes a workflow from the workflow engine.
	Delete(context.Context, DeleteOptions) error
	// Get gets a workflow from the workflow engine.
	Get(context.Context, GetOptions) (*v1alpha1.Workflow, error)
	// List lists all workflows from the workflow engine.
	List(context.Context, ListOptions) (*v1alpha1.WorkflowList, error)
}

// SubmitOptions is the options for submitting a workflow.
// WorkflowExecution's Edge WorkflowStageExecutions and their Edge WorkflowStepExecutions must be set.
type SubmitOptions struct {
	WorkflowExecution *model.WorkflowExecution
	SubjectID         object.ID
}

type ListOptions struct{}

type GetOptions struct {
	Workflow *model.Workflow
}

type DeleteOptions struct {
	Workflow *model.Workflow
}
