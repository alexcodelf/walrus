package k8s

import (
	"context"
	"errors"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
	"github.com/seal-io/walrus/pkg/terraform/util"
)

type K8sConvertorOptions struct {
	ConfigPath    string
	ConnSeparator string
	GetSecretName func(string) string
}

// K8sConvertor mutate the types.ConnectorTypeKubernetes connector to Kubernetes provider block.
type K8sConvertor string

func (m K8sConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Type == resourcehandler.ConnectorTypeKubernetes
}

func (m K8sConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
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

func (m K8sConvertor) toBlock(ctx context.Context, conn walruscore.Connector, opts any) (*block.Block, error) {
	convertOpts, ok := opts.(K8sConvertorOptions)
	if !ok {
		return nil, errors.New("invalid options type")
	}

	var (
		// NB(alex) the config path should keep the same with the secret mount path in deployer.
		configPath = convertOpts.ConfigPath + "/" + util.GetK8sSecretName(conn.Name)
		alias      = convertOpts.ConnSeparator + conn.Name
		attributes = map[string]any{
			"config_path": configPath,
			"alias":       alias,
		}
	)

	_, err := convertor.LoadK8sApiConfig(ctx, conn)
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
