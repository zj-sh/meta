// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/suite/suite_storage.proto

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

type StorageType int32

const (
	StorageType_STORAGE_TYPE_UNKNOWN StorageType = 0
	StorageType_STORAGE_TYPE_LOCAL   StorageType = 1
	StorageType_STORAGE_TYPE_S3      StorageType = 10
	StorageType_STORAGE_TYPE_OSS     StorageType = 11
	StorageType_STORAGE_TYPE_COS     StorageType = 12
	StorageType_STORAGE_TYPE_SFS     StorageType = 13
	StorageType_STORAGE_TYPE_QINIU   StorageType = 14
	StorageType_STORAGE_TYPE_PV      StorageType = 20
)

// Enum value maps for StorageType.
var (
	StorageType_name = map[int32]string{
		0:  "STORAGE_TYPE_UNKNOWN",
		1:  "STORAGE_TYPE_LOCAL",
		10: "STORAGE_TYPE_S3",
		11: "STORAGE_TYPE_OSS",
		12: "STORAGE_TYPE_COS",
		13: "STORAGE_TYPE_SFS",
		14: "STORAGE_TYPE_QINIU",
		20: "STORAGE_TYPE_PV",
	}
	StorageType_value = map[string]int32{
		"STORAGE_TYPE_UNKNOWN": 0,
		"STORAGE_TYPE_LOCAL":   1,
		"STORAGE_TYPE_S3":      10,
		"STORAGE_TYPE_OSS":     11,
		"STORAGE_TYPE_COS":     12,
		"STORAGE_TYPE_SFS":     13,
		"STORAGE_TYPE_QINIU":   14,
		"STORAGE_TYPE_PV":      20,
	}
)

func (x StorageType) Enum() *StorageType {
	p := new(StorageType)
	*p = x
	return p
}

func (x StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_suite_suite_storage_proto_enumTypes[0].Descriptor()
}

func (StorageType) Type() protoreflect.EnumType {
	return &file_v1_suite_suite_storage_proto_enumTypes[0]
}

func (x StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageType.Descriptor instead.
func (StorageType) EnumDescriptor() ([]byte, []int) {
	return file_v1_suite_suite_storage_proto_rawDescGZIP(), []int{0}
}

type StorageUploadStreamParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params   *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Filename string           `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Filesize int64            `protobuf:"varint,3,opt,name=filesize,proto3" json:"filesize,omitempty"`
	Md5      string           `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	Size     int64            `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Byte     []byte           `protobuf:"bytes,6,opt,name=byte,proto3" json:"byte,omitempty"`
	Done     bool             `protobuf:"varint,7,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *StorageUploadStreamParam) Reset() {
	*x = StorageUploadStreamParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageUploadStreamParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageUploadStreamParam) ProtoMessage() {}

func (x *StorageUploadStreamParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageUploadStreamParam.ProtoReflect.Descriptor instead.
func (*StorageUploadStreamParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_storage_proto_rawDescGZIP(), []int{0}
}

func (x *StorageUploadStreamParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *StorageUploadStreamParam) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *StorageUploadStreamParam) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

func (x *StorageUploadStreamParam) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *StorageUploadStreamParam) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StorageUploadStreamParam) GetByte() []byte {
	if x != nil {
		return x.Byte
	}
	return nil
}

func (x *StorageUploadStreamParam) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type StorageUploadFileParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params   *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Filename string           `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Filepath string           `protobuf:"bytes,3,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *StorageUploadFileParam) Reset() {
	*x = StorageUploadFileParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageUploadFileParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageUploadFileParam) ProtoMessage() {}

func (x *StorageUploadFileParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageUploadFileParam.ProtoReflect.Descriptor instead.
func (*StorageUploadFileParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StorageUploadFileParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *StorageUploadFileParam) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *StorageUploadFileParam) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

type StorageUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      base.RunStatus `protobuf:"varint,1,opt,name=status,proto3,enum=meta.v1.base.RunStatus" json:"status,omitempty"`
	Fileid      string         `protobuf:"bytes,2,opt,name=fileid,proto3" json:"fileid,omitempty"`
	Filepath    string         `protobuf:"bytes,3,opt,name=filepath,proto3" json:"filepath,omitempty"`
	StorageType StorageType    `protobuf:"varint,4,opt,name=storage_type,json=storageType,proto3,enum=meta.v1.suite.StorageType" json:"storage_type,omitempty"`
	Message     string         `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Outs        *OutResponse   `protobuf:"bytes,6,opt,name=outs,proto3" json:"outs,omitempty"`
}

func (x *StorageUploadResponse) Reset() {
	*x = StorageUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageUploadResponse) ProtoMessage() {}

func (x *StorageUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageUploadResponse.ProtoReflect.Descriptor instead.
func (*StorageUploadResponse) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_storage_proto_rawDescGZIP(), []int{2}
}

func (x *StorageUploadResponse) GetStatus() base.RunStatus {
	if x != nil {
		return x.Status
	}
	return base.RunStatus(0)
}

func (x *StorageUploadResponse) GetFileid() string {
	if x != nil {
		return x.Fileid
	}
	return ""
}

func (x *StorageUploadResponse) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *StorageUploadResponse) GetStorageType() StorageType {
	if x != nil {
		return x.StorageType
	}
	return StorageType_STORAGE_TYPE_UNKNOWN
}

func (x *StorageUploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StorageUploadResponse) GetOuts() *OutResponse {
	if x != nil {
		return x.Outs
	}
	return nil
}

type StorageDeleteParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params   *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Fileid   string           `protobuf:"bytes,2,opt,name=fileid,proto3" json:"fileid,omitempty"`
	Filepath string           `protobuf:"bytes,3,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *StorageDeleteParam) Reset() {
	*x = StorageDeleteParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageDeleteParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageDeleteParam) ProtoMessage() {}

