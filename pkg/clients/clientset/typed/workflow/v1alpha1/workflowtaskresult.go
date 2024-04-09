// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	workflowv1alpha1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/workflow/v1alpha1"
	scheme "github.com/seal-io/walrus/pkg/clients/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkflowTaskResultsGetter has a method to return a WorkflowTaskResultInterface.
// A group's client should implement this interface.
type WorkflowTaskResultsGetter interface {
	WorkflowTaskResults(namespace string) WorkflowTaskResultInterface
}

// WorkflowTaskResultInterface has methods to work with WorkflowTaskResult resources.
type WorkflowTaskResultInterface interface {
	Create(ctx context.Context, workflowTaskResult *v1alpha1.WorkflowTaskResult, opts v1.CreateOptions) (*v1alpha1.WorkflowTaskResult, error)
	Update(ctx context.Context, workflowTaskResult *v1alpha1.WorkflowTaskResult, opts v1.UpdateOptions) (*v1alpha1.WorkflowTaskResult, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WorkflowTaskResult, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WorkflowTaskResultList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkflowTaskResult, err error)
	Apply(ctx context.Context, workflowTaskResult *workflowv1alpha1.WorkflowTaskResultApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WorkflowTaskResult, err error)
	WorkflowTaskResultExpansion
}

// workflowTaskResults implements WorkflowTaskResultInterface
type workflowTaskResults struct {
	client rest.Interface
	ns     string
}

// newWorkflowTaskResults returns a WorkflowTaskResults
func newWorkflowTaskResults(c *ArgoprojV1alpha1Client, namespace string) *workflowTaskResults {
	return &workflowTaskResults{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the workflowTaskResult, and returns the corresponding workflowTaskResult object, and an error if there is any.
func (c *workflowTaskResults) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkflowTaskResult, err error) {
	result = &v1alpha1.WorkflowTaskResult{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkflowTaskResults that match those selectors.
func (c *workflowTaskResults) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkflowTaskResultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WorkflowTaskResultList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workflowTaskResults.
func (c *workflowTaskResults) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workflowTaskResult and creates it.  Returns the server's representation of the workflowTaskResult, and an error, if there is any.
func (c *workflowTaskResults) Create(ctx context.Context, workflowTaskResult *v1alpha1.WorkflowTaskResult, opts v1.CreateOptions) (result *v1alpha1.WorkflowTaskResult, err error) {
	result = &v1alpha1.WorkflowTaskResult{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workflowTaskResult).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workflowTaskResult and updates it. Returns the server's representation of the workflowTaskResult, and an error, if there is any.
func (c *workflowTaskResults) Update(ctx context.Context, workflowTaskResult *v1alpha1.WorkflowTaskResult, opts v1.UpdateOptions) (result *v1alpha1.WorkflowTaskResult, err error) {
	result = &v1alpha1.WorkflowTaskResult{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		Name(workflowTaskResult.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workflowTaskResult).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workflowTaskResult and deletes it. Returns an error if one occurs.
func (c *workflowTaskResults) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workflowTaskResults) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workflowtaskresults").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workflowTaskResult.
func (c *workflowTaskResults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkflowTaskResult, err error) {
	result = &v1alpha1.WorkflowTaskResult{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("workflowtaskresults").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied workflowTaskResult.
func (c *workflowTaskResults) Apply(ctx context.Context, workflowTaskResult *workflowv1alpha1.WorkflowTaskResultApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WorkflowTaskResult, err error) {
	if workflowTaskResult == nil {
		return nil, fmt.Errorf("workflowTaskResult provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(workflowTaskResult)
	if err != nil {
		return nil, err
	}
	name := workflowTaskResult.Name
	if name == nil {
		return nil, fmt.Errorf("workflowTaskResult.Name must be provided to Apply")
	}
	result = &v1alpha1.WorkflowTaskResult{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("workflowtaskresults").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
