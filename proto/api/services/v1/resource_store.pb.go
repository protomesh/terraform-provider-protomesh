// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/services/v1/resource_store.proto

package servicesv1

import (
	v1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PutResourceRequest object to put resources
// in the store (used for creating and updating resources).
type PutResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource to create/update.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *PutResourceRequest) Reset() {
	*x = PutResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResourceRequest) ProtoMessage() {}

func (x *PutResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResourceRequest.ProtoReflect.Descriptor instead.
func (*PutResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{0}
}

func (x *PutResourceRequest) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// PutResourceResponse returns the current resource version
// after the put operation.
type PutResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *v1.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PutResourceResponse) Reset() {
	*x = PutResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResourceResponse) ProtoMessage() {}

func (x *PutResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResourceResponse.ProtoReflect.Descriptor instead.
func (*PutResourceResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{1}
}

func (x *PutResourceResponse) GetVersion() *v1.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

// DropResourcesRequest drops multiple resources in the namespace.
type DropResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource IDs to drop.
	ResourceIds []string `protobuf:"bytes,1,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	// Namespace of the resources to drop (each resource is unique within its namespace).
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *DropResourcesRequest) Reset() {
	*x = DropResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropResourcesRequest) ProtoMessage() {}

func (x *DropResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropResourcesRequest.ProtoReflect.Descriptor instead.
func (*DropResourcesRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{2}
}

func (x *DropResourcesRequest) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *DropResourcesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// DropResourcesResponse response to drop resources operation.
type DropResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DropResourcesResponse) Reset() {
	*x = DropResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropResourcesResponse) ProtoMessage() {}

func (x *DropResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropResourcesResponse.ProtoReflect.Descriptor instead.
func (*DropResourcesResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{3}
}

// GetResourceRequest request to return just one resource.
type GetResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace of the specified resource.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Resource ID to bring.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetResourceRequest) Reset() {
	*x = GetResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRequest) ProtoMessage() {}

func (x *GetResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetResourceRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// GetResourceResponse response with the found resource.
// Returns gRPC not found status code when the resource
// is not present in the current store state.
type GetResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetResourceResponse) Reset() {
	*x = GetResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceResponse) ProtoMessage() {}

func (x *GetResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceResponse.ProtoReflect.Descriptor instead.
func (*GetResourceResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{5}
}

func (x *GetResourceResponse) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// WatchResourcesRequest stream resource state changes:
//
// When a resource is created or updated, it's returned
//
//	in the UpdatedResources field in the response.
//
// When a resource is dropped, it's returned
//
//	in the DroppedResources field in the response.
type WatchResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace to watch.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *WatchResourcesRequest) Reset() {
	*x = WatchResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResourcesRequest) ProtoMessage() {}

func (x *WatchResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResourcesRequest.ProtoReflect.Descriptor instead.
func (*WatchResourcesRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{6}
}

func (x *WatchResourcesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// WatchResourcesResponse response streamed from watch method call.
type WatchResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Created/Updated resources are returned in this field.
	UpdatedResources []*v1.Resource `protobuf:"bytes,1,rep,name=updated_resources,json=updatedResources,proto3" json:"updated_resources,omitempty"`
	// Dropped resources are returned in this field.
	DroppedResources []*v1.Resource `protobuf:"bytes,2,rep,name=dropped_resources,json=droppedResources,proto3" json:"dropped_resources,omitempty"`
	// Indicates an end of synchronization iteration (will wait next interval).
	EndOfList bool `protobuf:"varint,3,opt,name=end_of_list,json=endOfList,proto3" json:"end_of_list,omitempty"`
}

func (x *WatchResourcesResponse) Reset() {
	*x = WatchResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_resource_store_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResourcesResponse) ProtoMessage() {}

func (x *WatchResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_resource_store_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResourcesResponse.ProtoReflect.Descriptor instead.
func (*WatchResourcesResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_resource_store_proto_rawDescGZIP(), []int{7}
}

func (x *WatchResourcesResponse) GetUpdatedResources() []*v1.Resource {
	if x != nil {
		return x.UpdatedResources
	}
	return nil
}

func (x *WatchResourcesResponse) GetDroppedResources() []*v1.Resource {
	if x != nil {
		return x.DroppedResources
	}
	return nil
}

func (x *WatchResourcesResponse) GetEndOfList() bool {
	if x != nil {
		return x.EndOfList
	}
	return false
}

var File_api_services_v1_resource_store_proto protoreflect.FileDescriptor

var file_api_services_v1_resource_store_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x12, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x14, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x17, 0x0a,
	0x15, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x35, 0x0a, 0x15,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x11, 0x64, 0x72, 0x6f,
	0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0x9e, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x5e, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x2b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x05, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0xf9, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x6d,
	0x65, 0x73, 0x68, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_v1_resource_store_proto_rawDescOnce sync.Once
	file_api_services_v1_resource_store_proto_rawDescData = file_api_services_v1_resource_store_proto_rawDesc
)

func file_api_services_v1_resource_store_proto_rawDescGZIP() []byte {
	file_api_services_v1_resource_store_proto_rawDescOnce.Do(func() {
		file_api_services_v1_resource_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_v1_resource_store_proto_rawDescData)
	})
	return file_api_services_v1_resource_store_proto_rawDescData
}

var file_api_services_v1_resource_store_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_services_v1_resource_store_proto_goTypes = []interface{}{
	(*PutResourceRequest)(nil),     // 0: protomesh.services.v1.PutResourceRequest
	(*PutResourceResponse)(nil),    // 1: protomesh.services.v1.PutResourceResponse
	(*DropResourcesRequest)(nil),   // 2: protomesh.services.v1.DropResourcesRequest
	(*DropResourcesResponse)(nil),  // 3: protomesh.services.v1.DropResourcesResponse
	(*GetResourceRequest)(nil),     // 4: protomesh.services.v1.GetResourceRequest
	(*GetResourceResponse)(nil),    // 5: protomesh.services.v1.GetResourceResponse
	(*WatchResourcesRequest)(nil),  // 6: protomesh.services.v1.WatchResourcesRequest
	(*WatchResourcesResponse)(nil), // 7: protomesh.services.v1.WatchResourcesResponse
	(*v1.Resource)(nil),            // 8: protomesh.types.v1.Resource
	(*v1.Version)(nil),             // 9: protomesh.types.v1.Version
}
var file_api_services_v1_resource_store_proto_depIdxs = []int32{
	8, // 0: protomesh.services.v1.PutResourceRequest.resource:type_name -> protomesh.types.v1.Resource
	9, // 1: protomesh.services.v1.PutResourceResponse.version:type_name -> protomesh.types.v1.Version
	8, // 2: protomesh.services.v1.GetResourceResponse.resource:type_name -> protomesh.types.v1.Resource
	8, // 3: protomesh.services.v1.WatchResourcesResponse.updated_resources:type_name -> protomesh.types.v1.Resource
	8, // 4: protomesh.services.v1.WatchResourcesResponse.dropped_resources:type_name -> protomesh.types.v1.Resource
	0, // 5: protomesh.services.v1.ResourceStore.Put:input_type -> protomesh.services.v1.PutResourceRequest
	2, // 6: protomesh.services.v1.ResourceStore.Drop:input_type -> protomesh.services.v1.DropResourcesRequest
	4, // 7: protomesh.services.v1.ResourceStore.Get:input_type -> protomesh.services.v1.GetResourceRequest
	6, // 8: protomesh.services.v1.ResourceStore.Watch:input_type -> protomesh.services.v1.WatchResourcesRequest
	1, // 9: protomesh.services.v1.ResourceStore.Put:output_type -> protomesh.services.v1.PutResourceResponse
	3, // 10: protomesh.services.v1.ResourceStore.Drop:output_type -> protomesh.services.v1.DropResourcesResponse
	5, // 11: protomesh.services.v1.ResourceStore.Get:output_type -> protomesh.services.v1.GetResourceResponse
	7, // 12: protomesh.services.v1.ResourceStore.Watch:output_type -> protomesh.services.v1.WatchResourcesResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_services_v1_resource_store_proto_init() }
func file_api_services_v1_resource_store_proto_init() {
	if File_api_services_v1_resource_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_v1_resource_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResourceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResourceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropResourcesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropResourcesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResourcesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_services_v1_resource_store_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResourcesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_v1_resource_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_v1_resource_store_proto_goTypes,
		DependencyIndexes: file_api_services_v1_resource_store_proto_depIdxs,
		MessageInfos:      file_api_services_v1_resource_store_proto_msgTypes,
	}.Build()
	File_api_services_v1_resource_store_proto = out.File
	file_api_services_v1_resource_store_proto_rawDesc = nil
	file_api_services_v1_resource_store_proto_goTypes = nil
	file_api_services_v1_resource_store_proto_depIdxs = nil
}
