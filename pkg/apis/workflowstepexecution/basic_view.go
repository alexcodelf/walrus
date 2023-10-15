package workflowstepexecution

import "fmt"

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	fmt.Println("entity", entity)

	return nil
	// return h.modelClient.WorkflowExecutions().UpdateOne(entity).
	// 	Set(entity).
	// 	Exec(req.Context)
}
