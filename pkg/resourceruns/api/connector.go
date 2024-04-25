package api

import (
	"context"
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/system"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomConfig is the config of a custom connector.
// It is used to generate the custom connector config.
// E.g. A custom helm connector
//
//	configData := CustomConfig{
//		Attributes: map[string]any{
//			"access_url": "http://localhost:8080",
//		},
//		Dependencies: []Dependency{
//			{
//				Type: "kubernetes",
//				Label: []string{},
//				Attributes: map[string]any{
//					"config_path": "/home/user/.kube/config",
//				},
//			},
//		},
//	}
//
// This will generate the following terraform provider.
//
//	provider "helm" {
//		access_url = "http://localhost:8080"
//		kubernetes {
//			config_path = "/home/user/.kube/config"
//		}
//	}
type CustomConfig struct {
	// Attributes is the custom connector attribute
	// e.g. access_key, secret_key, etc.
	Attributes map[string]any `json:"attributes"`

	// TODO add block support, some custom connector may need Dependencies(blocks)
	// Dependencies is the dependencies of the custom connector.
	Dependencies []Dependency `json:"dependencies"`
}

// Dependency is the dependency of a custom connector.
type Dependency struct {
	Type       string         `json:"type"`
	Label      []string       `json:"label"`
	Attributes map[string]any `json:"attributes"`

	Children []Dependency `json:"children"`
}

// LoadCustomConfig loads the custom connector config from the connector.
func LoadCustomConfig(ctx context.Context, c walruscore.Connector) (*CustomConfig, error) {
	if c.Spec.Category != walruscore.ConnectorCategoryCustom {
		return nil, fmt.Errorf("connector type is not custom connector: %s", c.Name)
	}

	cc := &CustomConfig{
		Attributes: make(map[string]any),
	}

	loopbackKubeClient := system.LoopbackKubeClient.Get()

	// Get secret.
	sec, err := loopbackKubeClient.CoreV1().Secrets(c.Namespace).Get(ctx, c.Spec.SecretName, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	for k, d := range sec.Data {
		cc.Attributes[k] = d
	}

	return cc, nil
}
