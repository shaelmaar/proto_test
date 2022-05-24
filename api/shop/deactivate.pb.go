// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api/shop/deactivate.proto

package shop

import (
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

type DeactivateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xmlid string `protobuf:"bytes,1,opt,name=xmlid,proto3" json:"xmlid,omitempty"`
}

func (x *DeactivateRequest) Reset() {
	*x = DeactivateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_deactivate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateRequest) ProtoMessage() {}

func (x *DeactivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_deactivate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateRequest.ProtoReflect.Descriptor instead.
func (*DeactivateRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_deactivate_proto_rawDescGZIP(), []int{0}
}

func (x *DeactivateRequest) GetXmlid() string {
	if x != nil {
		return x.Xmlid
	}
	return ""
}

type DeactivateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeactivateResponse) Reset() {
	*x = DeactivateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_deactivate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateResponse) ProtoMessage() {}

func (x *DeactivateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_deactivate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateResponse.ProtoReflect.Descriptor instead.
func (*DeactivateResponse) Descriptor() ([]byte, []int) {
	return file_api_shop_deactivate_proto_rawDescGZIP(), []int{1}
}

func (x *DeactivateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_shop_deactivate_proto protoreflect.FileDescriptor

var file_api_shop_deactivate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x64, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x68, 0x6f,
	0x70, 0x73, 0x22, 0x29, 0x0a, 0x11, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x78, 0x6d, 0x6c, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x78, 0x6d, 0x6c, 0x69, 0x64, 0x22, 0x2e, 0x0a,
	0x12, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x51, 0x0a,
	0x0a, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x44,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x73, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x07, 0x5a, 0x05, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_shop_deactivate_proto_rawDescOnce sync.Once
	file_api_shop_deactivate_proto_rawDescData = file_api_shop_deactivate_proto_rawDesc
)

func file_api_shop_deactivate_proto_rawDescGZIP() []byte {
	file_api_shop_deactivate_proto_rawDescOnce.Do(func() {
		file_api_shop_deactivate_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shop_deactivate_proto_rawDescData)
	})
	return file_api_shop_deactivate_proto_rawDescData
}

var file_api_shop_deactivate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_shop_deactivate_proto_goTypes = []interface{}{
	(*DeactivateRequest)(nil),  // 0: shops.DeactivateRequest
	(*DeactivateResponse)(nil), // 1: shops.DeactivateResponse
}
var file_api_shop_deactivate_proto_depIdxs = []int32{
	0, // 0: shops.Deactivate.Deactivate:input_type -> shops.DeactivateRequest
	1, // 1: shops.Deactivate.Deactivate:output_type -> shops.DeactivateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_shop_deactivate_proto_init() }
func file_api_shop_deactivate_proto_init() {
	if File_api_shop_deactivate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_shop_deactivate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateRequest); i {
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
		file_api_shop_deactivate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateResponse); i {
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
			RawDescriptor: file_api_shop_deactivate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_shop_deactivate_proto_goTypes,
		DependencyIndexes: file_api_shop_deactivate_proto_depIdxs,
		MessageInfos:      file_api_shop_deactivate_proto_msgTypes,
	}.Build()
	File_api_shop_deactivate_proto = out.File
	file_api_shop_deactivate_proto_rawDesc = nil
	file_api_shop_deactivate_proto_goTypes = nil
	file_api_shop_deactivate_proto_depIdxs = nil
}
