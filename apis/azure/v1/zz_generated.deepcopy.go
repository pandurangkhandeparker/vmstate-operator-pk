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
func (in *PandurangAzureAVM) DeepCopyInto(out *PandurangAzureAVM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PandurangAzureAVM.
func (in *PandurangAzureAVM) DeepCopy() *PandurangAzureAVM {
	if in == nil {
		return nil
	}
	out := new(PandurangAzureAVM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PandurangAzureAVM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PandurangAzureAVMList) DeepCopyInto(out *PandurangAzureAVMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PandurangAzureAVM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PandurangAzureAVMList.
func (in *PandurangAzureAVMList) DeepCopy() *PandurangAzureAVMList {
	if in == nil {
		return nil
	}
	out := new(PandurangAzureAVMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PandurangAzureAVMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PandurangAzureAVMSpec) DeepCopyInto(out *PandurangAzureAVMSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PandurangAzureAVMSpec.
func (in *PandurangAzureAVMSpec) DeepCopy() *PandurangAzureAVMSpec {
	if in == nil {
		return nil
	}
	out := new(PandurangAzureAVMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PandurangAzureAVMStatus) DeepCopyInto(out *PandurangAzureAVMStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PandurangAzureAVMStatus.
func (in *PandurangAzureAVMStatus) DeepCopy() *PandurangAzureAVMStatus {
	if in == nil {
		return nil
	}
	out := new(PandurangAzureAVMStatus)
	in.DeepCopyInto(out)
	return out
}
