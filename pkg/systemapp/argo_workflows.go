package systemapp

import (
	"context"
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemapp/helm"
)

func installArgoWorkflows(ctx context.Context, cli *helm.Client, globalValuesContext map[string]any, disable sets.Set[string]) error {
	// NB: please update the following files if changed.
	// - hack/mirror/walrus-images.txt.
	// - pack/walrus/image/Dockerfile.
	// - github.com/seal-io/helm-charts/charts/walrus.

	name := "argo-workflows"
	version := "0.41.1"
	if disable.Has(name) {
		return nil
	}

	namespace := cli.Namespace()
	release := "walrus-workflows"
	file := filepath.Join(system.SubLibDir("charts"), fmt.Sprintf("%s.tgz", name))
	download := fmt.Sprintf("https://github.com/argoproj/argo-helm/releases/download/%[1]s-%[2]s/%[1]s-%[2]s.tgz", name, version)
	valuesTemplate := `
images:
  tag: "v3.5.0"
  pullPolicy: "IfNotPresent"

crds:
  install: true
  keep: true
  annotations:
    "{{ .ManagedLabel }}": "true"

createAggregateRoles: false

fullnameOverride: "{{ .Release }}"

namespaceOverride: "{{ .Namespace }}"

singleNamespace: false

workflow:
  rbac:
    create: true

controller:
  image:
    registry: "{{ .ImageRegistry }}"
    repository: "sealio/mirrored-workflow-controller"
  name: "controller"
  instanceID:
    enabled: true
    useReleaseName: true
  workflowNamespaces: []
  persistence: {}

executor:
  image:
    registry: "{{ .ImageRegistry }}"
    repository: "sealio/mirrored-argoexec"

server:
  enabled: false
`
	valuesContext := globalValuesContext
	valuesContext["Release"] = release
	valuesContext["Namespace"] = namespace

	chart := &helm.Chart{
		Name:            name,
		Version:         version,
		Release:         release,
		File:            file,
		FileDownloadURL: download,
		Values: helm.YamlTemplateChartValues{
			Template: valuesTemplate,
			Context:  valuesContext,
		},
	}
	_, err := cli.Install(ctx, chart)
	return err
}
