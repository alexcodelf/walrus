// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// ConnectorConfigEntryApplyConfiguration represents an declarative configuration of the ConnectorConfigEntry type for use
// with apply.
type ConnectorConfigEntryApplyConfiguration struct {
	Sensitive *bool   `json:"sensitive,omitempty"`
	Value     *string `json:"value,omitempty"`
}

// ConnectorConfigEntryApplyConfiguration constructs an declarative configuration of the ConnectorConfigEntry type for use with
// apply.
func ConnectorConfigEntry() *ConnectorConfigEntryApplyConfiguration {
	return &ConnectorConfigEntryApplyConfiguration{}
}

// WithSensitive sets the Sensitive field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Sensitive field is set to the value of the last call.
func (b *ConnectorConfigEntryApplyConfiguration) WithSensitive(value bool) *ConnectorConfigEntryApplyConfiguration {
	b.Sensitive = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *ConnectorConfigEntryApplyConfiguration) WithValue(value string) *ConnectorConfigEntryApplyConfiguration {
	b.Value = &value
	return b
}