// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ocis/services/search/v0/search.proto

package v0

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v0 "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/search/v0"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The maximum number of entries to return in the response
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A pagination token returned from a previous call to `Get`
	// that indicates from where search should continue
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Query     string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_search_v0_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_search_v0_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_ocis_services_search_v0_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*v0.Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalMatches  int32  `protobuf:"varint,3,opt,name=total_matches,json=totalMatches,proto3" json:"total_matches,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_search_v0_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_search_v0_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_ocis_services_search_v0_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetMatches() []*v0.Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *SearchResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *SearchResponse) GetTotalMatches() int32 {
	if x != nil {
		return x.TotalMatches
	}
	return 0
}

type SearchIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The maximum number of entries to return in the response
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A pagination token returned from a previous call to `Get`
	// that indicates from where search should continue
	PageToken string        `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Query     string        `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	Ref       *v0.Reference `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *SearchIndexRequest) Reset() {
	*x = SearchIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_search_v0_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchIndexRequest) ProtoMessage() {}

func (x *SearchIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_search_v0_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchIndexRequest.ProtoReflect.Descriptor instead.
func (*SearchIndexRequest) Descriptor() ([]byte, []int) {
	return file_ocis_services_search_v0_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchIndexRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchIndexRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *SearchIndexRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchIndexRequest) GetRef() *v0.Reference {
	if x != nil {
		return x.Ref
	}
	return nil
}

type SearchIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*v0.Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalMatches  int32  `protobuf:"varint,3,opt,name=total_matches,json=totalMatches,proto3" json:"total_matches,omitempty"`
}

func (x *SearchIndexResponse) Reset() {
	*x = SearchIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_search_v0_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchIndexResponse) ProtoMessage() {}

func (x *SearchIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_search_v0_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchIndexResponse.ProtoReflect.Descriptor instead.
func (*SearchIndexResponse) Descriptor() ([]byte, []int) {
	return file_ocis_services_search_v0_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchIndexResponse) GetMatches() []*v0.Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *SearchIndexResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *SearchIndexResponse) GetTotalMatches() int32 {
	if x != nil {
		return x.TotalMatches
	}
	return 0
}

type IndexSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IndexSpaceRequest) Reset() {
	*x = IndexSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_search_v0_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSpaceRequest) ProtoMessage() {}

func (x *IndexSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_search_v0_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSpaceRequest.ProtoReflect.Descriptor instead.
func (*IndexSpaceRequest) Descriptor() ([]byte, []int) {
	return file_ocis_services_search_v0_search_proto_rawDescGZIP(), []int{4}
}

func (x *IndexSpaceRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *IndexSpaceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type IndexSpaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexSpaceResponse) Reset() {
	*x = IndexSpaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_services_search_v0_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSpaceResponse) ProtoMessage() {}

func (x *IndexSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_services_search_v0_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSpaceResponse.ProtoReflect.Descriptor instead.
func (*IndexSpaceResponse) Descriptor() ([]byte, []int) {
	return file_ocis_services_search_v0_search_proto_rawDescGZIP(), []int{5}
}

var File_ocis_services_search_v0_search_proto protoreflect.FileDescriptor

var file_ocis_services_search_v0_search_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x1a,
	0x24, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x97, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x63, 0x69, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22,
	0xae, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x03, 0x72, 0x65, 0x66,
	0x22, 0x9c, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x63, 0x69, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22,
	0x47, 0x0a, 0x11, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9c,
	0x02, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x7b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x26, 0x2e, 0x6f, 0x63,
	0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x8c,
	0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2a, 0x2e,
	0x6f, 0x63, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x63, 0x69, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x32, 0x9d, 0x01,
	0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x8b, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x2e, 0x6f, 0x63, 0x69,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76,
	0x30, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x42, 0xdc, 0x02,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x30, 0x92, 0x41,
	0x9a, 0x02, 0x12, 0xb4, 0x01, 0x0a, 0x1e, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20,
	0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x20, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x20, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x47, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x20, 0x47, 0x6d, 0x62, 0x48, 0x12, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x1a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x40, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x42,
	0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2d, 0x32, 0x2e, 0x30, 0x12, 0x34, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x62,
	0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x72, 0x39, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x20, 0x4d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6f,
	0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocis_services_search_v0_search_proto_rawDescOnce sync.Once
	file_ocis_services_search_v0_search_proto_rawDescData = file_ocis_services_search_v0_search_proto_rawDesc
)

func file_ocis_services_search_v0_search_proto_rawDescGZIP() []byte {
	file_ocis_services_search_v0_search_proto_rawDescOnce.Do(func() {
		file_ocis_services_search_v0_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocis_services_search_v0_search_proto_rawDescData)
	})
	return file_ocis_services_search_v0_search_proto_rawDescData
}

