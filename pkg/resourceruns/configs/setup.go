package configs

import (
	"github.com/seal-io/walrus/pkg/resourceruns/config"
	"github.com/seal-io/walrus/pkg/resourceruns/configs/terraform"
)

// NewConfigurator creates a new configurator with the deployer type.
func NewConfigurator(templateFormat string) config.Configurator {
	switch templateFormat {
	case config.ConfigTypeTerraform:
		return terraform.NewConfigurator()
	default:
		return nil
	}
}
