// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: api/customer/customer.proto

package customer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetVerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telephone string `protobuf:"bytes,1,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
}

func (x *GetVerifyCodeReq) Reset() {
	*x = GetVerifyCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeReq) ProtoMessage() {}

func (x *GetVerifyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeReq.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeReq) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *GetVerifyCodeReq) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

type GetVerifyCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetVerifyCodeResp) Reset() {
	*x = GetVerifyCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeResp) ProtoMessage() {}

func (x *GetVerifyCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeResp.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeResp) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *GetVerifyCodeResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetVerifyCodeResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetVerifyCodeResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_api_customer_customer_proto protoreflect.FileDescriptor

var file_api_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x32, 0x7f, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x73,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2d, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x3b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_customer_customer_proto_rawDescOnce sync.Once
	file_api_customer_customer_proto_rawDescData = file_api_customer_customer_proto_rawDesc
)

func file_api_customer_customer_proto_rawDescGZIP() []byte {
	file_api_customer_customer_proto_rawDescOnce.Do(func() {
		file_api_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_customer_customer_proto_rawDescData)
	})
	return file_api_customer_customer_proto_rawDescData
}

var file_api_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_customer_customer_proto_goTypes = []interface{}{
	(*GetVerifyCodeReq)(nil),  // 0: api.customer.GetVerifyCodeReq
	(*GetVerifyCodeResp)(nil), // 1: api.customer.GetVerifyCodeResp
}
var file_api_customer_customer_proto_depIdxs = []int32{
	0, // 0: api.customer.Customer.GetVerifyCode:input_type -> api.customer.GetVerifyCodeReq
	1, // 1: api.customer.Customer.GetVerifyCode:output_type -> api.customer.GetVerifyCodeResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_customer_customer_proto_init() }
func file_api_customer_customer_proto_init() {
	if File_api_customer_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_customer_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeReq); i {
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
		file_api_customer_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeResp); i {
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
			RawDescriptor: file_api_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_customer_customer_proto_goTypes,
		DependencyIndexes: file_api_customer_customer_proto_depIdxs,
		MessageInfos:      file_api_customer_customer_proto_msgTypes,
	}.Build()
	File_api_customer_customer_proto = out.File
	file_api_customer_customer_proto_rawDesc = nil
	file_api_customer_customer_proto_goTypes = nil
	file_api_customer_customer_proto_depIdxs = nil
}
