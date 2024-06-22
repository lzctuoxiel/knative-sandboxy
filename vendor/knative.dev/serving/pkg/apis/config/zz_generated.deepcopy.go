// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

package config

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Defaults) DeepCopyInto(out *Defaults) {
	*out = *in
	if in.RevisionCPURequest != nil {
		in, out := &in.RevisionCPURequest, &out.RevisionCPURequest
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.RevisionCPULimit != nil {
		in, out := &in.RevisionCPULimit, &out.RevisionCPULimit
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.RevisionMemoryRequest != nil {
		in, out := &in.RevisionMemoryRequest, &out.RevisionMemoryRequest
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.RevisionMemoryLimit != nil {
		in, out := &in.RevisionMemoryLimit, &out.RevisionMemoryLimit
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Defaults.
func (in *Defaults) DeepCopy() *Defaults {
	if in == nil {
		return nil
	}
	out := new(Defaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}
