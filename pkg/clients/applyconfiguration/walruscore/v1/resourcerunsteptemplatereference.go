// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// ResourceRunStepTemplateReferenceApplyConfiguration represents an declarative configuration of the ResourceRunStepTemplateReference type for use
// with apply.
type ResourceRunStepTemplateReferenceApplyConfiguration struct {
	Namespace *string `json:"namespace,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// ResourceRunStepTemplateReferenceApplyConfiguration constructs an declarative configuration of the ResourceRunStepTemplateReference type for use with
// apply.
func ResourceRunStepTemplateReference() *ResourceRunStepTemplateReferenceApplyConfiguration {
	return &ResourceRunStepTemplateReferenceApplyConfiguration{}
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceRunStepTemplateReferenceApplyConfiguration) WithNamespace(value string) *ResourceRunStepTemplateReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceRunStepTemplateReferenceApplyConfiguration) WithName(value string) *ResourceRunStepTemplateReferenceApplyConfiguration {
	b.Name = &value
	return b
}
