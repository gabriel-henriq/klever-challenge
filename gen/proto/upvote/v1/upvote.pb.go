// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: proto/upvote/v1/upvote.proto

package upvotev1

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

type UpvoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpvoteRequest) Reset() {
	*x = UpvoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteRequest) ProtoMessage() {}

func (x *UpvoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpvoteRequest.ProtoReflect.Descriptor instead.
func (*UpvoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{0}
}

func (x *UpvoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpvoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Likes  int64  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *UpvoteResponse) Reset() {
	*x = UpvoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteResponse) ProtoMessage() {}

func (x *UpvoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpvoteResponse.ProtoReflect.Descriptor instead.
func (*UpvoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{1}
}

func (x *UpvoteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpvoteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpvoteResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpvoteResponse) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

type ListBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ListBooksRequest) Reset() {
	*x = ListBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksRequest) ProtoMessage() {}

func (x *ListBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksRequest.ProtoReflect.Descriptor instead.
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{2}
}

func (x *ListBooksRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ListBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Likes  int64  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{3}
}

func (x *ListBooksResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListBooksResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListBooksResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ListBooksResponse) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

var File_proto_upvote_v1_upvote_proto protoreflect.FileDescriptor

var file_proto_upvote_v1_upvote_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x0e, 0x55, 0x70,
	0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x67, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x32, 0x9c, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x75, 0x70,
	0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x6c, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x09, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x55, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x09, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x55, 0x70,
	0x76, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_upvote_v1_upvote_proto_rawDescOnce sync.Once
	file_proto_upvote_v1_upvote_proto_rawDescData = file_proto_upvote_v1_upvote_proto_rawDesc
)

func file_proto_upvote_v1_upvote_proto_rawDescGZIP() []byte {
	file_proto_upvote_v1_upvote_proto_rawDescOnce.Do(func() {
		file_proto_upvote_v1_upvote_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_upvote_v1_upvote_proto_rawDescData)
	})
	return file_proto_upvote_v1_upvote_proto_rawDescData
}

var file_proto_upvote_v1_upvote_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_upvote_v1_upvote_proto_goTypes = []interface{}{
	(*UpvoteRequest)(nil),     // 0: upvote.v1.UpvoteRequest
	(*UpvoteResponse)(nil),    // 1: upvote.v1.UpvoteResponse
	(*ListBooksRequest)(nil),  // 2: upvote.v1.ListBooksRequest
	(*ListBooksResponse)(nil), // 3: upvote.v1.ListBooksResponse
}
var file_proto_upvote_v1_upvote_proto_depIdxs = []int32{
	2, // 0: upvote.v1.UpvoteService.ListBooks:input_type -> upvote.v1.ListBooksRequest
	0, // 1: upvote.v1.UpvoteService.Upvote:input_type -> upvote.v1.UpvoteRequest
	3, // 2: upvote.v1.UpvoteService.ListBooks:output_type -> upvote.v1.ListBooksResponse
	1, // 3: upvote.v1.UpvoteService.Upvote:output_type -> upvote.v1.UpvoteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_upvote_v1_upvote_proto_init() }
func file_proto_upvote_v1_upvote_proto_init() {
	if File_proto_upvote_v1_upvote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_upvote_v1_upvote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpvoteRequest); i {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpvoteResponse); i {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksRequest); i {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksResponse); i {
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
			RawDescriptor: file_proto_upvote_v1_upvote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_upvote_v1_upvote_proto_goTypes,
		DependencyIndexes: file_proto_upvote_v1_upvote_proto_depIdxs,
		MessageInfos:      file_proto_upvote_v1_upvote_proto_msgTypes,
	}.Build()
	File_proto_upvote_v1_upvote_proto = out.File
	file_proto_upvote_v1_upvote_proto_rawDesc = nil
	file_proto_upvote_v1_upvote_proto_goTypes = nil
	file_proto_upvote_v1_upvote_proto_depIdxs = nil
}
