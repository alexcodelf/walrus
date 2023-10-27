package workflow

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

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
)

const (
	templateTypeStep  = "step"
	templateTypeStage = "stage"

	templateEnter = "enter"
	templateMain  = "main"
	templateExit  = "exit"

	workflowEntrypointTemplateName = "entrypoint"
	workflowEnterTemplateName      = "workflowEnter"
	workflowExitTemplateName       = "workflowExit"

	executionUpdateURL = "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
		"/workflows/{{workflow.parameters.workflowID}}" +
		"/executions/{{inputs.parameters.id}}"
	stageExecutionUpdateURL = "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
		"/workflows/{{workflow.parameters.workflowID}}" +
		"/executions/{{inputs.parameters.workflowExecutionID}}" +
		"/stage-executions/{{inputs.parameters.id}}"
	stepExecutionUpdateURL = "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
		"/workflows/{{workflow.parameters.workflowID}}" +
		"/executions/{{inputs.parameters.workflowExecutionID}}" +
		"/stage-executions/{{inputs.parameters.workflowStageExecutionID}}" +
		"/step-executions/{{inputs.parameters.id}}"
)

const (
	statusRequestBody = `{
	"id": "{{inputs.parameters.id}}",
	"status": "{{inputs.parameters.status}}"
}`
)

var (
	limit                     = intstr.FromInt(2)
	statusUpdateRetryStrategy = &v1alpha1.RetryStrategy{
		Limit:       &limit,
		RetryPolicy: v1alpha1.RetryPolicyOnFailure,
	}
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
) (*v1alpha1.Workflow, error) {
	// Prepare API token for terraform backend.
	const _1Day = 60 * 60 * 24

	s, err := auths.GetBotSubject(ctx, t.mc)
	if err != nil {
		return nil, err
	}

	// TODO using bot subject to create token.
	at, err := auths.CreateAccessToken(ctx,
		t.mc, s.ID, types.TokenKindDeployment, workflowExecution.ID.String(), pointer.Int(_1Day))
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

	workflowTemplates := make([]v1alpha1.Template, 0, len(wfTemplates)+2)
	for _, tpl := range wfTemplates {
		workflowTemplates = append(workflowTemplates, *tpl)
	}

	enterHandler := t.GetWorkflowExecutionEnterTemplate(workflowExecution)
	exitHandler := t.GetWorkflowExecutionExitTemplate(workflowExecution)
	workflowTemplates = append(workflowTemplates, *exitHandler)
	workflowTemplates = append(workflowTemplates, *enterHandler)

	wf := &v1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", workflowExecution.Name, workflowExecution.ID.String()),
		},
		Spec: v1alpha1.WorkflowSpec{
			Entrypoint: workflowEntrypointTemplateName,
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
				SecondsAfterCompletion: pointer.Int32(600),
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
	// Calculate the length of workflow templates.
	workflowTemplateLen := 2
	for _, stageExec := range stageExecutions {
		workflowTemplateLen += (len(stageExec.Edges.Steps) + 1) * 4
	}

	workflowTemplates := make([]*v1alpha1.Template, 0, workflowTemplateLen)
	tasks := make([]v1alpha1.DAGTask, len(stageExecutions)+1)
	entrypoint := &v1alpha1.Template{
		Name: workflowEntrypointTemplateName,
		DAG: &v1alpha1.DAGTemplate{
			Tasks: tasks,
		},
	}
	workflowTemplates = append(workflowTemplates, entrypoint)

	tasks[0] = v1alpha1.DAGTask{
		Name:     workflowEnterTemplateName,
		Template: workflowEnterTemplateName,
	}

	for i, stageExec := range stageExecutions {
		stageExtendTemplate, stageTemplates, err := t.GetStageExecutionExtendTemplates(ctx, stageExec)
		if err != nil {
			return nil, err
		}

		workflowTemplates = append(workflowTemplates, stageExtendTemplate)
		workflowTemplates = append(workflowTemplates, stageTemplates...)

		var dependencies []string

		// Add previous stage task as dependency.
		dependencies = append(dependencies, entrypoint.DAG.Tasks[i].Name)

		entrypoint.DAG.Tasks[i+1] = v1alpha1.DAGTask{
			Name:         templateName(stageExec.ID, templateTypeStage),
			Template:     stageExtendTemplate.Name,
			Dependencies: dependencies,
		}
	}

	return workflowTemplates, nil
}

func (t *TemplateManager) GetWorkflowExecutionStatusTemplate(
	name string,
	wf *model.WorkflowExecution,
) *v1alpha1.Template {
	status := "{{workflow.status}}"
	if name == workflowEnterTemplateName {
		status = "Running"
	}

	return &v1alpha1.Template{
		Name: name,
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "id",
					Value: v1alpha1.AnyStringPtr(wf.ID.String()),
				},
				{
					Name:  "status",
					Value: v1alpha1.AnyStringPtr(status),
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL: "{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}" +
				"/workflows/{{workflow.parameters.workflowID}}" +
				"/executions/{{inputs.parameters.id}}",
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
			Body:               statusRequestBody,
		},
		RetryStrategy: statusUpdateRetryStrategy,
	}
}

