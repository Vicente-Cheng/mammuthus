//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSAL) DeepCopyInto(out *FSAL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSAL.
func (in *FSAL) DeepCopy() *FSAL {
	if in == nil {
		return nil
	}
	out := new(FSAL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExport) DeepCopyInto(out *NFSExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExport.
func (in *NFSExport) DeepCopy() *NFSExport {
	if in == nil {
		return nil
	}
	out := new(NFSExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFSExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportList) DeepCopyInto(out *NFSExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFSExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportList.
func (in *NFSExportList) DeepCopy() *NFSExportList {
	if in == nil {
		return nil
	}
	out := new(NFSExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFSExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportSpec) DeepCopyInto(out *NFSExportSpec) {
	*out = *in
	if in.FSAL != nil {
		in, out := &in.FSAL, &out.FSAL
		*out = new(FSAL)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportSpec.
func (in *NFSExportSpec) DeepCopy() *NFSExportSpec {
	if in == nil {
		return nil
	}
	out := new(NFSExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportStatus) DeepCopyInto(out *NFSExportStatus) {
	*out = *in
	if in.FSAL != nil {
		in, out := &in.FSAL, &out.FSAL
		*out = new(FSAL)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportStatus.
func (in *NFSExportStatus) DeepCopy() *NFSExportStatus {
	if in == nil {
		return nil
	}
	out := new(NFSExportStatus)
	in.DeepCopyInto(out)
	return out
}
