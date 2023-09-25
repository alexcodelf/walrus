package servicerevision

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type IPlan interface {
	Plan(context.Context, PlanOptions) string
	LoadConfigs(context.Context, PlanOptions) (map[string][]byte, error)
}

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
