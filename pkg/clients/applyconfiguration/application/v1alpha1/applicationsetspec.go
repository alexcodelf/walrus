// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	applicationv1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

// ApplicationSetSpecApplyConfiguration represents an declarative configuration of the ApplicationSetSpec type for use
// with apply.
type ApplicationSetSpecApplyConfiguration struct {
	GoTemplate                   *bool                                                `json:"goTemplate,omitempty"`
	Generators                   []ApplicationSetGeneratorApplyConfiguration          `json:"generators,omitempty"`
	Template                     *ApplicationSetTemplateApplyConfiguration            `json:"template,omitempty"`
	SyncPolicy                   *ApplicationSetSyncPolicyApplyConfiguration          `json:"syncPolicy,omitempty"`
	Strategy                     *ApplicationSetStrategyApplyConfiguration            `json:"strategy,omitempty"`
	PreservedFields              *ApplicationPreservedFieldsApplyConfiguration        `json:"preservedFields,omitempty"`
	GoTemplateOptions            []string                                             `json:"goTemplateOptions,omitempty"`
	ApplyNestedSelectors         *bool                                                `json:"applyNestedSelectors,omitempty"`
	IgnoreApplicationDifferences *applicationv1alpha1.ApplicationSetIgnoreDifferences `json:"ignoreApplicationDifferences,omitempty"`
	TemplatePatch                *string                                              `json:"templatePatch,omitempty"`
}

// ApplicationSetSpecApplyConfiguration constructs an declarative configuration of the ApplicationSetSpec type for use with
// apply.
func ApplicationSetSpec() *ApplicationSetSpecApplyConfiguration {
	return &ApplicationSetSpecApplyConfiguration{}
}

// WithGoTemplate sets the GoTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GoTemplate field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithGoTemplate(value bool) *ApplicationSetSpecApplyConfiguration {
	b.GoTemplate = &value
	return b
}

// WithGenerators adds the given value to the Generators field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generators field.
func (b *ApplicationSetSpecApplyConfiguration) WithGenerators(values ...*ApplicationSetGeneratorApplyConfiguration) *ApplicationSetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerators")
		}
		b.Generators = append(b.Generators, *values[i])
	}
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithTemplate(value *ApplicationSetTemplateApplyConfiguration) *ApplicationSetSpecApplyConfiguration {
	b.Template = value
	return b
}

// WithSyncPolicy sets the SyncPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncPolicy field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithSyncPolicy(value *ApplicationSetSyncPolicyApplyConfiguration) *ApplicationSetSpecApplyConfiguration {
	b.SyncPolicy = value
	return b
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithStrategy(value *ApplicationSetStrategyApplyConfiguration) *ApplicationSetSpecApplyConfiguration {
	b.Strategy = value
	return b
}

// WithPreservedFields sets the PreservedFields field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreservedFields field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithPreservedFields(value *ApplicationPreservedFieldsApplyConfiguration) *ApplicationSetSpecApplyConfiguration {
	b.PreservedFields = value
	return b
}

// WithGoTemplateOptions adds the given value to the GoTemplateOptions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the GoTemplateOptions field.
func (b *ApplicationSetSpecApplyConfiguration) WithGoTemplateOptions(values ...string) *ApplicationSetSpecApplyConfiguration {
	for i := range values {
		b.GoTemplateOptions = append(b.GoTemplateOptions, values[i])
	}
	return b
}

// WithApplyNestedSelectors sets the ApplyNestedSelectors field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ApplyNestedSelectors field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithApplyNestedSelectors(value bool) *ApplicationSetSpecApplyConfiguration {
	b.ApplyNestedSelectors = &value
	return b
}

// WithIgnoreApplicationDifferences sets the IgnoreApplicationDifferences field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnoreApplicationDifferences field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithIgnoreApplicationDifferences(value applicationv1alpha1.ApplicationSetIgnoreDifferences) *ApplicationSetSpecApplyConfiguration {
	b.IgnoreApplicationDifferences = &value
	return b
}

// WithTemplatePatch sets the TemplatePatch field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TemplatePatch field is set to the value of the last call.
func (b *ApplicationSetSpecApplyConfiguration) WithTemplatePatch(value string) *ApplicationSetSpecApplyConfiguration {
	b.TemplatePatch = &value
	return b
}
