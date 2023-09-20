package workflow

import (
	"context"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
	"testing"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	"github.com/argoproj/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"
)

func TestSubmitWorkflow(t *testing.T) {
	// Get current user to determine home directory.
	usr, err := user.Current()
	checkErr(err)

	// Get kubeconfig file location.
	kubeconfig := flag.String(
		"kubeconfig",
		filepath.Join(usr.HomeDir, ".kube", "config"),
		"(optional) absolute path to the kubeconfig file",
	)
	flag.Parse()

	// Use the current context in kubeconfig.
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	checkErr(err)
	namespace := "default"

	// Create the workflow client.
	wfClient := wfclientset.NewForConfigOrDie(config).ArgoprojV1alpha1().Workflows(namespace)

	// Submit the hello world workflow.
	ctx := context.Background()
	createdWf, err := wfClient.Create(ctx, &helloWorldWorkflow, metav1.CreateOptions{})
	checkErr(err)
	fmt.Printf("Workflow %s submitted\n", createdWf.Name)

	// Wait for the workflow to complete.
	fieldSelector := fields.ParseSelectorOrDie(fmt.Sprintf("metadata.name=%s", createdWf.Name))
	watchIf, err := wfClient.Watch(
		ctx,
		metav1.ListOptions{FieldSelector: fieldSelector.String(), TimeoutSeconds: pointer.Int64Ptr(180)},
	)
	errors.CheckError(err)
	defer watchIf.Stop()
	for next := range watchIf.ResultChan() {
		wf, ok := next.Object.(*wfv1.Workflow)
		if !ok {
			continue
		}
		if !wf.Status.FinishedAt.IsZero() {
			fmt.Printf(
				"Workflow %s %s at %v. Message: %s.\n",
				wf.Name,
				wf.Status.Phase,
				wf.Status.FinishedAt,
				wf.Status.Message,
			)
			break
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
