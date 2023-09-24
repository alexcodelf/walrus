package workflow

import (
	"context"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

var helloWorldWorkflow = wfv1.Workflow{
	ObjectMeta: metav1.ObjectMeta{
		Name: "nameofit",
	},
	Spec: wfv1.WorkflowSpec{
		Entrypoint: "whalesay",
		Templates: []wfv1.Template{
			{
				Name: "whalesay",
				Container: &corev1.Container{
					Image:   "docker/whalesay:latest",
					Command: []string{"cowsay", "hello world"},
				},
			},
		},
	},
}

func NewClientSet(restConfig *rest.Config) (wfclientset.Interface, error) {
	return wfclientset.NewForConfig(restConfig)
}

// SubmitWorkflow creates a new workflow.
func SubmitWorkflow(
	ctx context.Context,
	wfClient wfclientset.Interface,
	namespace string,
	wf *wfv1.Workflow,
) (*wfv1.Workflow, error) {
	return wfClient.ArgoprojV1alpha1().Workflows(namespace).Create(ctx, wf, metav1.CreateOptions{})
}
