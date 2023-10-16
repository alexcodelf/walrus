package workflowstageexecution

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	return h.modelClient.WorkflowStageExecutions().UpdateOne(entity).
		SetRecord(req.Record).
		Exec(req.Context)
}
