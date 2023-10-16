package workflowstepexecution

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	return h.modelClient.WorkflowStepExecutions().UpdateOne(entity).
		SetRecord(req.Record).
		SetDuration(req.Duration).
		Exec(req.Context)
}
