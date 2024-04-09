// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	workflowv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// WorkflowStepApplyConfiguration represents an declarative configuration of the WorkflowStep type for use
// with apply.
type WorkflowStepApplyConfiguration struct {
	Name         *string                          `json:"name,omitempty"`
	Template     *string                          `json:"template,omitempty"`
	Inline       *TemplateApplyConfiguration      `json:"inline,omitempty"`
	Arguments    *ArgumentsApplyConfiguration     `json:"arguments,omitempty"`
	TemplateRef  *TemplateRefApplyConfiguration   `json:"templateRef,omitempty"`
	WithItems    []workflowv1alpha1.Item          `json:"withItems,omitempty"`
	WithParam    *string                          `json:"withParam,omitempty"`
	WithSequence *SequenceApplyConfiguration      `json:"withSequence,omitempty"`
	When         *string                          `json:"when,omitempty"`
	ContinueOn   *ContinueOnApplyConfiguration    `json:"continueOn,omitempty"`
	OnExit       *string                          `json:"onExit,omitempty"`
	Hooks        *workflowv1alpha1.LifecycleHooks `json:"hooks,omitempty"`
}

// WorkflowStepApplyConfiguration constructs an declarative configuration of the WorkflowStep type for use with
// apply.
func WorkflowStep() *WorkflowStepApplyConfiguration {
	return &WorkflowStepApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithName(value string) *WorkflowStepApplyConfiguration {
	b.Name = &value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithTemplate(value string) *WorkflowStepApplyConfiguration {
	b.Template = &value
	return b
}

// WithInline sets the Inline field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Inline field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithInline(value *TemplateApplyConfiguration) *WorkflowStepApplyConfiguration {
	b.Inline = value
	return b
}

// WithArguments sets the Arguments field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Arguments field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithArguments(value *ArgumentsApplyConfiguration) *WorkflowStepApplyConfiguration {
	b.Arguments = value
	return b
}

// WithTemplateRef sets the TemplateRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TemplateRef field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithTemplateRef(value *TemplateRefApplyConfiguration) *WorkflowStepApplyConfiguration {
	b.TemplateRef = value
	return b
}

// WithWithItems adds the given value to the WithItems field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the WithItems field.
func (b *WorkflowStepApplyConfiguration) WithWithItems(values ...workflowv1alpha1.Item) *WorkflowStepApplyConfiguration {
	for i := range values {
		b.WithItems = append(b.WithItems, values[i])
	}
	return b
}

// WithWithParam sets the WithParam field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WithParam field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithWithParam(value string) *WorkflowStepApplyConfiguration {
	b.WithParam = &value
	return b
}

// WithWithSequence sets the WithSequence field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WithSequence field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithWithSequence(value *SequenceApplyConfiguration) *WorkflowStepApplyConfiguration {
	b.WithSequence = value
	return b
}

// WithWhen sets the When field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the When field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithWhen(value string) *WorkflowStepApplyConfiguration {
	b.When = &value
	return b
}

// WithContinueOn sets the ContinueOn field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContinueOn field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithContinueOn(value *ContinueOnApplyConfiguration) *WorkflowStepApplyConfiguration {
	b.ContinueOn = value
	return b
}

// WithOnExit sets the OnExit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OnExit field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithOnExit(value string) *WorkflowStepApplyConfiguration {
	b.OnExit = &value
	return b
}

// WithHooks sets the Hooks field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hooks field is set to the value of the last call.
func (b *WorkflowStepApplyConfiguration) WithHooks(value workflowv1alpha1.LifecycleHooks) *WorkflowStepApplyConfiguration {
	b.Hooks = &value
	return b
}
