// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tv_service/private.proto

package tv_service

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

type StreamSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	StreamHost  *IPPort `protobuf:"bytes,2,req,name=stream_host,json=streamHost" json:"stream_host,omitempty"`
	ControlHost *IPPort `protobuf:"bytes,3,opt,name=control_host,json=controlHost" json:"control_host,omitempty"`
	Url         *string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (x *StreamSource) Reset() {
	*x = StreamSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_private_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSource) ProtoMessage() {}

func (x *StreamSource) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_private_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSource.ProtoReflect.Descriptor instead.
func (*StreamSource) Descriptor() ([]byte, []int) {
	return file_tv_service_private_proto_rawDescGZIP(), []int{0}
}

func (x *StreamSource) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *StreamSource) GetStreamHost() *IPPort {
	if x != nil {
		return x.StreamHost
	}
	return nil
}

func (x *StreamSource) GetControlHost() *IPPort {
	if x != nil {
		return x.ControlHost
	}
	return nil
}

func (x *StreamSource) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type ChannelSources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdChannel                *int32          `protobuf:"varint,1,req,name=id_channel,json=idChannel" json:"id_channel,omitempty"`
	IdOffset                 *int32          `protobuf:"varint,2,req,name=id_offset,json=idOffset" json:"id_offset,omitempty"`
	IdCompany                *int32          `protobuf:"varint,10,req,name=id_company,json=idCompany" json:"id_company,omitempty"`
	UdpMulticastGroup        []*StreamSource `protobuf:"bytes,3,rep,name=udp_multicast_group,json=udpMulticastGroup" json:"udp_multicast_group,omitempty"`
	SudpMulticastGroup       []*StreamSource `protobuf:"bytes,4,rep,name=sudp_multicast_group,json=sudpMulticastGroup" json:"sudp_multicast_group,omitempty"`
	UdpStreamer              []*StreamSource `protobuf:"bytes,5,rep,name=udp_streamer,json=udpStreamer" json:"udp_streamer,omitempty"`
	SudpStreamer             []*StreamSource `protobuf:"bytes,6,rep,name=sudp_streamer,json=sudpStreamer" json:"sudp_streamer,omitempty"`
	UdpHttpStreamer          []*StreamSource `protobuf:"bytes,7,rep,name=udp_http_streamer,json=udpHttpStreamer" json:"udp_http_streamer,omitempty"`
	SudpHttpStreamer         []*StreamSource `protobuf:"bytes,8,rep,name=sudp_http_streamer,json=sudpHttpStreamer" json:"sudp_http_streamer,omitempty"`
	CacheServer              []*StreamSource `protobuf:"bytes,9,rep,name=cache_server,json=cacheServer" json:"cache_server,omitempty"`
	HlsHttpStreamer          []*StreamSource `protobuf:"bytes,11,rep,name=hls_http_streamer,json=hlsHttpStreamer" json:"hls_http_streamer,omitempty"`
	HlsTimeshiftHttpStreamer []*StreamSource `protobuf:"bytes,12,rep,name=hls_timeshift_http_streamer,json=hlsTimeshiftHttpStreamer" json:"hls_timeshift_http_streamer,omitempty"`
	HlsDrmHttpsStreamer      []*StreamSource `protobuf:"bytes,13,rep,name=hls_drm_https_streamer,json=hlsDrmHttpsStreamer" json:"hls_drm_https_streamer,omitempty"`
	DashDrmHttpsStreamer     []*StreamSource `protobuf:"bytes,14,rep,name=dash_drm_https_streamer,json=dashDrmHttpsStreamer" json:"dash_drm_https_streamer,omitempty"`
	HlsAesHttpsStreamer      []*StreamSource `protobuf:"bytes,15,rep,name=hls_aes_https_streamer,json=hlsAesHttpsStreamer" json:"hls_aes_https_streamer,omitempty"`
}

func (x *ChannelSources) Reset() {
	*x = ChannelSources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_private_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelSources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelSources) ProtoMessage() {}

func (x *ChannelSources) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_private_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelSources.ProtoReflect.Descriptor instead.
func (*ChannelSources) Descriptor() ([]byte, []int) {
	return file_tv_service_private_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelSources) GetIdChannel() int32 {
	if x != nil && x.IdChannel != nil {
		return *x.IdChannel
	}
	return 0
}

