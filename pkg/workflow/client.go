package workflow

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// WorkflowClient is the interface that defines the operations of workflow engine.
type WorkflowClient interface {
	// Submit submits a workflow to the workflow engine.
	Submit(context.Context, SubmitOptions) error
	// Resume resumes a workflow step execution of a workflow execution..
	Resume(context.Context, ResumeOptions) error
	// Delete deletes a workflow from the workflow engine.
	Delete(context.Context, DeleteOptions) error
}

// SubmitOptions is the options for submitting a workflow.
// WorkflowExecution's Edge WorkflowStageExecutions and their Edge WorkflowStepExecutions must be set.
type SubmitOptions struct {
	WorkflowExecution *model.WorkflowExecution
	SubjectID         object.ID
}

type GetOptions struct {
	Workflow *model.WorkflowExecution
}

type DeleteOptions struct {
	Workflow *model.WorkflowExecution
}

type ResumeOptions struct {
	WorkflowStepExecution *model.WorkflowStepExecution
}
