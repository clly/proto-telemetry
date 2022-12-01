// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: echo/v1/echo.proto

package echov1

import (
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg     string            `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Num32   int32             `protobuf:"varint,2,opt,name=num32,proto3" json:"num32,omitempty"`
	Unum32  uint32            `protobuf:"varint,3,opt,name=unum32,proto3" json:"unum32,omitempty"`
	Num64   int64             `protobuf:"varint,4,opt,name=num64,proto3" json:"num64,omitempty"`
	Details *MessageDetails   `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	Meta    map[string]string `protobuf:"bytes,6,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *EchoRequest) GetNum32() int32 {
	if x != nil {
		return x.Num32
	}
	return 0
}

func (x *EchoRequest) GetUnum32() uint32 {
	if x != nil {
		return x.Unum32
	}
	return 0
}

func (x *EchoRequest) GetNum64() int64 {
	if x != nil {
		return x.Num64
	}
	return 0
}

func (x *EchoRequest) GetDetails() *MessageDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *EchoRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type MessageDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details string `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *MessageDetails) Reset() {
	*x = MessageDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDetails) ProtoMessage() {}

func (x *MessageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDetails.ProtoReflect.Descriptor instead.
func (*MessageDetails) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{1}
}

func (x *MessageDetails) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string             `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Now *datetime.DateTime `protobuf:"bytes,2,opt,name=now,proto3" json:"now,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{2}
}

func (x *EchoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *EchoResponse) GetNow() *datetime.DateTime {
	if x != nil {
		return x.Now
	}
	return nil
}

type EchoRequest_Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EchoRequest_Envelope) Reset() {
	*x = EchoRequest_Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest_Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest_Envelope) ProtoMessage() {}

func (x *EchoRequest_Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest_Envelope.ProtoReflect.Descriptor instead.
func (*EchoRequest_Envelope) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EchoRequest_Envelope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_echo_v1_echo_proto protoreflect.FileDescriptor

var file_echo_v1_echo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x75, 0x6d, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x75, 0x6d, 0x33,
	0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x75, 0x6d,
	0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x12,
	0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1e, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2a, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x49, 0x0a, 0x0c, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x27, 0x0a,
	0x03, 0x6e, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x03, 0x6e, 0x6f, 0x77, 0x32, 0x44, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x14, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9a, 0x01, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x45, 0x63,
	0x68, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6c, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x63,
	0x68, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x63, 0x68, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45,
	0x58, 0x58, 0xaa, 0x02, 0x07, 0x45, 0x63, 0x68, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x45,
	0x63, 0x68, 0x6f, 0x5f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x45, 0x63, 0x68, 0x6f, 0x5f, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x45, 0x63, 0x68, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_echo_v1_echo_proto_rawDescOnce sync.Once
	file_echo_v1_echo_proto_rawDescData = file_echo_v1_echo_proto_rawDesc
)

func file_echo_v1_echo_proto_rawDescGZIP() []byte {
	file_echo_v1_echo_proto_rawDescOnce.Do(func() {
		file_echo_v1_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_v1_echo_proto_rawDescData)
	})
	return file_echo_v1_echo_proto_rawDescData
}

var file_echo_v1_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_echo_v1_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),          // 0: echo.v1.EchoRequest
	(*MessageDetails)(nil),       // 1: echo.v1.MessageDetails
	(*EchoResponse)(nil),         // 2: echo.v1.EchoResponse
	(*EchoRequest_Envelope)(nil), // 3: echo.v1.EchoRequest.Envelope
	nil,                          // 4: echo.v1.EchoRequest.MetaEntry
	(*datetime.DateTime)(nil),    // 5: google.type.DateTime
}
var file_echo_v1_echo_proto_depIdxs = []int32{
	1, // 0: echo.v1.EchoRequest.details:type_name -> echo.v1.MessageDetails
	4, // 1: echo.v1.EchoRequest.meta:type_name -> echo.v1.EchoRequest.MetaEntry
	5, // 2: echo.v1.EchoResponse.now:type_name -> google.type.DateTime
	0, // 3: echo.v1.EchoService.Echo:input_type -> echo.v1.EchoRequest
	2, // 4: echo.v1.EchoService.Echo:output_type -> echo.v1.EchoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_echo_v1_echo_proto_init() }
func file_echo_v1_echo_proto_init() {
	if File_echo_v1_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_v1_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_echo_v1_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDetails); i {
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
		file_echo_v1_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
		file_echo_v1_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest_Envelope); i {
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
			RawDescriptor: file_echo_v1_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_v1_echo_proto_goTypes,
		DependencyIndexes: file_echo_v1_echo_proto_depIdxs,
		MessageInfos:      file_echo_v1_echo_proto_msgTypes,
	}.Build()
	File_echo_v1_echo_proto = out.File
	file_echo_v1_echo_proto_rawDesc = nil
	file_echo_v1_echo_proto_goTypes = nil
	file_echo_v1_echo_proto_depIdxs = nil
}
