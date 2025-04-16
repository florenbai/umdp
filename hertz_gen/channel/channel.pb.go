// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.22.2
// source: channel.proto

package channel

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	_ "umdp/hertz_gen/api"
	base "umdp/hertz_gen/base"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64           `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	ChannelName   string           `protobuf:"bytes,2,opt,name=channelName,proto3" form:"channelName" json:"channelName"`
	ChannelTag    string           `protobuf:"bytes,3,opt,name=channelTag,proto3" form:"channelTag" json:"channelTag"`
	ChannelConfig *structpb.Struct `protobuf:"bytes,4,opt,name=channelConfig,proto3" form:"channelConfig" json:"channelConfig"`
	ChannelStatus int32            `protobuf:"varint,5,opt,name=channelStatus,proto3" form:"channelStatus" json:"channelStatus"`
	CreatedAt     string           `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt" form:"createdAt" query:"createdAt"`
	UpdatedAt     string           `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt" form:"updatedAt" query:"updatedAt"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Channel) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *Channel) GetChannelTag() string {
	if x != nil {
		return x.ChannelTag
	}
	return ""
}

func (x *Channel) GetChannelConfig() *structpb.Struct {
	if x != nil {
		return x.ChannelConfig
	}
	return nil
}

func (x *Channel) GetChannelStatus() int32 {
	if x != nil {
		return x.ChannelStatus
	}
	return 0
}

func (x *Channel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Channel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ChannelMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	ChannelName string `protobuf:"bytes,2,opt,name=channelName,proto3" form:"channelName" json:"channelName"`
	ChannelTag  string `protobuf:"bytes,3,opt,name=channelTag,proto3" form:"channelTag" json:"channelTag"`
}

func (x *ChannelMap) Reset() {
	*x = ChannelMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMap) ProtoMessage() {}

func (x *ChannelMap) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMap.ProtoReflect.Descriptor instead.
func (*ChannelMap) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelMap) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChannelMap) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *ChannelMap) GetChannelTag() string {
	if x != nil {
		return x.ChannelTag
	}
	return ""
}

type ChannelListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64     `protobuf:"varint,1,opt,name=total,proto3" json:"total" form:"total" query:"total"`
	List  []*Channel `protobuf:"bytes,2,rep,name=list,proto3" json:"list" form:"list" query:"list"`
}

func (x *ChannelListResponse) Reset() {
	*x = ChannelListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelListResponse) ProtoMessage() {}

func (x *ChannelListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelListResponse.ProtoReflect.Descriptor instead.
func (*ChannelListResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ChannelListResponse) GetList() []*Channel {
	if x != nil {
		return x.List
	}
	return nil
}

type ChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelName   string           `protobuf:"bytes,1,opt,name=channelName,proto3" form:"channelName" json:"channelName"`
	ChannelTag    string           `protobuf:"bytes,2,opt,name=channelTag,proto3" form:"channelTag" json:"channelTag"`
	ChannelConfig *structpb.Struct `protobuf:"bytes,3,opt,name=channelConfig,proto3" form:"channelConfig" json:"channelConfig"`
	ChannelStatus int32            `protobuf:"varint,4,opt,name=channelStatus,proto3" form:"channelStatus" json:"channelStatus"`
}

func (x *ChannelRequest) Reset() {
	*x = ChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelRequest) ProtoMessage() {}

func (x *ChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelRequest.ProtoReflect.Descriptor instead.
func (*ChannelRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelRequest) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *ChannelRequest) GetChannelTag() string {
	if x != nil {
		return x.ChannelTag
	}
	return ""
}

func (x *ChannelRequest) GetChannelConfig() *structpb.Struct {
	if x != nil {
		return x.ChannelConfig
	}
	return nil
}

func (x *ChannelRequest) GetChannelStatus() int32 {
	if x != nil {
		return x.ChannelStatus
	}
	return 0
}

var File_channel_proto protoreflect.FileDescriptor

var file_channel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0xca, 0xbb, 0x18, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x67,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x67, 0x12, 0x50, 0x0a, 0x0d,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x11, 0xca, 0xbb,
	0x18, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37,
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x11, 0xca, 0xbb, 0x18, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x7f, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xca, 0xbb, 0x18, 0x0b, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x67, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x61, 0x67, 0x22, 0x51, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0xca, 0xbb, 0x18, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61,
	0x67, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x67, 0x12, 0x50, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x11, 0xca,
	0xbb, 0x18, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x37, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x11, 0xca, 0xbb, 0x18, 0x0d, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9a, 0x02, 0x0a, 0x0e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0c,
	0xd2, 0xc1, 0x18, 0x08, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x04,
	0x45, 0x64, 0x69, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0xda,
	0xc1, 0x18, 0x0c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x3a, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0xe2, 0xc1, 0x18, 0x0c, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x3a, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x19, 0x5a, 0x17, 0x6d, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x68,
	0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_channel_proto_rawDescOnce sync.Once
	file_channel_proto_rawDescData = file_channel_proto_rawDesc
)

func file_channel_proto_rawDescGZIP() []byte {
	file_channel_proto_rawDescOnce.Do(func() {
		file_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_rawDescData)
	})
	return file_channel_proto_rawDescData
}

var file_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_channel_proto_goTypes = []interface{}{
	(*Channel)(nil),             // 0: channel.Channel
	(*ChannelMap)(nil),          // 1: channel.ChannelMap
	(*ChannelListResponse)(nil), // 2: channel.ChannelListResponse
	(*ChannelRequest)(nil),      // 3: channel.ChannelRequest
	(*structpb.Struct)(nil),     // 4: google.protobuf.Struct
	(*base.BaseId)(nil),         // 5: base.BaseId
	(*base.BaseListReq)(nil),    // 6: base.BaseListReq
	(*base.BaseResp)(nil),       // 7: base.BaseResp
}
var file_channel_proto_depIdxs = []int32{
	4, // 0: channel.Channel.channelConfig:type_name -> google.protobuf.Struct
	0, // 1: channel.ChannelListResponse.list:type_name -> channel.Channel
	4, // 2: channel.ChannelRequest.channelConfig:type_name -> google.protobuf.Struct
	3, // 3: channel.ChannelService.Create:input_type -> channel.ChannelRequest
	3, // 4: channel.ChannelService.Edit:input_type -> channel.ChannelRequest
	5, // 5: channel.ChannelService.Delete:input_type -> base.BaseId
	6, // 6: channel.ChannelService.List:input_type -> base.BaseListReq
	7, // 7: channel.ChannelService.Create:output_type -> base.BaseResp
	7, // 8: channel.ChannelService.Edit:output_type -> base.BaseResp
	7, // 9: channel.ChannelService.Delete:output_type -> base.BaseResp
	2, // 10: channel.ChannelService.List:output_type -> channel.ChannelListResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_channel_proto_init() }
func file_channel_proto_init() {
	if File_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelMap); i {
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
		file_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelListResponse); i {
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
		file_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelRequest); i {
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
			RawDescriptor: file_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_channel_proto_goTypes,
		DependencyIndexes: file_channel_proto_depIdxs,
		MessageInfos:      file_channel_proto_msgTypes,
	}.Build()
	File_channel_proto = out.File
	file_channel_proto_rawDesc = nil
	file_channel_proto_goTypes = nil
	file_channel_proto_depIdxs = nil
}
