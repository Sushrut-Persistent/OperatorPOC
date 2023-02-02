//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSEC2) DeepCopyInto(out *SushrutAWSEC2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSEC2.
func (in *SushrutAWSEC2) DeepCopy() *SushrutAWSEC2 {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSEC2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SushrutAWSEC2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSEC2List) DeepCopyInto(out *SushrutAWSEC2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SushrutAWSEC2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSEC2List.
func (in *SushrutAWSEC2List) DeepCopy() *SushrutAWSEC2List {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSEC2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SushrutAWSEC2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSEC2Spec) DeepCopyInto(out *SushrutAWSEC2Spec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSEC2Spec.
func (in *SushrutAWSEC2Spec) DeepCopy() *SushrutAWSEC2Spec {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSEC2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSEC2Status) DeepCopyInto(out *SushrutAWSEC2Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSEC2Status.
func (in *SushrutAWSEC2Status) DeepCopy() *SushrutAWSEC2Status {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSEC2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSManager) DeepCopyInto(out *SushrutAWSManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSManager.
func (in *SushrutAWSManager) DeepCopy() *SushrutAWSManager {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SushrutAWSManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSManagerList) DeepCopyInto(out *SushrutAWSManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SushrutAWSManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSManagerList.
func (in *SushrutAWSManagerList) DeepCopy() *SushrutAWSManagerList {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SushrutAWSManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSManagerSpec) DeepCopyInto(out *SushrutAWSManagerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSManagerSpec.
func (in *SushrutAWSManagerSpec) DeepCopy() *SushrutAWSManagerSpec {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SushrutAWSManagerStatus) DeepCopyInto(out *SushrutAWSManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SushrutAWSManagerStatus.
func (in *SushrutAWSManagerStatus) DeepCopy() *SushrutAWSManagerStatus {
	if in == nil {
		return nil
	}
	out := new(SushrutAWSManagerStatus)
	in.DeepCopyInto(out)
	return out
}