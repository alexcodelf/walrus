// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// ContainerSetRetryStrategyApplyConfiguration represents an declarative configuration of the ContainerSetRetryStrategy type for use
// with apply.
type ContainerSetRetryStrategyApplyConfiguration struct {
	Duration *string             `json:"duration,omitempty"`
	Retries  *intstr.IntOrString `json:"retries,omitempty"`
}

// ContainerSetRetryStrategyApplyConfiguration constructs an declarative configuration of the ContainerSetRetryStrategy type for use with
// apply.
func ContainerSetRetryStrategy() *ContainerSetRetryStrategyApplyConfiguration {
	return &ContainerSetRetryStrategyApplyConfiguration{}
}

// WithDuration sets the Duration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Duration field is set to the value of the last call.
func (b *ContainerSetRetryStrategyApplyConfiguration) WithDuration(value string) *ContainerSetRetryStrategyApplyConfiguration {
	b.Duration = &value
	return b
}

// WithRetries sets the Retries field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Retries field is set to the value of the last call.
func (b *ContainerSetRetryStrategyApplyConfiguration) WithRetries(value intstr.IntOrString) *ContainerSetRetryStrategyApplyConfiguration {
	b.Retries = &value
	return b
}
