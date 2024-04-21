package walruscore

import (
	"context"
	"fmt"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/webhook"
)

// TemplateWebhook hooks a v1.Template object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="templates",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE","DELETE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10,subResources=["status"]
type TemplateWebhook struct {
	webhook.DefaultCustomValidator

	Client ctrlcli.Client
}

func (r *TemplateWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	r.Client = opts.Manager.GetClient()

	return &walruscore.Template{}, nil
}

var _ ctrlwebhook.CustomValidator = (*TemplateWebhook)(nil)

func (r *TemplateWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	tmpl := obj.(*walruscore.Template)

	// Validate VCS repository URL.
	_, err := sourceurl.ParseURLToSourceURL(tmpl.Spec.VCSRepository.URL)
	if err != nil {
		return nil, field.Invalid(
			field.NewPath("spec.vcsRepository.url"), tmpl.Spec.VCSRepository.URL, err.Error())
	}

	return nil, nil
}

func (r *TemplateWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	oldTmpl, newTmpl := oldObj.(*walruscore.Template), newObj.(*walruscore.Template)

	// Validate immutable fields.
	if err := r.validateImmutables(oldTmpl, newTmpl); err != nil {
		return nil, err
	}

	// Validate conditions.
	if err := r.validateConditions(oldTmpl, newTmpl); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *TemplateWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	tmpl := obj.(*walruscore.Template)

	// Cannot delete the template if it is in use.
	resCompsList := new(walruscore.ResourceComponentsList)
	err := r.Client.List(ctx, resCompsList,
		ctrlcli.MatchingFields{
			"status.template": kubemeta.GetNamespacedNameKey(tmpl),
		})
	if err != nil {
		return nil, kerrors.NewInternalError(fmt.Errorf("list related resource components: %w", err))
	}
	if len(resCompsList.Items) > 0 {
		return nil, field.Invalid(
			field.NewPath("metadata.name"), tmpl.GetName(), "blocked by in-used resource components")
	}

	return nil, nil
}

func (r *TemplateWebhook) validateImmutables(oldTmpl, newTmpl *walruscore.Template) error {
	if oldTmpl.Spec.TemplateFormat != newTmpl.Spec.TemplateFormat {
		return field.Invalid(
			field.NewPath("spec.templateFormat"), oldTmpl.Spec.TemplateFormat, "field is immutable")
	}
	if oldTmpl.Spec.VCSRepository.Platform != newTmpl.Spec.VCSRepository.Platform {
		return field.Invalid(
			field.NewPath("spec.vcsRepository.platform"), oldTmpl.Spec.VCSRepository.Platform, "field is immutable")
	}
	if oldTmpl.Spec.VCSRepository.URL != newTmpl.Spec.VCSRepository.URL {
		return field.Invalid(
			field.NewPath("spec.vcsRepository.url"), oldTmpl.Spec.VCSRepository.URL, "field is immutable")
	}
	return nil
}

func (r *TemplateWebhook) validateConditions(oldTmpl, newTmpl *walruscore.Template) error {
	switch {
	case oldTmpl.Generation != newTmpl.Generation:
		return nil
	case apistatus.TemplateConditionSynced.IsUnknown(newTmpl):
		if len(oldTmpl.Status.Conditions) == 0 ||
			apistatus.TemplateConditionSynced.IsTrueOrFalse(oldTmpl) {
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newTmpl.Status.Conditions, "previous syncing has not been completed yet")
	case apistatus.TemplateConditionDeleting.Exists(newTmpl):
		if newTmpl.DeletionTimestamp != nil {
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newTmpl.Status.Conditions, "deletion has not been requested yet")
	}
	return nil
}
