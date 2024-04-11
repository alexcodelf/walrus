// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	corev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// DaemonSetSpecApplyConfiguration represents an declarative configuration of the DaemonSetSpec type for use
// with apply.
type DaemonSetSpecApplyConfiguration struct {
	Selector             *v1.LabelSelectorApplyConfiguration        `json:"selector,omitempty"`
	Template             *corev1.PodTemplateSpecApplyConfiguration  `json:"template,omitempty"`
	UpdateStrategy       *DaemonSetUpdateStrategyApplyConfiguration `json:"updateStrategy,omitempty"`
	MinReadySeconds      *int32                                     `json:"minReadySeconds,omitempty"`
	RevisionHistoryLimit *int32                                     `json:"revisionHistoryLimit,omitempty"`
}

// DaemonSetSpecApplyConfiguration constructs an declarative configuration of the DaemonSetSpec type for use with
// apply.
func DaemonSetSpec() *DaemonSetSpecApplyConfiguration {
	return &DaemonSetSpecApplyConfiguration{}
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *DaemonSetSpecApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *DaemonSetSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *DaemonSetSpecApplyConfiguration) WithTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *DaemonSetSpecApplyConfiguration {
	b.Template = value
	return b
}

// WithUpdateStrategy sets the UpdateStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdateStrategy field is set to the value of the last call.
func (b *DaemonSetSpecApplyConfiguration) WithUpdateStrategy(value *DaemonSetUpdateStrategyApplyConfiguration) *DaemonSetSpecApplyConfiguration {
	b.UpdateStrategy = value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *DaemonSetSpecApplyConfiguration) WithMinReadySeconds(value int32) *DaemonSetSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithRevisionHistoryLimit sets the RevisionHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevisionHistoryLimit field is set to the value of the last call.
func (b *DaemonSetSpecApplyConfiguration) WithRevisionHistoryLimit(value int32) *DaemonSetSpecApplyConfiguration {
	b.RevisionHistoryLimit = &value
	return b
}