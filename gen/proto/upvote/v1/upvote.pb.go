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

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *CreateBookResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type DownvoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *DownvoteRequest) Reset() {
	*x = DownvoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownvoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownvoteRequest) ProtoMessage() {}

func (x *DownvoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownvoteRequest.ProtoReflect.Descriptor instead.
func (*DownvoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{2}
}

func (x *DownvoteRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type DownvoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId  string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author  string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Likes   int64  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
	Unlikes int64  `protobuf:"varint,5,opt,name=unlikes,proto3" json:"unlikes,omitempty"`
}

func (x *DownvoteResponse) Reset() {
	*x = DownvoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownvoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownvoteResponse) ProtoMessage() {}

func (x *DownvoteResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownvoteResponse.ProtoReflect.Descriptor instead.
func (*DownvoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{3}
}

func (x *DownvoteResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *DownvoteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DownvoteResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *DownvoteResponse) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *DownvoteResponse) GetUnlikes() int64 {
	if x != nil {
		return x.Unlikes
	}
	return 0
}

type UpvoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *UpvoteRequest) Reset() {
	*x = UpvoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteRequest) ProtoMessage() {}

func (x *UpvoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[4]
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
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{4}
}

func (x *UpvoteRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type UpvoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId  string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author  string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Likes   int64  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
	Unlikes int64  `protobuf:"varint,5,opt,name=unlikes,proto3" json:"unlikes,omitempty"`
}

func (x *UpvoteResponse) Reset() {
	*x = UpvoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteResponse) ProtoMessage() {}

func (x *UpvoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[5]
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
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{5}
}

func (x *UpvoteResponse) GetBookId() string {
	if x != nil {
		return x.BookId
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

func (x *UpvoteResponse) GetUnlikes() int64 {
	if x != nil {
		return x.Unlikes
	}
	return 0
}

type WatchBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *WatchBookRequest) Reset() {
	*x = WatchBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBookRequest) ProtoMessage() {}

func (x *WatchBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBookRequest.ProtoReflect.Descriptor instead.
func (*WatchBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{6}
}

func (x *WatchBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type WatchBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId  string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author  string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Likes   int64  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
	Unlikes int64  `protobuf:"varint,5,opt,name=unlikes,proto3" json:"unlikes,omitempty"`
}

func (x *WatchBookResponse) Reset() {
	*x = WatchBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_upvote_v1_upvote_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBookResponse) ProtoMessage() {}

func (x *WatchBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upvote_v1_upvote_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBookResponse.ProtoReflect.Descriptor instead.
func (*WatchBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_upvote_v1_upvote_proto_rawDescGZIP(), []int{7}
}

func (x *WatchBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *WatchBookResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WatchBookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *WatchBookResponse) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *WatchBookResponse) GetUnlikes() int64 {
	if x != nil {
		return x.Unlikes
	}
	return 0
}

var File_proto_upvote_v1_upvote_proto protoreflect.FileDescriptor

var file_proto_upvote_v1_upvote_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x5b, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x0f, 0x44, 0x6f, 0x77,
	0x6e, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x6e, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x22, 0x28, 0x0a, 0x0d, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x6e,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x8a, 0x01, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x32, 0xb0, 0x02, 0x0a,
	0x0d, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x75,
	0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x55,
	0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08,
	0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1b, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x6c, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x09, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x55,
	0x70, 0x76, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x55, 0x70, 0x76, 0x6f, 0x74,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_upvote_v1_upvote_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_upvote_v1_upvote_proto_goTypes = []interface{}{
	(*CreateBookRequest)(nil),  // 0: upvote.v1.CreateBookRequest
	(*CreateBookResponse)(nil), // 1: upvote.v1.CreateBookResponse
	(*DownvoteRequest)(nil),    // 2: upvote.v1.DownvoteRequest
	(*DownvoteResponse)(nil),   // 3: upvote.v1.DownvoteResponse
	(*UpvoteRequest)(nil),      // 4: upvote.v1.UpvoteRequest
	(*UpvoteResponse)(nil),     // 5: upvote.v1.UpvoteResponse
	(*WatchBookRequest)(nil),   // 6: upvote.v1.WatchBookRequest
	(*WatchBookResponse)(nil),  // 7: upvote.v1.WatchBookResponse
}
var file_proto_upvote_v1_upvote_proto_depIdxs = []int32{
	0, // 0: upvote.v1.UpvoteService.CreateBook:input_type -> upvote.v1.CreateBookRequest
	4, // 1: upvote.v1.UpvoteService.Upvote:input_type -> upvote.v1.UpvoteRequest
	2, // 2: upvote.v1.UpvoteService.Downvote:input_type -> upvote.v1.DownvoteRequest
	6, // 3: upvote.v1.UpvoteService.WatchBook:input_type -> upvote.v1.WatchBookRequest
	1, // 4: upvote.v1.UpvoteService.CreateBook:output_type -> upvote.v1.CreateBookResponse
	5, // 5: upvote.v1.UpvoteService.Upvote:output_type -> upvote.v1.UpvoteResponse
	3, // 6: upvote.v1.UpvoteService.Downvote:output_type -> upvote.v1.DownvoteResponse
	7, // 7: upvote.v1.UpvoteService.WatchBook:output_type -> upvote.v1.WatchBookResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
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
			switch v := v.(*CreateBookRequest); i {
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
			switch v := v.(*CreateBookResponse); i {
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
			switch v := v.(*DownvoteRequest); i {
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
			switch v := v.(*DownvoteResponse); i {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchBookRequest); i {
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
		file_proto_upvote_v1_upvote_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchBookResponse); i {
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
			NumMessages:   8,
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
