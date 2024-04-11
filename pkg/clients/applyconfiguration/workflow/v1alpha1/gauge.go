// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// GaugeApplyConfiguration represents an declarative configuration of the Gauge type for use
// with apply.
type GaugeApplyConfiguration struct {
	Value     *string                  `json:"value,omitempty"`
	Realtime  *bool                    `json:"realtime,omitempty"`
	Operation *v1alpha1.GaugeOperation `json:"operation,omitempty"`
}

// GaugeApplyConfiguration constructs an declarative configuration of the Gauge type for use with
// apply.
func Gauge() *GaugeApplyConfiguration {
	return &GaugeApplyConfiguration{}
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *GaugeApplyConfiguration) WithValue(value string) *GaugeApplyConfiguration {
	b.Value = &value
	return b
}

// WithRealtime sets the Realtime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Realtime field is set to the value of the last call.
func (b *GaugeApplyConfiguration) WithRealtime(value bool) *GaugeApplyConfiguration {
	b.Realtime = &value
	return b
}

// WithOperation sets the Operation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operation field is set to the value of the last call.
func (b *GaugeApplyConfiguration) WithOperation(value v1alpha1.GaugeOperation) *GaugeApplyConfiguration {
	b.Operation = &value
	return b
}