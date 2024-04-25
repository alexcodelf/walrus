// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// ResourceComponentDependencyApplyConfiguration represents an declarative configuration of the ResourceComponentDependency type for use
// with apply.
type ResourceComponentDependencyApplyConfiguration struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

// ResourceComponentDependencyApplyConfiguration constructs an declarative configuration of the ResourceComponentDependency type for use with
// apply.
func ResourceComponentDependency() *ResourceComponentDependencyApplyConfiguration {
	return &ResourceComponentDependencyApplyConfiguration{}
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *ResourceComponentDependencyApplyConfiguration) WithFrom(value string) *ResourceComponentDependencyApplyConfiguration {
	b.From = &value
	return b
}

// WithTo sets the To field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the To field is set to the value of the last call.
func (b *ResourceComponentDependencyApplyConfiguration) WithTo(value string) *ResourceComponentDependencyApplyConfiguration {
	b.To = &value
	return b
}