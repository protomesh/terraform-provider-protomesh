// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/types/v1/typed_resource.proto

package typesv1

import (
	_ "github.com/protomesh/protoc-gen-terraform/proto/terraform"
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

type TypedResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Spec:
	//
	//	*TypedResource_NetworkingNode
	//	*TypedResource_Trigger
	Spec isTypedResource_Spec `protobuf_oneof:"spec"`
}

func (x *TypedResource) Reset() {
	*x = TypedResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_types_v1_typed_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedResource) ProtoMessage() {}

func (x *TypedResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_types_v1_typed_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedResource.ProtoReflect.Descriptor instead.
func (*TypedResource) Descriptor() ([]byte, []int) {
	return file_api_types_v1_typed_resource_proto_rawDescGZIP(), []int{0}
}

func (x *TypedResource) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *TypedResource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TypedResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *TypedResource) GetSpec() isTypedResource_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (x *TypedResource) GetNetworkingNode() *NetworkingNode {
	if x, ok := x.GetSpec().(*TypedResource_NetworkingNode); ok {
		return x.NetworkingNode
	}
	return nil
}

func (x *TypedResource) GetTrigger() *Trigger {
	if x, ok := x.GetSpec().(*TypedResource_Trigger); ok {
		return x.Trigger
	}
	return nil
}

type isTypedResource_Spec interface {
	isTypedResource_Spec()
}

type TypedResource_NetworkingNode struct {
	NetworkingNode *NetworkingNode `protobuf:"bytes,4,opt,name=networking_node,json=networkingNode,proto3,oneof"`
}

type TypedResource_Trigger struct {
	Trigger *Trigger `protobuf:"bytes,5,opt,name=trigger,proto3,oneof"`
}

func (*TypedResource_NetworkingNode) isTypedResource_Spec() {}

func (*TypedResource_Trigger) isTypedResource_Spec() {}

var File_api_types_v1_typed_resource_proto protoreflect.FileDescriptor

var file_api_types_v1_typed_resource_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe7, 0x01, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65,
	0x48, 0x00, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x3a, 0x04, 0xba, 0xb9, 0x02,
	0x00, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x42, 0xe4, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x54, 0x79, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x54, 0x58, 0xaa, 0x02,
	0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x68, 0x5c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x6d, 0x65, 0x73, 0x68, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_types_v1_typed_resource_proto_rawDescOnce sync.Once
	file_api_types_v1_typed_resource_proto_rawDescData = file_api_types_v1_typed_resource_proto_rawDesc
)

func file_api_types_v1_typed_resource_proto_rawDescGZIP() []byte {
	file_api_types_v1_typed_resource_proto_rawDescOnce.Do(func() {
		file_api_types_v1_typed_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_types_v1_typed_resource_proto_rawDescData)
	})
	return file_api_types_v1_typed_resource_proto_rawDescData
}

var file_api_types_v1_typed_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_types_v1_typed_resource_proto_goTypes = []interface{}{
	(*TypedResource)(nil),  // 0: protomesh.types.v1.TypedResource
	(*NetworkingNode)(nil), // 1: protomesh.types.v1.NetworkingNode
	(*Trigger)(nil),        // 2: protomesh.types.v1.Trigger
}
var file_api_types_v1_typed_resource_proto_depIdxs = []int32{
	1, // 0: protomesh.types.v1.TypedResource.networking_node:type_name -> protomesh.types.v1.NetworkingNode
	2, // 1: protomesh.types.v1.TypedResource.trigger:type_name -> protomesh.types.v1.Trigger
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_types_v1_typed_resource_proto_init() }
func file_api_types_v1_typed_resource_proto_init() {
	if File_api_types_v1_typed_resource_proto != nil {
		return
	}
	file_api_types_v1_networking_proto_init()
	file_api_types_v1_process_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_types_v1_typed_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypedResource); i {
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
	file_api_types_v1_typed_resource_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TypedResource_NetworkingNode)(nil),
		(*TypedResource_Trigger)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_types_v1_typed_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_types_v1_typed_resource_proto_goTypes,
		DependencyIndexes: file_api_types_v1_typed_resource_proto_depIdxs,
		MessageInfos:      file_api_types_v1_typed_resource_proto_msgTypes,
	}.Build()
	File_api_types_v1_typed_resource_proto = out.File
	file_api_types_v1_typed_resource_proto_rawDesc = nil
	file_api_types_v1_typed_resource_proto_goTypes = nil
	file_api_types_v1_typed_resource_proto_depIdxs = nil
}