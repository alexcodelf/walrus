// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

// ApplicationSetSyncPolicyApplyConfiguration represents an declarative configuration of the ApplicationSetSyncPolicy type for use
// with apply.
type ApplicationSetSyncPolicyApplyConfiguration struct {
	PreserveResourcesOnDeletion *bool                            `json:"preserveResourcesOnDeletion,omitempty"`
	ApplicationsSync            *v1alpha1.ApplicationsSyncPolicy `json:"applicationsSync,omitempty"`
}

// ApplicationSetSyncPolicyApplyConfiguration constructs an declarative configuration of the ApplicationSetSyncPolicy type for use with
// apply.
func ApplicationSetSyncPolicy() *ApplicationSetSyncPolicyApplyConfiguration {
	return &ApplicationSetSyncPolicyApplyConfiguration{}
}

// WithPreserveResourcesOnDeletion sets the PreserveResourcesOnDeletion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreserveResourcesOnDeletion field is set to the value of the last call.
func (b *ApplicationSetSyncPolicyApplyConfiguration) WithPreserveResourcesOnDeletion(value bool) *ApplicationSetSyncPolicyApplyConfiguration {
	b.PreserveResourcesOnDeletion = &value
	return b
}

// WithApplicationsSync sets the ApplicationsSync field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ApplicationsSync field is set to the value of the last call.
func (b *ApplicationSetSyncPolicyApplyConfiguration) WithApplicationsSync(value v1alpha1.ApplicationsSyncPolicy) *ApplicationSetSyncPolicyApplyConfiguration {
	b.ApplicationsSync = &value
	return b
}
