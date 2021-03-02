// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: base/base.proto

package base

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type REPLY_STATUS int32

const (
	REPLY_STATUS_UNKNOWN REPLY_STATUS = 0
	REPLY_STATUS_SUCCESS REPLY_STATUS = 1
	REPLY_STATUS_FAILURE REPLY_STATUS = 2
)

// Enum value maps for REPLY_STATUS.
var (
	REPLY_STATUS_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILURE",
	}
	REPLY_STATUS_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILURE": 2,
	}
)

func (x REPLY_STATUS) Enum() *REPLY_STATUS {
	p := new(REPLY_STATUS)
	*p = x
	return p
}

func (x REPLY_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (REPLY_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_base_base_proto_enumTypes[0].Descriptor()
}

func (REPLY_STATUS) Type() protoreflect.EnumType {
	return &file_base_base_proto_enumTypes[0]
}

func (x REPLY_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use REPLY_STATUS.Descriptor instead.
func (REPLY_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{0}
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  REPLY_STATUS `protobuf:"varint,1,opt,name=status,proto3,enum=base.REPLY_STATUS" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetStatus() REPLY_STATUS {
	if x != nil {
		return x.Status
	}
	return REPLY_STATUS_UNKNOWN
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_base_base_proto protoreflect.FileDescriptor

var file_base_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52,
	0x45, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x35, 0x0a,
	0x0c, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x6e, 0x75, 0x76, 0x6a, 0x2f, 0x76, 0x68, 0x6f, 0x6a, 0x5f, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_base_proto_rawDescOnce sync.Once
	file_base_base_proto_rawDescData = file_base_base_proto_rawDesc
)

func file_base_base_proto_rawDescGZIP() []byte {
	file_base_base_proto_rawDescOnce.Do(func() {
		file_base_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_base_proto_rawDescData)
	})
	return file_base_base_proto_rawDescData
}

var file_base_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_base_base_proto_goTypes = []interface{}{
	(REPLY_STATUS)(0),    // 0: base.REPLY_STATUS
	(*BaseResponse)(nil), // 1: base.BaseResponse
}
var file_base_base_proto_depIdxs = []int32{
	0, // 0: base.BaseResponse.status:type_name -> base.REPLY_STATUS
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_base_base_proto_init() }
func file_base_base_proto_init() {
	if File_base_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
			RawDescriptor: file_base_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_base_proto_goTypes,
		DependencyIndexes: file_base_base_proto_depIdxs,
		EnumInfos:         file_base_base_proto_enumTypes,
		MessageInfos:      file_base_base_proto_msgTypes,
	}.Build()
	File_base_base_proto = out.File
	file_base_base_proto_rawDesc = nil
	file_base_base_proto_goTypes = nil
	file_base_base_proto_depIdxs = nil
}
