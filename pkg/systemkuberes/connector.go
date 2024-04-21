package systemkuberes

import (
	"context"
	"fmt"

	dtypes "github.com/docker/docker/api/types"
	dclient "github.com/docker/docker/client"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubeconfig"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/system"
)

func installDefaultKubernetesConnector(
	ctx context.Context,
	cli clientset.Interface,
	project string,
	envType walruscore.EnvironmentType,
) (*walrus.Connector, error) {
	connCli := cli.WalrusV1().Connectors(project)

	config, err := readLoopbackKubeConfig()
	if err != nil {
		return nil, fmt.Errorf("read kube config: %w", err)
	}

	eConn := &walrus.Connector{
		ObjectMeta: meta.ObjectMeta{
			Namespace: project,
			Name:      DefaultConnectorName,
		},
		Spec: walruscore.ConnectorSpec{
			ApplicableEnvironmentType: envType,
			Category:                  walruscore.ConnectorCategoryKubernetes,
			Type:                      resourcehandler.ConnectorTypeKubernetes,
			Description:               "Local Kubernetes",
			Config: walruscore.ConnectorConfig{
				Version: "v1",
				Data: map[string]walruscore.ConnectorConfigEntry{
					"kubeconfig": {
						Value:     config,
						Sensitive: true,
					},
				},
			},
		},
	}
	alignFn := func(aConn *walrus.Connector) (*walrus.Connector, bool, error) {
		aConn.Spec.Config = eConn.Spec.Config
		return aConn, false, nil
	}

	eConn, err = kubeclientset.Create(ctx, connCli, eConn,
		kubeclientset.WithUpdateIfExisted(alignFn))
	if err != nil {
		return nil, fmt.Errorf("install default kubernetes connector: %w", err)
	}

	return eConn, nil
}

func readLoopbackKubeConfig() (string, error) {
	kubeConfig := system.LoopbackKubeClientConfig.Get()

	kc, err := kubeconfig.ConvertRestConfigToApiConfig(&kubeConfig)
	if err != nil {
		return "", err
	}

	kcData, err := clientcmd.Write(kc)
	if err != nil {
		return "", err
	}

	return string(kcData), err
}

func installDefaultDockerConnector(
	ctx context.Context,
	cli clientset.Interface,
	project string,
	envType walruscore.EnvironmentType,
) (*walrus.Connector, error) {
	connCli := cli.WalrusV1().Connectors(project)

	eConn := &walrus.Connector{
		ObjectMeta: meta.ObjectMeta{
			Namespace: project,
			Name:      DefaultConnectorName,
		},
		Spec: walruscore.ConnectorSpec{
			ApplicableEnvironmentType: envType,
			Category:                  walruscore.ConnectorCategoryDocker,
			Type:                      resourcehandler.ConnectorTypeDocker,
			Description:               "Local Docker",
			Config: walruscore.ConnectorConfig{
				Version: "v1",
				Data:    map[string]walruscore.ConnectorConfigEntry{},
			},
		},
	}
	alignFn := func(aConn *walrus.Connector) (*walrus.Connector, bool, error) {
		aConn.Spec.Config = eConn.Spec.Config
		return aConn, false, nil
	}

	eConn, err := kubeclientset.Create(ctx, connCli, eConn,
		kubeclientset.WithUpdateIfExisted(alignFn))
	if err != nil {
		return nil, fmt.Errorf("install default docker connector: %w", err)
	}

	if err = applyLoopbackDockerNetwork(ctx); err != nil {
		return nil, fmt.Errorf("apply docker network: %w", err)
	}

	return eConn, nil
}

func applyLoopbackDockerNetwork(ctx context.Context) error {
	cli, err := dclient.NewClientWithOpts(dclient.FromEnv, dclient.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	nw := &dtypes.NetworkResource{
		Name: "local-walrus",
	}
	{
		nws, err := cli.NetworkList(ctx, dtypes.NetworkListOptions{})
		if err != nil {
			return err
		}
		for i := range nws {
			if nws[i].Name == nw.Name {
				nw = &nws[i]
				break
			}
		}
	}
	if nw.ID != "" {
		return nil
	}

	_, err = cli.NetworkCreate(ctx, nw.Name, dtypes.NetworkCreate{
		Driver: "bridge",
	})
	return err
}
