package docker

import (
	"context"
	"encoding/json"
	"errors"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	klog "k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/templates/api/property"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
)

type DockerConvertor string

func (m DockerConvertor) IsSupported(connector walruscore.Connector) bool {
	return connector.Spec.Type == resourcehandler.ConnectorTypeDocker
}

func (m DockerConvertor) ToBlocks(ctx context.Context, connectors walruscore.Connectors, opts convertor.Options) (block.Blocks, error) {
	var blocks block.Blocks

	for _, c := range connectors {
		if !m.IsSupported(c) {
			continue
		}

		b, err := toProviderBlock(ctx, string(m), c, opts)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, b)
	}

	return blocks, nil
}

func toProviderBlock(ctx context.Context, label string, conn walruscore.Connector, opts any) (*block.Block, error) {
	convertOpts, ok := opts.(convertor.ConvertOptions)
	if !ok {
		return nil, errors.New("invalid options type")
	}

	var (
		alias      = convertOpts.ConnSeparator + conn.Name
		attributes = map[string]any{
			"alias": alias,
		}
		err error
	)

	loopbackKubeClient := system.LoopbackKubeClient.Get()

	// Get secret.
	sec, err := loopbackKubeClient.CoreV1().Secrets(conn.Namespace).Get(ctx, conn.Spec.SecretName, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	for k, v := range sec.Data {
		attributes[k], _, err = property.GetString(json.RawMessage(v))
		if err != nil {
			klog.Infof("error get config data in connector %s:%s, %v", conn.Name, k, err)
		}
	}

	return &block.Block{
		Type:       block.TypeProvider,
		Attributes: attributes,
		Labels:     []string{label},
	}, nil
}
