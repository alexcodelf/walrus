package helm

import (
	"context"
	"errors"
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
	"github.com/seal-io/walrus/pkg/terraform/convertors/k8s"
	"github.com/seal-io/walrus/pkg/terraform/util"
)

// HelmConvertor mutate the types.ConnectorTypeKubernetes connector to TypeHelm provider block.
type HelmConvertor string

func (m HelmConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Type == resourcehandler.ConnectorTypeKubernetes
}

func (m HelmConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
	var blocks block.Blocks

	for _, c := range connectors {
		if !m.IsSupported(c) {
			continue
		}

		b, err := m.toBlock(ctx, c, opts)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, b)
	}

	return blocks, nil
}

func (m HelmConvertor) toBlock(ctx context.Context, conn walruscore.Connector, opts convertor.Options) (*block.Block, error) {
	k8sOpts, ok := opts.(k8s.K8sConvertorOptions)
	if !ok {
		return nil, errors.New("invalid k8s options")
	}

	if conn.Spec.Type != resourcehandler.ConnectorTypeKubernetes {
		return nil, fmt.Errorf("connector type is not k8s, connector: %s", conn.Name)
	}

	var (
		// NB(alex) the config path should keep the same with the secret mount path in deployer.
		configPath = k8sOpts.ConfigPath + "/" + util.GetK8sSecretName(conn.Name)
		alias      = k8sOpts.ConnSeparator + conn.Name
		attributes = map[string]any{
			"alias": alias,
		}
	)

	_, err := convertor.LoadK8sApiConfig(ctx, conn)
	if err != nil {
		return nil, err
	}

	// Helm provider need a kubernetes block.
	// It is not a regular attribute of the helm provider.
	// E.g.
	// Provider "helm" {
	// 	kubernetes {
	// 		config_path = "xxx"
	// 	}
	// }.

	var (
		helmBlock = &block.Block{
			Type:       block.TypeProvider,
			Attributes: attributes,
			// Convert the connector type to provider type.
			Labels: []string{string(m)},
		}
		k8sBlock = &block.Block{
			Type: block.TypeK8s,
			Attributes: map[string]any{
				"config_path": configPath,
			},
		}
	)

	helmBlock.AppendBlock(k8sBlock)

	return helmBlock, nil
}
