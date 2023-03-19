// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: test/v1/test.proto

package testv1

import (
	_ "github.com/clly/proto-telemetry/proto/telemetry/options/v1"
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

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *StringMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Int32Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num32 int32 `protobuf:"varint,1,opt,name=num32,proto3" json:"num32,omitempty"`
}

func (x *Int32Message) Reset() {
	*x = Int32Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Message) ProtoMessage() {}

func (x *Int32Message) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Message.ProtoReflect.Descriptor instead.
func (*Int32Message) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *Int32Message) GetNum32() int32 {
	if x != nil {
		return x.Num32
	}
	return 0
}

type Uint32Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unum32 uint32 `protobuf:"varint,1,opt,name=unum32,proto3" json:"unum32,omitempty"`
}

func (x *Uint32Message) Reset() {
	*x = Uint32Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint32Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint32Message) ProtoMessage() {}

func (x *Uint32Message) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint32Message.ProtoReflect.Descriptor instead.
func (*Uint32Message) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{2}
}

func (x *Uint32Message) GetUnum32() uint32 {
	if x != nil {
		return x.Unum32
	}
	return 0
}

type Int64Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num64 int64 `protobuf:"varint,1,opt,name=num64,proto3" json:"num64,omitempty"`
}

func (x *Int64Message) Reset() {
	*x = Int64Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Message) ProtoMessage() {}

func (x *Int64Message) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Message.ProtoReflect.Descriptor instead.
func (*Int64Message) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{3}
}

func (x *Int64Message) GetNum64() int64 {
	if x != nil {
		return x.Num64
	}
	return 0
}

type SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details *MessageDetails `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *SubMessage) Reset() {
	*x = SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessage) ProtoMessage() {}

func (x *SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessage.ProtoReflect.Descriptor instead.
func (*SubMessage) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{4}
}

func (x *SubMessage) GetDetails() *MessageDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

type MapMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapMessage) Reset() {
	*x = MapMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMessage) ProtoMessage() {}

func (x *MapMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMessage.ProtoReflect.Descriptor instead.
func (*MapMessage) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{5}
}

func (x *MapMessage) GetMeta() map[string]string {
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
		mi := &file_test_v1_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDetails) ProtoMessage() {}

func (x *MessageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[6]
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
	return file_test_v1_test_proto_rawDescGZIP(), []int{6}
}

func (x *MessageDetails) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type ExcludeField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NonMasked string `protobuf:"bytes,2,opt,name=non_masked,json=nonMasked,proto3" json:"non_masked,omitempty"`
}

func (x *ExcludeField) Reset() {
	*x = ExcludeField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExcludeField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludeField) ProtoMessage() {}

func (x *ExcludeField) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludeField.ProtoReflect.Descriptor instead.
func (*ExcludeField) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{7}
}

func (x *ExcludeField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExcludeField) GetNonMasked() string {
	if x != nil {
		return x.NonMasked
	}
	return ""
}

type ExcludeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string             `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Now *datetime.DateTime `protobuf:"bytes,2,opt,name=now,proto3" json:"now,omitempty"`
}

func (x *ExcludeMessage) Reset() {
	*x = ExcludeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExcludeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludeMessage) ProtoMessage() {}

func (x *ExcludeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludeMessage.ProtoReflect.Descriptor instead.
func (*ExcludeMessage) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{8}
}

func (x *ExcludeMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ExcludeMessage) GetNow() *datetime.DateTime {
	if x != nil {
		return x.Now
	}
	return nil
}

type RenameMessagePrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RenameMessagePrefix) Reset() {
	*x = RenameMessagePrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameMessagePrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameMessagePrefix) ProtoMessage() {}

func (x *RenameMessagePrefix) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameMessagePrefix.ProtoReflect.Descriptor instead.
func (*RenameMessagePrefix) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{9}
}

func (x *RenameMessagePrefix) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type NameField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NameField) Reset() {
	*x = NameField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameField) ProtoMessage() {}

func (x *NameField) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameField.ProtoReflect.Descriptor instead.
func (*NameField) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{10}
}

func (x *NameField) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SubMessage_Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SubMessage_Envelope) Reset() {
	*x = SubMessage_Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessage_Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessage_Envelope) ProtoMessage() {}

