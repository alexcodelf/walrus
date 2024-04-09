// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// RetryAffinityApplyConfiguration represents an declarative configuration of the RetryAffinity type for use
// with apply.
type RetryAffinityApplyConfiguration struct {
	NodeAntiAffinity *v1alpha1.RetryNodeAntiAffinity `json:"nodeAntiAffinity,omitempty"`
}

// RetryAffinityApplyConfiguration constructs an declarative configuration of the RetryAffinity type for use with
// apply.
func RetryAffinity() *RetryAffinityApplyConfiguration {
	return &RetryAffinityApplyConfiguration{}
}

// WithNodeAntiAffinity sets the NodeAntiAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeAntiAffinity field is set to the value of the last call.
func (b *RetryAffinityApplyConfiguration) WithNodeAntiAffinity(value v1alpha1.RetryNodeAntiAffinity) *RetryAffinityApplyConfiguration {
	b.NodeAntiAffinity = &value
	return b
}
