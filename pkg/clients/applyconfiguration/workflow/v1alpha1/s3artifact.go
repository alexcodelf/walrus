// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// S3ArtifactApplyConfiguration represents an declarative configuration of the S3Artifact type for use
// with apply.
type S3ArtifactApplyConfiguration struct {
	S3BucketApplyConfiguration `json:",inline"`
	Key                        *string `json:"key,omitempty"`
}

// S3ArtifactApplyConfiguration constructs an declarative configuration of the S3Artifact type for use with
// apply.
func S3Artifact() *S3ArtifactApplyConfiguration {
	return &S3ArtifactApplyConfiguration{}
}

// WithEndpoint sets the Endpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Endpoint field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithEndpoint(value string) *S3ArtifactApplyConfiguration {
	b.Endpoint = &value
	return b
}

// WithBucket sets the Bucket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bucket field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithBucket(value string) *S3ArtifactApplyConfiguration {
	b.Bucket = &value
	return b
}

// WithRegion sets the Region field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Region field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithRegion(value string) *S3ArtifactApplyConfiguration {
	b.Region = &value
	return b
}

// WithInsecure sets the Insecure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Insecure field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithInsecure(value bool) *S3ArtifactApplyConfiguration {
	b.Insecure = &value
	return b
}

// WithAccessKeySecret sets the AccessKeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AccessKeySecret field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithAccessKeySecret(value *v1.SecretKeySelectorApplyConfiguration) *S3ArtifactApplyConfiguration {
	b.AccessKeySecret = value
	return b
}

// WithSecretKeySecret sets the SecretKeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretKeySecret field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithSecretKeySecret(value *v1.SecretKeySelectorApplyConfiguration) *S3ArtifactApplyConfiguration {
	b.SecretKeySecret = value
	return b
}

// WithRoleARN sets the RoleARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoleARN field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithRoleARN(value string) *S3ArtifactApplyConfiguration {
	b.RoleARN = &value
	return b
}

// WithUseSDKCreds sets the UseSDKCreds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UseSDKCreds field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithUseSDKCreds(value bool) *S3ArtifactApplyConfiguration {
	b.UseSDKCreds = &value
	return b
}

// WithCreateBucketIfNotPresent sets the CreateBucketIfNotPresent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreateBucketIfNotPresent field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithCreateBucketIfNotPresent(value *CreateS3BucketOptionsApplyConfiguration) *S3ArtifactApplyConfiguration {
	b.CreateBucketIfNotPresent = value
	return b
}

// WithEncryptionOptions sets the EncryptionOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EncryptionOptions field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithEncryptionOptions(value *S3EncryptionOptionsApplyConfiguration) *S3ArtifactApplyConfiguration {
	b.EncryptionOptions = value
	return b
}

// WithCASecret sets the CASecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CASecret field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithCASecret(value *v1.SecretKeySelectorApplyConfiguration) *S3ArtifactApplyConfiguration {
	b.CASecret = value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *S3ArtifactApplyConfiguration) WithKey(value string) *S3ArtifactApplyConfiguration {
	b.Key = &value
	return b
}
