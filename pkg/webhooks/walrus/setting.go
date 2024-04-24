package walrus

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/webhook"
)

// SettingWebhook hooks a v1.Setting object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walrus.seal.io",version="v1",resource="settings",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type SettingWebhook struct {
	webhook.DefaultCustomValidator
}

func (r *SettingWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	return &walrus.Setting{}, nil
}

func (r *SettingWebhook) ValidateCreate(_ context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	set := obj.(*walrus.Setting)

	// Validate unsupported fields.
	if ws := r.validateUnsupports(set); ws != nil {
		return ws, nil
	}

	return nil, nil
}

func (r *SettingWebhook) ValidateUpdate(_ context.Context, _, newObj runtime.Object) (ctrladmission.Warnings, error) {
	set := newObj.(*walrus.Setting)

	// Validate unsupported fields.
	if ws := r.validateUnsupports(set); ws != nil {
		return ws, nil
	}

	return nil, nil
}

func (r *SettingWebhook) validateUnsupports(set *walrus.Setting) ctrladmission.Warnings {
	var ws ctrladmission.Warnings

	if len(set.Labels) != 0 {
		ws = append(ws, "labels are not supported")
	}

	if len(set.Finalizers) != 0 {
		ws = append(ws, "finalizers are not supported")
	}

	if len(set.OwnerReferences) != 0 {
		ws = append(ws, "ownerReferences are not supported")
	}

	if len(set.Annotations) != 1 {
		if _, ok := set.Annotations["kubectl.kubernetes.io/last-applied-configuration"]; !ok {
			ws = append(ws, "annotations are not supported, include 'kubectl.kubernetes.io/last-applied-configuration'")
		}
	}

	return ws
}
