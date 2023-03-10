// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tv_service/http_streamer_link.proto

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

type HttpStreamerLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link           *string `protobuf:"bytes,1,req,name=link" json:"link,omitempty"`
	Auth           *string `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	FirstBatchSize *uint32 `protobuf:"varint,3,opt,name=first_batch_size,json=firstBatchSize" json:"first_batch_size,omitempty"`
	ValidUntil     *uint64 `protobuf:"varint,4,opt,name=valid_until,json=validUntil" json:"valid_until,omitempty"`
	Ip             *uint32 `protobuf:"varint,5,opt,name=ip" json:"ip,omitempty"`
	TimeStart      *uint64 `protobuf:"varint,6,opt,name=time_start,json=timeStart" json:"time_start,omitempty"`
	TimeStop       *uint64 `protobuf:"varint,7,opt,name=time_stop,json=timeStop" json:"time_stop,omitempty"`
	Multistream    *bool   `protobuf:"varint,8,opt,name=multistream" json:"multistream,omitempty"`
	StreamName     *string `protobuf:"bytes,9,opt,name=stream_name,json=streamName" json:"stream_name,omitempty"`
	ContractId     *uint64 `protobuf:"varint,10,opt,name=contract_id,json=contractId" json:"contract_id,omitempty"`
	Locale         *string `protobuf:"bytes,11,opt,name=locale" json:"locale,omitempty"`
	StreamId       *int32  `protobuf:"varint,12,opt,name=stream_id,json=streamId" json:"stream_id,omitempty"`
}

func (x *HttpStreamerLink) Reset() {
	*x = HttpStreamerLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_http_streamer_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpStreamerLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpStreamerLink) ProtoMessage() {}

func (x *HttpStreamerLink) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_http_streamer_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpStreamerLink.ProtoReflect.Descriptor instead.
func (*HttpStreamerLink) Descriptor() ([]byte, []int) {
	return file_tv_service_http_streamer_link_proto_rawDescGZIP(), []int{0}
}

func (x *HttpStreamerLink) GetLink() string {
	if x != nil && x.Link != nil {
		return *x.Link
	}
	return ""
}

func (x *HttpStreamerLink) GetAuth() string {
	if x != nil && x.Auth != nil {
		return *x.Auth
	}
	return ""
}

func (x *HttpStreamerLink) GetFirstBatchSize() uint32 {
	if x != nil && x.FirstBatchSize != nil {
		return *x.FirstBatchSize
	}
	return 0
}

func (x *HttpStreamerLink) GetValidUntil() uint64 {
	if x != nil && x.ValidUntil != nil {
		return *x.ValidUntil
	}
	return 0
}

func (x *HttpStreamerLink) GetIp() uint32 {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return 0
}

func (x *HttpStreamerLink) GetTimeStart() uint64 {
	if x != nil && x.TimeStart != nil {
		return *x.TimeStart
	}
	return 0
}

func (x *HttpStreamerLink) GetTimeStop() uint64 {
	if x != nil && x.TimeStop != nil {
		return *x.TimeStop
	}
	return 0
}

func (x *HttpStreamerLink) GetMultistream() bool {
	if x != nil && x.Multistream != nil {
		return *x.Multistream
	}
	return false
}

func (x *HttpStreamerLink) GetStreamName() string {
	if x != nil && x.StreamName != nil {
		return *x.StreamName
	}
	return ""
}

func (x *HttpStreamerLink) GetContractId() uint64 {
	if x != nil && x.ContractId != nil {
		return *x.ContractId
	}
	return 0
}

func (x *HttpStreamerLink) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *HttpStreamerLink) GetStreamId() int32 {
	if x != nil && x.StreamId != nil {
		return *x.StreamId
	}
	return 0
}

var File_tv_service_http_streamer_link_proto protoreflect.FileDescriptor

var file_tv_service_http_streamer_link_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0xea, 0x02, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x28,
	0x0a, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x74,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65,
}

var (
	file_tv_service_http_streamer_link_proto_rawDescOnce sync.Once
	file_tv_service_http_streamer_link_proto_rawDescData = file_tv_service_http_streamer_link_proto_rawDesc
)

func file_tv_service_http_streamer_link_proto_rawDescGZIP() []byte {
	file_tv_service_http_streamer_link_proto_rawDescOnce.Do(func() {
		file_tv_service_http_streamer_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_tv_service_http_streamer_link_proto_rawDescData)
	})
	return file_tv_service_http_streamer_link_proto_rawDescData
}

var file_tv_service_http_streamer_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tv_service_http_streamer_link_proto_goTypes = []interface{}{
	(*HttpStreamerLink)(nil), // 0: tv_service.HttpStreamerLink
}
var file_tv_service_http_streamer_link_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tv_service_http_streamer_link_proto_init() }
func file_tv_service_http_streamer_link_proto_init() {
	if File_tv_service_http_streamer_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tv_service_http_streamer_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpStreamerLink); i {
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
			RawDescriptor: file_tv_service_http_streamer_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tv_service_http_streamer_link_proto_goTypes,
		DependencyIndexes: file_tv_service_http_streamer_link_proto_depIdxs,
		MessageInfos:      file_tv_service_http_streamer_link_proto_msgTypes,
	}.Build()
	File_tv_service_http_streamer_link_proto = out.File
	file_tv_service_http_streamer_link_proto_rawDesc = nil
	file_tv_service_http_streamer_link_proto_goTypes = nil
	file_tv_service_http_streamer_link_proto_depIdxs = nil
}
