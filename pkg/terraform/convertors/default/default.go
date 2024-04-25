package _default

import (
	"context"
	"errors"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	runapi "github.com/seal-io/walrus/pkg/resourceruns/api"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
)

// DefaultConvertor is the convertor for custom category connector.
type DefaultConvertor string

func (m DefaultConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Category == walruscore.ConnectorCategoryCustom &&
		connector.Spec.Type == string(m)
}

func (m DefaultConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
	toBlockOpts, ok := opts.(convertor.ConvertOptions)
	if !ok {
		return nil, errors.New("invalid convert options")
	}

	var blocks block.Blocks

	for _, conn := range connectors {
		if !m.IsSupported(conn) {
			continue
		}

		b, err := m.toBlock(ctx, conn, toBlockOpts)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, b)
	}

	return blocks, nil
}

func (m DefaultConvertor) toBlock(ctx context.Context, connector walruscore.Connector, opts convertor.ConvertOptions) (*block.Block, error) {
	customConfig, err := runapi.LoadCustomConfig(ctx, connector)
	if err != nil {
		return nil, err
	}

	attributes := customConfig.Attributes
	attributes["alias"] = opts.ConnSeparator + connector.Name

	return &block.Block{
		Type:       block.TypeProvider,
		Labels:     []string{string(m)},
		Attributes: attributes,
	}, nil
}
