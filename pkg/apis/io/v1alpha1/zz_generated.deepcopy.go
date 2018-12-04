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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FreeForm) DeepCopyInto(out *FreeForm) {
	*out = *in
	if in.json != nil {
		in, out := &in.json, &out.json
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FreeForm.
func (in *FreeForm) DeepCopy() *FreeForm {
	if in == nil {
		return nil
	}
	out := new(FreeForm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jaeger) DeepCopyInto(out *Jaeger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jaeger.
func (in *Jaeger) DeepCopy() *Jaeger {
	if in == nil {
		return nil
	}
	out := new(Jaeger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jaeger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerAgentSpec) DeepCopyInto(out *JaegerAgentSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerAgentSpec.
func (in *JaegerAgentSpec) DeepCopy() *JaegerAgentSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerAllInOneSpec) DeepCopyInto(out *JaegerAllInOneSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerAllInOneSpec.
func (in *JaegerAllInOneSpec) DeepCopy() *JaegerAllInOneSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerAllInOneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCassandraCreateSchemaSpec) DeepCopyInto(out *JaegerCassandraCreateSchemaSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCassandraCreateSchemaSpec.
func (in *JaegerCassandraCreateSchemaSpec) DeepCopy() *JaegerCassandraCreateSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCassandraCreateSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCollectorSpec) DeepCopyInto(out *JaegerCollectorSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCollectorSpec.
func (in *JaegerCollectorSpec) DeepCopy() *JaegerCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCommonSpec) DeepCopyInto(out *JaegerCommonSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCommonSpec.
func (in *JaegerCommonSpec) DeepCopy() *JaegerCommonSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerDependenciesSpec) DeepCopyInto(out *JaegerDependenciesSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerDependenciesSpec.
func (in *JaegerDependenciesSpec) DeepCopy() *JaegerDependenciesSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerDependenciesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerIngressSpec) DeepCopyInto(out *JaegerIngressSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerIngressSpec.
func (in *JaegerIngressSpec) DeepCopy() *JaegerIngressSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerList) DeepCopyInto(out *JaegerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jaeger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerList.
func (in *JaegerList) DeepCopy() *JaegerList {
	if in == nil {
		return nil
	}
	out := new(JaegerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JaegerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerQuerySpec) DeepCopyInto(out *JaegerQuerySpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerQuerySpec.
func (in *JaegerQuerySpec) DeepCopy() *JaegerQuerySpec {
	if in == nil {
		return nil
	}
	out := new(JaegerQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerSamplingSpec) DeepCopyInto(out *JaegerSamplingSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerSamplingSpec.
func (in *JaegerSamplingSpec) DeepCopy() *JaegerSamplingSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerSamplingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerSpec) DeepCopyInto(out *JaegerSpec) {
	*out = *in
	in.AllInOne.DeepCopyInto(&out.AllInOne)
	in.Query.DeepCopyInto(&out.Query)
	in.Collector.DeepCopyInto(&out.Collector)
	in.Agent.DeepCopyInto(&out.Agent)
	in.UI.DeepCopyInto(&out.UI)
	in.Sampling.DeepCopyInto(&out.Sampling)
	in.Storage.DeepCopyInto(&out.Storage)
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerSpec.
func (in *JaegerSpec) DeepCopy() *JaegerSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerStatus) DeepCopyInto(out *JaegerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerStatus.
func (in *JaegerStatus) DeepCopy() *JaegerStatus {
	if in == nil {
		return nil
	}
	out := new(JaegerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerStorageSpec) DeepCopyInto(out *JaegerStorageSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.CassandraCreateSchema.DeepCopyInto(&out.CassandraCreateSchema)
	out.SparkDependencies = in.SparkDependencies
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerStorageSpec.
func (in *JaegerStorageSpec) DeepCopy() *JaegerStorageSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerUISpec) DeepCopyInto(out *JaegerUISpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerUISpec.
func (in *JaegerUISpec) DeepCopy() *JaegerUISpec {
	if in == nil {
		return nil
	}
	out := new(JaegerUISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Options) DeepCopyInto(out *Options) {
	*out = *in
	if in.opts != nil {
		in, out := &in.opts, &out.opts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}
