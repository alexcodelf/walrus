// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// ArtifactoryAuthApplyConfiguration represents an declarative configuration of the ArtifactoryAuth type for use
// with apply.
type ArtifactoryAuthApplyConfiguration struct {
	UsernameSecret *v1.SecretKeySelectorApplyConfiguration `json:"usernameSecret,omitempty"`
	PasswordSecret *v1.SecretKeySelectorApplyConfiguration `json:"passwordSecret,omitempty"`
}

// ArtifactoryAuthApplyConfiguration constructs an declarative configuration of the ArtifactoryAuth type for use with
// apply.
func ArtifactoryAuth() *ArtifactoryAuthApplyConfiguration {
	return &ArtifactoryAuthApplyConfiguration{}
}

// WithUsernameSecret sets the UsernameSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UsernameSecret field is set to the value of the last call.
func (b *ArtifactoryAuthApplyConfiguration) WithUsernameSecret(value *v1.SecretKeySelectorApplyConfiguration) *ArtifactoryAuthApplyConfiguration {
	b.UsernameSecret = value
	return b
}

// WithPasswordSecret sets the PasswordSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PasswordSecret field is set to the value of the last call.
func (b *ArtifactoryAuthApplyConfiguration) WithPasswordSecret(value *v1.SecretKeySelectorApplyConfiguration) *ArtifactoryAuthApplyConfiguration {
	b.PasswordSecret = value
	return b
}
