package deployers

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/resourceruns/deployer"
	"github.com/seal-io/walrus/pkg/resourceruns/deployers/terraform"
)

func NewDeployer(opts deployer.CreateOptions) (deployer.Deployer, error) {
	switch opts.Type {
	case deployer.TypeTerraform:
		return terraform.NewDeployer(opts)
	default:
		return nil, fmt.Errorf("unknown deployer type: %s", opts.Type)
	}
}
