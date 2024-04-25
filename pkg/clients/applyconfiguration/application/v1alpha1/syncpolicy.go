// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	applicationv1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

// SyncPolicyApplyConfiguration represents an declarative configuration of the SyncPolicy type for use
// with apply.
type SyncPolicyApplyConfiguration struct {
	Automated                *SyncPolicyAutomatedApplyConfiguration      `json:"automated,omitempty"`
	SyncOptions              *applicationv1alpha1.SyncOptions            `json:"syncOptions,omitempty"`
	Retry                    *RetryStrategyApplyConfiguration            `json:"retry,omitempty"`
	ManagedNamespaceMetadata *ManagedNamespaceMetadataApplyConfiguration `json:"managedNamespaceMetadata,omitempty"`
}

// SyncPolicyApplyConfiguration constructs an declarative configuration of the SyncPolicy type for use with
// apply.
func SyncPolicy() *SyncPolicyApplyConfiguration {
	return &SyncPolicyApplyConfiguration{}
}

// WithAutomated sets the Automated field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Automated field is set to the value of the last call.
func (b *SyncPolicyApplyConfiguration) WithAutomated(value *SyncPolicyAutomatedApplyConfiguration) *SyncPolicyApplyConfiguration {
	b.Automated = value
	return b
}

// WithSyncOptions sets the SyncOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncOptions field is set to the value of the last call.
func (b *SyncPolicyApplyConfiguration) WithSyncOptions(value applicationv1alpha1.SyncOptions) *SyncPolicyApplyConfiguration {
	b.SyncOptions = &value
	return b
}

// WithRetry sets the Retry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Retry field is set to the value of the last call.
func (b *SyncPolicyApplyConfiguration) WithRetry(value *RetryStrategyApplyConfiguration) *SyncPolicyApplyConfiguration {
	b.Retry = value
	return b
}

// WithManagedNamespaceMetadata sets the ManagedNamespaceMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagedNamespaceMetadata field is set to the value of the last call.
func (b *SyncPolicyApplyConfiguration) WithManagedNamespaceMetadata(value *ManagedNamespaceMetadataApplyConfiguration) *SyncPolicyApplyConfiguration {
	b.ManagedNamespaceMetadata = value
	return b
}