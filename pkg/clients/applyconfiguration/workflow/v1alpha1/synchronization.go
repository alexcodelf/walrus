// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// SynchronizationApplyConfiguration represents an declarative configuration of the Synchronization type for use
// with apply.
type SynchronizationApplyConfiguration struct {
	Semaphore *SemaphoreRefApplyConfiguration `json:"semaphore,omitempty"`
	Mutex     *MutexApplyConfiguration        `json:"mutex,omitempty"`
}

// SynchronizationApplyConfiguration constructs an declarative configuration of the Synchronization type for use with
// apply.
func Synchronization() *SynchronizationApplyConfiguration {
	return &SynchronizationApplyConfiguration{}
}

// WithSemaphore sets the Semaphore field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Semaphore field is set to the value of the last call.
func (b *SynchronizationApplyConfiguration) WithSemaphore(value *SemaphoreRefApplyConfiguration) *SynchronizationApplyConfiguration {
	b.Semaphore = value
	return b
}

// WithMutex sets the Mutex field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mutex field is set to the value of the last call.
func (b *SynchronizationApplyConfiguration) WithMutex(value *MutexApplyConfiguration) *SynchronizationApplyConfiguration {
	b.Mutex = value
	return b
}
