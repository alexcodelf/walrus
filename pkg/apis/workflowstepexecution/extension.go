package workflowstepexecution

import (
	"context"
	"fmt"
	"io"

	"github.com/argoproj/argo-workflows/v3/pkg/apiclient"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/workflow"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func (h Handler) RouteLog(req RouteLogRequest) error {
	wse, err := h.modelClient.WorkflowStepExecutions().Query().
		Where(workflowstepexecution.ID(req.ID)).
		Only(req.Context)
	if err != nil {
		return err
	}

	workflowExec, err := h.modelClient.WorkflowExecutions().Query().
		Where(workflowexecution.ID(wse.WorkflowExecutionID)).
		Only(req.Context)
	if err != nil {
		return err
	}

	var (
		ctx context.Context
		out io.Writer
	)

	if req.Stream == nil {
		ctx = req.Stream
		out = req.Stream
	} else {
		ctx = req.Context
		out = req.Context.Writer
	}

	apiConfig := workflow.CreateKubeconfigFileForRestConfig(h.k8sConfig)
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

	return workflow.StreamWorkflowLogs(ctx, workflow.StreamLogsOptions{
		Workflow:  workflowExec.Name,
		ApiClient: apiClient,
		Selector:  fmt.Sprintf("step-execution-id=%s", wse.ID),
		LogOptions: &corev1.PodLogOptions{
			Container: "main",
		},
		Out: out,
	})
}
