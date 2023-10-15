package workflowstageexecution

import "fmt"

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	fmt.Println("entity", entity)
	return nil
}
