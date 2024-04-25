// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// ConnectorBindingSpecApplyConfiguration represents an declarative configuration of the ConnectorBindingSpec type for use
// with apply.
type ConnectorBindingSpecApplyConfiguration struct {
	Connector *ConnectorReferenceWithTypeApplyConfiguration `json:"connector,omitempty"`
}

// ConnectorBindingSpecApplyConfiguration constructs an declarative configuration of the ConnectorBindingSpec type for use with
// apply.
func ConnectorBindingSpec() *ConnectorBindingSpecApplyConfiguration {
	return &ConnectorBindingSpecApplyConfiguration{}
}

// WithConnector sets the Connector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Connector field is set to the value of the last call.
func (b *ConnectorBindingSpecApplyConfiguration) WithConnector(value *ConnectorReferenceWithTypeApplyConfiguration) *ConnectorBindingSpecApplyConfiguration {
	b.Connector = value
	return b
}