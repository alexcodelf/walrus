// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// GCSArtifactRepositoryApplyConfiguration represents an declarative configuration of the GCSArtifactRepository type for use
// with apply.
type GCSArtifactRepositoryApplyConfiguration struct {
	GCSBucketApplyConfiguration `json:",inline"`
	KeyFormat                   *string `json:"keyFormat,omitempty"`
}

// GCSArtifactRepositoryApplyConfiguration constructs an declarative configuration of the GCSArtifactRepository type for use with
// apply.
func GCSArtifactRepository() *GCSArtifactRepositoryApplyConfiguration {
	return &GCSArtifactRepositoryApplyConfiguration{}
}

// WithBucket sets the Bucket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bucket field is set to the value of the last call.
func (b *GCSArtifactRepositoryApplyConfiguration) WithBucket(value string) *GCSArtifactRepositoryApplyConfiguration {
	b.Bucket = &value
	return b
}

// WithServiceAccountKeySecret sets the ServiceAccountKeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountKeySecret field is set to the value of the last call.
func (b *GCSArtifactRepositoryApplyConfiguration) WithServiceAccountKeySecret(value *v1.SecretKeySelectorApplyConfiguration) *GCSArtifactRepositoryApplyConfiguration {
	b.ServiceAccountKeySecret = value
	return b
}

// WithKeyFormat sets the KeyFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeyFormat field is set to the value of the last call.
func (b *GCSArtifactRepositoryApplyConfiguration) WithKeyFormat(value string) *GCSArtifactRepositoryApplyConfiguration {
	b.KeyFormat = &value
	return b
}