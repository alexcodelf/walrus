// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// OSSArtifactApplyConfiguration represents an declarative configuration of the OSSArtifact type for use
// with apply.
type OSSArtifactApplyConfiguration struct {
	OSSBucketApplyConfiguration `json:",inline"`
	Key                         *string `json:"key,omitempty"`
}

// OSSArtifactApplyConfiguration constructs an declarative configuration of the OSSArtifact type for use with
// apply.
func OSSArtifact() *OSSArtifactApplyConfiguration {
	return &OSSArtifactApplyConfiguration{}
}

// WithEndpoint sets the Endpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Endpoint field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithEndpoint(value string) *OSSArtifactApplyConfiguration {
	b.Endpoint = &value
	return b
}

// WithBucket sets the Bucket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bucket field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithBucket(value string) *OSSArtifactApplyConfiguration {
	b.Bucket = &value
	return b
}

// WithAccessKeySecret sets the AccessKeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AccessKeySecret field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithAccessKeySecret(value *v1.SecretKeySelectorApplyConfiguration) *OSSArtifactApplyConfiguration {
	b.AccessKeySecret = value
	return b
}

// WithSecretKeySecret sets the SecretKeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretKeySecret field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithSecretKeySecret(value *v1.SecretKeySelectorApplyConfiguration) *OSSArtifactApplyConfiguration {
	b.SecretKeySecret = value
	return b
}

// WithCreateBucketIfNotPresent sets the CreateBucketIfNotPresent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreateBucketIfNotPresent field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithCreateBucketIfNotPresent(value bool) *OSSArtifactApplyConfiguration {
	b.CreateBucketIfNotPresent = &value
	return b
}

// WithSecurityToken sets the SecurityToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityToken field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithSecurityToken(value string) *OSSArtifactApplyConfiguration {
	b.SecurityToken = &value
	return b
}

// WithLifecycleRule sets the LifecycleRule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LifecycleRule field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithLifecycleRule(value *OSSLifecycleRuleApplyConfiguration) *OSSArtifactApplyConfiguration {
	b.LifecycleRule = value
	return b
}

// WithUseSDKCreds sets the UseSDKCreds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UseSDKCreds field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithUseSDKCreds(value bool) *OSSArtifactApplyConfiguration {
	b.UseSDKCreds = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *OSSArtifactApplyConfiguration) WithKey(value string) *OSSArtifactApplyConfiguration {
	b.Key = &value
	return b
}
