package workflow

import (
	"context"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func Create(ctx context.Context, mc model.ClientSet, wf *model.Workflow) error {
	return mc.WithTx(ctx, func(tx *model.Tx) error {
		workflow, err := tx.Workflows().Create().
			Set(wf).
			Save(ctx)
		if err != nil {
			return err
		}

		stages := make(model.WorkflowStages, 0, len(wf.Edges.Stages))

		for _, stage := range wf.Edges.Stages {
			stage.WorkflowID = workflow.ID

			st, err := tx.WorkflowStages().Create().
				Set(stage).
				Save(ctx)
			if err != nil {
				return err
			}

			stages = append(stages, st)
		}

		// Update workflow stages.
		workflow.Edges.Stages = stages
		wf = workflow

		return nil
	})
}

func Apply(ctx context.Context, mc model.ClientSet, clientConfig clientcmd.ClientConfig, wf *model.Workflow) error {
	client, err := NewArgoWorkflowClient(mc, clientConfig)
	if err != nil {
		return err
	}

	s := session.MustGetSubject(ctx)

	wfe, err := CreateWorkflowExecution(ctx, mc, wf)
	if err != nil {
		return err
	}

	return client.Submit(ctx, SubmitOptions{
		WorkflowExecution: wfe,
		SubjectID:         s.ID,
	})
}

func CreateWorkflowExecution(
	ctx context.Context,
	mc model.ClientSet,
	wf *model.Workflow,
) (*model.WorkflowExecution, error) {
	s := session.MustGetSubject(ctx)

	// TODO check if the workflow is already running.

	// Create workflow execution.
	progress := fmt.Sprintf("%d/%d", 0, len(wf.Edges.Stages))

	workflowExecution := &model.WorkflowExecution{
		Name:       fmt.Sprintf("%s-%d", wf.Name, time.Now().Unix()),
		ProjectID:  wf.ProjectID,
		WorkflowID: wf.ID,
		Progress:   progress,
		SubjectID:  s.ID,
	}

	status.WorkflowExecutionStatusPending.Unknown(workflowExecution, "")
	workflowExecution.Status.SetSummary(status.WalkWorkflowExecution(&workflowExecution.Status))

	entity, err := mc.WorkflowExecutions().Create().
		Set(workflowExecution).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stageExecutions := make(model.WorkflowStageExecutions, len(wf.Edges.Stages))

	for i, stage := range wf.Edges.Stages {
		// Create workflow stage execution.
		stageExecution, err := CreateWorkflowStageExecution(ctx, mc, stage, entity)
		if err != nil {
			return nil, err
		}

		stageExecutions[i] = stageExecution
	}

	entity.Edges.StageExecutions = stageExecutions

	return entity, nil
}

func CreateWorkflowStageExecution(
	ctx context.Context,
	mc model.ClientSet,
	stage *model.WorkflowStage,
	we *model.WorkflowExecution,
) (*model.WorkflowStageExecution, error) {
	stageExec := &model.WorkflowStageExecution{
		Name:                fmt.Sprintf("%s-%d", stage.Name, time.Now().Unix()),
		ProjectID:           stage.ProjectID,
		StageID:             stage.ID,
		WorkflowExecutionID: we.ID,
	}

	status.WorkflowStageStatusInitialized.Unknown(stageExec, "")
	stageExec.Status.SetSummary(status.WalkWorkflowStageExecution(&stageExec.Status))

	entity, err := mc.WorkflowStageExecutions().Create().
		Set(stageExec).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stepExecutions := make(model.WorkflowStepExecutions, len(stage.Edges.Steps))

	for i, step := range stage.Edges.Steps {
		// Create workflow step execution.
		stepExecution, err := CreateWorkflowStepExecution(ctx, mc, step, entity)
		if err != nil {
			return nil, err
		}

		stepExecutions[i] = stepExecution
	}

	entity.Edges.StepExecutions = stepExecutions

	return entity, nil
}

func CreateWorkflowStepExecution(
	ctx context.Context,
	mc model.ClientSet,
	step *model.WorkflowStep,
	wse *model.WorkflowStageExecution,
) (*model.WorkflowStepExecution, error) {
	stepExec := &model.WorkflowStepExecution{
		Name:                     fmt.Sprintf("%s-%d", step.Name, time.Now().Unix()),
		ProjectID:                step.ProjectID,
		WorkflowID:               step.WorkflowID,
		Type:                     step.Type,
		WorkflowStepID:           step.ID,
		WorkflowExecutionID:      wse.WorkflowExecutionID,
		WorkflowStageExecutionID: wse.ID,
		Spec:                     step.Spec,
	}

	status.WorkflowStepExecutionStatusPending.Unknown(stepExec, "")
	stepExec.Status.SetSummary(status.WalkWorkflowStepExecution(&stepExec.Status))

	return mc.WorkflowStepExecutions().Create().
		Set(stepExec).
		Save(ctx)
}

func CreateKubeconfigFileForRestConfig(restConfig *rest.Config) clientcmdapi.Config {
	clusters := make(map[string]*clientcmdapi.Cluster)
	clusters["default-cluster"] = &clientcmdapi.Cluster{
		Server:                   restConfig.Host,
		CertificateAuthorityData: restConfig.CAData,
	}
	contexts := make(map[string]*clientcmdapi.Context)
	contexts["default-context"] = &clientcmdapi.Context{
		Cluster:  "default-cluster",
		AuthInfo: "default-user",
	}
	authinfos := make(map[string]*clientcmdapi.AuthInfo)
	authinfos["default-user"] = &clientcmdapi.AuthInfo{
		ClientCertificateData: restConfig.CertData,
		ClientKeyData:         restConfig.KeyData,
	}
	clientConfig := clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clusters,
		Contexts:       contexts,
		CurrentContext: "default-context",
		AuthInfos:      authinfos,
	}

	return clientConfig
}
