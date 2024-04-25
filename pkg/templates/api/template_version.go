package api

import (
	"context"
	"fmt"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/system"
)

const (
	TemplateVesionSchemaTypeOrigin = iota
	TemplateVersionSchemaTypeOriginalUI
	TemplateVesionSchemaTypeUI
)

type TemplateVersionSchemas struct {
	Schema           *walruscore.Schema `json:"schema,omitempty"`
	OriginalUISchema *walruscore.Schema `json:"originalUISchema,omitempty"`
	UISchema         *walruscore.Schema `json:"uiSchema,omitempty"`
}

// GetTemplateVersionReferenceSchema get template version with target schema type with template version reference.
func GetTemplateVersionReferenceSchema(ctx context.Context, ref walruscore.TemplateReferenceWithVersion, schemaType int) (
	templateVersion *walruscore.TemplateVersion,
	schema *walruscore.Schema,
	err error,
) {
	LoopbackKubeClient := system.LoopbackKubeClient.Get()
	template, err := LoopbackKubeClient.WalruscoreV1().Templates(ref.Namespace).Get(ctx, ref.Name, meta.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	for i := range template.Status.Versions {
		v := &template.Status.Versions[i]
		if v.Version != ref.Version {
			continue
		}
		templateVersion = v
	}

	if templateVersion == nil {
		return nil, nil, fmt.Errorf("template %s version %s not found", ref.Name, ref.Version)
	}

	schema, err = GetTemplateVersionSchemaWithType(ctx, ref.Namespace, templateVersion, schemaType)
	if err != nil {
		return nil, nil, err
	}

	return
}

func GetTemplateVersionAllSchema(ctx context.Context, ref *walruscore.TemplateReferenceWithVersion) (
	templateVersion *walruscore.TemplateVersion,
	templateVersionSchemas *TemplateVersionSchemas,
	err error,
) {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	template, err := loopbackKubeClient.WalruscoreV1().Templates(ref.Namespace).Get(ctx, ref.Name, meta.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	for i := range template.Status.Versions {
		v := &template.Status.Versions[i]
		if v.Version != ref.Version {
			continue
		}
		templateVersion = v
	}

	if templateVersion == nil {
		return nil, nil, fmt.Errorf("template %s version %s not found", ref.Name, ref.Version)
	}

	schema, err := GetTemplateVersionSchemaWithType(ctx, ref.Namespace, templateVersion, TemplateVesionSchemaTypeOrigin)
	if err != nil {
		return nil, nil, err
	}

	originalUISchema, err := GetTemplateVersionSchemaWithType(ctx, ref.Namespace, templateVersion, TemplateVersionSchemaTypeOriginalUI)
	if err != nil {
		return nil, nil, err
	}

	uiSchema, err := GetTemplateVersionSchemaWithType(ctx, ref.Namespace, templateVersion, TemplateVesionSchemaTypeUI)
	if err != nil {
		return nil, nil, err
	}

	templateVersionSchemas = &TemplateVersionSchemas{
		Schema:           schema,
		OriginalUISchema: originalUISchema,
		UISchema:         uiSchema,
	}

	return
}

func GetTemplateVersionSchemaWithType(
	ctx context.Context,
	namespace string,
	templateVersion *walruscore.TemplateVersion,
	schemaType int,
) (*walruscore.Schema, error) {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	var name string
	switch schemaType {
	case TemplateVesionSchemaTypeOrigin:
		name = *templateVersion.TemplateSchemaName
	case TemplateVersionSchemaTypeOriginalUI:
		name = *templateVersion.OriginalUISchemaName
	case TemplateVesionSchemaTypeUI:
		name = *templateVersion.UISchemaName
	default:
		return nil, fmt.Errorf("unknown schema type %d", schemaType)
	}

	schema, err := loopbackKubeClient.WalruscoreV1().Schemas(namespace).Get(ctx, name, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	return schema, nil
}
