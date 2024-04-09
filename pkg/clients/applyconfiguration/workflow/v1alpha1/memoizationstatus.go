// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// MemoizationStatusApplyConfiguration represents an declarative configuration of the MemoizationStatus type for use
// with apply.
type MemoizationStatusApplyConfiguration struct {
	Hit       *bool   `json:"hit,omitempty"`
	Key       *string `json:"key,omitempty"`
	CacheName *string `json:"cacheName,omitempty"`
}

// MemoizationStatusApplyConfiguration constructs an declarative configuration of the MemoizationStatus type for use with
// apply.
func MemoizationStatus() *MemoizationStatusApplyConfiguration {
	return &MemoizationStatusApplyConfiguration{}
}

// WithHit sets the Hit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hit field is set to the value of the last call.
func (b *MemoizationStatusApplyConfiguration) WithHit(value bool) *MemoizationStatusApplyConfiguration {
	b.Hit = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *MemoizationStatusApplyConfiguration) WithKey(value string) *MemoizationStatusApplyConfiguration {
	b.Key = &value
	return b
}

// WithCacheName sets the CacheName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CacheName field is set to the value of the last call.
func (b *MemoizationStatusApplyConfiguration) WithCacheName(value string) *MemoizationStatusApplyConfiguration {
	b.CacheName = &value
	return b
}
