// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8dynamo) DeepCopyInto(out *K8dynamo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8dynamo.
func (in *K8dynamo) DeepCopy() *K8dynamo {
	if in == nil {
		return nil
	}
	out := new(K8dynamo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8dynamo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8dynamoList) DeepCopyInto(out *K8dynamoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K8dynamo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8dynamoList.
func (in *K8dynamoList) DeepCopy() *K8dynamoList {
	if in == nil {
		return nil
	}
	out := new(K8dynamoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8dynamoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8dynamoSpec) DeepCopyInto(out *K8dynamoSpec) {
	*out = *in
	if in.ReadUnit != nil {
		in, out := &in.ReadUnit, &out.ReadUnit
		*out = new(int32)
		**out = **in
	}
	if in.WriteUnit != nil {
		in, out := &in.WriteUnit, &out.WriteUnit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8dynamoSpec.
func (in *K8dynamoSpec) DeepCopy() *K8dynamoSpec {
	if in == nil {
		return nil
	}
	out := new(K8dynamoSpec)
	in.DeepCopyInto(out)
	return out
}
