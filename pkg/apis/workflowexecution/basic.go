package workflowexecution

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	return h.modelClient.WorkflowExecutions().UpdateOne(entity).
		SetDescription(req.Description).
		SetDuration(req.Duration).
		SetRecord(req.Record).
		Exec(req.Context)
}
