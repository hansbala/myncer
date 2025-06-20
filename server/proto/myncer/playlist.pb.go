// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: myncer/playlist.proto

package myncer_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Playlist struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	MusicSource *MusicSource           `protobuf:"bytes,1,opt,name=music_source,json=musicSource,proto3" json:"music_source,omitempty"`
	// Human readable name of the playlist as stored in the original datasource.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Human readable description of the playlist as stored in the original datasource.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// URL to the playlist image as stored in the original datasource.
	ImageUrl      string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"` // next: 5
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Playlist) Reset() {
	*x = Playlist{}
	mi := &file_myncer_playlist_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Playlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playlist) ProtoMessage() {}

func (x *Playlist) ProtoReflect() protoreflect.Message {
	mi := &file_myncer_playlist_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playlist.ProtoReflect.Descriptor instead.
func (*Playlist) Descriptor() ([]byte, []int) {
	return file_myncer_playlist_proto_rawDescGZIP(), []int{0}
}

func (x *Playlist) GetMusicSource() *MusicSource {
	if x != nil {
		return x.MusicSource
	}
	return nil
}

func (x *Playlist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Playlist) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Playlist) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_myncer_playlist_proto protoreflect.FileDescriptor

var file_myncer_playlist_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6d, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x1a,
	0x11, 0x6d, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x0c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6e, 0x73, 0x62, 0x61, 0x6c,
	0x61, 0x2f, 0x6d, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x79, 0x6e, 0x63, 0x65, 0x72, 0x3b, 0x6d, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_myncer_playlist_proto_rawDescOnce sync.Once
	file_myncer_playlist_proto_rawDescData []byte
)

func file_myncer_playlist_proto_rawDescGZIP() []byte {
	file_myncer_playlist_proto_rawDescOnce.Do(func() {
		file_myncer_playlist_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_myncer_playlist_proto_rawDesc), len(file_myncer_playlist_proto_rawDesc)))
	})
	return file_myncer_playlist_proto_rawDescData
}

var file_myncer_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_myncer_playlist_proto_goTypes = []any{
	(*Playlist)(nil),    // 0: myncer.Playlist
	(*MusicSource)(nil), // 1: myncer.MusicSource
}
var file_myncer_playlist_proto_depIdxs = []int32{
	1, // 0: myncer.Playlist.music_source:type_name -> myncer.MusicSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_myncer_playlist_proto_init() }
func file_myncer_playlist_proto_init() {
	if File_myncer_playlist_proto != nil {
		return
	}
	file_myncer_sync_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_myncer_playlist_proto_rawDesc), len(file_myncer_playlist_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_myncer_playlist_proto_goTypes,
		DependencyIndexes: file_myncer_playlist_proto_depIdxs,
		MessageInfos:      file_myncer_playlist_proto_msgTypes,
	}.Build()
	File_myncer_playlist_proto = out.File
	file_myncer_playlist_proto_goTypes = nil
	file_myncer_playlist_proto_depIdxs = nil
}
