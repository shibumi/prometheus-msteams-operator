//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMSTeamsBridge) DeepCopyInto(out *PrometheusMSTeamsBridge) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMSTeamsBridge.
func (in *PrometheusMSTeamsBridge) DeepCopy() *PrometheusMSTeamsBridge {
	if in == nil {
		return nil
	}
	out := new(PrometheusMSTeamsBridge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusMSTeamsBridge) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMSTeamsBridgeList) DeepCopyInto(out *PrometheusMSTeamsBridgeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusMSTeamsBridge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMSTeamsBridgeList.
func (in *PrometheusMSTeamsBridgeList) DeepCopy() *PrometheusMSTeamsBridgeList {
	if in == nil {
		return nil
	}
	out := new(PrometheusMSTeamsBridgeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusMSTeamsBridgeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMSTeamsBridgeSpec) DeepCopyInto(out *PrometheusMSTeamsBridgeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMSTeamsBridgeSpec.
func (in *PrometheusMSTeamsBridgeSpec) DeepCopy() *PrometheusMSTeamsBridgeSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusMSTeamsBridgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMSTeamsBridgeStatus) DeepCopyInto(out *PrometheusMSTeamsBridgeStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMSTeamsBridgeStatus.
func (in *PrometheusMSTeamsBridgeStatus) DeepCopy() *PrometheusMSTeamsBridgeStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusMSTeamsBridgeStatus)
	in.DeepCopyInto(out)
	return out
}
