package deployer

import (
	"context"
	"io"

	wf "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

const (
	DeployerServiceAccountName = "walrus-deployer"
)

// Type indicates the type of Deployer,
// e.g. Terraform, KubeVela, etc.
type Type = string

const (
	TypeTerraform Type = walruscore.TemplateFormatTerraform
)

// Deployer holds the actions that a deployer must satisfy.
type Deployer interface {
	// Type returns Type.
	Type() Type

	// Apply create workflow template that creates/updates the resources of the given ResourceRun,
	// also cleans stale resources.
	Apply(context.Context, *walruscore.ResourceRun, ApplyOptions) (*wf.Template, error)

	// Destroy create workflow template that cleans all resources of the given ResourceRun.
	Destroy(context.Context, *walruscore.ResourceRun, DestroyOptions) (*wf.Template, error)

	// Plan create workflow template that plans the resources of the given ResourceRun.
	Plan(context.Context, *walruscore.ResourceRun, PlanOptions) (*wf.Template, error)
}

// ApplyOptions holds the options of Deployer's Apply action.
type ApplyOptions struct{}

// DestroyOptions holds the options of Deployer's Destroy action.
type DestroyOptions struct{}

// PlanOptions holds the options of Deployer's Plan action.
type PlanOptions struct{}

type CreateTemplateOptions struct {
	ResourceRunStepType walruscore.ResourceRunStepType

	Image      string
	Env        []corev1.EnvVar
	DockerMode bool

	Token     string
	ServerURL string
}

type StreamLogsOptions struct {
	Cli             *coreclient.CoreV1Client
	ResourceRunName string
	JobType         string
	Out             io.Writer
}
