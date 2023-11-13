// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: archive/v1/archive.proto

package archivePb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateArchiveMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchiveEventId string `protobuf:"bytes,1,opt,name=archive_event_id,json=archiveEventId,proto3" json:"archive_event_id,omitempty"`
}

func (x *CreateArchiveMetaInfo) Reset() {
	*x = CreateArchiveMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_v1_archive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArchiveMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArchiveMetaInfo) ProtoMessage() {}

func (x *CreateArchiveMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_archive_v1_archive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArchiveMetaInfo.ProtoReflect.Descriptor instead.
func (*CreateArchiveMetaInfo) Descriptor() ([]byte, []int) {
	return file_archive_v1_archive_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArchiveMetaInfo) GetArchiveEventId() string {
	if x != nil {
		return x.ArchiveEventId
	}
	return ""
}

type GetArchiveMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchiveId   string `protobuf:"bytes,1,opt,name=archive_id,json=archiveId,proto3" json:"archive_id,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Size        int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetArchiveMetaInfo) Reset() {
	*x = GetArchiveMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_v1_archive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchiveMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchiveMetaInfo) ProtoMessage() {}

func (x *GetArchiveMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_archive_v1_archive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchiveMetaInfo.ProtoReflect.Descriptor instead.
func (*GetArchiveMetaInfo) Descriptor() ([]byte, []int) {
	return file_archive_v1_archive_proto_rawDescGZIP(), []int{1}
}

func (x *GetArchiveMetaInfo) GetArchiveId() string {
	if x != nil {
		return x.ArchiveId
	}
	return ""
}

func (x *GetArchiveMetaInfo) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *GetArchiveMetaInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type CreateArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*CreateArchiveRequest_Meta
	//	*CreateArchiveRequest_Chunk
	Value isCreateArchiveRequest_Value `protobuf_oneof:"value"`
}

func (x *CreateArchiveRequest) Reset() {
	*x = CreateArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_v1_archive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArchiveRequest) ProtoMessage() {}

func (x *CreateArchiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archive_v1_archive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArchiveRequest.ProtoReflect.Descriptor instead.
func (*CreateArchiveRequest) Descriptor() ([]byte, []int) {
	return file_archive_v1_archive_proto_rawDescGZIP(), []int{2}
}

func (m *CreateArchiveRequest) GetValue() isCreateArchiveRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *CreateArchiveRequest) GetMeta() *CreateArchiveMetaInfo {
	if x, ok := x.GetValue().(*CreateArchiveRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *CreateArchiveRequest) GetChunk() []byte {
	if x, ok := x.GetValue().(*CreateArchiveRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isCreateArchiveRequest_Value interface {
	isCreateArchiveRequest_Value()
}

type CreateArchiveRequest_Meta struct {
	Meta *CreateArchiveMetaInfo `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type CreateArchiveRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*CreateArchiveRequest_Meta) isCreateArchiveRequest_Value() {}

func (*CreateArchiveRequest_Chunk) isCreateArchiveRequest_Value() {}

type GetArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchiveEventId string `protobuf:"bytes,1,opt,name=archive_event_id,json=archiveEventId,proto3" json:"archive_event_id,omitempty"`
}

func (x *GetArchiveRequest) Reset() {
	*x = GetArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_v1_archive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchiveRequest) ProtoMessage() {}

func (x *GetArchiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archive_v1_archive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchiveRequest.ProtoReflect.Descriptor instead.
func (*GetArchiveRequest) Descriptor() ([]byte, []int) {
	return file_archive_v1_archive_proto_rawDescGZIP(), []int{3}
}

func (x *GetArchiveRequest) GetArchiveEventId() string {
	if x != nil {
		return x.ArchiveEventId
	}
	return ""
}

type GetArchiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*GetArchiveResponse_Meta
	//	*GetArchiveResponse_Chunk
	Value isGetArchiveResponse_Value `protobuf_oneof:"value"`
}

func (x *GetArchiveResponse) Reset() {
	*x = GetArchiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_v1_archive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchiveResponse) ProtoMessage() {}

