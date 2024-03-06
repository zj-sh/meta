// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/suite/suite_source_pull.proto

package suite

import (
	base "github.com/zj-sh/meta/v1/base"
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

type SourcePullCloneParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Dir    string           `protobuf:"bytes,2,opt,name=dir,proto3" json:"dir,omitempty"`
	Url    string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Key    string           `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Branch string           `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	Tag    string           `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	Commit string           `protobuf:"bytes,7,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *SourcePullCloneParam) Reset() {
	*x = SourcePullCloneParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_source_pull_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePullCloneParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePullCloneParam) ProtoMessage() {}

func (x *SourcePullCloneParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_source_pull_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePullCloneParam.ProtoReflect.Descriptor instead.
func (*SourcePullCloneParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_source_pull_proto_rawDescGZIP(), []int{0}
}

func (x *SourcePullCloneParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *SourcePullCloneParam) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *SourcePullCloneParam) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SourcePullCloneParam) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SourcePullCloneParam) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *SourcePullCloneParam) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *SourcePullCloneParam) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

type SourcePullCreateParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Dir    string           `protobuf:"bytes,2,opt,name=dir,proto3" json:"dir,omitempty"`
	Url    string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Key    string           `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Branch string           `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	Tag    string           `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	Commit string           `protobuf:"bytes,7,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *SourcePullCreateParam) Reset() {
	*x = SourcePullCreateParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_source_pull_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePullCreateParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePullCreateParam) ProtoMessage() {}

func (x *SourcePullCreateParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_source_pull_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePullCreateParam.ProtoReflect.Descriptor instead.
func (*SourcePullCreateParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_source_pull_proto_rawDescGZIP(), []int{1}
}

func (x *SourcePullCreateParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *SourcePullCreateParam) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *SourcePullCreateParam) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SourcePullCreateParam) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SourcePullCreateParam) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *SourcePullCreateParam) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *SourcePullCreateParam) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

type SourcePullMergeParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Dir    string           `protobuf:"bytes,2,opt,name=dir,proto3" json:"dir,omitempty"`
	Url    string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Key    string           `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	From   string           `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To     string           `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *SourcePullMergeParam) Reset() {
	*x = SourcePullMergeParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_source_pull_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePullMergeParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePullMergeParam) ProtoMessage() {}

func (x *SourcePullMergeParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_source_pull_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePullMergeParam.ProtoReflect.Descriptor instead.
func (*SourcePullMergeParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_source_pull_proto_rawDescGZIP(), []int{2}
}

func (x *SourcePullMergeParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *SourcePullMergeParam) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *SourcePullMergeParam) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SourcePullMergeParam) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SourcePullMergeParam) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SourcePullMergeParam) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type SourcePullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  base.RunStatus `protobuf:"varint,1,opt,name=status,proto3,enum=meta.v1.base.RunStatus" json:"status,omitempty"`
	Log     *base.Log      `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	Message string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Outs    *OutResponse   `protobuf:"bytes,4,opt,name=outs,proto3" json:"outs,omitempty"`
}

func (x *SourcePullResponse) Reset() {
	*x = SourcePullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_source_pull_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePullResponse) ProtoMessage() {}

func (x *SourcePullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_source_pull_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePullResponse.ProtoReflect.Descriptor instead.
func (*SourcePullResponse) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_source_pull_proto_rawDescGZIP(), []int{3}
}

func (x *SourcePullResponse) GetStatus() base.RunStatus {
	if x != nil {
		return x.Status
	}
	return base.RunStatus(0)
}

func (x *SourcePullResponse) GetLog() *base.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *SourcePullResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SourcePullResponse) GetOuts() *OutResponse {
	if x != nil {
		return x.Outs
	}
	return nil
}

var File_v1_suite_suite_source_pull_proto protoreflect.FileDescriptor

var file_v1_suite_suite_source_pull_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74,
	0x65, 0x1a, 0x12, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76,
	0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x14, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0xc1, 0x01,
	0x0a, 0x15, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x75, 0x6c, 0x6c,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x12, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6f, 0x75, 0x74, 0x73, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73,
	0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_suite_suite_source_pull_proto_rawDescOnce sync.Once
	file_v1_suite_suite_source_pull_proto_rawDescData = file_v1_suite_suite_source_pull_proto_rawDesc
)

func file_v1_suite_suite_source_pull_proto_rawDescGZIP() []byte {
	file_v1_suite_suite_source_pull_proto_rawDescOnce.Do(func() {
		file_v1_suite_suite_source_pull_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_suite_suite_source_pull_proto_rawDescData)
	})
	return file_v1_suite_suite_source_pull_proto_rawDescData
}

var file_v1_suite_suite_source_pull_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_suite_suite_source_pull_proto_goTypes = []interface{}{
	(*SourcePullCloneParam)(nil),  // 0: meta.v1.suite.SourcePullCloneParam
	(*SourcePullCreateParam)(nil), // 1: meta.v1.suite.SourcePullCreateParam
	(*SourcePullMergeParam)(nil),  // 2: meta.v1.suite.SourcePullMergeParam
	(*SourcePullResponse)(nil),    // 3: meta.v1.suite.SourcePullResponse
	(*base.JsonString)(nil),       // 4: meta.v1.base.JsonString
	(base.RunStatus)(0),           // 5: meta.v1.base.RunStatus
	(*base.Log)(nil),              // 6: meta.v1.base.Log
	(*OutResponse)(nil),           // 7: meta.v1.suite.OutResponse
}
var file_v1_suite_suite_source_pull_proto_depIdxs = []int32{
	4, // 0: meta.v1.suite.SourcePullCloneParam.params:type_name -> meta.v1.base.JsonString
	4, // 1: meta.v1.suite.SourcePullCreateParam.params:type_name -> meta.v1.base.JsonString
	4, // 2: meta.v1.suite.SourcePullMergeParam.params:type_name -> meta.v1.base.JsonString
	5, // 3: meta.v1.suite.SourcePullResponse.status:type_name -> meta.v1.base.RunStatus
	6, // 4: meta.v1.suite.SourcePullResponse.log:type_name -> meta.v1.base.Log
	7, // 5: meta.v1.suite.SourcePullResponse.outs:type_name -> meta.v1.suite.OutResponse
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_suite_suite_source_pull_proto_init() }
func file_v1_suite_suite_source_pull_proto_init() {
	if File_v1_suite_suite_source_pull_proto != nil {
		return
	}
	file_v1_suite_meta_running_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_suite_suite_source_pull_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourcePullCloneParam); i {
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
		file_v1_suite_suite_source_pull_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourcePullCreateParam); i {
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
		file_v1_suite_suite_source_pull_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourcePullMergeParam); i {
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
		file_v1_suite_suite_source_pull_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourcePullResponse); i {
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
			RawDescriptor: file_v1_suite_suite_source_pull_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_suite_suite_source_pull_proto_goTypes,
		DependencyIndexes: file_v1_suite_suite_source_pull_proto_depIdxs,
		MessageInfos:      file_v1_suite_suite_source_pull_proto_msgTypes,
	}.Build()
	File_v1_suite_suite_source_pull_proto = out.File
	file_v1_suite_suite_source_pull_proto_rawDesc = nil
	file_v1_suite_suite_source_pull_proto_goTypes = nil
	file_v1_suite_suite_source_pull_proto_depIdxs = nil
}
