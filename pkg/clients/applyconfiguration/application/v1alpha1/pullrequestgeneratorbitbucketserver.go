// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// PullRequestGeneratorBitbucketServerApplyConfiguration represents an declarative configuration of the PullRequestGeneratorBitbucketServer type for use
// with apply.
type PullRequestGeneratorBitbucketServerApplyConfiguration struct {
	Project   *string                                     `json:"project,omitempty"`
	Repo      *string                                     `json:"repo,omitempty"`
	API       *string                                     `json:"api,omitempty"`
	BasicAuth *BasicAuthBitbucketServerApplyConfiguration `json:"basicAuth,omitempty"`
}

// PullRequestGeneratorBitbucketServerApplyConfiguration constructs an declarative configuration of the PullRequestGeneratorBitbucketServer type for use with
// apply.
func PullRequestGeneratorBitbucketServer() *PullRequestGeneratorBitbucketServerApplyConfiguration {
	return &PullRequestGeneratorBitbucketServerApplyConfiguration{}
}

// WithProject sets the Project field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Project field is set to the value of the last call.
func (b *PullRequestGeneratorBitbucketServerApplyConfiguration) WithProject(value string) *PullRequestGeneratorBitbucketServerApplyConfiguration {
	b.Project = &value
	return b
}

// WithRepo sets the Repo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Repo field is set to the value of the last call.
func (b *PullRequestGeneratorBitbucketServerApplyConfiguration) WithRepo(value string) *PullRequestGeneratorBitbucketServerApplyConfiguration {
	b.Repo = &value
	return b
}

// WithAPI sets the API field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the API field is set to the value of the last call.
func (b *PullRequestGeneratorBitbucketServerApplyConfiguration) WithAPI(value string) *PullRequestGeneratorBitbucketServerApplyConfiguration {
	b.API = &value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *PullRequestGeneratorBitbucketServerApplyConfiguration) WithBasicAuth(value *BasicAuthBitbucketServerApplyConfiguration) *PullRequestGeneratorBitbucketServerApplyConfiguration {
	b.BasicAuth = value
	return b
}
