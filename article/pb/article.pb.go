// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: article.proto

package pb

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

// やり取りする記事の型定義
type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author  string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_article_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// CREATEとUPDATEにおける入力のための型
type ArticleInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author  string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ArticleInput) Reset() {
	*x = ArticleInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleInput) ProtoMessage() {}

func (x *ArticleInput) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleInput.ProtoReflect.Descriptor instead.
func (*ArticleInput) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleInput) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ArticleInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleInput) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// CREATEする記事の内容(author, title, content)を受け取る
type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleInput *ArticleInput `protobuf:"bytes,1,opt,name=articleInput,proto3" json:"articleInput,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{2}
}

func (x *CreateArticleRequest) GetArticleInput() *ArticleInput {
	if x != nil {
		return x.ArticleInput
	}
	return nil
}

// CREATEした記事を返す
type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{3}
}

func (x *CreateArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

// READする記事のIDを受け取る
type ReadArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadArticleRequest) Reset() {
	*x = ReadArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadArticleRequest) ProtoMessage() {}

func (x *ReadArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadArticleRequest.ProtoReflect.Descriptor instead.
func (*ReadArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{4}
}

func (x *ReadArticleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// READする記事のIDを返す
type ReadArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *ReadArticleResponse) Reset() {
	*x = ReadArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadArticleResponse) ProtoMessage() {}

func (x *ReadArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadArticleResponse.ProtoReflect.Descriptor instead.
func (*ReadArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{5}
}

func (x *ReadArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

// UPDATEする記事のIDと変更内容を受け取る
type UpdateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ArticleInput *ArticleInput `protobuf:"bytes,2,opt,name=articleInput,proto3" json:"articleInput,omitempty"`
}

func (x *UpdateArticleRequest) Reset() {
	*x = UpdateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleRequest) ProtoMessage() {}

func (x *UpdateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateArticleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateArticleRequest) GetArticleInput() *ArticleInput {
	if x != nil {
		return x.ArticleInput
	}
	return nil
}

// UPDATEした記事を返す
type UpdateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *UpdateArticleResponse) Reset() {
	*x = UpdateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleResponse) ProtoMessage() {}

func (x *UpdateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleResponse.ProtoReflect.Descriptor instead.
func (*UpdateArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

// UPDATEする記事のIDを受け取る
type DeleteArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArticleRequest) Reset() {
	*x = DeleteArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleRequest) ProtoMessage() {}

func (x *DeleteArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteArticleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// UPDATEした記事のIDを返す
type DeleteArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArticleResponse) Reset() {
	*x = DeleteArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleResponse) ProtoMessage() {}

func (x *DeleteArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleResponse.ProtoReflect.Descriptor instead.
func (*DeleteArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteArticleResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_article_proto protoreflect.FileDescriptor

var file_article_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x61, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x0c, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0c, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x41, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x22, 0x61, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0c,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x26, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xca, 0x02,
	0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x1b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_article_proto_rawDescOnce sync.Once
	file_article_proto_rawDescData = file_article_proto_rawDesc
)

func file_article_proto_rawDescGZIP() []byte {
	file_article_proto_rawDescOnce.Do(func() {
		file_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_proto_rawDescData)
	})
	return file_article_proto_rawDescData
}

var file_article_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_article_proto_goTypes = []interface{}{
	(*Article)(nil),               // 0: article.Article
	(*ArticleInput)(nil),          // 1: article.ArticleInput
	(*CreateArticleRequest)(nil),  // 2: article.CreateArticleRequest
	(*CreateArticleResponse)(nil), // 3: article.CreateArticleResponse
	(*ReadArticleRequest)(nil),    // 4: article.ReadArticleRequest
	(*ReadArticleResponse)(nil),   // 5: article.ReadArticleResponse
	(*UpdateArticleRequest)(nil),  // 6: article.UpdateArticleRequest
	(*UpdateArticleResponse)(nil), // 7: article.UpdateArticleResponse
	(*DeleteArticleRequest)(nil),  // 8: article.DeleteArticleRequest
	(*DeleteArticleResponse)(nil), // 9: article.DeleteArticleResponse
}
var file_article_proto_depIdxs = []int32{
	1, // 0: article.CreateArticleRequest.articleInput:type_name -> article.ArticleInput
	0, // 1: article.CreateArticleResponse.article:type_name -> article.Article
	0, // 2: article.ReadArticleResponse.article:type_name -> article.Article
	1, // 3: article.UpdateArticleRequest.articleInput:type_name -> article.ArticleInput
	0, // 4: article.UpdateArticleResponse.article:type_name -> article.Article
	2, // 5: article.ArticleService.CreateArticle:input_type -> article.CreateArticleRequest
	4, // 6: article.ArticleService.ReadArticle:input_type -> article.ReadArticleRequest
	6, // 7: article.ArticleService.UpdateArticle:input_type -> article.UpdateArticleRequest
	8, // 8: article.ArticleService.DeleteArticle:input_type -> article.DeleteArticleRequest
	3, // 9: article.ArticleService.CreateArticle:output_type -> article.CreateArticleResponse
	5, // 10: article.ArticleService.ReadArticle:output_type -> article.ReadArticleResponse
	7, // 11: article.ArticleService.UpdateArticle:output_type -> article.UpdateArticleResponse
	9, // 12: article.ArticleService.DeleteArticle:output_type -> article.DeleteArticleResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_article_proto_init() }
func file_article_proto_init() {
	if File_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleInput); i {
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
		file_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadArticleRequest); i {
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
		file_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadArticleResponse); i {
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
		file_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleRequest); i {
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
		file_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleResponse); i {
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
		file_article_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleRequest); i {
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
		file_article_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleResponse); i {
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
			RawDescriptor: file_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_proto_goTypes,
		DependencyIndexes: file_article_proto_depIdxs,
		MessageInfos:      file_article_proto_msgTypes,
	}.Build()
	File_article_proto = out.File
	file_article_proto_rawDesc = nil
	file_article_proto_goTypes = nil
	file_article_proto_depIdxs = nil
}
