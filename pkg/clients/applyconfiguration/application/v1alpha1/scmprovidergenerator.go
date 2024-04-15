// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// SCMProviderGeneratorApplyConfiguration represents an declarative configuration of the SCMProviderGenerator type for use
// with apply.
type SCMProviderGeneratorApplyConfiguration struct {
	Github              *SCMProviderGeneratorGithubApplyConfiguration          `json:"github,omitempty"`
	Gitlab              *SCMProviderGeneratorGitlabApplyConfiguration          `json:"gitlab,omitempty"`
	Bitbucket           *SCMProviderGeneratorBitbucketApplyConfiguration       `json:"bitbucket,omitempty"`
	BitbucketServer     *SCMProviderGeneratorBitbucketServerApplyConfiguration `json:"bitbucketServer,omitempty"`
	Gitea               *SCMProviderGeneratorGiteaApplyConfiguration           `json:"gitea,omitempty"`
	AzureDevOps         *SCMProviderGeneratorAzureDevOpsApplyConfiguration     `json:"azureDevOps,omitempty"`
	Filters             []SCMProviderGeneratorFilterApplyConfiguration         `json:"filters,omitempty"`
	CloneProtocol       *string                                                `json:"cloneProtocol,omitempty"`
	RequeueAfterSeconds *int64                                                 `json:"requeueAfterSeconds,omitempty"`
	Template            *ApplicationSetTemplateApplyConfiguration              `json:"template,omitempty"`
	Values              map[string]string                                      `json:"values,omitempty"`
	AWSCodeCommit       *SCMProviderGeneratorAWSCodeCommitApplyConfiguration   `json:"awsCodeCommit,omitempty"`
}

// SCMProviderGeneratorApplyConfiguration constructs an declarative configuration of the SCMProviderGenerator type for use with
// apply.
func SCMProviderGenerator() *SCMProviderGeneratorApplyConfiguration {
	return &SCMProviderGeneratorApplyConfiguration{}
}

// WithGithub sets the Github field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Github field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithGithub(value *SCMProviderGeneratorGithubApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.Github = value
	return b
}

// WithGitlab sets the Gitlab field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gitlab field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithGitlab(value *SCMProviderGeneratorGitlabApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.Gitlab = value
	return b
}

// WithBitbucket sets the Bitbucket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bitbucket field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithBitbucket(value *SCMProviderGeneratorBitbucketApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.Bitbucket = value
	return b
}

// WithBitbucketServer sets the BitbucketServer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BitbucketServer field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithBitbucketServer(value *SCMProviderGeneratorBitbucketServerApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.BitbucketServer = value
	return b
}

// WithGitea sets the Gitea field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gitea field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithGitea(value *SCMProviderGeneratorGiteaApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.Gitea = value
	return b
}

// WithAzureDevOps sets the AzureDevOps field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AzureDevOps field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithAzureDevOps(value *SCMProviderGeneratorAzureDevOpsApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.AzureDevOps = value
	return b
}

// WithFilters adds the given value to the Filters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Filters field.
func (b *SCMProviderGeneratorApplyConfiguration) WithFilters(values ...*SCMProviderGeneratorFilterApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFilters")
		}
		b.Filters = append(b.Filters, *values[i])
	}
	return b
}

// WithCloneProtocol sets the CloneProtocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CloneProtocol field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithCloneProtocol(value string) *SCMProviderGeneratorApplyConfiguration {
	b.CloneProtocol = &value
	return b
}

// WithRequeueAfterSeconds sets the RequeueAfterSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RequeueAfterSeconds field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithRequeueAfterSeconds(value int64) *SCMProviderGeneratorApplyConfiguration {
	b.RequeueAfterSeconds = &value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithTemplate(value *ApplicationSetTemplateApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.Template = value
	return b
}

// WithValues puts the entries into the Values field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Values field,
// overwriting an existing map entries in Values field with the same key.
func (b *SCMProviderGeneratorApplyConfiguration) WithValues(entries map[string]string) *SCMProviderGeneratorApplyConfiguration {
	if b.Values == nil && len(entries) > 0 {
		b.Values = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Values[k] = v
	}
	return b
}

// WithAWSCodeCommit sets the AWSCodeCommit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AWSCodeCommit field is set to the value of the last call.
func (b *SCMProviderGeneratorApplyConfiguration) WithAWSCodeCommit(value *SCMProviderGeneratorAWSCodeCommitApplyConfiguration) *SCMProviderGeneratorApplyConfiguration {
	b.AWSCodeCommit = value
	return b
}
