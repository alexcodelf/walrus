// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	corev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WorkflowStatusApplyConfiguration represents an declarative configuration of the WorkflowStatus type for use
// with apply.
type WorkflowStatusApplyConfiguration struct {
	Phase                       *v1alpha1.WorkflowPhase                        `json:"phase,omitempty"`
	StartedAt                   *v1.Time                                       `json:"startedAt,omitempty"`
	FinishedAt                  *v1.Time                                       `json:"finishedAt,omitempty"`
	EstimatedDuration           *v1alpha1.EstimatedDuration                    `json:"estimatedDuration,omitempty"`
	Progress                    *v1alpha1.Progress                             `json:"progress,omitempty"`
	Message                     *string                                        `json:"message,omitempty"`
	CompressedNodes             *string                                        `json:"compressedNodes,omitempty"`
	Nodes                       *v1alpha1.Nodes                                `json:"nodes,omitempty"`
	OffloadNodeStatusVersion    *string                                        `json:"offloadNodeStatusVersion,omitempty"`
	StoredTemplates             map[string]TemplateApplyConfiguration          `json:"storedTemplates,omitempty"`
	PersistentVolumeClaims      []corev1.VolumeApplyConfiguration              `json:"persistentVolumeClaims,omitempty"`
	Outputs                     *OutputsApplyConfiguration                     `json:"outputs,omitempty"`
	Conditions                  *v1alpha1.Conditions                           `json:"conditions,omitempty"`
	ResourcesDuration           *v1alpha1.ResourcesDuration                    `json:"resourcesDuration,omitempty"`
	StoredWorkflowSpec          *WorkflowSpecApplyConfiguration                `json:"storedWorkflowTemplateSpec,omitempty"`
	Synchronization             *SynchronizationStatusApplyConfiguration       `json:"synchronization,omitempty"`
	ArtifactRepositoryRef       *ArtifactRepositoryRefStatusApplyConfiguration `json:"artifactRepositoryRef,omitempty"`
	ArtifactGCStatus            *ArtGCStatusApplyConfiguration                 `json:"artifactGCStatus,omitempty"`
	TaskResultsCompletionStatus map[string]bool                                `json:"taskResultsCompletionStatus,omitempty"`
}

// WorkflowStatusApplyConfiguration constructs an declarative configuration of the WorkflowStatus type for use with
// apply.
func WorkflowStatus() *WorkflowStatusApplyConfiguration {
	return &WorkflowStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithPhase(value v1alpha1.WorkflowPhase) *WorkflowStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithStartedAt sets the StartedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartedAt field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithStartedAt(value v1.Time) *WorkflowStatusApplyConfiguration {
	b.StartedAt = &value
	return b
}

// WithFinishedAt sets the FinishedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FinishedAt field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithFinishedAt(value v1.Time) *WorkflowStatusApplyConfiguration {
	b.FinishedAt = &value
	return b
}

// WithEstimatedDuration sets the EstimatedDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EstimatedDuration field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithEstimatedDuration(value v1alpha1.EstimatedDuration) *WorkflowStatusApplyConfiguration {
	b.EstimatedDuration = &value
	return b
}

// WithProgress sets the Progress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Progress field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithProgress(value v1alpha1.Progress) *WorkflowStatusApplyConfiguration {
	b.Progress = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithMessage(value string) *WorkflowStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithCompressedNodes sets the CompressedNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompressedNodes field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithCompressedNodes(value string) *WorkflowStatusApplyConfiguration {
	b.CompressedNodes = &value
	return b
}

// WithNodes sets the Nodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Nodes field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithNodes(value v1alpha1.Nodes) *WorkflowStatusApplyConfiguration {
	b.Nodes = &value
	return b
}

// WithOffloadNodeStatusVersion sets the OffloadNodeStatusVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OffloadNodeStatusVersion field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithOffloadNodeStatusVersion(value string) *WorkflowStatusApplyConfiguration {
	b.OffloadNodeStatusVersion = &value
	return b
}

// WithStoredTemplates puts the entries into the StoredTemplates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the StoredTemplates field,
// overwriting an existing map entries in StoredTemplates field with the same key.
func (b *WorkflowStatusApplyConfiguration) WithStoredTemplates(entries map[string]TemplateApplyConfiguration) *WorkflowStatusApplyConfiguration {
	if b.StoredTemplates == nil && len(entries) > 0 {
		b.StoredTemplates = make(map[string]TemplateApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.StoredTemplates[k] = v
	}
	return b
}

// WithPersistentVolumeClaims adds the given value to the PersistentVolumeClaims field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PersistentVolumeClaims field.
func (b *WorkflowStatusApplyConfiguration) WithPersistentVolumeClaims(values ...*corev1.VolumeApplyConfiguration) *WorkflowStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPersistentVolumeClaims")
		}
		b.PersistentVolumeClaims = append(b.PersistentVolumeClaims, *values[i])
	}
	return b
}

// WithOutputs sets the Outputs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Outputs field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithOutputs(value *OutputsApplyConfiguration) *WorkflowStatusApplyConfiguration {
	b.Outputs = value
	return b
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithConditions(value v1alpha1.Conditions) *WorkflowStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithResourcesDuration sets the ResourcesDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourcesDuration field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithResourcesDuration(value v1alpha1.ResourcesDuration) *WorkflowStatusApplyConfiguration {
	b.ResourcesDuration = &value
	return b
}

// WithStoredWorkflowSpec sets the StoredWorkflowSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StoredWorkflowSpec field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithStoredWorkflowSpec(value *WorkflowSpecApplyConfiguration) *WorkflowStatusApplyConfiguration {
	b.StoredWorkflowSpec = value
	return b
}

// WithSynchronization sets the Synchronization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Synchronization field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithSynchronization(value *SynchronizationStatusApplyConfiguration) *WorkflowStatusApplyConfiguration {
	b.Synchronization = value
	return b
}

// WithArtifactRepositoryRef sets the ArtifactRepositoryRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArtifactRepositoryRef field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithArtifactRepositoryRef(value *ArtifactRepositoryRefStatusApplyConfiguration) *WorkflowStatusApplyConfiguration {
	b.ArtifactRepositoryRef = value
	return b
}

// WithArtifactGCStatus sets the ArtifactGCStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArtifactGCStatus field is set to the value of the last call.
func (b *WorkflowStatusApplyConfiguration) WithArtifactGCStatus(value *ArtGCStatusApplyConfiguration) *WorkflowStatusApplyConfiguration {
	b.ArtifactGCStatus = value
	return b
}

// WithTaskResultsCompletionStatus puts the entries into the TaskResultsCompletionStatus field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the TaskResultsCompletionStatus field,
// overwriting an existing map entries in TaskResultsCompletionStatus field with the same key.
func (b *WorkflowStatusApplyConfiguration) WithTaskResultsCompletionStatus(entries map[string]bool) *WorkflowStatusApplyConfiguration {
	if b.TaskResultsCompletionStatus == nil && len(entries) > 0 {
		b.TaskResultsCompletionStatus = make(map[string]bool, len(entries))
	}
	for k, v := range entries {
		b.TaskResultsCompletionStatus[k] = v
	}
	return b
}