func (x *SubMessage_Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessage_Envelope.ProtoReflect.Descriptor instead.
func (*SubMessage_Envelope) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SubMessage_Envelope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SubMessage_Envelope_Letter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents string `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *SubMessage_Envelope_Letter) Reset() {
	*x = SubMessage_Envelope_Letter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessage_Envelope_Letter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessage_Envelope_Letter) ProtoMessage() {}

func (x *SubMessage_Envelope_Letter) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessage_Envelope_Letter.ProtoReflect.Descriptor instead.
func (*SubMessage_Envelope_Letter) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{4, 0, 0}
}

func (x *SubMessage_Envelope_Letter) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

var File_test_v1_test_proto protoreflect.FileDescriptor

var file_test_v1_test_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x24,
	0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e,
	0x75, 0x6d, 0x33, 0x32, 0x22, 0x27, 0x0a, 0x0d, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x22, 0x24, 0x0a,
	0x0c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x75,
	0x6d, 0x36, 0x34, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x44, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x24, 0x0a, 0x06, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x0a, 0x4d, 0x61,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x2a, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x47,
	0x0a, 0x0c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x18,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xf1,
	0x04, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f,
	0x6e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x0e, 0x45, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x03, 0x6e,
	0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x03, 0x6e, 0x6f, 0x77, 0x3a, 0x04, 0xd0, 0xf1, 0x04, 0x01, 0x22, 0x30, 0x0a, 0x13, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x3a, 0x07, 0xda, 0xf1, 0x04, 0x03, 0x70, 0x66, 0x78, 0x22, 0x2a, 0x0a, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x8a, 0xf1, 0x04, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0xee, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6c, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x65, 0x73, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x54, 0x58, 0xaa, 0x02, 0x15, 0x4f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x4f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x4f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x54, 0x65, 0x73, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x17, 0x4f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x3a,
	0x3a, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_test_v1_test_proto_rawDescOnce sync.Once
	file_test_v1_test_proto_rawDescData = file_test_v1_test_proto_rawDesc
)

func file_test_v1_test_proto_rawDescGZIP() []byte {
	file_test_v1_test_proto_rawDescOnce.Do(func() {
		file_test_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_v1_test_proto_rawDescData)
	})
	return file_test_v1_test_proto_rawDescData
}

var file_test_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_test_v1_test_proto_goTypes = []interface{}{
	(*StringMessage)(nil),              // 0: opentelemetry.test.v1.StringMessage
	(*Int32Message)(nil),               // 1: opentelemetry.test.v1.Int32Message
	(*Uint32Message)(nil),              // 2: opentelemetry.test.v1.Uint32Message
	(*Int64Message)(nil),               // 3: opentelemetry.test.v1.Int64Message
	(*SubMessage)(nil),                 // 4: opentelemetry.test.v1.SubMessage
	(*MapMessage)(nil),                 // 5: opentelemetry.test.v1.MapMessage
	(*MessageDetails)(nil),             // 6: opentelemetry.test.v1.MessageDetails
	(*ExcludeField)(nil),               // 7: opentelemetry.test.v1.ExcludeField
	(*ExcludeMessage)(nil),             // 8: opentelemetry.test.v1.ExcludeMessage
	(*RenameMessagePrefix)(nil),        // 9: opentelemetry.test.v1.RenameMessagePrefix
	(*NameField)(nil),                  // 10: opentelemetry.test.v1.NameField
	(*SubMessage_Envelope)(nil),        // 11: opentelemetry.test.v1.SubMessage.Envelope
	(*SubMessage_Envelope_Letter)(nil), // 12: opentelemetry.test.v1.SubMessage.Envelope.Letter
	nil,                                // 13: opentelemetry.test.v1.MapMessage.MetaEntry
	(*datetime.DateTime)(nil),          // 14: google.type.DateTime
}
var file_test_v1_test_proto_depIdxs = []int32{
	6,  // 0: opentelemetry.test.v1.SubMessage.details:type_name -> opentelemetry.test.v1.MessageDetails
	13, // 1: opentelemetry.test.v1.MapMessage.meta:type_name -> opentelemetry.test.v1.MapMessage.MetaEntry
	14, // 2: opentelemetry.test.v1.ExcludeMessage.now:type_name -> google.type.DateTime
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_test_v1_test_proto_init() }
func file_test_v1_test_proto_init() {
	if File_test_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
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
		file_test_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Message); i {
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
		file_test_v1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint32Message); i {
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
		file_test_v1_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Message); i {
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
		file_test_v1_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessage); i {
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
		file_test_v1_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMessage); i {
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
		file_test_v1_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_test_v1_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExcludeField); i {
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
		file_test_v1_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExcludeMessage); i {
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
		file_test_v1_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameMessagePrefix); i {
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
		file_test_v1_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameField); i {
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
		file_test_v1_test_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessage_Envelope); i {
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
		file_test_v1_test_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessage_Envelope_Letter); i {
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
			RawDescriptor: file_test_v1_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_v1_test_proto_goTypes,
		DependencyIndexes: file_test_v1_test_proto_depIdxs,
		MessageInfos:      file_test_v1_test_proto_msgTypes,
	}.Build()
	File_test_v1_test_proto = out.File
	file_test_v1_test_proto_rawDesc = nil
	file_test_v1_test_proto_goTypes = nil
	file_test_v1_test_proto_depIdxs = nil
}
