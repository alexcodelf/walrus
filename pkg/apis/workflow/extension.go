package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	pkgworkflow "github.com/seal-io/walrus/pkg/workflow"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func (h Handler) CollectionRouteGetTestRequest(req CollectionRouteGetTestRequest) (any, error) {
	return nil, nil
}

func (h Handler) RouteApplyRequest(req RouteApplyRequest) (any, error) {
	wf, err := h.modelClient.Workflows().Query().
		Where(workflow.ID(req.ID)).
		WithStages(func(wsq *model.WorkflowStageQuery) {
			wsq.WithSteps()
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	apiConfig := CreateKubeconfigFileForRestConfig(h.k8sConfig)
	clientConfig := clientcmd.NewDefaultClientConfig(apiConfig, nil)

	err = h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		return pkgworkflow.Apply(req.Context, h.modelClient, clientConfig, wf)
	})

	return nil, err
}

func CreateKubeconfigFileForRestConfig(restConfig *rest.Config) clientcmdapi.Config {
	clusters := make(map[string]*clientcmdapi.Cluster)
	clusters["default-cluster"] = &clientcmdapi.Cluster{
		Server:                   restConfig.Host,
		CertificateAuthorityData: restConfig.CAData,
	}
	contexts := make(map[string]*clientcmdapi.Context)
	contexts["default-context"] = &clientcmdapi.Context{
		Cluster:  "default-cluster",
		AuthInfo: "default-user",
	}
	authinfos := make(map[string]*clientcmdapi.AuthInfo)
	authinfos["default-user"] = &clientcmdapi.AuthInfo{
		ClientCertificateData: restConfig.CertData,
		ClientKeyData:         restConfig.KeyData,
	}
	clientConfig := clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clusters,
		Contexts:       contexts,
		CurrentContext: "default-context",
		AuthInfos:      authinfos,
	}

	return clientConfig
}
