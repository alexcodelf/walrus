package workflow

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-workflows/v3/pkg/apiclient"
	"github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"
	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/k8s"
	"github.com/seal-io/walrus/utils/log"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client is the interface that defines the operations of workflow engine.
type Client interface {
	// Submit submits a workflow to the workflow engine.
	Submit(context.Context, SubmitOptions) error
	// Resume resumes a workflow step execution of a workflow execution..
	Resume(context.Context, ResumeOptions) error
	// Resubmit resubmits a workflow to the workflow engine.
	Resubmit(context.Context, ResubmitOptions) error
	// Delete deletes a workflow from the workflow engine.
	Delete(context.Context, DeleteOptions) error
}

type (
	SubmitOptions struct {
		WorkflowExecution *model.WorkflowExecution
		SubjectID         object.ID
	}
	GetOptions struct {
		Workflow *model.WorkflowExecution
	}

	DeleteOptions struct {
		Workflow *model.WorkflowExecution
	}

	// SubmitOptions is the options for submitting a workflow.
	// WorkflowExecution's Edge WorkflowStageExecutions and their Edge WorkflowStepExecutions must be set.
	ResumeOptions struct {
		// "workflow" or "step".
		Type string
		// Only used when Type is "workflow".
		WorkflowExecution *model.WorkflowExecution
		// Only used when Type is "step".
		WorkflowStepExecution *model.WorkflowStepExecution
	}

	ResubmitOptions struct {
		WorkflowExecution *model.WorkflowExecution
	}
)

type ArgoWorkflowClient struct {
	Logger log.Logger
	mc     model.ClientSet
	tm     *TemplateManager
	// Argo workflow clientset.
	apiClient *ArgoAPIClient
}

func NewArgoWorkflowClient(mc model.ClientSet, restCfg *rest.Config) (Client, error) {
	apiClient, err := NewArgoAPIClient(restCfg)
	if err != nil {
		return nil, err
	}

	return &ArgoWorkflowClient{
		Logger:    log.WithName("workflow-service"),
		mc:        mc,
		tm:        NewTemplateManager(mc),
		apiClient: apiClient,
	}, nil
}

func (s *ArgoWorkflowClient) Submit(ctx context.Context, opts SubmitOptions) error {
	wf, err := s.tm.GetWorkflowExecutionWorkflow(ctx, opts.WorkflowExecution, opts.SubjectID)
	if err != nil {
		return err
	}

	_, err = s.apiClient.NewWorkflowServiceClient().CreateWorkflow(s.apiClient.Ctx, &workflow.WorkflowCreateRequest{
		Namespace: types.WalrusWorkflowNamespace,
		Workflow:  wf,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArgoWorkflowClient) Resume(ctx context.Context, opts ResumeOptions) error {
	subjectID := session.MustGetSubject(ctx)

	fmt.Println(subjectID)

	workflowExecution, err := s.mc.WorkflowExecutions().Query().
		Where(workflowexecution.ID(opts.WorkflowStepExecution.WorkflowExecutionID)).
		Only(ctx)
	if err != nil {
		return err
	}

	_, err = s.apiClient.NewWorkflowServiceClient().ResumeWorkflow(s.apiClient.Ctx, &workflow.WorkflowResumeRequest{
		Name:              workflowExecution.Name,
		Namespace:         types.WalrusWorkflowNamespace,
		NodeFieldSelector: fmt.Sprintf("templateName=suspend-%s", opts.WorkflowStepExecution.ID.String()),
	})

	return err
}

func (s *ArgoWorkflowClient) Resubmit(ctx context.Context, opts ResubmitOptions) error {
	_, err := s.apiClient.NewWorkflowServiceClient().ResubmitWorkflow(s.apiClient.Ctx, &workflow.WorkflowResubmitRequest{
		Name:      opts.WorkflowExecution.Name,
		Namespace: types.WalrusWorkflowNamespace,
	})

	return err
}

func (s *ArgoWorkflowClient) Delete(ctx context.Context, opts DeleteOptions) error {
	_, err := s.apiClient.NewWorkflowServiceClient().DeleteWorkflow(s.apiClient.Ctx, &workflow.WorkflowDeleteRequest{
		Name:      opts.Workflow.Name,
		Namespace: types.WalrusWorkflowNamespace,
	})
	if err != nil && !kerrors.IsNotFound(err) {
		return err
	}

	return nil
}

type ArgoAPIClient struct {
	apiclient.Client

	Ctx context.Context
}

func NewArgoAPIClient(restCfg *rest.Config) (*ArgoAPIClient, error) {
	apiConfig := k8s.ToClientCmdApiConfig(restCfg)
	clientConfig := clientcmd.NewDefaultClientConfig(apiConfig, nil)

	ctx, apiClient, err := apiclient.NewClientFromOpts(apiclient.Opts{
		ClientConfigSupplier: func() clientcmd.ClientConfig {
			return clientConfig
		},
	})
	if err != nil {
		return nil, err
	}

	return &ArgoAPIClient{
		Client: apiClient,
		Ctx:    ctx,
	}, nil
}