func (x *StorageDeleteParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageDeleteParam.ProtoReflect.Descriptor instead.
func (*StorageDeleteParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_storage_proto_rawDescGZIP(), []int{3}
}

func (x *StorageDeleteParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *StorageDeleteParam) GetFileid() string {
	if x != nil {
		return x.Fileid
	}
	return ""
}

func (x *StorageDeleteParam) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

var File_v1_suite_suite_storage_proto protoreflect.FileDescriptor

var file_v1_suite_suite_storage_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x12, 0x76,
	0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2,
	0x01, 0x0a, 0x18, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x79, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x79, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0x85, 0x02, 0x0a, 0x15, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x4f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6f, 0x75, 0x74, 0x73,
	0x22, 0x7a, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2a, 0xc3, 0x01, 0x0a,
	0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x33, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4f, 0x53, 0x53, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x4f,
	0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x53, 0x10, 0x0c, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x46, 0x53, 0x10, 0x0d, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x49, 0x4e, 0x49, 0x55, 0x10, 0x0e, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x56,
	0x10, 0x14, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x75, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_suite_suite_storage_proto_rawDescOnce sync.Once
	file_v1_suite_suite_storage_proto_rawDescData = file_v1_suite_suite_storage_proto_rawDesc
)

func file_v1_suite_suite_storage_proto_rawDescGZIP() []byte {
	file_v1_suite_suite_storage_proto_rawDescOnce.Do(func() {
		file_v1_suite_suite_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_suite_suite_storage_proto_rawDescData)
	})
	return file_v1_suite_suite_storage_proto_rawDescData
}

var file_v1_suite_suite_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_suite_suite_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_suite_suite_storage_proto_goTypes = []interface{}{
	(StorageType)(0),                 // 0: meta.v1.suite.StorageType
	(*StorageUploadStreamParam)(nil), // 1: meta.v1.suite.StorageUploadStreamParam
	(*StorageUploadFileParam)(nil),   // 2: meta.v1.suite.StorageUploadFileParam
	(*StorageUploadResponse)(nil),    // 3: meta.v1.suite.StorageUploadResponse
	(*StorageDeleteParam)(nil),       // 4: meta.v1.suite.StorageDeleteParam
	(*base.JsonString)(nil),          // 5: meta.v1.base.JsonString
	(base.RunStatus)(0),              // 6: meta.v1.base.RunStatus
	(*OutResponse)(nil),              // 7: meta.v1.suite.OutResponse
}
var file_v1_suite_suite_storage_proto_depIdxs = []int32{
	5, // 0: meta.v1.suite.StorageUploadStreamParam.params:type_name -> meta.v1.base.JsonString
	5, // 1: meta.v1.suite.StorageUploadFileParam.params:type_name -> meta.v1.base.JsonString
	6, // 2: meta.v1.suite.StorageUploadResponse.status:type_name -> meta.v1.base.RunStatus
	0, // 3: meta.v1.suite.StorageUploadResponse.storage_type:type_name -> meta.v1.suite.StorageType
	7, // 4: meta.v1.suite.StorageUploadResponse.outs:type_name -> meta.v1.suite.OutResponse
	5, // 5: meta.v1.suite.StorageDeleteParam.params:type_name -> meta.v1.base.JsonString
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_suite_suite_storage_proto_init() }
func file_v1_suite_suite_storage_proto_init() {
	if File_v1_suite_suite_storage_proto != nil {
		return
	}
	file_v1_suite_meta_running_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_suite_suite_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageUploadStreamParam); i {
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
		file_v1_suite_suite_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageUploadFileParam); i {
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
		file_v1_suite_suite_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageUploadResponse); i {
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
		file_v1_suite_suite_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageDeleteParam); i {
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
			RawDescriptor: file_v1_suite_suite_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_suite_suite_storage_proto_goTypes,
		DependencyIndexes: file_v1_suite_suite_storage_proto_depIdxs,
		EnumInfos:         file_v1_suite_suite_storage_proto_enumTypes,
		MessageInfos:      file_v1_suite_suite_storage_proto_msgTypes,
	}.Build()
	File_v1_suite_suite_storage_proto = out.File
	file_v1_suite_suite_storage_proto_rawDesc = nil
	file_v1_suite_suite_storage_proto_goTypes = nil
	file_v1_suite_suite_storage_proto_depIdxs = nil
}
