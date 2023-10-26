package workflow

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
	"github.com/seal-io/walrus/pkg/auths"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step"
	steptypes "github.com/seal-io/walrus/pkg/workflow/step/types"
	"github.com/seal-io/walrus/utils/pointer"
	"github.com/seal-io/walrus/utils/strs"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	templateTypeStep  = "step"
	templateTypeStage = "stage"

	templateEnter = "enter"
	templateMain  = "main"
	templateExit  = "exit"

	workflowEntrypointTemplateName = "entrypoint"
	workflowExitTemplateName       = "workflowExit"
)

// TemplateManager is the manager of workflow templates.
// It generates templates for workflow with workflow, stage and step executions.
type TemplateManager struct {
	mc model.ClientSet
}

func NewTemplateManager(mc model.ClientSet) *TemplateManager {
	return &TemplateManager{
		mc: mc,
	}
}

// GetWorkflowExecutionWorkflow returns an argo workflow for a workflow execution.
func (t *TemplateManager) GetWorkflowExecutionWorkflow(
	ctx context.Context,
	workflowExecution *model.WorkflowExecution,
	subjectID object.ID,
) (*v1alpha1.Workflow, error) {
	// Prepare API token for terraform backend.
	const _1Day = 60 * 60 * 24

	// TODO using bot subject to create token.
	at, err := auths.CreateAccessToken(ctx,
		t.mc, subjectID, types.TokenKindDeployment, workflowExecution.ID.String(), pointer.Int(_1Day))
	if err != nil {
		return nil, err
	}

	token := at.AccessToken

	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, t.mc)
	if err != nil {
		return nil, err
	}

	if serverAddress == "" {
		return nil, errors.New("server address is empty")
	}

	wfTemplates, err := t.GetWorkflowExecutionTemplates(ctx, workflowExecution.Edges.Stages)
	if err != nil {
		return nil, err
	}

	workflowTemplates := make([]v1alpha1.Template, 0, len(wfTemplates)+1)
	for _, tpl := range wfTemplates {
		workflowTemplates = append(workflowTemplates, *tpl)
	}

	exitHandler := t.GetWorkflowExecutionExitTemplate(workflowExecution)
	workflowTemplates = append(workflowTemplates, *exitHandler)

	wf := &v1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", workflowExecution.Name, workflowExecution.ID.String()),
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
						Value: v1alpha1.AnyStringPtr(workflowExecution.ProjectID.String()),
					},
					{
						Name:  "workflowID",
						Value: v1alpha1.AnyStringPtr(workflowExecution.WorkflowID.String()),
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

	return wf, nil
}

