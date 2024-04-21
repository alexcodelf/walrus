package k8s

import (
	"fmt"
	"time"

	"github.com/seal-io/utils/stringx"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
)

// GetConfig returns the rest.Config with the given model,
// by default, the rest.Config configures with 15s timeout/16 qps/64 burst,
// please modify the default configuration with ConfigOption as need.
func GetConfig(connCfg walrus.ConnectorConfig, opts ...func(*rest.Config)) (restConfig *rest.Config, err error) {
	apiConfig, _, err := LoadApiConfig(connCfg)
	if err != nil {
		return nil, err
	}

	restConfig, err = clientcmd.
		NewNonInteractiveClientConfig(*apiConfig, "", &clientcmd.ConfigOverrides{}, nil).
		ClientConfig()
	if err != nil {
		err = fmt.Errorf("cannot construct rest config from api config: %w", err)
		return
	}
	restConfig.Timeout = 15 * time.Second

	for _, opt := range opts {
		opt(restConfig)
	}

	return
}

// LoadApiConfig returns the api.Config with the given model.
func LoadApiConfig(connCfg walrus.ConnectorConfig) (apiConfig *clientcmdapi.Config, raw string, err error) {
	version := connCfg.Status.Version

	switch version {
	default:
		return nil, "", fmt.Errorf("unknown config version: %v", version)
	case "v1":
		// {
		//      "configVersion": "v1",
		//      "configData": {
		//          "kubeconfig": "..."
		//      }
		// }.
		raw, err = loadRawConfigV1(connCfg.Status.Data)
		if err != nil {
			return nil, "", fmt.Errorf("error load config from connector %s: %w", connCfg.Name, err)
		}

		apiConfig, err = loadApiConfigV1(raw)
		if err != nil {
			return nil, "", fmt.Errorf("error load version %s config: %w", version, err)
		}
	}

	return
}

func loadRawConfigV1(data map[string][]byte) (string, error) {
	// {
	//      "kubeconfig": "..."
	// }.
	kubeconfigText, ok := data["kubeconfig"]

	if !ok {
		return "", fmt.Errorf(`failed to get "kubeconfig"`)
	}

	return string(kubeconfigText), nil
}

func loadApiConfigV1(kubeconfigText string) (*clientcmdapi.Config, error) {
	config, err := clientcmd.Load(stringx.ToBytes(&kubeconfigText))
	if err != nil {
		return nil, fmt.Errorf("error load api config: %w", err)
	}

	err = clientcmd.Validate(*config)
	if err != nil {
		return nil, fmt.Errorf("error validate api config: %w", err)
	}

	return config, nil
}
