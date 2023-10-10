package step

import (
	"context"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/workflow/step/types"
)

type ApprovalStepManager struct {
	mc model.ClientSet
}

func NewApprovalStepManager(opts types.CreateOptions) types.StepManager {
	return &ApprovalStepManager{
		mc: opts.ModelClient,
	}
}

func (m *ApprovalStepManager) GenerateTemplate(
	ctx context.Context,
	stepExec *model.WorkflowStepExecution,
) (*v1alpha1.Template, error) {
	return nil, nil
}
