package walrus

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/webhook"
)

// VariableWebhook hooks a v1.Variable object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walrus.seal.io",version="v1",resource="variables",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type VariableWebhook struct {
	webhook.DefaultCustomValidator
}

func (r *VariableWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	return &walrus.Variable{}, nil
}

func (r *VariableWebhook) ValidateCreate(_ context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	vra := obj.(*walrus.Variable)

	// Validate unsupported fields.
	if ws := r.validateUnsupports(vra); ws != nil {
		return ws, nil
	}

	return nil, nil
}

func (r *VariableWebhook) ValidateUpdate(_ context.Context, _, newObj runtime.Object) (ctrladmission.Warnings, error) {
	vra := newObj.(*walrus.Variable)

	// Validate unsupported fields.
	if ws := r.validateUnsupports(vra); ws != nil {
		return ws, nil
	}

	return nil, nil
}

func (r *VariableWebhook) validateUnsupports(vra *walrus.Variable) ctrladmission.Warnings {
	var ws ctrladmission.Warnings

	if len(vra.Labels) != 0 {
		ws = append(ws, "labels are not supported")
	}

	if len(vra.Finalizers) != 0 {
		ws = append(ws, "finalizers are not supported")
	}

	if len(vra.OwnerReferences) != 0 {
		ws = append(ws, "ownerReferences are not supported")
	}

	if len(vra.Annotations) != 1 {
		if _, ok := vra.Annotations["kubectl.kubernetes.io/last-applied-configuration"]; !ok {
			ws = append(ws, "annotations are not supported, include 'kubectl.kubernetes.io/last-applied-configuration'")
		}
	}

	return ws
}
