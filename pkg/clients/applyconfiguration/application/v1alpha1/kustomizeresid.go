// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// KustomizeResIdApplyConfiguration represents an declarative configuration of the KustomizeResId type for use
// with apply.
type KustomizeResIdApplyConfiguration struct {
	KustomizeGvkApplyConfiguration `json:",omitempty,inline"`
	Name                           *string `json:"name,omitempty"`
	Namespace                      *string `json:"namespace,omitempty"`
}

// KustomizeResIdApplyConfiguration constructs an declarative configuration of the KustomizeResId type for use with
// apply.
func KustomizeResId() *KustomizeResIdApplyConfiguration {
	return &KustomizeResIdApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *KustomizeResIdApplyConfiguration) WithGroup(value string) *KustomizeResIdApplyConfiguration {
	b.Group = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *KustomizeResIdApplyConfiguration) WithVersion(value string) *KustomizeResIdApplyConfiguration {
	b.Version = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *KustomizeResIdApplyConfiguration) WithKind(value string) *KustomizeResIdApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *KustomizeResIdApplyConfiguration) WithName(value string) *KustomizeResIdApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *KustomizeResIdApplyConfiguration) WithNamespace(value string) *KustomizeResIdApplyConfiguration {
	b.Namespace = &value
	return b
}
