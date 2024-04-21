package google

import (
	"context"
	"fmt"

	"github.com/seal-io/utils/stringx"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	gtypes "github.com/seal-io/walrus/pkg/resourcehandlers/google/types"
)

const OperatorType = resourcehandler.ConnectorTypeGoogle

func New(_ context.Context, connCfg walrus.ConnectorConfig) (resourcehandler.ResourceHandler, error) {
	cred, err := gtypes.GetCredential(connCfg.Status.Data)
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       connCfg.Name,
		cred:       cred,
		identifier: stringx.SumBySHA256("google:", cred.Project, cred.Region, cred.Zone),
	}, nil
}

type Operator struct {
	name       string
	cred       *gtypes.Credential
	identifier string
}

func (op Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (op Operator) IsConnected(ctx context.Context) error {
	service, err := compute.NewService(ctx, option.WithCredentialsJSON([]byte(op.cred.Credentials)))
	if err != nil {
		return err
	}

	_, err = service.Regions.List(op.cred.Project).Do()
	if err != nil {
		return fmt.Errorf("error connect to google cloud: %w", err)
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
