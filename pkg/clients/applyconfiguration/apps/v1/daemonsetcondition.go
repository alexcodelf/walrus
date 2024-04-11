// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DaemonSetConditionApplyConfiguration represents an declarative configuration of the DaemonSetCondition type for use
// with apply.
type DaemonSetConditionApplyConfiguration struct {
	Type               *v1.DaemonSetConditionType `json:"type,omitempty"`
	Status             *corev1.ConditionStatus    `json:"status,omitempty"`
	LastTransitionTime *metav1.Time               `json:"lastTransitionTime,omitempty"`
	Reason             *string                    `json:"reason,omitempty"`
	Message            *string                    `json:"message,omitempty"`
}

// DaemonSetConditionApplyConfiguration constructs an declarative configuration of the DaemonSetCondition type for use with
// apply.
func DaemonSetCondition() *DaemonSetConditionApplyConfiguration {
	return &DaemonSetConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DaemonSetConditionApplyConfiguration) WithType(value v1.DaemonSetConditionType) *DaemonSetConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *DaemonSetConditionApplyConfiguration) WithStatus(value corev1.ConditionStatus) *DaemonSetConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *DaemonSetConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *DaemonSetConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *DaemonSetConditionApplyConfiguration) WithReason(value string) *DaemonSetConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *DaemonSetConditionApplyConfiguration) WithMessage(value string) *DaemonSetConditionApplyConfiguration {
	b.Message = &value
	return b
}