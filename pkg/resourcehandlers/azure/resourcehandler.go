package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	aztypes "github.com/seal-io/walrus/pkg/resourcehandlers/azure/types"
)

const OperatorType = resourcehandler.ConnectorTypeAzure

func New(_ context.Context, connCfg walrus.ConnectorConfig) (resourcehandler.ResourceHandler, error) {
	cred, err := aztypes.GetCredential(connCfg.Status.Data)
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       connCfg.Name,
		cred:       cred,
		identifier: stringx.SumBySHA256("azure:", cred.SubscriptionID, cred.TenantID, cred.ClientID),
	}, nil
}

type Operator struct {
	name       string
	cred       *aztypes.Credential
	identifier string
}

func (o Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (o Operator) IsConnected(ctx context.Context) error {
	cred, err := azidentity.NewClientSecretCredential(o.cred.TenantID, o.cred.ClientID, o.cred.ClientSecret, nil)
	if err != nil {
		return err
	}

	clientFactory, err := armresources.NewClientFactory(o.cred.SubscriptionID, cred, nil)
	if err != nil {
		return err
	}

	client := clientFactory.NewResourceGroupsClient()

	pager := client.NewListPager(nil)

	_, err = pager.NextPage(ctx)
	if err != nil {
		return fmt.Errorf("error connect to azure: %w", err)
	}

	return nil
}

func (o Operator) Burst() int {
	return 100
}

func (o Operator) ID() string {
	return o.identifier
}

func (o Operator) GetKeys(ctx context.Context, resComps *walruscore.ResourceComponents) (*walruscore.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (o Operator) GetStatus(ctx context.Context, resComps *walruscore.ResourceComponents) ([]meta.Condition, error) {
	// TODO: Implement this method after resource is migrated.

	return nil, nil
}

func (o Operator) GetComponents(ctx context.Context, resComps *walruscore.ResourceComponents) ([]*walruscore.ResourceComponents, error) {
	return nil, nil
}

func (o Operator) Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error {
	return nil
}

func (o Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	return nil
}