var file_ocis_services_search_v0_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ocis_services_search_v0_search_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),       // 0: ocis.services.search.v0.SearchRequest
	(*SearchResponse)(nil),      // 1: ocis.services.search.v0.SearchResponse
	(*SearchIndexRequest)(nil),  // 2: ocis.services.search.v0.SearchIndexRequest
	(*SearchIndexResponse)(nil), // 3: ocis.services.search.v0.SearchIndexResponse
	(*IndexSpaceRequest)(nil),   // 4: ocis.services.search.v0.IndexSpaceRequest
	(*IndexSpaceResponse)(nil),  // 5: ocis.services.search.v0.IndexSpaceResponse
	(*v0.Match)(nil),            // 6: ocis.messages.search.v0.Match
	(*v0.Reference)(nil),        // 7: ocis.messages.search.v0.Reference
}
var file_ocis_services_search_v0_search_proto_depIdxs = []int32{
	6, // 0: ocis.services.search.v0.SearchResponse.matches:type_name -> ocis.messages.search.v0.Match
	7, // 1: ocis.services.search.v0.SearchIndexRequest.ref:type_name -> ocis.messages.search.v0.Reference
	6, // 2: ocis.services.search.v0.SearchIndexResponse.matches:type_name -> ocis.messages.search.v0.Match
	0, // 3: ocis.services.search.v0.SearchProvider.Search:input_type -> ocis.services.search.v0.SearchRequest
	4, // 4: ocis.services.search.v0.SearchProvider.IndexSpace:input_type -> ocis.services.search.v0.IndexSpaceRequest
	2, // 5: ocis.services.search.v0.IndexProvider.Search:input_type -> ocis.services.search.v0.SearchIndexRequest
	1, // 6: ocis.services.search.v0.SearchProvider.Search:output_type -> ocis.services.search.v0.SearchResponse
	5, // 7: ocis.services.search.v0.SearchProvider.IndexSpace:output_type -> ocis.services.search.v0.IndexSpaceResponse
	3, // 8: ocis.services.search.v0.IndexProvider.Search:output_type -> ocis.services.search.v0.SearchIndexResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ocis_services_search_v0_search_proto_init() }
func file_ocis_services_search_v0_search_proto_init() {
	if File_ocis_services_search_v0_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocis_services_search_v0_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_ocis_services_search_v0_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_ocis_services_search_v0_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchIndexRequest); i {
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
		file_ocis_services_search_v0_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchIndexResponse); i {
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
		file_ocis_services_search_v0_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSpaceRequest); i {
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
		file_ocis_services_search_v0_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSpaceResponse); i {
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
			RawDescriptor: file_ocis_services_search_v0_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_ocis_services_search_v0_search_proto_goTypes,
		DependencyIndexes: file_ocis_services_search_v0_search_proto_depIdxs,
		MessageInfos:      file_ocis_services_search_v0_search_proto_msgTypes,
	}.Build()
	File_ocis_services_search_v0_search_proto = out.File
	file_ocis_services_search_v0_search_proto_rawDesc = nil
	file_ocis_services_search_v0_search_proto_goTypes = nil
	file_ocis_services_search_v0_search_proto_depIdxs = nil
}
