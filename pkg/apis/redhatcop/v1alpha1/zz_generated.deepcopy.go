// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clair) DeepCopyInto(out *Clair) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clair.
func (in *Clair) DeepCopy() *Clair {
	if in == nil {
		return nil
	}
	out := new(Clair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRegistryBackendSource) DeepCopyInto(out *LocalRegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRegistryBackendSource.
func (in *LocalRegistryBackendSource) DeepCopy() *LocalRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(LocalRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Quay) DeepCopyInto(out *Quay) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryBackends != nil {
		in, out := &in.RegistryBackends, &out.RegistryBackends
		*out = make([]RegistryBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RegistryStorage != nil {
		in, out := &in.RegistryStorage, &out.RegistryStorage
		*out = new(RegistryStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Quay.
func (in *Quay) DeepCopy() *Quay {
	if in == nil {
		return nil
	}
	out := new(Quay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystem) DeepCopyInto(out *QuayEcosystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystem.
func (in *QuayEcosystem) DeepCopy() *QuayEcosystem {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuayEcosystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemCondition) DeepCopyInto(out *QuayEcosystemCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemCondition.
func (in *QuayEcosystemCondition) DeepCopy() *QuayEcosystemCondition {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemList) DeepCopyInto(out *QuayEcosystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuayEcosystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemList.
func (in *QuayEcosystemList) DeepCopy() *QuayEcosystemList {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuayEcosystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemSpec) DeepCopyInto(out *QuayEcosystemSpec) {
	*out = *in
	if in.Quay != nil {
		in, out := &in.Quay, &out.Quay
		*out = new(Quay)
		(*in).DeepCopyInto(*out)
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = new(Redis)
		(*in).DeepCopyInto(*out)
	}
	if in.Clair != nil {
		in, out := &in.Clair, &out.Clair
		*out = new(Clair)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemSpec.
func (in *QuayEcosystemSpec) DeepCopy() *QuayEcosystemSpec {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemStatus) DeepCopyInto(out *QuayEcosystemStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]QuayEcosystemCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemStatus.
func (in *QuayEcosystemStatus) DeepCopy() *QuayEcosystemStatus {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryBackend) DeepCopyInto(out *RegistryBackend) {
	*out = *in
	in.RegistryBackendSource.DeepCopyInto(&out.RegistryBackendSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryBackend.
func (in *RegistryBackend) DeepCopy() *RegistryBackend {
	if in == nil {
		return nil
	}
	out := new(RegistryBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryBackendSource) DeepCopyInto(out *RegistryBackendSource) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalRegistryBackendSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryBackendSource.
func (in *RegistryBackendSource) DeepCopy() *RegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(RegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStorage) DeepCopyInto(out *RegistryStorage) {
	*out = *in
	if in.PersistentVolumeAccessModes != nil {
		in, out := &in.PersistentVolumeAccessModes, &out.PersistentVolumeAccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStorage.
func (in *RegistryStorage) DeepCopy() *RegistryStorage {
	if in == nil {
		return nil
	}
	out := new(RegistryStorage)
	in.DeepCopyInto(out)
	return out
}