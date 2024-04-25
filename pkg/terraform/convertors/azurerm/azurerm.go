package azurerm

import (
	"context"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
)

type AzureRMConvertor string

func (m AzureRMConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Type == resourcehandler.ConnectorTypeAzure
}

func (m AzureRMConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
	var blocks block.Blocks

	for _, c := range connectors {
		if !m.IsSupported(c) {
			continue
		}

		b, err := convertor.ToCloudProviderBlock(ctx, string(m), c, opts)
		if err != nil {
			return nil, err
		}

		b.AppendBlock(&block.Block{
			Type: block.TypeFeatures,
		})

		blocks = append(blocks, b)
	}

	return blocks, nil
}
