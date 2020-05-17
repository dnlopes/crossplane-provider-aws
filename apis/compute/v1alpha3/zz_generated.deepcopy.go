// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha3

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCluster) DeepCopyInto(out *EKSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCluster.
func (in *EKSCluster) DeepCopy() *EKSCluster {
	if in == nil {
		return nil
	}
	out := new(EKSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterClass) DeepCopyInto(out *EKSClusterClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterClass.
func (in *EKSClusterClass) DeepCopy() *EKSClusterClass {
	if in == nil {
		return nil
	}
	out := new(EKSClusterClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterClassList) DeepCopyInto(out *EKSClusterClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSClusterClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterClassList.
func (in *EKSClusterClassList) DeepCopy() *EKSClusterClassList {
	if in == nil {
		return nil
	}
	out := new(EKSClusterClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterClassSpecTemplate) DeepCopyInto(out *EKSClusterClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.EKSClusterParameters.DeepCopyInto(&out.EKSClusterParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterClassSpecTemplate.
func (in *EKSClusterClassSpecTemplate) DeepCopy() *EKSClusterClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(EKSClusterClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterList) DeepCopyInto(out *EKSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterList.
func (in *EKSClusterList) DeepCopy() *EKSClusterList {
	if in == nil {
		return nil
	}
	out := new(EKSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterParameters) DeepCopyInto(out *EKSClusterParameters) {
	*out = *in
	if in.RoleARNRef != nil {
		in, out := &in.RoleARNRef, &out.RoleARNRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.RoleARNSelector != nil {
		in, out := &in.RoleARNSelector, &out.RoleARNSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1alpha1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1alpha1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.WorkerNodes.DeepCopyInto(&out.WorkerNodes)
	if in.MapRoles != nil {
		in, out := &in.MapRoles, &out.MapRoles
		*out = make([]MapRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MapUsers != nil {
		in, out := &in.MapUsers, &out.MapUsers
		*out = make([]MapUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterParameters.
func (in *EKSClusterParameters) DeepCopy() *EKSClusterParameters {
	if in == nil {
		return nil
	}
	out := new(EKSClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterSpec) DeepCopyInto(out *EKSClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.EKSClusterParameters.DeepCopyInto(&out.EKSClusterParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterSpec.
func (in *EKSClusterSpec) DeepCopy() *EKSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EKSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterStatus) DeepCopyInto(out *EKSClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterStatus.
func (in *EKSClusterStatus) DeepCopy() *EKSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EKSClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapRole) DeepCopyInto(out *MapRole) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapRole.
func (in *MapRole) DeepCopy() *MapRole {
	if in == nil {
		return nil
	}
	out := new(MapRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapUser) DeepCopyInto(out *MapUser) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapUser.
func (in *MapUser) DeepCopy() *MapUser {
	if in == nil {
		return nil
	}
	out := new(MapUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerNodesSpec) DeepCopyInto(out *WorkerNodesSpec) {
	*out = *in
	if in.NodeAutoScalingGroupMinSize != nil {
		in, out := &in.NodeAutoScalingGroupMinSize, &out.NodeAutoScalingGroupMinSize
		*out = new(int)
		**out = **in
	}
	if in.NodeAutoScalingGroupMaxSize != nil {
		in, out := &in.NodeAutoScalingGroupMaxSize, &out.NodeAutoScalingGroupMaxSize
		*out = new(int)
		**out = **in
	}
	if in.NodeVolumeSize != nil {
		in, out := &in.NodeVolumeSize, &out.NodeVolumeSize
		*out = new(int)
		**out = **in
	}
	if in.ClusterControlPlaneSecurityGroupRef != nil {
		in, out := &in.ClusterControlPlaneSecurityGroupRef, &out.ClusterControlPlaneSecurityGroupRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ClusterControlPlaneSecurityGroupSelector != nil {
		in, out := &in.ClusterControlPlaneSecurityGroupSelector, &out.ClusterControlPlaneSecurityGroupSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerNodesSpec.
func (in *WorkerNodesSpec) DeepCopy() *WorkerNodesSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerNodesSpec)
	in.DeepCopyInto(out)
	return out
}
