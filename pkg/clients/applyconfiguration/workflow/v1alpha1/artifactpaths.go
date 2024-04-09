// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// ArtifactPathsApplyConfiguration represents an declarative configuration of the ArtifactPaths type for use
// with apply.
type ArtifactPathsApplyConfiguration struct {
	ArtifactApplyConfiguration `json:",inline"`
}

// ArtifactPathsApplyConfiguration constructs an declarative configuration of the ArtifactPaths type for use with
// apply.
func ArtifactPaths() *ArtifactPathsApplyConfiguration {
	return &ArtifactPathsApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithName(value string) *ArtifactPathsApplyConfiguration {
	b.Name = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithPath(value string) *ArtifactPathsApplyConfiguration {
	b.Path = &value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithMode(value int32) *ArtifactPathsApplyConfiguration {
	b.Mode = &value
	return b
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithFrom(value string) *ArtifactPathsApplyConfiguration {
	b.From = &value
	return b
}

// WithArchiveLogs sets the ArchiveLogs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArchiveLogs field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithArchiveLogs(value bool) *ArtifactPathsApplyConfiguration {
	b.ArchiveLogs = &value
	return b
}

// WithS3 sets the S3 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the S3 field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithS3(value *S3ArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.S3 = value
	return b
}

// WithGit sets the Git field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Git field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithGit(value *GitArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.Git = value
	return b
}

// WithHTTP sets the HTTP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTP field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithHTTP(value *HTTPArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.HTTP = value
	return b
}

// WithArtifactory sets the Artifactory field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Artifactory field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithArtifactory(value *ArtifactoryArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.Artifactory = value
	return b
}

// WithHDFS sets the HDFS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HDFS field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithHDFS(value *HDFSArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.HDFS = value
	return b
}

// WithRaw sets the Raw field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Raw field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithRaw(value *RawArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.Raw = value
	return b
}

// WithOSS sets the OSS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSS field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithOSS(value *OSSArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.OSS = value
	return b
}

// WithGCS sets the GCS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GCS field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithGCS(value *GCSArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.GCS = value
	return b
}

// WithAzure sets the Azure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Azure field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithAzure(value *AzureArtifactApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.Azure = value
	return b
}

// WithGlobalName sets the GlobalName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GlobalName field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithGlobalName(value string) *ArtifactPathsApplyConfiguration {
	b.GlobalName = &value
	return b
}

// WithArchive sets the Archive field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Archive field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithArchive(value *ArchiveStrategyApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.Archive = value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithOptional(value bool) *ArtifactPathsApplyConfiguration {
	b.Optional = &value
	return b
}

// WithSubPath sets the SubPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubPath field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithSubPath(value string) *ArtifactPathsApplyConfiguration {
	b.SubPath = &value
	return b
}

// WithRecurseMode sets the RecurseMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RecurseMode field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithRecurseMode(value bool) *ArtifactPathsApplyConfiguration {
	b.RecurseMode = &value
	return b
}

// WithFromExpression sets the FromExpression field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FromExpression field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithFromExpression(value string) *ArtifactPathsApplyConfiguration {
	b.FromExpression = &value
	return b
}

// WithArtifactGC sets the ArtifactGC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArtifactGC field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithArtifactGC(value *ArtifactGCApplyConfiguration) *ArtifactPathsApplyConfiguration {
	b.ArtifactGC = value
	return b
}

// WithDeleted sets the Deleted field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Deleted field is set to the value of the last call.
func (b *ArtifactPathsApplyConfiguration) WithDeleted(value bool) *ArtifactPathsApplyConfiguration {
	b.Deleted = &value
	return b
}
