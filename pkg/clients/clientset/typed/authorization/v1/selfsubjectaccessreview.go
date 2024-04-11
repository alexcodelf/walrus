// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	"context"

	scheme "github.com/seal-io/walrus/pkg/clients/clientset/scheme"
	v1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// SelfSubjectAccessReviewsGetter has a method to return a SelfSubjectAccessReviewInterface.
// A group's client should implement this interface.
type SelfSubjectAccessReviewsGetter interface {
	SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface
}

// SelfSubjectAccessReviewInterface has methods to work with SelfSubjectAccessReview resources.
type SelfSubjectAccessReviewInterface interface {
	Create(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.CreateOptions) (*v1.SelfSubjectAccessReview, error)
	SelfSubjectAccessReviewExpansion
}

// selfSubjectAccessReviews implements SelfSubjectAccessReviewInterface
type selfSubjectAccessReviews struct {
	client rest.Interface
}

// newSelfSubjectAccessReviews returns a SelfSubjectAccessReviews
func newSelfSubjectAccessReviews(c *AuthorizationV1Client) *selfSubjectAccessReviews {
	return &selfSubjectAccessReviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a selfSubjectAccessReview and creates it.  Returns the server's representation of the selfSubjectAccessReview, and an error, if there is any.
func (c *selfSubjectAccessReviews) Create(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.CreateOptions) (result *v1.SelfSubjectAccessReview, err error) {
	result = &v1.SelfSubjectAccessReview{}
	err = c.client.Post().
		Resource("selfsubjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectAccessReview).
		Do(ctx).
		Into(result)
	return
}