// GetWorkflowExecutionTemplates get workflow execution templates.
func (t *TemplateManager) GetWorkflowExecutionTemplates(
	ctx context.Context,
	stageExecutions model.WorkflowStageExecutions,
) ([]*v1alpha1.Template, error) {
	workflowTemplateLen := 1
	for _, stageExec := range stageExecutions {
		workflowTemplateLen += (len(stageExec.Edges.Steps) + 1) * 3
	}

	workflowTemplates := make([]*v1alpha1.Template, 0, workflowTemplateLen)
	entrypoint := &v1alpha1.Template{
		Name: workflowEntrypointTemplateName,
		DAG:  &v1alpha1.DAGTemplate{},
	}
	workflowTemplates = append(workflowTemplates, entrypoint)

	for i, stageExec := range stageExecutions {
		stageTemplates, err := t.GetStageExecutionTemplates(ctx, stageExec)
		if err != nil {
			return nil, err
		}

		workflowTemplates = append(workflowTemplates, stageTemplates...)

		var dependencies []string
		if i > 0 {
			// Add previous stage task as dependency.
			dependencies = append(dependencies, entrypoint.DAG.Tasks[i-1].Name)
		}

		var (
			taskName         = templateName(stageExec.ID, templateTypeStage)
			taskTemplateName = templateName(stageExec.ID, templateTypeStage)
		)

		entrypoint.DAG.Tasks = append(entrypoint.DAG.Tasks, v1alpha1.DAGTask{
			Name:         taskName,
			Template:     taskTemplateName,
			Dependencies: dependencies,
			Hooks: v1alpha1.LifecycleHooks{
				templateEnter: v1alpha1.LifecycleHook{
					Template:   statusTemplateName(stageExec.ID, templateTypeStage, templateEnter),
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
				templateExit: v1alpha1.LifecycleHook{
					Template: statusTemplateName(stageExec.ID, templateTypeStage, templateExit),
					Expression: fmt.Sprintf(
						"tasks['%s'].status == \"Succeeded\" || "+
							"tasks['%s'].status == \"Failed\" || "+
							"tasks['%s'].status == \"Error\"",
						taskName, taskName, taskName),
					Arguments: v1alpha1.Arguments{
						Parameters: []v1alpha1.Parameter{
							{
								Name:  "status",
								Value: v1alpha1.AnyStringPtr(fmt.Sprintf("{{tasks.%s.status}}", taskName)),
							},
						},
					},
				},
			},
		})
	}

	return workflowTemplates, nil
}

// getExitTemplate returns template for workflow exit handler.
func (t *TemplateManager) GetWorkflowExecutionExitTemplate(wf *model.WorkflowExecution) *v1alpha1.Template {
	limit := intstr.FromInt(3)
	return &v1alpha1.Template{
		Name: workflowExitTemplateName,
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
		RetryStrategy: &v1alpha1.RetryStrategy{
			RetryPolicy: v1alpha1.RetryPolicyOnFailure,
			Limit:       &limit,
		},
	}
}

// GetStageExecutionTemplates extends one stage execution to three stage executions,
// enter template, main template, and exit template and its step templates.
func (t *TemplateManager) GetStageExecutionTemplates(
	ctx context.Context,
	stageExecution *model.WorkflowStageExecution,
) ([]*v1alpha1.Template, error) {
	var (
		enterTemplate = t.GetStageExecutionEnterTemplate(stageExecution)
		exitTemplate  = t.GetStageExecutionExitTemplate(stageExecution)
		stageTemplate = &v1alpha1.Template{
			Name: templateName(stageExecution.ID, templateTypeStage),
			DAG:  &v1alpha1.DAGTemplate{},
		}

		// stageTemplates stores tempalate of the stage own.
		stageTemplates = []*v1alpha1.Template{
			enterTemplate,
			stageTemplate,
			exitTemplate,
		}

		// stageStepTemplates stores all step templates of the stage.
		stageStepTemplates = make([]*v1alpha1.Template, 0)
	)

	tasks := make([]v1alpha1.DAGTask, 0, len(stageExecution.Edges.Steps))

	// Get step templates with step executions.
	for _, stepExecution := range stageExecution.Edges.Steps {
		extendTemplate, stepTemplates, err := t.GetStepExecutionExtendTemplates(ctx, stepExecution)
		if err != nil {
			return nil, err
		}

		stageStepTemplates = append(stageStepTemplates, stepTemplates...)
		stageStepTemplates = append(stageStepTemplates, extendTemplate)

		tasks = append(tasks, v1alpha1.DAGTask{
			Name:     "step-execution-" + stepExecution.ID.String(),
			Template: extendTemplate.Name,
		})
	}

	stageTemplate.DAG.Tasks = tasks

	// Add step templates to templates list.
	stageTemplates = append(stageTemplates, stageStepTemplates...)

	return stageTemplates, nil
}

// GetStageExecutionStatusTemplate returns the status template of a stage execution.
// The status template handler sync the status of the stage execution to "Running", "Succeeded" or "Failed".
func (t *TemplateManager) GetStageExecutionStatusTemplate(
	name string,
	stageExecution *model.WorkflowStageExecution,
) *v1alpha1.Template {
	limit := intstr.FromInt(3)
	return &v1alpha1.Template{
		Name: name,
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
		RetryStrategy: &v1alpha1.RetryStrategy{
			RetryPolicy: v1alpha1.RetryPolicyOnFailure,
			Limit:       &limit,
		},
	}
}

// GetStageExecutionEnterTemplate returns the enter template of a stage execution.
// The template handler sync the status of the stage execution to "Running".
func (t *TemplateManager) GetStageExecutionEnterTemplate(
	stageExecution *model.WorkflowStageExecution,
) *v1alpha1.Template {
	return t.GetStageExecutionStatusTemplate(
		statusTemplateName(
			stageExecution.ID,
			templateTypeStage,
			templateEnter,
		),
		stageExecution,
	)
}

// GetStageExecutionExitTemplate returns the exit template of a stage execution.
// The template handler sync the status of the stage execution to "Succeeded" or "Failed".
func (t *TemplateManager) GetStageExecutionExitTemplate(
	stageExecution *model.WorkflowStageExecution,
) *v1alpha1.Template {
	return t.GetStageExecutionStatusTemplate(
		statusTemplateName(
			stageExecution.ID,
			templateTypeStage,
			templateExit,
		),
		stageExecution,
	)
}

// GetStepExecutionExtendTemplates extends one step execution to three step executions, enter template, main template,
// exit step template, which are used to update the status of the step execution.
func (t *TemplateManager) GetStepExecutionExtendTemplates(
	ctx context.Context,
	stepExecution *model.WorkflowStepExecution,
) (extendTemplate *v1alpha1.Template, stepTemplates []*v1alpha1.Template, err error) {
	stepTemplates, err = t.GetStepExecutionTemplates(ctx, stepExecution)
	if err != nil {
		return nil, nil, err
	}

	// Extend one step template to three step templates, enter template, main template,
	// and exit template.
	extendTemplate = &v1alpha1.Template{
		Name: fmt.Sprintf("step-execution-%s-", stepExecution.ID.String()),
		Steps: []v1alpha1.ParallelSteps{
			{
				Steps: []v1alpha1.WorkflowStep{
					{
						Name:     statusTemplateName(stepExecution.ID, templateTypeStep, templateEnter),
						Template: stepTemplates[0].Name,
						Arguments: v1alpha1.Arguments{
							Parameters: []v1alpha1.Parameter{
								{
									Name:  "status",
									Value: v1alpha1.AnyStringPtr("Running"),
								},
							},
						},
					},
				},
			},
			{
				Steps: []v1alpha1.WorkflowStep{
					{
						Name:     statusTemplateName(stepExecution.ID, templateTypeStep, templateMain),
						Template: stepTemplates[1].Name,
					},
				},
			},
			{
				Steps: []v1alpha1.WorkflowStep{
					{
						Name:     statusTemplateName(stepExecution.ID, templateTypeStep, templateExit),
						Template: stepTemplates[2].Name,
						Arguments: v1alpha1.Arguments{
							Parameters: []v1alpha1.Parameter{
								{
									Name: "status",
									Value: v1alpha1.AnyStringPtr(
										fmt.Sprintf("{{steps.%s.status}}",
											statusTemplateName(stepExecution.ID, templateTypeStep, templateMain))),
								},
							},
						},
					},
				},
			},
		},
	}

	return
}

// GetStepExecutionTemplates extends one step execution to three step executions, enter template, main template,
// exit step template, which are used to update the status of the step execution.
func (t *TemplateManager) GetStepExecutionTemplates(
	ctx context.Context,
	stepExecution *model.WorkflowStepExecution,
) ([]*v1alpha1.Template, error) {
	// Get enter template.
	enterTemplate := t.stepExecutionEnterTemplate(stepExecution)
	exitTemplate := t.stepExecutionExitTemplate(stepExecution)

	stepService, err := step.GetStepManager(steptypes.CreateOptions{
		Type:        steptypes.Type(stepExecution.Type),
		ModelClient: t.mc,
	})
	if err != nil {
		return nil, err
	}

	// Generate service template.
	mainTemplate, err := stepService.GenerateTemplate(ctx, stepExecution)
	if err != nil {
		return nil, err
	}

	return []*v1alpha1.Template{
		enterTemplate,
		mainTemplate,
		exitTemplate,
	}, nil
}

// GetStepExecutionStatusTemplate returns the status template of a step execution.
// The status template handler sync the status of the step execution to
// "Running", "Succeeded" or "Failed".
func (t *TemplateManager) GetStepExecutionStatusTemplate(
	name string,
	stepExecution *model.WorkflowStepExecution,
) *v1alpha1.Template {
	limit := intstr.FromInt(3)
	return &v1alpha1.Template{
		Name: name,
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(stepExecution.ID.String()),
				},
				{
					Name:  "workflowStageExecutionID",
					Value: v1alpha1.AnyStringPtr(stepExecution.WorkflowStageExecutionID.String()),
				},
				{
					Name:  "workflowExecutionID",
					Value: v1alpha1.AnyStringPtr(stepExecution.WorkflowExecutionID.String()),
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
		RetryStrategy: &v1alpha1.RetryStrategy{
			RetryPolicy: v1alpha1.RetryPolicyOnFailure,
			Limit:       &limit,
		},
	}
}

// stepExecutionEnterTemplate returns the enter template of a step execution. The enter template is used to
// update the status of the step execution to "Running".
func (t *TemplateManager) stepExecutionEnterTemplate(stepExecution *model.WorkflowStepExecution) *v1alpha1.Template {
	return t.GetStepExecutionStatusTemplate(
		statusTemplateName(
			stepExecution.ID,
			templateTypeStep,
			templateEnter,
		),
		stepExecution,
	)
}

// stepExecutionExitTemplate returns the exit template of a step execution. The exit template is used to
// update the status of the step execution to "Succeeded" or "Failed".
func (t *TemplateManager) stepExecutionExitTemplate(stepExecution *model.WorkflowStepExecution) *v1alpha1.Template {
	return t.GetStepExecutionStatusTemplate(
		statusTemplateName(
			stepExecution.ID,
			templateTypeStep,
			templateExit,
		),
		stepExecution,
	)
}

func statusTemplateName(id object.ID, templateType, stage string) string {
	return strs.Join("-", templateName(id, templateType), stage)
}

func templateName(id object.ID, templateType string) string {
	return strs.Join("-", templateType, id.String())
}
