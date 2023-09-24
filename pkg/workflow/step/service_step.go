package step

import (
	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
)

type ServiceStep struct {
	WorkflowStep *model.WorkflowStep

	Service model.Service
}

func (s *ServiceStep) Generate() (*wfv1.Template, error) {
	return &wfv1.Template{
		Name:   s.WorkflowStep.Name,
		Inputs: wfv1.Inputs{},
	}, nil
}

// Walrus Step 类型 自动根据用户填入的 service template version 获取输入输出
// 怎么获取输入输出呢？

func (s *ServiceStep) GetInputs() (wfv1.Inputs, error) {
	// get terraform apply files

	return wfv1.Inputs{}, nil
}
