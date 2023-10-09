package workflow

import (
	"context"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
)

var testWorkflow = &model.Workflow{
	Name:        "test",
	DisplayName: "test",
}

var testStage = &model.WorkflowStage{
	Name: "test",
}

func GenerateWorkflow(ctx context.Context, wf *model.Workflow) (*v1alpha1.Workflow, error) {
	return nil, nil
}
