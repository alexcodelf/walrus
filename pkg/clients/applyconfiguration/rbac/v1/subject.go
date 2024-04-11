// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// SubjectApplyConfiguration represents an declarative configuration of the Subject type for use
// with apply.
type SubjectApplyConfiguration struct {
	Kind      *string `json:"kind,omitempty"`
	APIGroup  *string `json:"apiGroup,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// SubjectApplyConfiguration constructs an declarative configuration of the Subject type for use with
// apply.
func Subject() *SubjectApplyConfiguration {
	return &SubjectApplyConfiguration{}
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *SubjectApplyConfiguration) WithKind(value string) *SubjectApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIGroup sets the APIGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIGroup field is set to the value of the last call.
func (b *SubjectApplyConfiguration) WithAPIGroup(value string) *SubjectApplyConfiguration {
	b.APIGroup = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SubjectApplyConfiguration) WithName(value string) *SubjectApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *SubjectApplyConfiguration) WithNamespace(value string) *SubjectApplyConfiguration {
	b.Namespace = &value
	return b
}