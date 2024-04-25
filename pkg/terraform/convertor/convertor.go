package convertor

import (
	"context"
	"errors"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	opk8s "github.com/seal-io/walrus/pkg/resourcehandlers/k8s"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/terraform/block"
)

type (
	Options = any
	// Convertor converts the connector to provider block.
	// E.g. ConnectorType(kubernetes) connector to ProviderType(kubernetes) provider block.
	// ConnectorType(kubernetes) connector to ProviderType(helm) provider block.
	Convertor interface {
		// IsSupported checks if the connector is supported by the convertor.
		IsSupported(walruscore.Connector) bool
		// ToBlocks converts the connectors to provider blocks.
		ToBlocks(context.Context, walruscore.Connectors, Options) (block.Blocks, error)
	}
)

type ConvertOptions struct {
	SecretMountPath string
	ConnSeparator   string
	Providers       []string
}

func ToCloudProviderBlock(ctx context.Context, label string, conn walruscore.Connector, opts any) (*block.Block, error) {
	convertOpts, ok := opts.(ConvertOptions)
	if !ok {
		return nil, errors.New("invalid options type")
	}

	var (
		alias      = convertOpts.ConnSeparator + conn.Name
		attributes = map[string]any{
			"alias": alias,
		}
	)

	loopbackKubeClient := system.LoopbackKubeClient.Get()
	sec, err := loopbackKubeClient.CoreV1().Secrets(conn.Namespace).Get(ctx, conn.Spec.SecretName, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	for k, v := range sec.Data {
		attributes[k] = v
	}

	return &block.Block{
		Type:       block.TypeProvider,
		Attributes: attributes,
		Labels:     []string{label},
	}, nil
}

func LoadK8sApiConfig(ctx context.Context, conn walruscore.Connector) (*clientcmdapi.Config, error) {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	// Get secret.
	sec, err := loopbackKubeClient.CoreV1().Secrets(conn.Namespace).Get(ctx, conn.Spec.SecretName, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	connCfg := walrus.ConnectorConfig{
		ObjectMeta: meta.ObjectMeta{
			Namespace: conn.Namespace,
			Name:      conn.Name,
		},
		Status: walrus.ConnectorConfigStatus{
			ApplicableEnvironmentType: conn.Spec.ApplicableEnvironmentType,
			Category:                  conn.Spec.Category,
			Type:                      conn.Spec.Type,
			Version:                   conn.Spec.Config.Version,
			Data:                      sec.Data,
			ConditionSummary:          conn.Status.ConditionSummary,
		},
	}

	apiCfg, _, err := opk8s.LoadApiConfig(connCfg)
	if err != nil {
		return nil, err
	}

	return apiCfg, nil
}