func (x *ChannelSources) GetIdOffset() int32 {
	if x != nil && x.IdOffset != nil {
		return *x.IdOffset
	}
	return 0
}

func (x *ChannelSources) GetIdCompany() int32 {
	if x != nil && x.IdCompany != nil {
		return *x.IdCompany
	}
	return 0
}

func (x *ChannelSources) GetUdpMulticastGroup() []*StreamSource {
	if x != nil {
		return x.UdpMulticastGroup
	}
	return nil
}

func (x *ChannelSources) GetSudpMulticastGroup() []*StreamSource {
	if x != nil {
		return x.SudpMulticastGroup
	}
	return nil
}

func (x *ChannelSources) GetUdpStreamer() []*StreamSource {
	if x != nil {
		return x.UdpStreamer
	}
	return nil
}

func (x *ChannelSources) GetSudpStreamer() []*StreamSource {
	if x != nil {
		return x.SudpStreamer
	}
	return nil
}

func (x *ChannelSources) GetUdpHttpStreamer() []*StreamSource {
	if x != nil {
		return x.UdpHttpStreamer
	}
	return nil
}

func (x *ChannelSources) GetSudpHttpStreamer() []*StreamSource {
	if x != nil {
		return x.SudpHttpStreamer
	}
	return nil
}

func (x *ChannelSources) GetCacheServer() []*StreamSource {
	if x != nil {
		return x.CacheServer
	}
	return nil
}

func (x *ChannelSources) GetHlsHttpStreamer() []*StreamSource {
	if x != nil {
		return x.HlsHttpStreamer
	}
	return nil
}

func (x *ChannelSources) GetHlsTimeshiftHttpStreamer() []*StreamSource {
	if x != nil {
		return x.HlsTimeshiftHttpStreamer
	}
	return nil
}

func (x *ChannelSources) GetHlsDrmHttpsStreamer() []*StreamSource {
	if x != nil {
		return x.HlsDrmHttpsStreamer
	}
	return nil
}

func (x *ChannelSources) GetDashDrmHttpsStreamer() []*StreamSource {
	if x != nil {
		return x.DashDrmHttpsStreamer
	}
	return nil
}

func (x *ChannelSources) GetHlsAesHttpsStreamer() []*StreamSource {
	if x != nil {
		return x.HlsAesHttpsStreamer
	}
	return nil
}

type ChannelSorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name      *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	IdChannel []int32 `protobuf:"varint,3,rep,name=id_channel,json=idChannel" json:"id_channel,omitempty"`
}

func (x *ChannelSorting) Reset() {
	*x = ChannelSorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_private_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelSorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelSorting) ProtoMessage() {}

func (x *ChannelSorting) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_private_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelSorting.ProtoReflect.Descriptor instead.
func (*ChannelSorting) Descriptor() ([]byte, []int) {
	return file_tv_service_private_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelSorting) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ChannelSorting) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ChannelSorting) GetIdChannel() []int32 {
	if x != nil {
		return x.IdChannel
	}
	return nil
}

var File_tv_service_private_proto protoreflect.FileDescriptor

var file_tv_service_private_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x18, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74,
	0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x50, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0xd6, 0x07, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0a, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x48, 0x0a,
	0x13, 0x75, 0x64, 0x70, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x11, 0x75, 0x64, 0x70, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x14, 0x73, 0x75, 0x64, 0x70, 0x5f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x12, 0x73, 0x75, 0x64, 0x70, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x3b, 0x0a, 0x0c, 0x75, 0x64, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x0b, 0x75, 0x64, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x12, 0x3d, 0x0a, 0x0d, 0x73, 0x75, 0x64, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x0c, 0x73, 0x75, 0x64, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12,
	0x44, 0x0a, 0x11, 0x75, 0x64, 0x70, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x0f, 0x75, 0x64, 0x70, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x12, 0x73, 0x75, 0x64, 0x70, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x10, 0x73, 0x75, 0x64,
	0x70, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x3b, 0x0a,
	0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x11, 0x68, 0x6c,
	0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0f, 0x68, 0x6c, 0x73, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x12, 0x57, 0x0a, 0x1b, 0x68, 0x6c, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x66,
	0x74, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x18, 0x68, 0x6c, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x66, 0x74, 0x48, 0x74, 0x74,
	0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x16, 0x68, 0x6c, 0x73,
	0x5f, 0x64, 0x72, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x13, 0x68, 0x6c, 0x73, 0x44, 0x72, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x73,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x17, 0x64, 0x61, 0x73, 0x68,
	0x5f, 0x64, 0x72, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x14, 0x64, 0x61, 0x73, 0x68, 0x44, 0x72, 0x6d, 0x48, 0x74, 0x74, 0x70,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x16, 0x68, 0x6c, 0x73,
	0x5f, 0x61, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x13, 0x68, 0x6c, 0x73, 0x41, 0x65, 0x73, 0x48, 0x74, 0x74, 0x70, 0x73,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x74, 0x76,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65,
}

