package k8s

import (
	"context"
	"time"

	"github.com/seal-io/utils/stringx"
	"github.com/seal-io/utils/waitx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	dynamicclient "k8s.io/client-go/dynamic"
	batchclient "k8s.io/client-go/kubernetes/typed/batch/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	networkingclient "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
)

const OperatorType = resourcehandler.ConnectorTypeKubernetes

func New(_ context.Context, connCfg walrus.ConnectorConfig) (resourcehandler.ResourceHandler, error) {
	restConfig, err := GetConfig(connCfg, func(c *rest.Config) {
		c.Timeout = 0
	})
	if err != nil {
		return nil, err
	}

	restCli, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, err
	}

	coreCli, err := coreclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	batchCli, err := batchclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	networkingCli, err := networkingclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	dynamicCli, err := dynamicclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	op := Operator{
		Logger:        klog.Background().WithName("resourcehandlers").WithName("k8s"),
		Identifier:    stringx.SumBySHA256("k8s:", restConfig.Host, restConfig.APIPath),
		RestConfig:    restConfig,
		CoreCli:       coreCli,
		BatchCli:      batchCli,
		NetworkingCli: networkingCli,
		DynamicCli:    dynamicCli,
	}

	return op, nil
}

type Operator struct {
	Logger        klog.Logger
	Identifier    string
	RestConfig    *rest.Config
	CoreCli       *coreclient.CoreV1Client
	BatchCli      *batchclient.BatchV1Client
	NetworkingCli *networkingclient.NetworkingV1Client
	DynamicCli    *dynamicclient.DynamicClient
}

func (Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (op Operator) IsConnected(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := waitx.PollUntilContextCancel(ctx, time.Second, true,
		func(ctx context.Context) error {
			return IsConnected(context.TODO(), op.CoreCli.RESTClient())
		},
	)

	return err
}

func (op Operator) Burst() int {
	if op.RestConfig.Burst == 0 {
		return rest.DefaultBurst
	}

	return op.RestConfig.Burst
}

func (op Operator) ID() string {
	return op.Identifier
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
