// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// S3EncryptionOptionsApplyConfiguration represents an declarative configuration of the S3EncryptionOptions type for use
// with apply.
type S3EncryptionOptionsApplyConfiguration struct {
	KmsKeyId                    *string                                 `json:"kmsKeyId,omitempty"`
	KmsEncryptionContext        *string                                 `json:"kmsEncryptionContext,omitempty"`
	EnableEncryption            *bool                                   `json:"enableEncryption,omitempty"`
	ServerSideCustomerKeySecret *v1.SecretKeySelectorApplyConfiguration `json:"serverSideCustomerKeySecret,omitempty"`
}

// S3EncryptionOptionsApplyConfiguration constructs an declarative configuration of the S3EncryptionOptions type for use with
// apply.
func S3EncryptionOptions() *S3EncryptionOptionsApplyConfiguration {
	return &S3EncryptionOptionsApplyConfiguration{}
}

// WithKmsKeyId sets the KmsKeyId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KmsKeyId field is set to the value of the last call.
func (b *S3EncryptionOptionsApplyConfiguration) WithKmsKeyId(value string) *S3EncryptionOptionsApplyConfiguration {
	b.KmsKeyId = &value
	return b
}

// WithKmsEncryptionContext sets the KmsEncryptionContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KmsEncryptionContext field is set to the value of the last call.
func (b *S3EncryptionOptionsApplyConfiguration) WithKmsEncryptionContext(value string) *S3EncryptionOptionsApplyConfiguration {
	b.KmsEncryptionContext = &value
	return b
}

// WithEnableEncryption sets the EnableEncryption field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableEncryption field is set to the value of the last call.
func (b *S3EncryptionOptionsApplyConfiguration) WithEnableEncryption(value bool) *S3EncryptionOptionsApplyConfiguration {
	b.EnableEncryption = &value
	return b
}

// WithServerSideCustomerKeySecret sets the ServerSideCustomerKeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerSideCustomerKeySecret field is set to the value of the last call.
func (b *S3EncryptionOptionsApplyConfiguration) WithServerSideCustomerKeySecret(value *v1.SecretKeySelectorApplyConfiguration) *S3EncryptionOptionsApplyConfiguration {
	b.ServerSideCustomerKeySecret = value
	return b
}
