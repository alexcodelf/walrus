package servicerevision

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type IPlan interface {
	// Plan plans the revision.
	LoadPlan(context.Context, PlanOptions) ([]byte, error)
	// LoadConfigs loads the plan configs of the plan options.
	LoadConfigs(context.Context, PlanOptions) (map[string][]byte, error)
	// LoadConnectorConfigs loads the connector configs of the plan options.
	// Some connectors may be required to deploy the service.
	LoadConnectorConfigs(model.Connectors) (map[string][]byte, error)
}

// PlanOptions are the options for planning a revision.
type PlanOptions struct {
	// SecretMountPath of the deploy job.
	SecretMountPath string

	ServiceRevision *model.ServiceRevision
	Connectors      model.Connectors
	ProjectID       object.ID
	EnvironmentID   object.ID
	SubjectID       object.ID
	// Metadata.
	ProjectName          string
	EnvironmentName      string
	ServiceName          string
	ServiceID            object.ID
	ManagedNamespaceName string
}

// NewPlan creates a new plan with the plan type.
func NewPlan(planType string, mc model.ClientSet) IPlan {
	switch planType {
	case types.DeployerTypeTF:
		return NewTerraformPlan(mc)
	default:
		return nil
	}
}
