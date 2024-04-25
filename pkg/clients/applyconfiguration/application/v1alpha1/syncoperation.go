// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	applicationv1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

// SyncOperationApplyConfiguration represents an declarative configuration of the SyncOperation type for use
// with apply.
type SyncOperationApplyConfiguration struct {
	Revision     *string                                   `json:"revision,omitempty"`
	Prune        *bool                                     `json:"prune,omitempty"`
	DryRun       *bool                                     `json:"dryRun,omitempty"`
	SyncStrategy *SyncStrategyApplyConfiguration           `json:"syncStrategy,omitempty"`
	Resources    []SyncOperationResourceApplyConfiguration `json:"resources,omitempty"`
	Source       *ApplicationSourceApplyConfiguration      `json:"source,omitempty"`
	Manifests    []string                                  `json:"manifests,omitempty"`
	SyncOptions  *applicationv1alpha1.SyncOptions          `json:"syncOptions,omitempty"`
	Sources      *applicationv1alpha1.ApplicationSources   `json:"sources,omitempty"`
	Revisions    []string                                  `json:"revisions,omitempty"`
}

// SyncOperationApplyConfiguration constructs an declarative configuration of the SyncOperation type for use with
// apply.
func SyncOperation() *SyncOperationApplyConfiguration {
	return &SyncOperationApplyConfiguration{}
}

// WithRevision sets the Revision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Revision field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithRevision(value string) *SyncOperationApplyConfiguration {
	b.Revision = &value
	return b
}

// WithPrune sets the Prune field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Prune field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithPrune(value bool) *SyncOperationApplyConfiguration {
	b.Prune = &value
	return b
}

// WithDryRun sets the DryRun field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DryRun field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithDryRun(value bool) *SyncOperationApplyConfiguration {
	b.DryRun = &value
	return b
}

// WithSyncStrategy sets the SyncStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncStrategy field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithSyncStrategy(value *SyncStrategyApplyConfiguration) *SyncOperationApplyConfiguration {
	b.SyncStrategy = value
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *SyncOperationApplyConfiguration) WithResources(values ...*SyncOperationResourceApplyConfiguration) *SyncOperationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResources")
		}
		b.Resources = append(b.Resources, *values[i])
	}
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithSource(value *ApplicationSourceApplyConfiguration) *SyncOperationApplyConfiguration {
	b.Source = value
	return b
}

// WithManifests adds the given value to the Manifests field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Manifests field.
func (b *SyncOperationApplyConfiguration) WithManifests(values ...string) *SyncOperationApplyConfiguration {
	for i := range values {
		b.Manifests = append(b.Manifests, values[i])
	}
	return b
}

// WithSyncOptions sets the SyncOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncOptions field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithSyncOptions(value applicationv1alpha1.SyncOptions) *SyncOperationApplyConfiguration {
	b.SyncOptions = &value
	return b
}

// WithSources sets the Sources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Sources field is set to the value of the last call.
func (b *SyncOperationApplyConfiguration) WithSources(value applicationv1alpha1.ApplicationSources) *SyncOperationApplyConfiguration {
	b.Sources = &value
	return b
}

// WithRevisions adds the given value to the Revisions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Revisions field.
func (b *SyncOperationApplyConfiguration) WithRevisions(values ...string) *SyncOperationApplyConfiguration {
	for i := range values {
		b.Revisions = append(b.Revisions, values[i])
	}
	return b
}