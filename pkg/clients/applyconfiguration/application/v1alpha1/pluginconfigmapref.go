// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// PluginConfigMapRefApplyConfiguration represents an declarative configuration of the PluginConfigMapRef type for use
// with apply.
type PluginConfigMapRefApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// PluginConfigMapRefApplyConfiguration constructs an declarative configuration of the PluginConfigMapRef type for use with
// apply.
func PluginConfigMapRef() *PluginConfigMapRefApplyConfiguration {
	return &PluginConfigMapRefApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PluginConfigMapRefApplyConfiguration) WithName(value string) *PluginConfigMapRefApplyConfiguration {
	b.Name = &value
	return b
}
