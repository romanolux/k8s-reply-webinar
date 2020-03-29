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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	k8dynamov1 "k8s-webinar/pkg/apis/k8dynamo/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeK8dynamos implements K8dynamoInterface
type FakeK8dynamos struct {
	Fake *FakeStormV1
	ns   string
}

var k8dynamosResource = schema.GroupVersionResource{Group: "storm.com", Version: "v1", Resource: "k8dynamos"}

var k8dynamosKind = schema.GroupVersionKind{Group: "storm.com", Version: "v1", Kind: "K8dynamo"}

// Get takes name of the k8dynamo, and returns the corresponding k8dynamo object, and an error if there is any.
func (c *FakeK8dynamos) Get(ctx context.Context, name string, options v1.GetOptions) (result *k8dynamov1.K8dynamo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(k8dynamosResource, c.ns, name), &k8dynamov1.K8dynamo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8dynamov1.K8dynamo), err
}

// List takes label and field selectors, and returns the list of K8dynamos that match those selectors.
func (c *FakeK8dynamos) List(ctx context.Context, opts v1.ListOptions) (result *k8dynamov1.K8dynamoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(k8dynamosResource, k8dynamosKind, c.ns, opts), &k8dynamov1.K8dynamoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &k8dynamov1.K8dynamoList{ListMeta: obj.(*k8dynamov1.K8dynamoList).ListMeta}
	for _, item := range obj.(*k8dynamov1.K8dynamoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested k8dynamos.
func (c *FakeK8dynamos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(k8dynamosResource, c.ns, opts))

}

// Create takes the representation of a k8dynamo and creates it.  Returns the server's representation of the k8dynamo, and an error, if there is any.
func (c *FakeK8dynamos) Create(ctx context.Context, k8dynamo *k8dynamov1.K8dynamo, opts v1.CreateOptions) (result *k8dynamov1.K8dynamo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(k8dynamosResource, c.ns, k8dynamo), &k8dynamov1.K8dynamo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8dynamov1.K8dynamo), err
}

// Update takes the representation of a k8dynamo and updates it. Returns the server's representation of the k8dynamo, and an error, if there is any.
func (c *FakeK8dynamos) Update(ctx context.Context, k8dynamo *k8dynamov1.K8dynamo, opts v1.UpdateOptions) (result *k8dynamov1.K8dynamo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(k8dynamosResource, c.ns, k8dynamo), &k8dynamov1.K8dynamo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8dynamov1.K8dynamo), err
}

// Delete takes name of the k8dynamo and deletes it. Returns an error if one occurs.
func (c *FakeK8dynamos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(k8dynamosResource, c.ns, name), &k8dynamov1.K8dynamo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeK8dynamos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(k8dynamosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &k8dynamov1.K8dynamoList{})
	return err
}

// Patch applies the patch and returns the patched k8dynamo.
func (c *FakeK8dynamos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *k8dynamov1.K8dynamo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(k8dynamosResource, c.ns, name, pt, data, subresources...), &k8dynamov1.K8dynamo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8dynamov1.K8dynamo), err
}