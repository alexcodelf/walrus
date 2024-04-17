// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	walruscorev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/walruscore/v1"
	scheme "github.com/seal-io/walrus/pkg/clients/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceComponentsGetter has a method to return a ResourceComponentsInterface.
// A group's client should implement this interface.
type ResourceComponentsGetter interface {
	ResourceComponents(namespace string) ResourceComponentsInterface
}

// ResourceComponentsInterface has methods to work with ResourceComponents resources.
type ResourceComponentsInterface interface {
	Create(ctx context.Context, resourceComponents *v1.ResourceComponents, opts metav1.CreateOptions) (*v1.ResourceComponents, error)
	Update(ctx context.Context, resourceComponents *v1.ResourceComponents, opts metav1.UpdateOptions) (*v1.ResourceComponents, error)
	UpdateStatus(ctx context.Context, resourceComponents *v1.ResourceComponents, opts metav1.UpdateOptions) (*v1.ResourceComponents, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ResourceComponents, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ResourceComponentsList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceComponents, err error)
	Apply(ctx context.Context, resourceComponents *walruscorev1.ResourceComponentsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceComponents, err error)
	ApplyStatus(ctx context.Context, resourceComponents *walruscorev1.ResourceComponentsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceComponents, err error)
	ResourceComponentsExpansion
}

// resourceComponents implements ResourceComponentsInterface
type resourceComponents struct {
	client rest.Interface
	ns     string
}

// newResourceComponents returns a ResourceComponents
func newResourceComponents(c *WalruscoreV1Client, namespace string) *resourceComponents {
	return &resourceComponents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceComponents, and returns the corresponding resourceComponents object, and an error if there is any.
func (c *resourceComponents) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ResourceComponents, err error) {
	result = &v1.ResourceComponents{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceComponents that match those selectors.
func (c *resourceComponents) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ResourceComponentsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ResourceComponentsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcecomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceComponents.
func (c *resourceComponents) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcecomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceComponents and creates it.  Returns the server's representation of the resourceComponents, and an error, if there is any.
func (c *resourceComponents) Create(ctx context.Context, resourceComponents *v1.ResourceComponents, opts metav1.CreateOptions) (result *v1.ResourceComponents, err error) {
	result = &v1.ResourceComponents{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcecomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceComponents).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceComponents and updates it. Returns the server's representation of the resourceComponents, and an error, if there is any.
func (c *resourceComponents) Update(ctx context.Context, resourceComponents *v1.ResourceComponents, opts metav1.UpdateOptions) (result *v1.ResourceComponents, err error) {
	result = &v1.ResourceComponents{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(resourceComponents.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceComponents).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *resourceComponents) UpdateStatus(ctx context.Context, resourceComponents *v1.ResourceComponents, opts metav1.UpdateOptions) (result *v1.ResourceComponents, err error) {
	result = &v1.ResourceComponents{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(resourceComponents.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceComponents).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceComponents and deletes it. Returns an error if one occurs.
func (c *resourceComponents) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceComponents) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcecomponents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceComponents.
func (c *resourceComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceComponents, err error) {
	result = &v1.ResourceComponents{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceComponents.
func (c *resourceComponents) Apply(ctx context.Context, resourceComponents *walruscorev1.ResourceComponentsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceComponents, err error) {
	if resourceComponents == nil {
		return nil, fmt.Errorf("resourceComponents provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(resourceComponents)
	if err != nil {
		return nil, err
	}
	name := resourceComponents.Name
	if name == nil {
		return nil, fmt.Errorf("resourceComponents.Name must be provided to Apply")
	}
	result = &v1.ResourceComponents{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *resourceComponents) ApplyStatus(ctx context.Context, resourceComponents *walruscorev1.ResourceComponentsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceComponents, err error) {
	if resourceComponents == nil {
		return nil, fmt.Errorf("resourceComponents provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(resourceComponents)
	if err != nil {
		return nil, err
	}

	name := resourceComponents.Name
	if name == nil {
		return nil, fmt.Errorf("resourceComponents.Name must be provided to Apply")
	}

	result = &v1.ResourceComponents{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("resourcecomponents").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
