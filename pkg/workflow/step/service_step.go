package step

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/vars"
)

var SkipTLSVerify = vars.SetOnce[bool]{}

// ConfigService is service to generate service configs.
type ConfigService struct {
	mc model.ClientSet
}

// NewConfigService
func NewConfigService(mc model.ClientSet) *ConfigService {
	return &ConfigService{
		mc: mc,
	}
}

type configServiceOptions struct {
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

func (cs ConfigService) loadConfigsBytes(ctx context.Context)
