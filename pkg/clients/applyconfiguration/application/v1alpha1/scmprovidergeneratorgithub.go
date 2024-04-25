// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// SCMProviderGeneratorGithubApplyConfiguration represents an declarative configuration of the SCMProviderGeneratorGithub type for use
// with apply.
type SCMProviderGeneratorGithubApplyConfiguration struct {
	Organization  *string                      `json:"organization,omitempty"`
	API           *string                      `json:"api,omitempty"`
	TokenRef      *SecretRefApplyConfiguration `json:"tokenRef,omitempty"`
	AppSecretName *string                      `json:"appSecretName,omitempty"`
	AllBranches   *bool                        `json:"allBranches,omitempty"`
}

// SCMProviderGeneratorGithubApplyConfiguration constructs an declarative configuration of the SCMProviderGeneratorGithub type for use with
// apply.
func SCMProviderGeneratorGithub() *SCMProviderGeneratorGithubApplyConfiguration {
	return &SCMProviderGeneratorGithubApplyConfiguration{}
}

// WithOrganization sets the Organization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Organization field is set to the value of the last call.
func (b *SCMProviderGeneratorGithubApplyConfiguration) WithOrganization(value string) *SCMProviderGeneratorGithubApplyConfiguration {
	b.Organization = &value
	return b
}

// WithAPI sets the API field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the API field is set to the value of the last call.
func (b *SCMProviderGeneratorGithubApplyConfiguration) WithAPI(value string) *SCMProviderGeneratorGithubApplyConfiguration {
	b.API = &value
	return b
}

// WithTokenRef sets the TokenRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TokenRef field is set to the value of the last call.
func (b *SCMProviderGeneratorGithubApplyConfiguration) WithTokenRef(value *SecretRefApplyConfiguration) *SCMProviderGeneratorGithubApplyConfiguration {
	b.TokenRef = value
	return b
}

// WithAppSecretName sets the AppSecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AppSecretName field is set to the value of the last call.
func (b *SCMProviderGeneratorGithubApplyConfiguration) WithAppSecretName(value string) *SCMProviderGeneratorGithubApplyConfiguration {
	b.AppSecretName = &value
	return b
}

// WithAllBranches sets the AllBranches field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllBranches field is set to the value of the last call.
func (b *SCMProviderGeneratorGithubApplyConfiguration) WithAllBranches(value bool) *SCMProviderGeneratorGithubApplyConfiguration {
	b.AllBranches = &value
	return b
}