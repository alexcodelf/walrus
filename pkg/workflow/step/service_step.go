package step

import (
	"github.com/seal-io/walrus/pkg/dao/model"
)

// ConfigService is service to generate service configs.
type ConfigService struct {
	mc model.ClientSet
}

// NewConfigService.
func NewConfigService(mc model.ClientSet) *ConfigService {
	return &ConfigService{
		mc: mc,
	}
}
