package aws

import (
	"context"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
)

type AWSConvertor string

func (m AWSConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Type == resourcehandler.ConnectorTypeAWS
}

func (m AWSConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
	var blocks block.Blocks

	for _, c := range connectors {
		if !m.IsSupported(c) {
			continue
		}

		b, err := convertor.ToCloudProviderBlock(ctx, string(m), c, opts)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, b)
	}

	return blocks, nil
}
