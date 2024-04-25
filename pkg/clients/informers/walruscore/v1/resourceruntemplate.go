// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	walruscorev1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	clientset "github.com/seal-io/walrus/pkg/clients/clientset"
	internalinterfaces "github.com/seal-io/walrus/pkg/clients/informers/internalinterfaces"
	v1 "github.com/seal-io/walrus/pkg/clients/listers/walruscore/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceRunTemplateInformer provides access to a shared informer and lister for
// ResourceRunTemplates.
type ResourceRunTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ResourceRunTemplateLister
}

type resourceRunTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceRunTemplateInformer constructs a new informer for ResourceRunTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceRunTemplateInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceRunTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceRunTemplateInformer constructs a new informer for ResourceRunTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceRunTemplateInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WalruscoreV1().ResourceRunTemplates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WalruscoreV1().ResourceRunTemplates(namespace).Watch(context.TODO(), options)
			},
		},
		&walruscorev1.ResourceRunTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceRunTemplateInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceRunTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceRunTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&walruscorev1.ResourceRunTemplate{}, f.defaultInformer)
}

func (f *resourceRunTemplateInformer) Lister() v1.ResourceRunTemplateLister {
	return v1.NewResourceRunTemplateLister(f.Informer().GetIndexer())
}