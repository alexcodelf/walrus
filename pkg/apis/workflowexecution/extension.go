package workflowexecution

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	pkgworkflow "github.com/seal-io/walrus/pkg/workflow"
)

func (h Handler) RouteResubmitRequest(req RouteResubmitRequest) error {
	entity, err := h.modelClient.WorkflowExecutions().Query().
		Where(workflowexecution.ID(req.ID)).
		Only(req.Context)
	if err != nil {
		return err
	}

	if status.WorkflowExecutionStatusPending.IsUnknown(entity) ||
		status.WorkflowExecutionStatusRunning.IsUnknown(entity) {
		return fmt.Errorf("workflow execution is pending or running")
	}

	return h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		return pkgworkflow.Resubmit(req.Context, h.modelClient, h.k8sConfig, entity)
	})
}
