package walruscore

import (
	"context"
	"reflect"
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/webhook"
)

// CatalogWebhook hooks a v1.Catalog object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="catalogs",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10,subResources=["status"]
type CatalogWebhook struct {
	webhook.DefaultCustomValidator
}

func (r *CatalogWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	return &walruscore.Catalog{}, nil
}

var _ ctrlwebhook.CustomValidator = (*CatalogWebhook)(nil)

func (r *CatalogWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	cat := obj.(*walruscore.Catalog)

	// Validate filters.
	if err := r.validateFilters(cat); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *CatalogWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	oldCat, newCat := oldObj.(*walruscore.Catalog), newObj.(*walruscore.Catalog)

	// Validate immutable fields.
	if err := r.validateImmutables(oldCat, newCat); err != nil {
		return nil, err
	}

	// Validate filters.
	if err := r.validateFilters(newCat); err != nil {
		return nil, err
	}

	// Validate conditions.
	if err := r.validateConditions(oldCat, newCat); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *CatalogWebhook) validateFilters(cat *walruscore.Catalog) error {
	filters := cat.Spec.Filters
	if filters == nil {
		return nil
	}

	if filters.IncludeExpression != "" {
		if _, err := regexp.Compile(filters.IncludeExpression); err != nil {
			return field.Invalid(
				field.NewPath("spec.filters.includeExpression"), filters.IncludeExpression, err.Error())
		}
	}
	if filters.ExcludeExpression != "" {
		if _, err := regexp.Compile(filters.ExcludeExpression); err != nil {
			return field.Invalid(
				field.NewPath("spec.filters.excludeExpression"), filters.ExcludeExpression, err.Error())
		}
	}
	return nil
}

func (r *CatalogWebhook) validateImmutables(oldCat, newCat *walruscore.Catalog) error {
	if oldCat.Spec.Builtin != newCat.Spec.Builtin {
		return field.Invalid(
			field.NewPath("spec.builtin"), oldCat.Spec.Builtin, "field is immutable")
	}
	if oldCat.Spec.TemplateFormat != newCat.Spec.TemplateFormat {
		return field.Invalid(
			field.NewPath("spec.templateFormat"), oldCat.Spec.TemplateFormat, "field is immutable")
	}
	if !reflect.DeepEqual(oldCat.Spec.Filters, newCat.Spec.Filters) {
		return field.Invalid(
			field.NewPath("spec.filters"), oldCat.Spec.Filters, "field is immutable")
	}
	if !reflect.DeepEqual(oldCat.Spec.VCSSource, newCat.Spec.VCSSource) {
		return field.Invalid(
			field.NewPath("spec.vcsSource"), oldCat.Spec.VCSSource, "field is immutable")
	}
	return nil
}

func (r *CatalogWebhook) validateConditions(oldCat, newCat *walruscore.Catalog) error {
	switch {
	case oldCat.Generation != newCat.Generation:
		return nil
	case apistatus.CatalogConditionFetched.IsUnknown(newCat):
		if len(oldCat.Status.Conditions) == 0 ||
			apistatus.CatalogConditionFetched.IsFalse(oldCat) ||
			apistatus.CatalogConditionSyncedTemplates.IsTrueOrFalse(oldCat) {
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newCat.Status.Conditions, "previous fetching/syncing templates has not been completed yet")
	case apistatus.CatalogConditionSyncedTemplates.IsUnknown(newCat):
		if apistatus.CatalogConditionFetched.IsTrue(oldCat) {
			// Resync if previous fetching has been completed yet.
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newCat.Status.Conditions, "previous fetching has not been completed yet")
	case apistatus.CatalogConditionDeleting.Exists(newCat):
		if newCat.DeletionTimestamp != nil {
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newCat.Status.Conditions, "deletion has not been requested yet")
	}
	return nil
}
