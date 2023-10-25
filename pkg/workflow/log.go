package workflow

import (
	"context"
	"errors"
	"io"

	"github.com/argoproj/argo-workflows/v3/pkg/apiclient"
	"github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"
	corev1 "k8s.io/api/core/v1"

	"github.com/seal-io/walrus/pkg/dao/types"
)

// StreamLogsOptions contains options for streaming workflow logs.
type StreamLogsOptions struct {
	Workflow   string
	PodName    string
	Grep       string
	Selector   string
	ApiClient  apiclient.Client
	LogOptions *corev1.PodLogOptions
	Out        io.Writer
}

// StreamWorkflowLogs streams workflow logs.
// With selector step-execution-id=stepExecutionID it can filter logs by step name.
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

func GetWorkflowTemplateLogs() {}
