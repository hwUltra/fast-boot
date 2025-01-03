// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: pb/oms.proto

package omsPb

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

// ------------------------------------
// Common Message
// ------------------------------------
type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{0}
}

func (x *IdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids string `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsReq) Reset() {
	*x = IdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsReq) ProtoMessage() {}

func (x *IdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsReq.ProtoReflect.Descriptor instead.
func (*IdsReq) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{1}
}

func (x *IdsReq) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type AnyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnyReq) Reset() {
	*x = AnyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyReq) ProtoMessage() {}

func (x *AnyReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyReq.ProtoReflect.Descriptor instead.
func (*AnyReq) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{2}
}

type IdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdResp) Reset() {
	*x = IdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResp) ProtoMessage() {}

func (x *IdResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResp.ProtoReflect.Descriptor instead.
func (*IdResp) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{3}
}

func (x *IdResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SuccessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SuccessResp) Reset() {
	*x = SuccessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResp) ProtoMessage() {}

func (x *SuccessResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResp.ProtoReflect.Descriptor instead.
func (*SuccessResp) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{4}
}

type SuccessIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SuccessIdResp) Reset() {
	*x = SuccessIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessIdResp) ProtoMessage() {}

func (x *SuccessIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessIdResp.ProtoReflect.Descriptor instead.
func (*SuccessIdResp) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessIdResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords string `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Status   int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	PageNum  int64  `protobuf:"varint,3,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{6}
}

func (x *ListReq) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ListReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{7}
}

type OrderListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderListResp) Reset() {
	*x = OrderListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResp) ProtoMessage() {}

func (x *OrderListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResp.ProtoReflect.Descriptor instead.
func (*OrderListResp) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{8}
}

type OrderAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderAddReq) Reset() {
	*x = OrderAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddReq) ProtoMessage() {}

func (x *OrderAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddReq.ProtoReflect.Descriptor instead.
func (*OrderAddReq) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{9}
}

type OrderUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderUpdateReq) Reset() {
	*x = OrderUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_oms_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateReq) ProtoMessage() {}

func (x *OrderUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_oms_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateReq.ProtoReflect.Descriptor instead.
func (*OrderUpdateReq) Descriptor() ([]byte, []int) {
	return file_pb_oms_proto_rawDescGZIP(), []int{10}
}

var File_pb_oms_proto protoreflect.FileDescriptor

var file_pb_oms_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6f, 0x6d, 0x73, 0x50, 0x62, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a,
	0x0a, 0x06, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x08, 0x0a, 0x06, 0x41, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x22, 0x18, 0x0a, 0x06, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0d,
	0x0a, 0x0b, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x0a,
	0x0d, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73,
	0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0x0a,
	0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x22, 0x10, 0x0a, 0x0e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x32, 0xd0,
	0x01, 0x0a, 0x03, 0x6f, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47,
	0x65, 0x74, 0x12, 0x0c, 0x2e, 0x6f, 0x6d, 0x73, 0x50, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x6f, 0x6d, 0x73, 0x50, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x6f, 0x6d,
	0x73, 0x50, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6f, 0x6d,
	0x73, 0x50, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x34, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x12, 0x12, 0x2e,
	0x6f, 0x6d, 0x73, 0x50, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x6f, 0x6d, 0x73, 0x50, 0x62, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x6f, 0x6d, 0x73, 0x50, 0x62, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x6f, 0x6d, 0x73, 0x50, 0x62, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6f, 0x6d, 0x73, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_oms_proto_rawDescOnce sync.Once
	file_pb_oms_proto_rawDescData = file_pb_oms_proto_rawDesc
)

func file_pb_oms_proto_rawDescGZIP() []byte {
	file_pb_oms_proto_rawDescOnce.Do(func() {
		file_pb_oms_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_oms_proto_rawDescData)
	})
	return file_pb_oms_proto_rawDescData
}

var file_pb_oms_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pb_oms_proto_goTypes = []interface{}{
	(*IdReq)(nil),          // 0: omsPb.IdReq
	(*IdsReq)(nil),         // 1: omsPb.IdsReq
	(*AnyReq)(nil),         // 2: omsPb.AnyReq
	(*IdResp)(nil),         // 3: omsPb.IdResp
	(*SuccessResp)(nil),    // 4: omsPb.SuccessResp
	(*SuccessIdResp)(nil),  // 5: omsPb.SuccessIdResp
	(*ListReq)(nil),        // 6: omsPb.ListReq
	(*Order)(nil),          // 7: omsPb.Order
	(*OrderListResp)(nil),  // 8: omsPb.OrderListResp
	(*OrderAddReq)(nil),    // 9: omsPb.OrderAddReq
	(*OrderUpdateReq)(nil), // 10: omsPb.OrderUpdateReq
}
var file_pb_oms_proto_depIdxs = []int32{
	0,  // 0: omsPb.oms.OrderGet:input_type -> omsPb.IdReq
	6,  // 1: omsPb.oms.OrderList:input_type -> omsPb.ListReq
	9,  // 2: omsPb.oms.OrderAdd:input_type -> omsPb.OrderAddReq
	10, // 3: omsPb.oms.OrderUpdate:input_type -> omsPb.OrderUpdateReq
	7,  // 4: omsPb.oms.OrderGet:output_type -> omsPb.Order
	8,  // 5: omsPb.oms.OrderList:output_type -> omsPb.OrderListResp
	5,  // 6: omsPb.oms.OrderAdd:output_type -> omsPb.SuccessIdResp
	4,  // 7: omsPb.oms.OrderUpdate:output_type -> omsPb.SuccessResp
	4,  // [4:8] is the sub-list for method output_type
	0,  // [0:4] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pb_oms_proto_init() }
func file_pb_oms_proto_init() {
	if File_pb_oms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_oms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_pb_oms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdsReq); i {
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
		file_pb_oms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyReq); i {
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
		file_pb_oms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdResp); i {
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
		file_pb_oms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResp); i {
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
		file_pb_oms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessIdResp); i {
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
		file_pb_oms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_pb_oms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_pb_oms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResp); i {
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
		file_pb_oms_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddReq); i {
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
		file_pb_oms_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdateReq); i {
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
			RawDescriptor: file_pb_oms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_oms_proto_goTypes,
		DependencyIndexes: file_pb_oms_proto_depIdxs,
		MessageInfos:      file_pb_oms_proto_msgTypes,
	}.Build()
	File_pb_oms_proto = out.File
	file_pb_oms_proto_rawDesc = nil
	file_pb_oms_proto_goTypes = nil
	file_pb_oms_proto_depIdxs = nil
}