func (x *GetArchiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_archive_v1_archive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchiveResponse.ProtoReflect.Descriptor instead.
func (*GetArchiveResponse) Descriptor() ([]byte, []int) {
	return file_archive_v1_archive_proto_rawDescGZIP(), []int{4}
}

func (m *GetArchiveResponse) GetValue() isGetArchiveResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *GetArchiveResponse) GetMeta() *GetArchiveMetaInfo {
	if x, ok := x.GetValue().(*GetArchiveResponse_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *GetArchiveResponse) GetChunk() []byte {
	if x, ok := x.GetValue().(*GetArchiveResponse_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isGetArchiveResponse_Value interface {
	isGetArchiveResponse_Value()
}

type GetArchiveResponse_Meta struct {
	Meta *GetArchiveMetaInfo `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type GetArchiveResponse_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*GetArchiveResponse_Meta) isGetArchiveResponse_Value() {}

func (*GetArchiveResponse_Chunk) isGetArchiveResponse_Value() {}

var File_archive_v1_archive_proto protoreflect.FileDescriptor

var file_archive_v1_archive_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x61, 0x74, 0x65,
	0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x6a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x7a, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x61,
	0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xce,
	0x01, 0x0a, 0x0e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x57, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x12, 0x2a, 0x2e, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x63, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x61, 0x74, 0x65, 0x6c,
	0x6c, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x69,
	0x74, 0x79, 0x42, 0x65, 0x61, 0x72, 0x33, 0x2f, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_archive_v1_archive_proto_rawDescOnce sync.Once
	file_archive_v1_archive_proto_rawDescData = file_archive_v1_archive_proto_rawDesc
)

func file_archive_v1_archive_proto_rawDescGZIP() []byte {
	file_archive_v1_archive_proto_rawDescOnce.Do(func() {
		file_archive_v1_archive_proto_rawDescData = protoimpl.X.CompressGZIP(file_archive_v1_archive_proto_rawDescData)
	})
	return file_archive_v1_archive_proto_rawDescData
}

var file_archive_v1_archive_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_archive_v1_archive_proto_goTypes = []interface{}{
	(*CreateArchiveMetaInfo)(nil), // 0: satellite.archive.v1.CreateArchiveMetaInfo
	(*GetArchiveMetaInfo)(nil),    // 1: satellite.archive.v1.GetArchiveMetaInfo
	(*CreateArchiveRequest)(nil),  // 2: satellite.archive.v1.CreateArchiveRequest
	(*GetArchiveRequest)(nil),     // 3: satellite.archive.v1.GetArchiveRequest
	(*GetArchiveResponse)(nil),    // 4: satellite.archive.v1.GetArchiveResponse
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_archive_v1_archive_proto_depIdxs = []int32{
	0, // 0: satellite.archive.v1.CreateArchiveRequest.meta:type_name -> satellite.archive.v1.CreateArchiveMetaInfo
	1, // 1: satellite.archive.v1.GetArchiveResponse.meta:type_name -> satellite.archive.v1.GetArchiveMetaInfo
	2, // 2: satellite.archive.v1.ArchiveService.CreateArchive:input_type -> satellite.archive.v1.CreateArchiveRequest
	3, // 3: satellite.archive.v1.ArchiveService.GetArchive:input_type -> satellite.archive.v1.GetArchiveRequest
	5, // 4: satellite.archive.v1.ArchiveService.CreateArchive:output_type -> google.protobuf.Empty
	4, // 5: satellite.archive.v1.ArchiveService.GetArchive:output_type -> satellite.archive.v1.GetArchiveResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_archive_v1_archive_proto_init() }
func file_archive_v1_archive_proto_init() {
	if File_archive_v1_archive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_archive_v1_archive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArchiveMetaInfo); i {
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
		file_archive_v1_archive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchiveMetaInfo); i {
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
		file_archive_v1_archive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArchiveRequest); i {
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
		file_archive_v1_archive_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchiveRequest); i {
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
		file_archive_v1_archive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchiveResponse); i {
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
	file_archive_v1_archive_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CreateArchiveRequest_Meta)(nil),
		(*CreateArchiveRequest_Chunk)(nil),
	}
	file_archive_v1_archive_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetArchiveResponse_Meta)(nil),
		(*GetArchiveResponse_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_archive_v1_archive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_archive_v1_archive_proto_goTypes,
		DependencyIndexes: file_archive_v1_archive_proto_depIdxs,
		MessageInfos:      file_archive_v1_archive_proto_msgTypes,
	}.Build()
	File_archive_v1_archive_proto = out.File
	file_archive_v1_archive_proto_rawDesc = nil
	file_archive_v1_archive_proto_goTypes = nil
	file_archive_v1_archive_proto_depIdxs = nil
}
