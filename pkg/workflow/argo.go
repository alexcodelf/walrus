package workflow

import (
	"context"
	"errors"
	"os/user"
	"path/filepath"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	wfv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
	"github.com/seal-io/walrus/pkg/auths"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step"
	steptypes "github.com/seal-io/walrus/pkg/workflow/step/types"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/pointer"
)

const (
	WorkflowFlowNamespace = "walrus-system"
)

type ArgoWorkflowClient struct {
	Logger log.Logger
	mc     model.ClientSet
	// Argo workflow clientset.
	cs wfv1alpha1.WorkflowInterface
}

func NewArgoWorkflowClient(mc model.ClientSet, config *rest.Config) (WorkflowClient, error) {
	// get current user to determine home directory
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	// get kubeconfig file location
	kubeconfig := filepath.Join(usr.HomeDir, ".kube", "config")

	// use the current context in kubeconfig
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	cs, err := wfclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &ArgoWorkflowClient{
		Logger: log.WithName("workflow-service"),
		mc:     mc,
		cs:     cs.ArgoprojV1alpha1().Workflows("argo"),
	}, nil
}

func (s *ArgoWorkflowClient) Submit(ctx context.Context, opts SubmitOptions) error {
	// Prepare API token for terraform backend.
	const _1Day = 60 * 60 * 24

	at, err := auths.CreateAccessToken(ctx,
		s.mc, opts.SubjectID, types.TokenKindDeployment, opts.WorkflowExecution.ID.String(), pointer.Int(_1Day))
	if err != nil {
		return err
	}

	token := at.AccessToken
	// token := "bzAk_gT4dY_Rel45_fKlmEWqRzuOhdVD4QAdSDnXLezi97JRjX6UlLqsuh5Y0efeuYE4JZhDAr9nnXaI_CHIqzDbsEk_vNh1ygz9UJ4jdy5ClELsNvWt6DFCrX9v0qE4NIk"
	tlsVerify := apiconfig.TlsCertified.Get()

	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, s.mc)
	if err != nil {
		return err
	}

	if serverAddress == "" {
		return errors.New("server address is empty")
	}

	stepService, err := step.GetStepManager(steptypes.CreateOptions{
		Type:        steptypes.StepTypeService,
		ModelClient: s.mc,
	})
	if err != nil {
		return err
	}

	swe := &model.WorkflowStepExecution{
		ProjectID:   "480121971162504616",
		Name:        "test-svc-demo",
		Description: "test-svc-demo",
		Type:        steptypes.StepTypeService.String(),
		Spec: map[string]any{
			"projectID":     "480121971162504616",
			"environmentID": "480123459335118248",
			"project": map[string]any{
				"id": "480121971162504616",
			},
			"environment": map[string]any{
				"id": "480123459335118248",
			},
			"name": "sfsf",
			"attributes": map[string]any{
				"env":            map[string]any{},
				"name":           "",
				"image":          "nginx",
				"ports":          []int{80},
				"replicas":       1,
				"limit_cpu":      "",
				"namespace":      "",
				"request_cpu":    "0.1",
				"limit_memory":   "",
				"request_memory": "128Mi",
			},
			"template": map[string]any{
				"name":    "webservice",
				"version": "v0.0.3",
			},
			"jobType": "apply",
		},
	}

	// Generate service template.
	serviceTemplate, err := stepService.GenerateTemplate(ctx, swe)
	if err != nil {
		return err
	}

	testWf := v1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name: opts.WorkflowExecution.Name,
		},
		Spec: v1alpha1.WorkflowSpec{
			Entrypoint: "main",
			Arguments: v1alpha1.Arguments{
				Parameters: []v1alpha1.Parameter{
					{
						Name:  "server",
						Value: v1alpha1.AnyStringPtr(serverAddress),
					},
					{
						Name:  "token",
						Value: v1alpha1.AnyStringPtr(token),
					},
					{
						Name:  "tlsVerify",
						Value: v1alpha1.AnyStringPtr(tlsVerify),
					},
				},
			},
			Templates: []v1alpha1.Template{
				{
					Name: "main",
					DAG: &v1alpha1.DAGTemplate{
						Tasks: []v1alpha1.DAGTask{
							{
								Name:     serviceTemplate.Name + "-task",
								Template: serviceTemplate.Name,
							},
						},
					},
				},
				*serviceTemplate,
			},
		},
	}

	_, err = s.cs.Create(ctx, &testWf, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArgoWorkflowClient) Delete(ctx context.Context, opts DeleteOptions) error {
	// 1. Delete workflow from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowClient) Get(ctx context.Context, opts GetOptions) (*v1alpha1.Workflow, error) {
	// 1. Get workflow from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}

func (s *ArgoWorkflowClient) List(ctx context.Context, opts ListOptions) (*v1alpha1.WorkflowList, error) {
	// 1. List workflows from workflow engine.
	// 2. Update workflow status.
	panic("not implemented")
}
