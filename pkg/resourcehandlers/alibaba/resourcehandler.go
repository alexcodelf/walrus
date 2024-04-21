package alibaba

import (
	"context"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba/resourceexec"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba/resourcelog"
	"github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

const OperatorType = resourcehandler.ConnectorTypeAlibabaCloud

func New(_ context.Context, connCfg walrus.ConnectorConfig) (resourcehandler.ResourceHandler, error) {
	cred, err := types.GetCredential(connCfg.Status.Data)
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       connCfg.Name,
		cred:       cred,
		identifier: stringx.SumBySHA256("alibaba:", cred.AccessKey, cred.AccessSecret),
	}, nil
}

type Operator struct {
	name       string
	cred       *types.Credential
	identifier string
}

func (op Operator) IsConnected(ctx context.Context) error {
	client, err := ecs.NewClientWithAccessKey(
		op.cred.Region,
		op.cred.AccessKey,
		op.cred.AccessSecret,
	)
	if err != nil {
		return fmt.Errorf("error create alibaba client %s: %w", op.name, err)
	}

	// Use DescribeRegion API to check reachable and user has access to region.
	// https://www.alibabacloud.com/help/en/elastic-compute-service/latest/regions-describeregions
	req := ecs.CreateDescribeRegionsRequest()
	req.Scheme = "HTTPS"

	_, err = client.DescribeRegions(req)
	if err != nil {
		return fmt.Errorf("error connect to %s: %w", op.name, err)
	}

	return nil
}

func (op Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (op Operator) Burst() int {
	return 200
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
	newCtx := context.WithValue(ctx, types.CredentialKey, op.cred)
	return resourcelog.Log(newCtx, key, opts)
}

func (op Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	newCtx := context.WithValue(ctx, types.CredentialKey, op.cred)
	return resourceexec.Exec(newCtx, key, opts)
}
