package workflowstepexecution

import (
	"context"
	"fmt"
	"io"

	"github.com/argoproj/argo-workflows/v3/pkg/apiclient"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/k8s"
	"github.com/seal-io/walrus/pkg/workflow"
	"github.com/seal-io/walrus/utils/strs"
)

func (h Handler) RouteLog(req RouteLogRequest) error {
	var (
		ctx context.Context
		out io.Writer
	)

	if req.Stream != nil {
		ctx = req.Stream
		out = req.Stream
	} else {
		ctx = req.Context
		out = req.Context.Writer
	}

	wse, err := h.modelClient.WorkflowStepExecutions().Query().
		Where(workflowstepexecution.ID(req.ID)).
		Only(ctx)
	if err != nil {
		return err
	}

	workflowExec, err := h.modelClient.WorkflowExecutions().Query().
		Where(workflowexecution.ID(wse.WorkflowExecutionID)).
		Only(ctx)
	if err != nil {
		return err
	}

	apiConfig := k8s.ToClientCmdApiConfig(h.k8sConfig)
	clientConfig := clientcmd.NewDefaultClientConfig(apiConfig, nil)

	ctx, apiClient, err := apiclient.NewClientFromOpts(apiclient.Opts{
		ClientConfigSupplier: func() clientcmd.ClientConfig {
			return clientConfig
		},
		Context: ctx,
	})
	if err != nil {
		return err
	}

	return workflow.StreamWorkflowLogs(ctx, workflow.LogOptions{
		Workflow:  strs.Join(workflowExec.Name, workflowExec.ID.String()),
		ApiClient: apiClient,
		Selector:  fmt.Sprintf("step-execution-id=%s", wse.ID),
		LogOptions: &corev1.PodLogOptions{
			Container: "main",
		},
		Out: out,
	})
}

func (h Handler) RouteApprove(req RouteApproveRequest) error {
	stepExecution, err := h.modelClient.WorkflowStepExecutions().Query().
		Where(workflowstepexecution.ID(req.ID)).
		Only(req.Context)
	if err != nil {
		return err
	}

	client, err := workflow.NewArgoWorkflowClient(h.modelClient, h.k8sConfig)
	if err != nil {
		return err
	}

	return client.Resume(req.Context, workflow.ResumeOptions{
		WorkflowStepExecution: stepExecution,
	})
}
