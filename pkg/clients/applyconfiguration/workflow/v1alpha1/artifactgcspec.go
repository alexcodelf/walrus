// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// ArtifactGCSpecApplyConfiguration represents an declarative configuration of the ArtifactGCSpec type for use
// with apply.
type ArtifactGCSpecApplyConfiguration struct {
	ArtifactsByNode map[string]ArtifactNodeSpecApplyConfiguration `json:"artifactsByNode,omitempty"`
}

// ArtifactGCSpecApplyConfiguration constructs an declarative configuration of the ArtifactGCSpec type for use with
// apply.
func ArtifactGCSpec() *ArtifactGCSpecApplyConfiguration {
	return &ArtifactGCSpecApplyConfiguration{}
}

// WithArtifactsByNode puts the entries into the ArtifactsByNode field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ArtifactsByNode field,
// overwriting an existing map entries in ArtifactsByNode field with the same key.
func (b *ArtifactGCSpecApplyConfiguration) WithArtifactsByNode(entries map[string]ArtifactNodeSpecApplyConfiguration) *ArtifactGCSpecApplyConfiguration {
	if b.ArtifactsByNode == nil && len(entries) > 0 {
		b.ArtifactsByNode = make(map[string]ArtifactNodeSpecApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.ArtifactsByNode[k] = v
	}
	return b
}