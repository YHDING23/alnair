// +build !ignore_autogenerated

/*


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
func (in *ElasticHorovodJob) DeepCopyInto(out *ElasticHorovodJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticHorovodJob.
func (in *ElasticHorovodJob) DeepCopy() *ElasticHorovodJob {
	if in == nil {
		return nil
	}
	out := new(ElasticHorovodJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticHorovodJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticHorovodJobList) DeepCopyInto(out *ElasticHorovodJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticHorovodJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticHorovodJobList.
func (in *ElasticHorovodJobList) DeepCopy() *ElasticHorovodJobList {
	if in == nil {
		return nil
	}
	out := new(ElasticHorovodJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticHorovodJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticHorovodJobSpec) DeepCopyInto(out *ElasticHorovodJobSpec) {
	*out = *in
	in.LauncherSpec.DeepCopyInto(&out.LauncherSpec)
	in.WorkersSpec.DeepCopyInto(&out.WorkersSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticHorovodJobSpec.
func (in *ElasticHorovodJobSpec) DeepCopy() *ElasticHorovodJobSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticHorovodJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticHorovodJobStatus) DeepCopyInto(out *ElasticHorovodJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticHorovodJobStatus.
func (in *ElasticHorovodJobStatus) DeepCopy() *ElasticHorovodJobStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticHorovodJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticHorovodWorkersSpec) DeepCopyInto(out *ElasticHorovodWorkersSpec) {
	*out = *in
	if in.TargetReplicas != nil {
		in, out := &in.TargetReplicas, &out.TargetReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticHorovodWorkersSpec.
func (in *ElasticHorovodWorkersSpec) DeepCopy() *ElasticHorovodWorkersSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticHorovodWorkersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorovodJob) DeepCopyInto(out *HorovodJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorovodJob.
func (in *HorovodJob) DeepCopy() *HorovodJob {
	if in == nil {
		return nil
	}
	out := new(HorovodJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorovodJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorovodJobList) DeepCopyInto(out *HorovodJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HorovodJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorovodJobList.
func (in *HorovodJobList) DeepCopy() *HorovodJobList {
	if in == nil {
		return nil
	}
	out := new(HorovodJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorovodJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorovodJobSpec) DeepCopyInto(out *HorovodJobSpec) {
	*out = *in
	in.LauncherSpec.DeepCopyInto(&out.LauncherSpec)
	in.WorkersSpec.DeepCopyInto(&out.WorkersSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorovodJobSpec.
func (in *HorovodJobSpec) DeepCopy() *HorovodJobSpec {
	if in == nil {
		return nil
	}
	out := new(HorovodJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorovodJobStatus) DeepCopyInto(out *HorovodJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorovodJobStatus.
func (in *HorovodJobStatus) DeepCopy() *HorovodJobStatus {
	if in == nil {
		return nil
	}
	out := new(HorovodJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorovodLauncherSpec) DeepCopyInto(out *HorovodLauncherSpec) {
	*out = *in
	if in.PythonCommand != nil {
		in, out := &in.PythonCommand, &out.PythonCommand
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorovodLauncherSpec.
func (in *HorovodLauncherSpec) DeepCopy() *HorovodLauncherSpec {
	if in == nil {
		return nil
	}
	out := new(HorovodLauncherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorovodWorkersSpec) DeepCopyInto(out *HorovodWorkersSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorovodWorkersSpec.
func (in *HorovodWorkersSpec) DeepCopy() *HorovodWorkersSpec {
	if in == nil {
		return nil
	}
	out := new(HorovodWorkersSpec)
	in.DeepCopyInto(out)
	return out
}
