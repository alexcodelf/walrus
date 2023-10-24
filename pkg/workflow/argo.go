package workflow

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/argoproj/argo-workflows/v3/pkg/apiclient"
	"github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	corev1 "k8s.io/api/core/v1"

	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
	"github.com/seal-io/walrus/pkg/auths"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step"
	steptypes "github.com/seal-io/walrus/pkg/workflow/step/types"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/pointer"
)

const (
	beforeTemplateKey = "before"
	afterTemplateKey  = "after"
	mainTemplateKey   = "main"
)

type ArgoWorkflowClient struct {
	Logger log.Logger
	mc     model.ClientSet
	// Argo workflow clientset.
	apiClient apiclient.Client
	ctx       context.Context
}

func NewArgoWorkflowClient(mc model.ClientSet, clientConfig clientcmd.ClientConfig) (WorkflowClient, error) {
	ctx, apiClient, err := apiclient.NewClientFromOpts(apiclient.Opts{
		ClientConfigSupplier: func() clientcmd.ClientConfig {
			return clientConfig
		},
	})
	if err != nil {
		return nil, err
	}

	return &ArgoWorkflowClient{
		Logger:    log.WithName("workflow-service"),
		mc:        mc,
		apiClient: apiClient,
		ctx:       ctx,
	}, nil
}

