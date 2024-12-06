// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: most_likes_articles.proto

package blog

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

type TopArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *TopArticlesRequest) Reset() {
	*x = TopArticlesRequest{}
	mi := &file_most_likes_articles_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopArticlesRequest) ProtoMessage() {}

func (x *TopArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_most_likes_articles_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopArticlesRequest.ProtoReflect.Descriptor instead.
func (*TopArticlesRequest) Descriptor() ([]byte, []int) {
	return file_most_likes_articles_proto_rawDescGZIP(), []int{0}
}

func (x *TopArticlesRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TopArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *TopArticlesResponse) Reset() {
	*x = TopArticlesResponse{}
	mi := &file_most_likes_articles_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopArticlesResponse) ProtoMessage() {}

func (x *TopArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_most_likes_articles_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopArticlesResponse.ProtoReflect.Descriptor instead.
func (*TopArticlesResponse) Descriptor() ([]byte, []int) {
	return file_most_likes_articles_proto_rawDescGZIP(), []int{1}
}

func (x *TopArticlesResponse) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Likes int32  `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	mi := &file_most_likes_articles_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_most_likes_articles_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_most_likes_articles_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

var File_most_likes_articles_proto protoreflect.FileDescriptor

var file_most_likes_articles_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x43, 0x0a, 0x13, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x32, 0x5d, 0x0a, 0x0e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x1b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54, 0x6f, 0x70, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x61, 0x6d, 0x69, 0x72, 0x75,
	0x70, 0x2f, 0x77, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_most_likes_articles_proto_rawDescOnce sync.Once
	file_most_likes_articles_proto_rawDescData = file_most_likes_articles_proto_rawDesc
)

func file_most_likes_articles_proto_rawDescGZIP() []byte {
	file_most_likes_articles_proto_rawDescOnce.Do(func() {
		file_most_likes_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_most_likes_articles_proto_rawDescData)
	})
	return file_most_likes_articles_proto_rawDescData
}

var file_most_likes_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_most_likes_articles_proto_goTypes = []any{
	(*TopArticlesRequest)(nil),  // 0: article.TopArticlesRequest
	(*TopArticlesResponse)(nil), // 1: article.TopArticlesResponse
	(*Article)(nil),             // 2: article.Article
}
var file_most_likes_articles_proto_depIdxs = []int32{
	2, // 0: article.TopArticlesResponse.articles:type_name -> article.Article
	0, // 1: article.ArticleService.GetTopArticles:input_type -> article.TopArticlesRequest
	1, // 2: article.ArticleService.GetTopArticles:output_type -> article.TopArticlesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_most_likes_articles_proto_init() }
func file_most_likes_articles_proto_init() {
	if File_most_likes_articles_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_most_likes_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_most_likes_articles_proto_goTypes,
		DependencyIndexes: file_most_likes_articles_proto_depIdxs,
		MessageInfos:      file_most_likes_articles_proto_msgTypes,
	}.Build()
	File_most_likes_articles_proto = out.File
	file_most_likes_articles_proto_rawDesc = nil
	file_most_likes_articles_proto_goTypes = nil
	file_most_likes_articles_proto_depIdxs = nil
}
