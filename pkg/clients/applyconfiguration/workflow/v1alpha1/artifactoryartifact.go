// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// ArtifactoryArtifactApplyConfiguration represents an declarative configuration of the ArtifactoryArtifact type for use
// with apply.
type ArtifactoryArtifactApplyConfiguration struct {
	URL                               *string `json:"url,omitempty"`
	ArtifactoryAuthApplyConfiguration `json:",inline"`
}

// ArtifactoryArtifactApplyConfiguration constructs an declarative configuration of the ArtifactoryArtifact type for use with
// apply.
func ArtifactoryArtifact() *ArtifactoryArtifactApplyConfiguration {
	return &ArtifactoryArtifactApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *ArtifactoryArtifactApplyConfiguration) WithURL(value string) *ArtifactoryArtifactApplyConfiguration {
	b.URL = &value
	return b
}

// WithUsernameSecret sets the UsernameSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UsernameSecret field is set to the value of the last call.
func (b *ArtifactoryArtifactApplyConfiguration) WithUsernameSecret(value *v1.SecretKeySelectorApplyConfiguration) *ArtifactoryArtifactApplyConfiguration {
	b.UsernameSecret = value
	return b
}

// WithPasswordSecret sets the PasswordSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PasswordSecret field is set to the value of the last call.
func (b *ArtifactoryArtifactApplyConfiguration) WithPasswordSecret(value *v1.SecretKeySelectorApplyConfiguration) *ArtifactoryArtifactApplyConfiguration {
	b.PasswordSecret = value
	return b
}
