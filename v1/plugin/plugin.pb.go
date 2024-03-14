// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/plugin/plugin.proto

package plugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_UNKNOWN   Type = 0
	Type_SHELL     Type = 1
	Type_EXECFILE  Type = 2
	Type_CONTAINER Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "SHELL",
		2: "EXECFILE",
		3: "CONTAINER",
	}
	Type_value = map[string]int32{
		"UNKNOWN":   0,
		"SHELL":     1,
		"EXECFILE":  2,
		"CONTAINER": 3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_plugin_plugin_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_v1_plugin_plugin_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name" yaml:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	// @gotags: json:"fullName" yaml:"fullName"
	FullName string `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName" yaml:"fullName"`
	// @gotags: json:"version" yaml:"version"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" yaml:"version"`
	// @gotags: json:"type" yaml:"type"
	Type Type `protobuf:"varint,4,opt,name=type,proto3,enum=meta.v1.plugin.Type" json:"type" yaml:"type"`
	// @gotags: json:"prescript,omitempty" yaml:"prescript,omitempty"
	Prescript string `protobuf:"bytes,5,opt,name=prescript,proto3" json:"prescript,omitempty" yaml:"prescript,omitempty"`
	// @gotags: json:"image,omitempty" yaml:"image,omitempty"
	Image *Image `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty" yaml:"image,omitempty"`
	// @gotags: json:"effects,omitempty" yaml:"effects,omitempty"
	Effects []int64 `protobuf:"varint,7,rep,packed,name=effects,proto3" json:"effects,omitempty" yaml:"effects,omitempty"`
	// @gotags: json:"support,omitempty" yaml:"support,omitempty"
	Support *Support `protobuf:"bytes,8,opt,name=support,proto3" json:"support,omitempty" yaml:"support,omitempty"`
	// @gotags: json:"metadata,omitempty" yaml:"metadata,omitempty"
	Metadata *Metadata `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// @gotags: json:"repo,omitempty" yaml:"repo,omitempty"
	Repo string `protobuf:"bytes,10,opt,name=repo,proto3" json:"repo,omitempty" yaml:"repo,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Plugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugin) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_UNKNOWN
}

func (x *Plugin) GetPrescript() string {
	if x != nil {
		return x.Prescript
	}
	return ""
}

func (x *Plugin) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Plugin) GetEffects() []int64 {
	if x != nil {
		return x.Effects
	}
	return nil
}

func (x *Plugin) GetSupport() *Support {
	if x != nil {
		return x.Support
	}
	return nil
}

func (x *Plugin) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Plugin) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name,omitempty" yaml:"name,omitempty"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	// @gotags: json:"tag,omitempty" yaml:"tag,omitempty"
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty" yaml:"tag,omitempty"`
	// @gotags: json:"username,omitempty" yaml:"username,omitempty"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty" yaml:"username,omitempty"`
	// @gotags: json:"password,omitempty" yaml:"password,omitempty"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty" yaml:"password,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Image) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Image) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Support struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"
	RuntimeVersion string `protobuf:"bytes,1,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
	// @gotags: json:"os,omitempty" yaml:"os,omitempty"
	Os []string `protobuf:"bytes,2,rep,name=os,proto3" json:"os,omitempty" yaml:"os,omitempty"`
	// @gotags: json:"arch,omitempty" yaml:"arch,omitempty"
	Arch []string `protobuf:"bytes,3,rep,name=arch,proto3" json:"arch,omitempty" yaml:"arch,omitempty"`
	// @gotags: json:"dependencies,omitempty" yaml:"dependencies,omitempty"
	Dependencies []*Dependency `protobuf:"bytes,4,rep,name=dependencies,proto3" json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}

func (x *Support) Reset() {
	*x = Support{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Support) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Support) ProtoMessage() {}

func (x *Support) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Support.ProtoReflect.Descriptor instead.
func (*Support) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *Support) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *Support) GetOs() []string {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *Support) GetArch() []string {
	if x != nil {
		return x.Arch
	}
	return nil
}

func (x *Support) GetDependencies() []*Dependency {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name,omitempty" yaml:"name,omitempty"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	// @gotags: json:"version,omitempty" yaml:"version,omitempty"
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" yaml:"version,omitempty"`
	// @gotags: json:"remote,omitempty" yaml:"remote,omitempty"
	Remote string `protobuf:"bytes,3,opt,name=remote,proto3" json:"remote,omitempty" yaml:"remote,omitempty"`
}

func (x *Dependency) Reset() {
	*x = Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependency) ProtoMessage() {}

func (x *Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependency.ProtoReflect.Descriptor instead.
func (*Dependency) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *Dependency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dependency) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Dependency) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

type Maintainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name,omitempty" yaml:"name,omitempty"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	// @gotags: json:"email,omitempty" yaml:"email,omitempty"
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" yaml:"email,omitempty"`
	// @gotags: json:"home,omitempty" yaml:"home,omitempty"
	Home string `protobuf:"bytes,3,opt,name=home,proto3" json:"home,omitempty" yaml:"home,omitempty"`
}

func (x *Maintainer) Reset() {
	*x = Maintainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maintainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maintainer) ProtoMessage() {}

