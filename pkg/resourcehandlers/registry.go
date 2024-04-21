package resourcehandlers

import (
	"context"
	"fmt"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba"
	"github.com/seal-io/walrus/pkg/resourcehandlers/aws"
	"github.com/seal-io/walrus/pkg/resourcehandlers/azure"
	"github.com/seal-io/walrus/pkg/resourcehandlers/docker"
	"github.com/seal-io/walrus/pkg/resourcehandlers/google"
	"github.com/seal-io/walrus/pkg/resourcehandlers/k8s"
	"github.com/seal-io/walrus/pkg/resourcehandlers/unknown"
	"github.com/seal-io/walrus/pkg/system"
)

// NB(thxCode): Register creators below.
var creators = map[resourcehandler.Type]_ResourceHandlerCreator{
	k8s.OperatorType:     k8s.New,
	aws.OperatorType:     aws.New,
	alibaba.OperatorType: alibaba.New,
	azure.OperatorType:   azure.New,
	google.OperatorType:  google.New,
	docker.OperatorType:  docker.New,
}

type _ResourceHandlerCreator func(context.Context, walrus.ConnectorConfig) (resourcehandler.ResourceHandler, error)

// Get returns resourcehandler.ResourceHandler with the given options.
func Get(ctx context.Context, opts resourcehandler.GetOptions) (resourcehandler.ResourceHandler, error) {
	create, exist := creators[opts.Type]
	if !exist {
		return unknown.Operator{}, nil
	}

	loopbackKubeCli := system.LoopbackKubeClient.Get()
	connConfig, err := loopbackKubeCli.WalrusV1().Connectors(opts.Connector.Namespace).
		GetConfig(ctx, opts.Connector.Name, meta.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get connector config: %w", err)
	}

	op, err := create(ctx, *connConfig)
	if err != nil {
		return nil, fmt.Errorf("create resourcehandlers for type %s: %w", opts.Type, err)
	}

	return op, nil
}
