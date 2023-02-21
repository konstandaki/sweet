// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tv_service/channel.proto

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

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    *int32       `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                  *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Icon                  []byte       `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	Epg                   []*EpgRecord `protobuf:"bytes,4,rep,name=epg" json:"epg,omitempty"`
	Offset                []int32      `protobuf:"varint,5,rep,name=offset" json:"offset,omitempty"`
	Available             *bool        `protobuf:"varint,6,opt,name=available,def=1" json:"available,omitempty"`
	Hidden                *bool        `protobuf:"varint,7,req,name=hidden" json:"hidden,omitempty"`
	Category              []int32      `protobuf:"varint,8,rep,name=category" json:"category,omitempty"`
	HiddenInFullList      *bool        `protobuf:"varint,9,opt,name=hidden_in_full_list,json=hiddenInFullList" json:"hidden_in_full_list,omitempty"`
	IconBig               []byte       `protobuf:"bytes,10,opt,name=icon_big,json=iconBig" json:"icon_big,omitempty"`
	Catchup               *bool        `protobuf:"varint,11,opt,name=catchup" json:"catchup,omitempty"`
	CatchupDuration       *int32       `protobuf:"varint,12,opt,name=catchup_duration,json=catchupDuration" json:"catchup_duration,omitempty"`
	IconUrl               *string      `protobuf:"bytes,13,opt,name=icon_url,json=iconUrl" json:"icon_url,omitempty"`
	Tariffs               []int32      `protobuf:"varint,14,rep,name=tariffs" json:"tariffs,omitempty"`
	Drm                   *bool        `protobuf:"varint,15,opt,name=drm" json:"drm,omitempty"`
	OwnerId               *int32       `protobuf:"varint,16,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	Slug                  *string      `protobuf:"bytes,17,opt,name=slug" json:"slug,omitempty"`
	Number                *int32       `protobuf:"varint,18,opt,name=number" json:"number,omitempty"`
	BannerUrl             *string      `protobuf:"bytes,19,opt,name=banner_url,json=bannerUrl" json:"banner_url,omitempty"`
	DarkThemeIconUrl      *string      `protobuf:"bytes,20,opt,name=dark_theme_icon_url,json=darkThemeIconUrl" json:"dark_theme_icon_url,omitempty"`
	Colour                *string      `protobuf:"bytes,21,opt,name=colour" json:"colour,omitempty"`
	IconV2Url             *string      `protobuf:"bytes,22,opt,name=icon_v2_url,json=iconV2Url" json:"icon_v2_url,omitempty"`
	Translations          []string     `protobuf:"bytes,23,rep,name=translations" json:"translations,omitempty"`
	AvailabilityInfo      *string      `protobuf:"bytes,38,opt,name=availability_info,json=availabilityInfo" json:"availability_info,omitempty"`
	AvailabilityInfoColor *string      `protobuf:"bytes,44,opt,name=availability_info_color,json=availabilityInfoColor" json:"availability_info_color,omitempty"`
	// 100+ reserved for flags
	RewindDisabled      *bool `protobuf:"varint,100,opt,name=rewind_disabled,json=rewindDisabled" json:"rewind_disabled,omitempty"`
	FastForwardDisabled *bool `protobuf:"varint,101,opt,name=fast_forward_disabled,json=fastForwardDisabled" json:"fast_forward_disabled,omitempty"`
	SpeedupDisabled     *bool `protobuf:"varint,102,opt,name=speedup_disabled,json=speedupDisabled" json:"speedup_disabled,omitempty"`
	// 200+ reserved for messages
	RewindMessage      *string `protobuf:"bytes,200,opt,name=rewind_message,json=rewindMessage" json:"rewind_message,omitempty"`
	FastForwardMessage *string `protobuf:"bytes,201,opt,name=fast_forward_message,json=fastForwardMessage" json:"fast_forward_message,omitempty"`
	SpeedupMessage     *string `protobuf:"bytes,202,opt,name=speedup_message,json=speedupMessage" json:"speedup_message,omitempty"`
}

// Default values for Channel fields.
const (
	Default_Channel_Available = bool(true)
)

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_channel_proto_msgTypes[0]
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
	return file_tv_service_channel_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Channel) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Channel) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *Channel) GetEpg() []*EpgRecord {
	if x != nil {
		return x.Epg
	}
	return nil
}

func (x *Channel) GetOffset() []int32 {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *Channel) GetAvailable() bool {
	if x != nil && x.Available != nil {
		return *x.Available
	}
	return Default_Channel_Available
}

func (x *Channel) GetHidden() bool {
	if x != nil && x.Hidden != nil {
		return *x.Hidden
	}
	return false
}

func (x *Channel) GetCategory() []int32 {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Channel) GetHiddenInFullList() bool {
	if x != nil && x.HiddenInFullList != nil {
		return *x.HiddenInFullList
	}
	return false
}

func (x *Channel) GetIconBig() []byte {
	if x != nil {
		return x.IconBig
	}
	return nil
}

func (x *Channel) GetCatchup() bool {
	if x != nil && x.Catchup != nil {
		return *x.Catchup
	}
	return false
}

func (x *Channel) GetCatchupDuration() int32 {
	if x != nil && x.CatchupDuration != nil {
		return *x.CatchupDuration
	}
	return 0
}

func (x *Channel) GetIconUrl() string {
	if x != nil && x.IconUrl != nil {
		return *x.IconUrl
	}
	return ""
}

func (x *Channel) GetTariffs() []int32 {
	if x != nil {
		return x.Tariffs
	}
	return nil
}

func (x *Channel) GetDrm() bool {
	if x != nil && x.Drm != nil {
		return *x.Drm
	}
	return false
}

func (x *Channel) GetOwnerId() int32 {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return 0
}

func (x *Channel) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

func (x *Channel) GetNumber() int32 {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return 0
}

func (x *Channel) GetBannerUrl() string {
	if x != nil && x.BannerUrl != nil {
		return *x.BannerUrl
	}
	return ""
}

func (x *Channel) GetDarkThemeIconUrl() string {
	if x != nil && x.DarkThemeIconUrl != nil {
		return *x.DarkThemeIconUrl
	}
	return ""
}

func (x *Channel) GetColour() string {
	if x != nil && x.Colour != nil {
		return *x.Colour
	}
	return ""
}

func (x *Channel) GetIconV2Url() string {
	if x != nil && x.IconV2Url != nil {
		return *x.IconV2Url
	}
	return ""
}

func (x *Channel) GetTranslations() []string {
	if x != nil {
		return x.Translations
	}
	return nil
}

func (x *Channel) GetAvailabilityInfo() string {
	if x != nil && x.AvailabilityInfo != nil {
		return *x.AvailabilityInfo
	}
	return ""
}

func (x *Channel) GetAvailabilityInfoColor() string {
	if x != nil && x.AvailabilityInfoColor != nil {
		return *x.AvailabilityInfoColor
	}
	return ""
}

func (x *Channel) GetRewindDisabled() bool {
	if x != nil && x.RewindDisabled != nil {
		return *x.RewindDisabled
	}
	return false
}

func (x *Channel) GetFastForwardDisabled() bool {
	if x != nil && x.FastForwardDisabled != nil {
		return *x.FastForwardDisabled
	}
	return false
}

func (x *Channel) GetSpeedupDisabled() bool {
	if x != nil && x.SpeedupDisabled != nil {
		return *x.SpeedupDisabled
	}
	return false
}

func (x *Channel) GetRewindMessage() string {
	if x != nil && x.RewindMessage != nil {
		return *x.RewindMessage
	}
	return ""
}

func (x *Channel) GetFastForwardMessage() string {
	if x != nil && x.FastForwardMessage != nil {
		return *x.FastForwardMessage
	}
	return ""
}

func (x *Channel) GetSpeedupMessage() string {
	if x != nil && x.SpeedupMessage != nil {
		return *x.SpeedupMessage
	}
	return ""
}

var File_tv_service_channel_proto protoreflect.FileDescriptor

var file_tv_service_channel_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x14, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x65, 0x70, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x08, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x03, 0x65, 0x70, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x70, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x03, 0x65, 0x70, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x22, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18,
	0x07, 0x20, 0x02, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x13, 0x68, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x49, 0x6e,
	0x46, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e,
	0x5f, 0x62, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e,
	0x42, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x72, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x64, 0x72, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x61, 0x72, 0x6b, 0x5f, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x61, 0x72, 0x6b, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x63, 0x6f,
	0x6e, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x12, 0x1e, 0x0a, 0x0b,
	0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x56, 0x32, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x17, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2b, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a,
	0x17, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x5f,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x32,
	0x0a, 0x15, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x66,
	0x61, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x70, 0x65, 0x65, 0x64, 0x75, 0x70, 0x5f, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x75, 0x70, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xc9, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x61, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xca, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x65, 0x64, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x77, 0x65,
	0x65, 0x74, 0x2e, 0x74, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
}

var (
	file_tv_service_channel_proto_rawDescOnce sync.Once
	file_tv_service_channel_proto_rawDescData = file_tv_service_channel_proto_rawDesc
)

func file_tv_service_channel_proto_rawDescGZIP() []byte {
	file_tv_service_channel_proto_rawDescOnce.Do(func() {
		file_tv_service_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_tv_service_channel_proto_rawDescData)
	})
	return file_tv_service_channel_proto_rawDescData
}

var file_tv_service_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tv_service_channel_proto_goTypes = []interface{}{
	(*Channel)(nil),   // 0: tv_service.Channel
	(*EpgRecord)(nil), // 1: tv_service.EpgRecord
}
var file_tv_service_channel_proto_depIdxs = []int32{
	1, // 0: tv_service.Channel.epg:type_name -> tv_service.EpgRecord
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tv_service_channel_proto_init() }
func file_tv_service_channel_proto_init() {
	if File_tv_service_channel_proto != nil {
		return
	}
	file_tv_service_epg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tv_service_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tv_service_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tv_service_channel_proto_goTypes,
		DependencyIndexes: file_tv_service_channel_proto_depIdxs,
		MessageInfos:      file_tv_service_channel_proto_msgTypes,
	}.Build()
	File_tv_service_channel_proto = out.File
	file_tv_service_channel_proto_rawDesc = nil
	file_tv_service_channel_proto_goTypes = nil
	file_tv_service_channel_proto_depIdxs = nil
}