var (
	file_tv_service_private_proto_rawDescOnce sync.Once
	file_tv_service_private_proto_rawDescData = file_tv_service_private_proto_rawDesc
)

func file_tv_service_private_proto_rawDescGZIP() []byte {
	file_tv_service_private_proto_rawDescOnce.Do(func() {
		file_tv_service_private_proto_rawDescData = protoimpl.X.CompressGZIP(file_tv_service_private_proto_rawDescData)
	})
	return file_tv_service_private_proto_rawDescData
}

var file_tv_service_private_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tv_service_private_proto_goTypes = []interface{}{
	(*StreamSource)(nil),   // 0: tv_service.StreamSource
	(*ChannelSources)(nil), // 1: tv_service.ChannelSources
	(*ChannelSorting)(nil), // 2: tv_service.ChannelSorting
	(*IPPort)(nil),         // 3: tv_service.IPPort
}
var file_tv_service_private_proto_depIdxs = []int32{
	3,  // 0: tv_service.StreamSource.stream_host:type_name -> tv_service.IPPort
	3,  // 1: tv_service.StreamSource.control_host:type_name -> tv_service.IPPort
	0,  // 2: tv_service.ChannelSources.udp_multicast_group:type_name -> tv_service.StreamSource
	0,  // 3: tv_service.ChannelSources.sudp_multicast_group:type_name -> tv_service.StreamSource
	0,  // 4: tv_service.ChannelSources.udp_streamer:type_name -> tv_service.StreamSource
	0,  // 5: tv_service.ChannelSources.sudp_streamer:type_name -> tv_service.StreamSource
	0,  // 6: tv_service.ChannelSources.udp_http_streamer:type_name -> tv_service.StreamSource
	0,  // 7: tv_service.ChannelSources.sudp_http_streamer:type_name -> tv_service.StreamSource
	0,  // 8: tv_service.ChannelSources.cache_server:type_name -> tv_service.StreamSource
	0,  // 9: tv_service.ChannelSources.hls_http_streamer:type_name -> tv_service.StreamSource
	0,  // 10: tv_service.ChannelSources.hls_timeshift_http_streamer:type_name -> tv_service.StreamSource
	0,  // 11: tv_service.ChannelSources.hls_drm_https_streamer:type_name -> tv_service.StreamSource
	0,  // 12: tv_service.ChannelSources.dash_drm_https_streamer:type_name -> tv_service.StreamSource
	0,  // 13: tv_service.ChannelSources.hls_aes_https_streamer:type_name -> tv_service.StreamSource
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_tv_service_private_proto_init() }
func file_tv_service_private_proto_init() {
	if File_tv_service_private_proto != nil {
		return
	}
	file_tv_service_servers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tv_service_private_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamSource); i {
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
		file_tv_service_private_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelSources); i {
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
		file_tv_service_private_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelSorting); i {
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
			RawDescriptor: file_tv_service_private_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tv_service_private_proto_goTypes,
		DependencyIndexes: file_tv_service_private_proto_depIdxs,
		MessageInfos:      file_tv_service_private_proto_msgTypes,
	}.Build()
	File_tv_service_private_proto = out.File
	file_tv_service_private_proto_rawDesc = nil
	file_tv_service_private_proto_goTypes = nil
	file_tv_service_private_proto_depIdxs = nil
}