func (t *TemplateManager) GetWorkflowExecutionEnterTemplate(wf *model.WorkflowExecution) *v1alpha1.Template {
	return t.GetWorkflowExecutionStatusTemplate(workflowEnterTemplateName, wf)
}

// getExitTemplate returns template for workflow exit handler.
func (t *TemplateManager) GetWorkflowExecutionExitTemplate(wf *model.WorkflowExecution) *v1alpha1.Template {
	return t.GetWorkflowExecutionStatusTemplate(workflowExitTemplateName, wf)
}

// GetStageExecutionExtendTemplates extends one stage execution to three stage executions,
// enter template, main template, and exit template and its step templates.
// The extend templates are used to manager lifecycle of the stage execution.
func (t *TemplateManager) GetStageExecutionExtendTemplates(
	ctx context.Context,
	stageExecution *model.WorkflowStageExecution,
) (extendTemplate *v1alpha1.Template, stageTemplates []*v1alpha1.Template, err error) {
	stageTemplates, err = t.GetStageExecutionTemplates(ctx, stageExecution)
	if err != nil {
		return nil, nil, err
	}

	main := statusTemplateName(stageExecution.ID, templateTypeStage, templateMain)

	// Extend one stage template to three stage templates, enter template, main template,
	// and exit template.
	extendTemplate = &v1alpha1.Template{
		Name: fmt.Sprintf("stage-execution-%s-extend", stageExecution.ID.String()),
		Steps: []v1alpha1.ParallelSteps{
			{
				Steps: []v1alpha1.WorkflowStep{
					{
						Name:     statusTemplateName(stageExecution.ID, templateTypeStage, templateEnter),
						Template: stageTemplates[0].Name,
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
						Name:     main,
						Template: stageTemplates[1].Name,
						Hooks: v1alpha1.LifecycleHooks{
							"failed": getLifecycleHook(
								stageTemplates[2].Name,
								"Failed",
								fmt.Sprintf("steps[\"%s\"].status == \"Failed\"",
									main,
								),
							),
							"succeeded": getLifecycleHook(
								stageTemplates[2].Name,
								"Succeeded",
								fmt.Sprintf("steps[\"%s\"].status == \"Succeeded\"",
									main,
								),
							),
							"error": getLifecycleHook(
								stageTemplates[2].Name,
								"Error",
								fmt.Sprintf("steps[\"%s\"].status == \"Error\"",
									main,
								),
							),
						},
					},
				},
			},
		},
	}

	return extendTemplate, stageTemplates, nil
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

		// StageTemplates stores tempalate of the stage own.
		stageTemplates = []*v1alpha1.Template{
			enterTemplate,
			stageTemplate,
			exitTemplate,
		}

		// StageStepTemplates stores all step templates of the stage.
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
			URL:    stageExecutionUpdateURL,
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
			TimeoutSeconds:     pointer.Int64(5),
			Body:               statusRequestBody,
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
		},
		RetryStrategy: statusUpdateRetryStrategy,
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
// The extend templates are used to manager lifecycle of the step execution.
func (t *TemplateManager) GetStepExecutionExtendTemplates(
	ctx context.Context,
	stepExecution *model.WorkflowStepExecution,
) (extendTemplate *v1alpha1.Template, stepTemplates []*v1alpha1.Template, err error) {
	stepTemplates, err = t.GetStepExecutionTemplates(ctx, stepExecution)
	if err != nil {
		return nil, nil, err
	}

	main := statusTemplateName(stepExecution.ID, templateTypeStep, templateMain)
	// Extend one step template to three step templates, enter template, main template,
	// and exit template.
	extendTemplate = &v1alpha1.Template{
		Name: fmt.Sprintf("step-execution-%s-extend", stepExecution.ID.String()),
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
						Hooks: v1alpha1.LifecycleHooks{
							"succeeded": getLifecycleHook(
								stepTemplates[2].Name,
								"Succeeded",
								fmt.Sprintf("steps[\"%s\"].status == \"Succeeded\"",
									main,
								),
							),
							"failed": getLifecycleHook(
								stepTemplates[2].Name,
								"Failed",
								fmt.Sprintf("steps[\"%s\"].status == \"Failed\"",
									main,
								),
							),
							"error": getLifecycleHook(
								stepTemplates[2].Name,
								"Error",
								fmt.Sprintf("steps[\"%s\"].status == \"Error\"",
									main,
								),
							),
						},
					},
				},
			},
		},
	}

	return extendTemplate, stepTemplates, nil
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
			URL:    stepExecutionUpdateURL,
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
			TimeoutSeconds:     pointer.Int64(5),
			Body:               statusRequestBody,
			SuccessCondition:   "response.statusCode >= 200 && response.statusCode < 300",
			InsecureSkipVerify: !apiconfig.TlsCertified.Get(),
		},
		RetryStrategy: statusUpdateRetryStrategy,
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

// getLifecycleHook returns a lifecycle hook of target tasks or steps.
func getLifecycleHook(templateName, status, expression string) v1alpha1.LifecycleHook {
	return v1alpha1.LifecycleHook{
		Template: templateName,
		Arguments: v1alpha1.Arguments{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "status",
					Value: v1alpha1.AnyStringPtr(status),
				},
			},
		},
		Expression: expression,
	}
}
