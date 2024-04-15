// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApplicationConditionApplyConfiguration represents an declarative configuration of the ApplicationCondition type for use
// with apply.
type ApplicationConditionApplyConfiguration struct {
	Type               *string  `json:"type,omitempty"`
	Message            *string  `json:"message,omitempty"`
	LastTransitionTime *v1.Time `json:"lastTransitionTime,omitempty"`
}

// ApplicationConditionApplyConfiguration constructs an declarative configuration of the ApplicationCondition type for use with
// apply.
func ApplicationCondition() *ApplicationConditionApplyConfiguration {
	return &ApplicationConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ApplicationConditionApplyConfiguration) WithType(value string) *ApplicationConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ApplicationConditionApplyConfiguration) WithMessage(value string) *ApplicationConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ApplicationConditionApplyConfiguration) WithLastTransitionTime(value v1.Time) *ApplicationConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
