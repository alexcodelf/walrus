package workflow

import (
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/utils/log"
)

// WorkflowService is the interface that defines the operations of workflow engine.
type WorkflowService interface {
	// Submit submits a workflow to the workflow engine.
	Submit(*model.Workflow) error
	// Delete deletes a workflow from the workflow engine.
	Delete(*model.Workflow) error
	// Get gets a workflow from the workflow engine.
	Get(*model.Workflow) (*v1alpha1.Workflow, error)
	// List lists all workflows from the workflow engine.
	List() (*v1alpha1.WorkflowList, error)
}

const (
	WorkflowFlowNamespace = "walrus-system"
)

type ArgoWorkflowService struct {
	Logger log.Logger
	mc     *model.Client
	// Argo workflow clientset.
	cs wfclientset.Interface
}

func NewArgoWorkflowService(mc *model.Client, config *rest.Config) (*ArgoWorkflowService, error) {
	cs, err := wfclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &ArgoWorkflowService{
		Logger: log.WithName("workflow-service"),
		mc:     mc,
		cs:     cs,
	}, nil
}

func (s *ArgoWorkflowService) Submit(wf *model.Workflow) error {
	// 1. Convert walrus model.Workflow to argo workflow config.
	// 2. Submit workflow to workflow engine.
	// 3. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowService) Delete(wf *model.Workflow) error {
	// 1. Delete workflow from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowService) Get(wf *model.Workflow) (*v1alpha1.Workflow, error) {
	// 1. Get workflow from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowService) List() (*v1alpha1.WorkflowList, error) {
	// 1. List workflows from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}
