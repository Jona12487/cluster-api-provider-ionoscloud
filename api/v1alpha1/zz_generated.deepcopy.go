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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudCluster) DeepCopyInto(out *IONOSCloudCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudCluster.
func (in *IONOSCloudCluster) DeepCopy() *IONOSCloudCluster {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterIdentity) DeepCopyInto(out *IONOSCloudClusterIdentity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterIdentity.
func (in *IONOSCloudClusterIdentity) DeepCopy() *IONOSCloudClusterIdentity {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudClusterIdentity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterIdentityList) DeepCopyInto(out *IONOSCloudClusterIdentityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IONOSCloudClusterIdentity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterIdentityList.
func (in *IONOSCloudClusterIdentityList) DeepCopy() *IONOSCloudClusterIdentityList {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterIdentityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudClusterIdentityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterIdentitySpec) DeepCopyInto(out *IONOSCloudClusterIdentitySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterIdentitySpec.
func (in *IONOSCloudClusterIdentitySpec) DeepCopy() *IONOSCloudClusterIdentitySpec {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterIdentitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterIdentityStatus) DeepCopyInto(out *IONOSCloudClusterIdentityStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterIdentityStatus.
func (in *IONOSCloudClusterIdentityStatus) DeepCopy() *IONOSCloudClusterIdentityStatus {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterIdentityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterList) DeepCopyInto(out *IONOSCloudClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IONOSCloudCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterList.
func (in *IONOSCloudClusterList) DeepCopy() *IONOSCloudClusterList {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterSpec) DeepCopyInto(out *IONOSCloudClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if in.Lans != nil {
		in, out := &in.Lans, &out.Lans
		*out = make([]IONOSLanSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LoadBalancer = in.LoadBalancer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterSpec.
func (in *IONOSCloudClusterSpec) DeepCopy() *IONOSCloudClusterSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterStatus) DeepCopyInto(out *IONOSCloudClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterStatus.
func (in *IONOSCloudClusterStatus) DeepCopy() *IONOSCloudClusterStatus {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterTemplate) DeepCopyInto(out *IONOSCloudClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterTemplate.
func (in *IONOSCloudClusterTemplate) DeepCopy() *IONOSCloudClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterTemplateList) DeepCopyInto(out *IONOSCloudClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IONOSCloudClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterTemplateList.
func (in *IONOSCloudClusterTemplateList) DeepCopy() *IONOSCloudClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterTemplateResource) DeepCopyInto(out *IONOSCloudClusterTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterTemplateResource.
func (in *IONOSCloudClusterTemplateResource) DeepCopy() *IONOSCloudClusterTemplateResource {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudClusterTemplateSpec) DeepCopyInto(out *IONOSCloudClusterTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudClusterTemplateSpec.
func (in *IONOSCloudClusterTemplateSpec) DeepCopy() *IONOSCloudClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachine) DeepCopyInto(out *IONOSCloudMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachine.
func (in *IONOSCloudMachine) DeepCopy() *IONOSCloudMachine {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineList) DeepCopyInto(out *IONOSCloudMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IONOSCloudMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineList.
func (in *IONOSCloudMachineList) DeepCopy() *IONOSCloudMachineList {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineSpec) DeepCopyInto(out *IONOSCloudMachineSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(int32)
		**out = **in
	}
	if in.CpuFamily != nil {
		in, out := &in.CpuFamily, &out.CpuFamily
		*out = new(string)
		**out = **in
	}
	if in.Ram != nil {
		in, out := &in.Ram, &out.Ram
		*out = new(int32)
		**out = **in
	}
	in.BootVolume.DeepCopyInto(&out.BootVolume)
	if in.Nics != nil {
		in, out := &in.Nics, &out.Nics
		*out = make([]IONOSNicSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineSpec.
func (in *IONOSCloudMachineSpec) DeepCopy() *IONOSCloudMachineSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineStatus) DeepCopyInto(out *IONOSCloudMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineStatus.
func (in *IONOSCloudMachineStatus) DeepCopy() *IONOSCloudMachineStatus {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineTemplate) DeepCopyInto(out *IONOSCloudMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineTemplate.
func (in *IONOSCloudMachineTemplate) DeepCopy() *IONOSCloudMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineTemplateList) DeepCopyInto(out *IONOSCloudMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IONOSCloudMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineTemplateList.
func (in *IONOSCloudMachineTemplateList) DeepCopy() *IONOSCloudMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IONOSCloudMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineTemplateResource) DeepCopyInto(out *IONOSCloudMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineTemplateResource.
func (in *IONOSCloudMachineTemplateResource) DeepCopy() *IONOSCloudMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSCloudMachineTemplateSpec) DeepCopyInto(out *IONOSCloudMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSCloudMachineTemplateSpec.
func (in *IONOSCloudMachineTemplateSpec) DeepCopy() *IONOSCloudMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSCloudMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSLanRefSpec) DeepCopyInto(out *IONOSLanRefSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSLanRefSpec.
func (in *IONOSLanRefSpec) DeepCopy() *IONOSLanRefSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSLanRefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSLanSpec) DeepCopyInto(out *IONOSLanSpec) {
	*out = *in
	if in.LanID != nil {
		in, out := &in.LanID, &out.LanID
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSLanSpec.
func (in *IONOSLanSpec) DeepCopy() *IONOSLanSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSLanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSLoadBalancerSpec) DeepCopyInto(out *IONOSLoadBalancerSpec) {
	*out = *in
	out.ListenerLanRef = in.ListenerLanRef
	out.TargetLanRef = in.TargetLanRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSLoadBalancerSpec.
func (in *IONOSLoadBalancerSpec) DeepCopy() *IONOSLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSNicSpec) DeepCopyInto(out *IONOSNicSpec) {
	*out = *in
	out.LanRef = in.LanRef
	if in.PrimaryIP != nil {
		in, out := &in.PrimaryIP, &out.PrimaryIP
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSNicSpec.
func (in *IONOSNicSpec) DeepCopy() *IONOSNicSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSNicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IONOSVolumeSpec) DeepCopyInto(out *IONOSVolumeSpec) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IONOSVolumeSpec.
func (in *IONOSVolumeSpec) DeepCopy() *IONOSVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(IONOSVolumeSpec)
	in.DeepCopyInto(out)
	return out
}
