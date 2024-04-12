// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName specifies the group name used to register the objects.
const GroupName = "walrus.seal.io"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

// SchemeGroupVersionResource is a convenience method to return the GroupVersionResource for a resource.
func SchemeGroupVersionResource(resource string) schema.GroupVersionResource {
	return SchemeGroupVersion.WithResource(resource)
}

// SchemeResource takes an unqualified resource and returns a Group qualified GroupResource.
func SchemeResource(resource string) schema.GroupResource {
	return SchemeGroupVersionResource(resource).GroupResource()
}

// SchemeGroupVersionKind is a convenience method to return the GroupVersionKind for a kind.
func SchemeGroupVersionKind(kind string) schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind(kind)
}

// SchemeKind takes an unqualified kind and returns a Group qualified GroupKind.
func SchemeKind(kind string) schema.GroupKind {
	return SchemeGroupVersionKind(kind).GroupKind()
}

var (
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	// Depreciated: use Install instead
	AddToScheme = localSchemeBuilder.AddToScheme
	Install     = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Catalog{},
		&CatalogList{},
		&Connector{},
		&ConnectorBinding{},
		&ConnectorBindingList{},
		&ConnectorList{},
		&Environment{},
		&EnvironmentList{},
		&FileExample{},
		&FileExampleList{},
		&Project{},
		&ProjectList{},
		&ProjectSubjects{},
		&Resource{},
		&ResourceComponents{},
		&ResourceComponentsList{},
		&ResourceDefinition{},
		&ResourceDefinitionList{},
		&ResourceList{},
		&ResourceRun{},
		&ResourceRunList{},
		&Schema{},
		&SchemaList{},
		&Setting{},
		&SettingList{},
		&Subject{},
		&SubjectList{},
		&SubjectLogin{},
		&SubjectProvider{},
		&SubjectProviderList{},
		&SubjectToken{},
		&Template{},
		&TemplateList{},
		&Variable{},
		&VariableList{},
	)
	// AddToGroupVersion allows the serialization of client types like ListOptions.
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
