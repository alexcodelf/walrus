package unknown

import (
	"context"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
)

const OperatorType = "Unknown"

type Operator struct{}

func (Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (Operator) IsConnected(ctx context.Context) error {
	return nil
}

func (op Operator) Burst() int {
	return 100
}

func (op Operator) ID() string {
	return ""
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

func (Operator) Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error {
	return nil
}

func (Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	return nil
}
