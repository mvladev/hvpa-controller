/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	autoscalingv1alpha1 "k8s.io/autoscaler/hvpa-controller/pkg/apis/autoscaling/v1alpha1"
	versioned "k8s.io/autoscaler/hvpa-controller/pkg/client/clientset/versioned"
	internalinterfaces "k8s.io/autoscaler/hvpa-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/autoscaler/hvpa-controller/pkg/client/listers/autoscaling/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
)

// HvpaInformer provides access to a shared informer and lister for
// Hvpas.
type HvpaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HvpaLister
}

type hvpaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHvpaInformer constructs a new informer for Hvpa type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHvpaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHvpaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHvpaInformer constructs a new informer for Hvpa type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHvpaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha1().Hvpas(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha1().Hvpas(namespace).Watch(options)
			},
		},
		&autoscalingv1alpha1.Hvpa{},
		resyncPeriod,
		indexers,
	)
}

func (f *hvpaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHvpaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hvpaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autoscalingv1alpha1.Hvpa{}, f.defaultInformer)
}

func (f *hvpaInformer) Lister() v1alpha1.HvpaLister {
	return v1alpha1.NewHvpaLister(f.Informer().GetIndexer())
}
