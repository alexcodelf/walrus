// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// ApplicationPreservedFieldsApplyConfiguration represents an declarative configuration of the ApplicationPreservedFields type for use
// with apply.
type ApplicationPreservedFieldsApplyConfiguration struct {
	Annotations []string `json:"annotations,omitempty"`
	Labels      []string `json:"labels,omitempty"`
}

// ApplicationPreservedFieldsApplyConfiguration constructs an declarative configuration of the ApplicationPreservedFields type for use with
// apply.
func ApplicationPreservedFields() *ApplicationPreservedFieldsApplyConfiguration {
	return &ApplicationPreservedFieldsApplyConfiguration{}
}

// WithAnnotations adds the given value to the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Annotations field.
func (b *ApplicationPreservedFieldsApplyConfiguration) WithAnnotations(values ...string) *ApplicationPreservedFieldsApplyConfiguration {
	for i := range values {
		b.Annotations = append(b.Annotations, values[i])
	}
	return b
}

// WithLabels adds the given value to the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Labels field.
func (b *ApplicationPreservedFieldsApplyConfiguration) WithLabels(values ...string) *ApplicationPreservedFieldsApplyConfiguration {
	for i := range values {
		b.Labels = append(b.Labels, values[i])
	}
	return b
}