func (s *ArgoWorkflowClient) Submit(ctx context.Context, opts SubmitOptions) error {
	// Prepare API token for terraform backend.
	const _1Day = 60 * 60 * 24

	at, err := auths.CreateAccessToken(ctx,
		s.mc, opts.SubjectID, types.TokenKindDeployment, opts.WorkflowExecution.ID.String(), pointer.Int(_1Day))
	if err != nil {
		return err
	}

	token := at.AccessToken

	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, s.mc)
	if err != nil {
		return err
	}

	if serverAddress == "" {
		return errors.New("server address is empty")
	}

	wfTemplates, err := s.GenerateWorkflowTemplateEntrypoint(ctx, opts.WorkflowExecution.Edges.Stages)
	if err != nil {
		return err
	}

	workflowTemplates := make([]v1alpha1.Template, 0, len(wfTemplates)+1)
	for _, tpl := range wfTemplates {
		workflowTemplates = append(workflowTemplates, *tpl)
	}

	exitHandler := getExitTemplate(opts.WorkflowExecution)
	workflowTemplates = append(workflowTemplates, *exitHandler)

	wf := &v1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name: opts.WorkflowExecution.Name,
		},
		Spec: v1alpha1.WorkflowSpec{
			Entrypoint: "entrypoint",
			OnExit:     exitHandler.Name,
			Arguments: v1alpha1.Arguments{
				Parameters: []v1alpha1.Parameter{
					{
						Name:  "server",
						Value: v1alpha1.AnyStringPtr(serverAddress),
					},
					{
						Name:  "projectID",
						Value: v1alpha1.AnyStringPtr(opts.WorkflowExecution.ProjectID.String()),
					},
					{
						Name:  "workflowID",
						Value: v1alpha1.AnyStringPtr(opts.WorkflowExecution.WorkflowID.String()),
					},
					{
						Name:  "token",
						Value: v1alpha1.AnyStringPtr(token),
					},
					{
						Name:  "tlsVerify",
						Value: v1alpha1.AnyStringPtr(apiconfig.TlsCertified.Get()),
					},
				},
			},
			SecurityContext: &corev1.PodSecurityContext{
				RunAsNonRoot: pointer.Bool(true),
				RunAsUser:    pointer.Int64(1000),
			},
			TTLStrategy: &v1alpha1.TTLStrategy{
				SecondsAfterCompletion: pointer.Int32(86400),
			},
			PodGC: &v1alpha1.PodGC{
				Strategy: v1alpha1.PodGCOnWorkflowCompletion,
			},
			ServiceAccountName: types.WorkflowServiceAccountName,
			Templates:          workflowTemplates,
		},
	}

	_, err = s.apiClient.NewWorkflowServiceClient().CreateWorkflow(s.ctx, &workflow.WorkflowCreateRequest{
		Namespace: types.WalrusWorkflowNamespace,
		Workflow:  wf,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArgoWorkflowClient) Resume(ctx context.Context, opts ResumeOptions) error {
	// TODO get subjects from spec.
	workflowExecution, err := s.mc.WorkflowExecutions().Query().
		Where(workflowexecution.ID(opts.WorkflowStepExecution.WorkflowExecutionID)).
		Only(ctx)
	if err != nil {
		return err
	}

	_, err = s.apiClient.NewWorkflowServiceClient().ResumeWorkflow(s.ctx, &workflow.WorkflowResumeRequest{
		Name:              workflowExecution.Name,
		Namespace:         types.WalrusWorkflowNamespace,
		NodeFieldSelector: fmt.Sprintf("templateName=suspend-%s", opts.WorkflowStepExecution.ID.String()),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArgoWorkflowClient) Delete(ctx context.Context, opts DeleteOptions) error {
	_, err := s.apiClient.NewWorkflowServiceClient().DeleteWorkflow(s.ctx, &workflow.WorkflowDeleteRequest{
		Name:      opts.Workflow.Name,
		Namespace: types.WalrusWorkflowNamespace,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArgoWorkflowClient) GenerateWorkflowTemplateEntrypoint(
	ctx context.Context,
	stageExecutions model.WorkflowStageExecutions,
) ([]*v1alpha1.Template, error) {
	workflowTplLen := 1
	for _, stageExec := range stageExecutions {
		workflowTplLen += (len(stageExec.Edges.Steps) + 1) * 3
	}

	workflowTemplates := make([]*v1alpha1.Template, 0, workflowTplLen)

	entrypoint := &v1alpha1.Template{
		Name: "entrypoint",
		DAG:  &v1alpha1.DAGTemplate{},
	}

	taskTemplates := make([]*v1alpha1.Template, 0)

	for i, stageExec := range stageExecutions {
		stageTemplates, err := s.GenerateStageTemplates(ctx, stageExec)
		if err != nil {
			return nil, err
		}

		taskTemplates = append(taskTemplates, stageTemplates...)

		var dependencies []string
		if i > 0 {
			// Add previous stage task as dependency.
			dependencies = append(dependencies, entrypoint.DAG.Tasks[i-1].Name)
		}

		var (
			taskName         = "stage-execution-" + stageExec.ID.String()
			taskTemplateName = "stage-execution-" + stageExec.ID.String()
		)

		entrypoint.DAG.Tasks = append(entrypoint.DAG.Tasks, v1alpha1.DAGTask{
			Name:         taskName,
			Template:     taskTemplateName,
			Dependencies: dependencies,
			Hooks: v1alpha1.LifecycleHooks{
				"running": v1alpha1.LifecycleHook{
					Template:   taskTemplateName + "-before",
					Expression: fmt.Sprintf("tasks['%s'].status == \"Running\"", taskName),
					Arguments: v1alpha1.Arguments{
						Parameters: []v1alpha1.Parameter{
							{
								Name:  "status",
								Value: v1alpha1.AnyStringPtr("Running"),
							},
						},
					},
				},
				"succeeded": v1alpha1.LifecycleHook{
					Template:   taskTemplateName + "-after",
					Expression: fmt.Sprintf("tasks['%s'].status == \"Succeeded\"", taskName),
					Arguments: v1alpha1.Arguments{
						Parameters: []v1alpha1.Parameter{
							{
								Name:  "status",
								Value: v1alpha1.AnyStringPtr("Succeeded"),
							},
						},
					},
				},
				"failed": v1alpha1.LifecycleHook{
					Template:   taskTemplateName + "-after",
					Expression: fmt.Sprintf("tasks['%s'].status == \"Failed\"", taskName),
					Arguments: v1alpha1.Arguments{
						Parameters: []v1alpha1.Parameter{
							{
								Name:  "status",
								Value: v1alpha1.AnyStringPtr("Failed"),
							},
						},
					},
				},
			},
		})
	}

	workflowTemplates = append(workflowTemplates, taskTemplates...)
	workflowTemplates = append(workflowTemplates, entrypoint)

	return workflowTemplates, nil
}

// GenerateStageTemplates generates templates for workflow stage.
// It generates before, after and main templates,
// where main template is a DAG template with tasks for each step templates.
func (s *ArgoWorkflowClient) GenerateStageTemplates(
	ctx context.Context,
	stageExecution *model.WorkflowStageExecution,
) ([]*v1alpha1.Template, error) {
	stageTemplate := &v1alpha1.Template{
		Name: "stage-execution-" + stageExecution.ID.String(),
		DAG:  &v1alpha1.DAGTemplate{},
	}

	beforeTemplate := &v1alpha1.Template{
		Name: "stage-execution-" + stageExecution.ID.String() + "-before",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(stageExecution.ID.String()),
				},
				{
					Name: "status",
				},
				{
					Name:  "workflowExecutionID",
					Value: v1alpha1.AnyStringPtr(stageExecution.WorkflowExecutionID.String()),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL: "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
				"/workflows/{{workflow.parameters.workflowID}}" +
				"/executions/{{inputs.parameters.workflowExecutionID}}" +
				"/stage-executions/{{inputs.parameters.id}}",
			Method: http.MethodPut,
			Headers: v1alpha1.HTTPHeaders{
				{
					Name:  "Authorization",
					Value: "Bearer {{workflow.parameters.token}}",
				},
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			},
			TimeoutSeconds: pointer.Int64(5),
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
		},
	}
	afterTemplate := &v1alpha1.Template{
		Name: "stage-execution-" + stageExecution.ID.String() + "-after",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(stageExecution.ID.String()),
				},
				{
					Name: "status",
				},
				{
					Name:  "workflowExecutionID",
					Value: v1alpha1.AnyStringPtr(stageExecution.WorkflowExecutionID.String()),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL: "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
				"/workflows/{{workflow.parameters.workflowID}}" +
				"/executions/{{inputs.parameters.workflowExecutionID}}" +
				"/stage-executions/{{inputs.parameters.id}}",
			Method: http.MethodPut,
			Headers: v1alpha1.HTTPHeaders{
				{
					Name:  "Authorization",
					Value: "Bearer {{workflow.parameters.token}}",
				},
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			},
			TimeoutSeconds: pointer.Int64(5),
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
		},
	}

	stageTemplates := []*v1alpha1.Template{
		beforeTemplate,
		stageTemplate,
		afterTemplate,
	}

	tasks := make([]v1alpha1.DAGTask, 0, len(stageExecution.Edges.Steps))
	stepTemplates := make([]*v1alpha1.Template, 0)

	for _, stepExec := range stageExecution.Edges.Steps {
		stepTemplateMap, err := s.GenerateStepTemplate(ctx, stepExec)
		if err != nil {
			return nil, err
		}

		for key, stepTemplate := range stepTemplateMap {
			stepTemplates = append(stepTemplates, stepTemplate)

			if key == mainTemplateKey {
				taskName := "step-execution-" + stepExec.ID.String()
				tasks = append(tasks, v1alpha1.DAGTask{
					Name:     taskName,
					Template: stepTemplate.Name,
					Hooks: v1alpha1.LifecycleHooks{
						"running": v1alpha1.LifecycleHook{
							Template:   stepTemplateMap[beforeTemplateKey].Name,
							Expression: fmt.Sprintf("tasks['%s'].status == \"Running\"", taskName),
							Arguments: v1alpha1.Arguments{
								Parameters: []v1alpha1.Parameter{
									{
										Name:  "status",
										Value: v1alpha1.AnyStringPtr("Running"),
									},
								},
							},
						},
						"succeeded": v1alpha1.LifecycleHook{
							Template:   stepTemplateMap[afterTemplateKey].Name,
							Expression: fmt.Sprintf("tasks['%s'].status == \"Succeeded\"", taskName),
							Arguments: v1alpha1.Arguments{
								Parameters: []v1alpha1.Parameter{
									{
										Name:  "status",
										Value: v1alpha1.AnyStringPtr("Succeeded"),
									},
								},
							},
						},
						"failed": v1alpha1.LifecycleHook{
							Template:   stepTemplateMap[afterTemplateKey].Name,
							Expression: fmt.Sprintf("tasks['%s'].status == \"Failed\"", taskName),
							Arguments: v1alpha1.Arguments{
								Parameters: []v1alpha1.Parameter{
									{
										Name:  "status",
										Value: v1alpha1.AnyStringPtr("Failed"),
									},
								},
							},
						},
					},
					// TODO add dependencies.
				})
			}
		}
	}

	stageTemplate.DAG.Tasks = tasks

	// Add step templates to templates list.
	stageTemplates = append(stageTemplates, stepTemplates...)

	return stageTemplates, nil
}

func (s *ArgoWorkflowClient) GenerateStepTemplate(
	ctx context.Context,
	wse *model.WorkflowStepExecution,
) (map[string]*v1alpha1.Template, error) {
	beforeTemplate := &v1alpha1.Template{
		Name: "step-execution-" + wse.ID.String() + "-before",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(wse.ID.String()),
				},
				{
					Name:  "workflowStageExecutionID",
					Value: v1alpha1.AnyStringPtr(wse.WorkflowStageExecutionID.String()),
				},
				{
					Name:  "workflowExecutionID",
					Value: v1alpha1.AnyStringPtr(wse.WorkflowExecutionID.String()),
				},
				{
					Name: "status",
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL: "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
				"/workflows/{{workflow.parameters.workflowID}}" +
				"/executions/{{inputs.parameters.workflowExecutionID}}" +
				"/stage-executions/{{inputs.parameters.workflowStageExecutionID}}" +
				"/step-executions/{{inputs.parameters.id}}",
			Method: http.MethodPut,
			Headers: v1alpha1.HTTPHeaders{
				{
					Name:  "Authorization",
					Value: "Bearer {{workflow.parameters.token}}",
				},
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			},
			TimeoutSeconds: pointer.Int64(5),
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
		},
	}
	afterTemplate := &v1alpha1.Template{
		Name: "step-execution-" + wse.ID.String() + "-after",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(wse.ID.String()),
				},
				{
					Name: "status",
				},
				{
					Name:  "workflowStageExecutionID",
					Value: v1alpha1.AnyStringPtr(wse.WorkflowStageExecutionID.String()),
				},
				{
					Name:  "workflowExecutionID",
					Value: v1alpha1.AnyStringPtr(wse.WorkflowExecutionID.String()),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL: "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
				"/workflows/{{workflow.parameters.workflowID}}" +
				"/executions/{{inputs.parameters.workflowExecutionID}}" +
				"/stage-executions/{{inputs.parameters.workflowStageExecutionID}}" +
				"/step-executions/{{inputs.parameters.id}}",
			Method: http.MethodPut,
			Headers: v1alpha1.HTTPHeaders{
				{
					Name:  "Authorization",
					Value: "Bearer {{workflow.parameters.token}}",
				},
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			},
			TimeoutSeconds: pointer.Int64(5),
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
		},
	}

	stepService, err := step.GetStepManager(steptypes.CreateOptions{
		Type:        steptypes.Type(wse.Type),
		ModelClient: s.mc,
	})
	if err != nil {
		return nil, err
	}

	// Generate service template.
	serviceTemplate, err := stepService.GenerateTemplate(ctx, wse)
	if err != nil {
		return nil, err
	}

	return map[string]*v1alpha1.Template{
		beforeTemplateKey: beforeTemplate,
		afterTemplateKey:  afterTemplate,
		mainTemplateKey:   serviceTemplate,
	}, nil
}

func getExitTemplate(wf *model.WorkflowExecution) *v1alpha1.Template {
	return &v1alpha1.Template{
		Name: "notify",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "workflowExecutionID",
					Value: v1alpha1.AnyStringPtr(wf.ID.String()),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL: "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
				"/workflows/{{workflow.parameters.workflowID}}" +
				"/executions/{{inputs.parameters.workflowExecutionID}}",
			Method: http.MethodPut,
			Headers: v1alpha1.HTTPHeaders{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
				{
					Name:  "Authorization",
					Value: "Bearer {{workflow.parameters.token}}",
				},
			},
			TimeoutSeconds:     pointer.Int64(5),
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			Body: `{
				"status": "{{workflow.status}}",
				"id": "{{inputs.parameters.workflowExecutionID}}"
			}`,
		},
	}
}

type StreamLogsOptions struct {
	Workflow   string
	PodName    string
	Grep       string
	Selector   string
	ApiClient  apiclient.Client
	LogOptions *corev1.PodLogOptions

	Out io.Writer
}

func StreamWorkflowLogs(
	ctx context.Context,
	opts StreamLogsOptions,
) error {
	serviceClient := opts.ApiClient.NewWorkflowServiceClient()

	stream, err := serviceClient.WorkflowLogs(ctx, &workflow.WorkflowLogRequest{
		Name:       opts.Workflow,
		Namespace:  types.WalrusWorkflowNamespace,
		PodName:    opts.PodName,
		LogOptions: opts.LogOptions,
		Selector:   opts.Selector,
		Grep:       opts.Grep,
	})
	if err != nil {
		return err
	}

	for {
		event, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		_, err = opts.Out.Write([]byte(event.Content + "\n"))
		if err != nil {
			return err
		}
	}
}
