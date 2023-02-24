package dao

import (
	"errors"

	"github.com/seal-io/seal/pkg/dao/model"
)

func ConnectorCreates(mc model.ClientSet, input ...*model.Connector) ([]*model.ConnectorCreate, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid input: empty list")
	}

	var rrs = make([]*model.ConnectorCreate, len(input))
	for i := range input {
		r := input[i]
		if r == nil {
			return nil, errors.New("invalid input: nil entity")
		}

		// required.
		var c = mc.Connectors().Create().
			SetName(r.Name).
			SetType(r.Type).
			SetConfigVersion(r.ConfigVersion).
			SetConfigData(r.ConfigData).
			SetEnableFinOps(r.EnableFinOps)

		// optional.
		c.SetDescription(r.Description)
		if r.Labels != nil {
			c.SetLabels(r.Labels)
		}
		rrs[i] = c
	}
	return rrs, nil
}

func ConnectorUpdate(mc model.ClientSet, input *model.Connector) (*model.ConnectorUpdateOne, error) {
	if input == nil {
		return nil, errors.New("invalid input: nil entity")
	}

	// predicated.
	if input.ID == "" {
		return nil, errors.New("invalid input: illegal predicates")
	}

	// conditional.
	var c = mc.Connectors().UpdateOne(input).
		SetDescription(input.Description).
		SetEnableFinOps(input.EnableFinOps).
		SetStatus(input.Status).
		SetStatusMessage(input.StatusMessage)
	if input.Name != "" {
		c.SetName(input.Name)
	}
	if input.Labels != nil {
		c.SetLabels(input.Labels)
	}
	if input.ConfigVersion != "" {
		c.SetConfigVersion(input.ConfigVersion)
	}
	if input.ConfigData != nil {
		c.SetConfigData(input.ConfigData)
	}
	return c, nil
}