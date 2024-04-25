package kubectl

import (
	"context"
	"errors"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
	"github.com/seal-io/walrus/pkg/terraform/convertors/k8s"
	"github.com/seal-io/walrus/pkg/terraform/util"
)

type KubectlConvertor string

func (m KubectlConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Type == resourcehandler.ConnectorTypeKubernetes
}

func (m KubectlConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
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

func (m KubectlConvertor) toBlock(ctx context.Context, connector walruscore.Connector, opts any) (*block.Block, error) {
	convertOpts, ok := opts.(k8s.K8sConvertorOptions)
	if !ok {
		return nil, errors.New("invalid options type")
	}

	var (
		// NB(alex) the config path should keep the same with the secret mount path in deployer.
		configPath = convertOpts.ConfigPath + "/" + util.GetK8sSecretName(connector.Name)
		alias      = convertOpts.ConnSeparator + connector.Name
		attributes = map[string]any{
			"config_path": configPath,
			"alias":       alias,
		}
	)

	_, err := convertor.LoadK8sApiConfig(ctx, connector)
	if err != nil {
		return nil, err
	}

	return &block.Block{
		Type:       block.TypeProvider,
		Attributes: attributes,
		// Convert the connector type to provider type.
		Labels: []string{string(m)},
	}, nil
}
