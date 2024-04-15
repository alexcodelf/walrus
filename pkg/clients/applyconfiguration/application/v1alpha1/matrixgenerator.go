// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// MatrixGeneratorApplyConfiguration represents an declarative configuration of the MatrixGenerator type for use
// with apply.
type MatrixGeneratorApplyConfiguration struct {
	Generators []ApplicationSetNestedGeneratorApplyConfiguration `json:"generators,omitempty"`
	Template   *ApplicationSetTemplateApplyConfiguration         `json:"template,omitempty"`
}

// MatrixGeneratorApplyConfiguration constructs an declarative configuration of the MatrixGenerator type for use with
// apply.
func MatrixGenerator() *MatrixGeneratorApplyConfiguration {
	return &MatrixGeneratorApplyConfiguration{}
}

// WithGenerators adds the given value to the Generators field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generators field.
func (b *MatrixGeneratorApplyConfiguration) WithGenerators(values ...*ApplicationSetNestedGeneratorApplyConfiguration) *MatrixGeneratorApplyConfiguration {
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
func (b *MatrixGeneratorApplyConfiguration) WithTemplate(value *ApplicationSetTemplateApplyConfiguration) *MatrixGeneratorApplyConfiguration {
	b.Template = value
	return b
}