func (x *Maintainer) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maintainer.ProtoReflect.Descriptor instead.
func (*Maintainer) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *Maintainer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Maintainer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Maintainer) GetHome() string {
	if x != nil {
		return x.Home
	}
	return ""
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"icon,omitempty" yaml:"icon,omitempty"
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty" yaml:"icon,omitempty"`
	// @gotags: json:"description,omitempty" yaml:"description,omitempty"
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description,omitempty"`
	// @gotags: json:"digest,omitempty" yaml:"digest,omitempty"
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty" yaml:"digest,omitempty"`
	// @gotags: json:"keywords,omitempty" yaml:"keywords,omitempty"
	Keywords []string `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty" yaml:"keywords,omitempty"`
	// @gotags: json:"maintainers,omitempty" yaml:"maintainers,omitempty"
	Maintainers []*Maintainer `protobuf:"bytes,5,rep,name=maintainers,proto3" json:"maintainers,omitempty" yaml:"maintainers,omitempty"`
	// @gotags: json:"generated,omitempty" yaml:"generated,omitempty"
	Generated *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=generated,proto3" json:"generated,omitempty" yaml:"generated,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *Metadata) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Metadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Metadata) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Metadata) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Metadata) GetMaintainers() []*Maintainer {
	if x != nil {
		return x.Maintainers
	}
	return nil
}

func (x *Metadata) GetGenerated() *timestamppb.Timestamp {
	if x != nil {
		return x.Generated
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name,omitempty" yaml:"name,omitempty"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	// @gotags: json:"platform,omitempty" yaml:"platform,omitempty"
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty" yaml:"platform,omitempty"`
	// @gotags: json:"ip,omitempty" yaml:"ip,omitempty"
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty" yaml:"ip,omitempty"`
	// @gotags: json:"generated,omitempty" yaml:"generated,omitempty"
	Generated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=generated,proto3" json:"generated,omitempty" yaml:"generated,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_plugin_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_v1_plugin_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_v1_plugin_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Author) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Author) GetGenerated() *timestamppb.Timestamp {
	if x != nil {
		return x.Generated
	}
	return nil
}

var File_v1_plugin_plugin_proto protoreflect.FileDescriptor

var file_v1_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x02, 0x0a, 0x06, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x12, 0x31, 0x0a,
	0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x65, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x3e, 0x0a, 0x0c, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x0a, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x4a, 0x0a,
	0x0a, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x3c, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2a, 0x3b, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x58, 0x45, 0x43, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_plugin_plugin_proto_rawDescOnce sync.Once
	file_v1_plugin_plugin_proto_rawDescData = file_v1_plugin_plugin_proto_rawDesc
)

func file_v1_plugin_plugin_proto_rawDescGZIP() []byte {
	file_v1_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_v1_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_plugin_plugin_proto_rawDescData)
	})
	return file_v1_plugin_plugin_proto_rawDescData
}

var file_v1_plugin_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_plugin_plugin_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: meta.v1.plugin.Type
	(*Plugin)(nil),                // 1: meta.v1.plugin.Plugin
	(*Image)(nil),                 // 2: meta.v1.plugin.Image
	(*Support)(nil),               // 3: meta.v1.plugin.Support
	(*Dependency)(nil),            // 4: meta.v1.plugin.Dependency
	(*Maintainer)(nil),            // 5: meta.v1.plugin.Maintainer
	(*Metadata)(nil),              // 6: meta.v1.plugin.Metadata
	(*Author)(nil),                // 7: meta.v1.plugin.Author
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_v1_plugin_plugin_proto_depIdxs = []int32{
	0, // 0: meta.v1.plugin.Plugin.type:type_name -> meta.v1.plugin.Type
	2, // 1: meta.v1.plugin.Plugin.image:type_name -> meta.v1.plugin.Image
	3, // 2: meta.v1.plugin.Plugin.support:type_name -> meta.v1.plugin.Support
	6, // 3: meta.v1.plugin.Plugin.metadata:type_name -> meta.v1.plugin.Metadata
	4, // 4: meta.v1.plugin.Support.dependencies:type_name -> meta.v1.plugin.Dependency
	5, // 5: meta.v1.plugin.Metadata.maintainers:type_name -> meta.v1.plugin.Maintainer
	8, // 6: meta.v1.plugin.Metadata.generated:type_name -> google.protobuf.Timestamp
	8, // 7: meta.v1.plugin.Author.generated:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_v1_plugin_plugin_proto_init() }
func file_v1_plugin_plugin_proto_init() {
	if File_v1_plugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_plugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_v1_plugin_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_v1_plugin_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Support); i {
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
		file_v1_plugin_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependency); i {
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
		file_v1_plugin_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maintainer); i {
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
		file_v1_plugin_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_v1_plugin_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
			RawDescriptor: file_v1_plugin_plugin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_v1_plugin_plugin_proto_depIdxs,
		EnumInfos:         file_v1_plugin_plugin_proto_enumTypes,
		MessageInfos:      file_v1_plugin_plugin_proto_msgTypes,
	}.Build()
	File_v1_plugin_plugin_proto = out.File
	file_v1_plugin_plugin_proto_rawDesc = nil
	file_v1_plugin_plugin_proto_goTypes = nil
	file_v1_plugin_plugin_proto_depIdxs = nil
}