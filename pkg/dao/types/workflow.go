package types

import (
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

type WorkflowStepTemplateSchema = map[string]any

const (
	WorkflowTypeBasic       = "Basic"
	WorkflowTypeDAG         = "DAG"
	WorkflowTypeParentChild = "ParentChild"
)

type RetryStrategy = v1alpha1.RetryStrategy
