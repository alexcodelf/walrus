package workflow

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/user"
	"path/filepath"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	wfv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
	"github.com/seal-io/walrus/pkg/auths"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step"
	steptypes "github.com/seal-io/walrus/pkg/workflow/step/types"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/pointer"
)

const (
	WorkflowFlowNamespace = "walrus-system"

	beforeTemplateKey = "before"
	afterTemplateKey  = "after"
	mainTemplateKey   = "main"
)

type ArgoWorkflowClient struct {
	Logger log.Logger
	mc     model.ClientSet
	// Argo workflow clientset.
	cs wfv1alpha1.WorkflowInterface
}

func NewArgoWorkflowClient(mc model.ClientSet, config *rest.Config) (WorkflowClient, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	// get kubeconfig file location
	kubeconfig := filepath.Join(usr.HomeDir, ".kube", "config")

	// use the current context in kubeconfig
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	cs, err := wfclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &ArgoWorkflowClient{
		Logger: log.WithName("workflow-service"),
		mc:     mc,
		cs:     cs.ArgoprojV1alpha1().Workflows("argo"),
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
	tlsVerify := apiconfig.TlsCertified.Get()

	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, s.mc)
	if err != nil {
		return err
	}

	if serverAddress == "" {
		return errors.New("server address is empty")
	}

	wfTemplates, err := s.GenerateWorkflowTemplateEntrypoint(ctx, opts.WorkflowExecution.Edges.WorkflowStageExecutions)
	if err != nil {
		return err
	}

	workflowTemplates := make([]v1alpha1.Template, 0, len(wfTemplates))
	for _, tpl := range wfTemplates {
		workflowTemplates = append(workflowTemplates, *tpl)
	}

	wf := v1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name: opts.WorkflowExecution.Name,
		},
		Spec: v1alpha1.WorkflowSpec{
			Entrypoint: "entrypoint",
			Arguments: v1alpha1.Arguments{
				Parameters: []v1alpha1.Parameter{
					{
						Name:  "server",
						Value: v1alpha1.AnyStringPtr(serverAddress),
					},
					{
						Name:  "token",
						Value: v1alpha1.AnyStringPtr(token),
					},
					{
						Name:  "tlsVerify",
						Value: v1alpha1.AnyStringPtr(tlsVerify),
					},
				},
			},
			Templates: workflowTemplates,
		},
	}

	_, err = s.cs.Create(ctx, &wf, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArgoWorkflowClient) Delete(ctx context.Context, opts DeleteOptions) error {
	// 1. Delete workflow from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowClient) Get(ctx context.Context, opts GetOptions) (*v1alpha1.Workflow, error) {
	// 1. Get workflow from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowClient) List(ctx context.Context, opts ListOptions) (*v1alpha1.WorkflowList, error) {
	// 1. List workflows from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowClient) GenerateWorkflowTemplateEntrypoint(
	ctx context.Context,
	stageExecutions model.WorkflowStageExecutions,
) ([]*v1alpha1.Template, error) {
	workflowTplLen := 1
	for _, stageExec := range stageExecutions {
		workflowTplLen += (len(stageExec.Edges.StepExecutions) + 1) * 3
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
			taskName         = stageExec.ID.String() + "-stage-task"
			taskTemplateName = stageExec.ID.String()
		)
		entrypoint.DAG.Tasks = append(entrypoint.DAG.Tasks, v1alpha1.DAGTask{
			Name:         taskName,
			Template:     taskTemplateName,
			Dependencies: dependencies,
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
		Name: stageExecution.ID.String(),
		DAG:  &v1alpha1.DAGTemplate{},
	}

	beforeTemplate := &v1alpha1.Template{
		Name: stageExecution.ID.String() + "-before",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(stageExecution.ID.String()),
				},
				{
					Name:  "status",
					Value: v1alpha1.AnyStringPtr("running"),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL:    "{{workflow.parameters.server}}/v1/workflow-stage-executions",
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
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition: "status == '200'",
		},
	}
	afterTemplate := &v1alpha1.Template{
		Name: stageExecution.ID.String() + "-after",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(stageExecution.ID.String()),
				},
				{
					Name:  "status",
					Value: v1alpha1.AnyStringPtr("ready"),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL:    "{{workflow.parameters.server}}/v1/workflow-stage-executions",
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
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition: "status == '200'",
		},
	}

	stageTemplates := []*v1alpha1.Template{
		beforeTemplate,
		stageTemplate,
		afterTemplate,
	}

	tasks := make([]v1alpha1.DAGTask, 0, len(stageExecution.Edges.StepExecutions))
	stepTemplates := make([]*v1alpha1.Template, 0)

	for _, stepExec := range stageExecution.Edges.StepExecutions {
		stepTemplateMap, err := s.GenerateStepTemplate(ctx, stepExec)
		if err != nil {
			return nil, err
		}

		for key, stepTemplate := range stepTemplateMap {
			stepTemplates = append(stepTemplates, stepTemplate)
			if key == mainTemplateKey {
				taskName := stepTemplate.Name + "-task"
				tasks = append(tasks, v1alpha1.DAGTask{
					Name:     taskName,
					Template: stepTemplate.Name,
					Hooks: v1alpha1.LifecycleHooks{
						"running": v1alpha1.LifecycleHook{
							Template:   beforeTemplateKey,
							Expression: fmt.Sprintf("tasks['%s'].status==\"Running\"", taskName),
						},
						"finished": v1alpha1.LifecycleHook{
							Template: afterTemplateKey,
							Expression: fmt.Sprintf("tasks['%s'].status==\"Succeeded\" ||"+
								" tasks['%s'].status==\"Failed\"", taskName, taskName),
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
		Name: wse.ID.String() + "-before",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(wse.ID.String()),
				},
				{
					Name:  "status",
					Value: v1alpha1.AnyStringPtr("running"),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL:    "{{workflow.parameters.server}}/v1/workflow-step-executions",
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
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition: "status == '200'",
		},
	}
	afterTemplate := &v1alpha1.Template{
		Name: wse.ID.String() + "-after",
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(wse.ID.String()),
				},
				{
					Name:  "status",
					Value: v1alpha1.AnyStringPtr("ready"),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL:    "{{workflow.parameters.server}}/v1/workflow-step-executions",
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
			Body: `{
				"id": "{{inputs.parameters.id}}",
				"status": "{{inputs.parameters.status}}"
			}`,
			SuccessCondition: "status == '200'",
		},
	}

	stepService, err := step.GetStepManager(steptypes.CreateOptions{
		Type:        steptypes.StepTypeService,
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
