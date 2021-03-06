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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/romanolux/k8s-reply-webinar/pkg/apis/k8dynamo/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// K8dynamoLister helps list K8dynamos.
// All objects returned here must be treated as read-only.
type K8dynamoLister interface {
	// List lists all K8dynamos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.K8dynamo, err error)
	// K8dynamos returns an object that can list and get K8dynamos.
	K8dynamos(namespace string) K8dynamoNamespaceLister
	K8dynamoListerExpansion
}

// k8dynamoLister implements the K8dynamoLister interface.
type k8dynamoLister struct {
	indexer cache.Indexer
}

// NewK8dynamoLister returns a new K8dynamoLister.
func NewK8dynamoLister(indexer cache.Indexer) K8dynamoLister {
	return &k8dynamoLister{indexer: indexer}
}

// List lists all K8dynamos in the indexer.
func (s *k8dynamoLister) List(selector labels.Selector) (ret []*v1.K8dynamo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.K8dynamo))
	})
	return ret, err
}

// K8dynamos returns an object that can list and get K8dynamos.
func (s *k8dynamoLister) K8dynamos(namespace string) K8dynamoNamespaceLister {
	return k8dynamoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// K8dynamoNamespaceLister helps list and get K8dynamos.
// All objects returned here must be treated as read-only.
type K8dynamoNamespaceLister interface {
	// List lists all K8dynamos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.K8dynamo, err error)
	// Get retrieves the K8dynamo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.K8dynamo, error)
	K8dynamoNamespaceListerExpansion
}

// k8dynamoNamespaceLister implements the K8dynamoNamespaceLister
// interface.
type k8dynamoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all K8dynamos in the indexer for a given namespace.
func (s k8dynamoNamespaceLister) List(selector labels.Selector) (ret []*v1.K8dynamo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.K8dynamo))
	})
	return ret, err
}

// Get retrieves the K8dynamo from the indexer for a given namespace and name.
func (s k8dynamoNamespaceLister) Get(name string) (*v1.K8dynamo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("k8dynamo"), name)
	}
	return obj.(*v1.K8dynamo), nil
}
