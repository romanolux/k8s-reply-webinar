/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	k8dynamov1 "github.com/romanolux/k8s-reply-webinar/pkg/apis/k8dynamo/v1"
	versioned "github.com/romanolux/k8s-reply-webinar/pkg/client/clientset/versioned"
	internalinterfaces "github.com/romanolux/k8s-reply-webinar/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/romanolux/k8s-reply-webinar/pkg/client/listers/k8dynamo/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// K8dynamoInformer provides access to a shared informer and lister for
// K8dynamos.
type K8dynamoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.K8dynamoLister
}

type k8dynamoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewK8dynamoInformer constructs a new informer for K8dynamo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewK8dynamoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredK8dynamoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredK8dynamoInformer constructs a new informer for K8dynamo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredK8dynamoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StormV1().K8dynamos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StormV1().K8dynamos(namespace).Watch(context.TODO(), options)
			},
		},
		&k8dynamov1.K8dynamo{},
		resyncPeriod,
		indexers,
	)
}

func (f *k8dynamoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredK8dynamoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *k8dynamoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&k8dynamov1.K8dynamo{}, f.defaultInformer)
}

func (f *k8dynamoInformer) Lister() v1.K8dynamoLister {
	return v1.NewK8dynamoLister(f.Informer().GetIndexer())
}
