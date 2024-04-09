// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	workflowv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	clientset "github.com/seal-io/walrus/pkg/clients/clientset"
	internalinterfaces "github.com/seal-io/walrus/pkg/clients/informers/internalinterfaces"
	v1alpha1 "github.com/seal-io/walrus/pkg/clients/listers/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WorkflowArtifactGCTaskInformer provides access to a shared informer and lister for
// WorkflowArtifactGCTasks.
type WorkflowArtifactGCTaskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WorkflowArtifactGCTaskLister
}

type workflowArtifactGCTaskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkflowArtifactGCTaskInformer constructs a new informer for WorkflowArtifactGCTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkflowArtifactGCTaskInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkflowArtifactGCTaskInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkflowArtifactGCTaskInformer constructs a new informer for WorkflowArtifactGCTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkflowArtifactGCTaskInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().WorkflowArtifactGCTasks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().WorkflowArtifactGCTasks(namespace).Watch(context.TODO(), options)
			},
		},
		&workflowv1alpha1.WorkflowArtifactGCTask{},
		resyncPeriod,
		indexers,
	)
}

func (f *workflowArtifactGCTaskInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkflowArtifactGCTaskInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workflowArtifactGCTaskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workflowv1alpha1.WorkflowArtifactGCTask{}, f.defaultInformer)
}

func (f *workflowArtifactGCTaskInformer) Lister() v1alpha1.WorkflowArtifactGCTaskLister {
	return v1alpha1.NewWorkflowArtifactGCTaskLister(f.Informer().GetIndexer())
}
