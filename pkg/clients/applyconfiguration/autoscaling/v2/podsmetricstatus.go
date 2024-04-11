// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v2

// PodsMetricStatusApplyConfiguration represents an declarative configuration of the PodsMetricStatus type for use
// with apply.
type PodsMetricStatusApplyConfiguration struct {
	Metric  *MetricIdentifierApplyConfiguration  `json:"metric,omitempty"`
	Current *MetricValueStatusApplyConfiguration `json:"current,omitempty"`
}

// PodsMetricStatusApplyConfiguration constructs an declarative configuration of the PodsMetricStatus type for use with
// apply.
func PodsMetricStatus() *PodsMetricStatusApplyConfiguration {
	return &PodsMetricStatusApplyConfiguration{}
}

// WithMetric sets the Metric field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Metric field is set to the value of the last call.
func (b *PodsMetricStatusApplyConfiguration) WithMetric(value *MetricIdentifierApplyConfiguration) *PodsMetricStatusApplyConfiguration {
	b.Metric = value
	return b
}

// WithCurrent sets the Current field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Current field is set to the value of the last call.
func (b *PodsMetricStatusApplyConfiguration) WithCurrent(value *MetricValueStatusApplyConfiguration) *PodsMetricStatusApplyConfiguration {
	b.Current = value
	return b
}