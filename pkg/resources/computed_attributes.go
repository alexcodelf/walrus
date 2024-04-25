package resources

import (
	"context"
	"fmt"

	"github.com/seal-io/utils/json"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	resapi "github.com/seal-io/walrus/pkg/resources/api"
	templateapi "github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/api/property"
	"github.com/seal-io/walrus/pkg/templates/openapi"
)

// GenComputedAttributes generate computed attributes for resource.
// Required:
// obj.Namespace, obj.Spec.Attributes
// obj.Spec.TemplateVersion/obj.Spec.ResourceDefinitionMatchingRule.
func GenComputedAttributes(
	ctx context.Context,
	obj *walruscore.Resource,
) (property.Values, error) {
	// Check.
	if obj.Spec.Template == nil &&
		obj.Status.ResourceDefinitionMatchingRule == nil {
		return nil, fmt.Errorf("failed to generate computed attributes, " +
			"both template and resource definition matching rule id are empty")
	}

	// Compute.
	var computedAttrs property.Values

	// Walrus Context.
	wctx := *resapi.NewContext().
		SetProject(obj.Status.Project).
		SetEnvironment(obj.Namespace).
		SetResource(obj.Name)

	switch {
	case obj.Spec.Template != nil:
		attrs, err := property.AttributesToPropertyValues(obj.Spec.Attributes)
		if err != nil {
			return nil, err
		}

		computedAttrs, err = computedAttributeWithTemplate(ctx, wctx, attrs, obj.Spec.Template)
		if err != nil {
			return nil, err
		}

	case obj.Status.ResourceDefinitionMatchingRule != nil:
		// TODO (alex): not implemented
		// support compute attributes from matching rule.
		var rule *walruscore.ResourceDefinitionMatchingRule
		attrs, err := property.AttributesToPropertyValues(obj.Spec.Attributes)
		if err != nil {
			return nil, err
		}

		computedAttrs, err = computedAttributeWithResourceDefinition(ctx, wctx, attrs, rule)
		if err != nil {
			return nil, err
		}
	}

	return computedAttrs, nil
}

// computedAttributeWithTemplate generate computed attribute from template.
func computedAttributeWithTemplate(
	ctx context.Context,
	wctx resapi.Context,
	attrs property.Values,
	ref *walruscore.TemplateReferenceWithVersion,
) (property.Values, error) {
	var (
		err       error
		wctxByte  []byte
		attrsByte []byte
	)

	wctxByte, err = json.Marshal(map[string]any{"context": wctx})
	if err != nil {
		return nil, err
	}

	_, schemaGroups, err := templateapi.GetTemplateVersionAllSchema(ctx, ref)
	if err != nil {
		return nil, err
	}

	var schema *templateapi.TemplateSchema
	if err := json.Unmarshal(schemaGroups.Schema.Status.Value.Raw, &schema); err != nil {
		return nil, err
	}

	var uiSchema *templateapi.TemplateSchema
	if err := json.Unmarshal(schemaGroups.Schema.Status.Value.Raw, &uiSchema); err != nil {
		return nil, err
	}

	attrsByte, err = openapi.GenSchemaDefaultWithAttribute(
		ctx,
		uiSchema.VariableSchema(),
		attrs,
		schema.DefaultValue)
	if err != nil {
		return nil, err
	}

	merged, err := json.ApplyPatches(wctxByte, attrsByte)
	if err != nil {
		return nil, err
	}

	var ca property.Values

	err = json.Unmarshal(merged, &ca)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

// computedAttributeWithResourceDefinition computed attribute with resource definition.
// required: rule.Spec.Template
func computedAttributeWithResourceDefinition(
	ctx context.Context,
	wctx resapi.Context,
	attrs property.Values,
	rule *walruscore.ResourceDefinitionMatchingRule,
) (property.Values, error) {
	// TODO (alex) support compute attributes from matching rule.
	return nil, nil
}
