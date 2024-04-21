package docker

import (
	"context"
	"errors"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
)

const OperatorType = resourcehandler.ConnectorTypeDocker

func New(_ context.Context, connCfg walrus.ConnectorConfig) (resourcehandler.ResourceHandler, error) {
	host := string(connCfg.Status.Data["host"])
	if host == "" {
		return nil, errors.New("host is empty")
	}

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       connCfg.Name,
		identifier: stringx.SumBySHA256("docker:", host),
		client:     cli,
	}, nil
}

type Operator struct {
	name       string
	identifier string
	client     *client.Client
}

func (op Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (op Operator) IsConnected(ctx context.Context) error {
	if _, err := op.client.ServerVersion(ctx); err != nil {
		return fmt.Errorf("error connect to docker daemon: %w", err)
	}

	return nil
}

func (op Operator) Burst() int {
	return 100
}

func (op Operator) ID() string {
	return op.identifier
}

func (op Operator) GetKeys(ctx context.Context, resComps *walruscore.ResourceComponents) (*walruscore.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (op Operator) GetStatus(ctx context.Context, resComps *walruscore.ResourceComponents) ([]meta.Condition, error) {
	// TODO: Implement this method after resource is migrated.

	return nil, nil
}

func (op Operator) GetComponents(ctx context.Context, resComps *walruscore.ResourceComponents) ([]*walruscore.ResourceComponents, error) {
	return nil, nil
}

func (op Operator) Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error {
	return nil
}

func (op Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	return nil
}
