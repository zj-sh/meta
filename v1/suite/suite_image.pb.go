// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/suite/suite_image.proto

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

type ImageBuildParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params       *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                     // docker image name
	Tag          string           `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`                                       // docker image tag
	Repo         string           `protobuf:"bytes,4,opt,name=repo,proto3" json:"repo,omitempty"`                                     // docker registry repo
	RegistryAuth string           `protobuf:"bytes,5,opt,name=registry_auth,json=registryAuth,proto3" json:"registry_auth,omitempty"` // docker registry auth
	Dockerfile   string           `protobuf:"bytes,6,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`                         // dockerfile path
	ContextPath  string           `protobuf:"bytes,7,opt,name=context_path,json=contextPath,proto3" json:"context_path,omitempty"`    // docker build 执行目录
	Arch         []base.BuildArch `protobuf:"varint,8,rep,packed,name=arch,proto3,enum=meta.v1.base.BuildArch" json:"arch,omitempty"` // 构建架构
	AutoPush     bool             `protobuf:"varint,9,opt,name=auto_push,json=autoPush,proto3" json:"auto_push,omitempty"`            // 是否自动推送
}

func (x *ImageBuildParam) Reset() {
	*x = ImageBuildParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageBuildParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageBuildParam) ProtoMessage() {}

func (x *ImageBuildParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageBuildParam.ProtoReflect.Descriptor instead.
func (*ImageBuildParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_image_proto_rawDescGZIP(), []int{0}
}

func (x *ImageBuildParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ImageBuildParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImageBuildParam) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ImageBuildParam) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *ImageBuildParam) GetRegistryAuth() string {
	if x != nil {
		return x.RegistryAuth
	}
	return ""
}

func (x *ImageBuildParam) GetDockerfile() string {
	if x != nil {
		return x.Dockerfile
	}
	return ""
}

func (x *ImageBuildParam) GetContextPath() string {
	if x != nil {
		return x.ContextPath
	}
	return ""
}

func (x *ImageBuildParam) GetArch() []base.BuildArch {
	if x != nil {
		return x.Arch
	}
	return nil
}

func (x *ImageBuildParam) GetAutoPush() bool {
	if x != nil {
		return x.AutoPush
	}
	return false
}

type ImageBuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  base.RunStatus `protobuf:"varint,1,opt,name=status,proto3,enum=meta.v1.base.RunStatus" json:"status,omitempty"`
	Log     *base.Log      `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	Message string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Outs    *OutResponse   `protobuf:"bytes,4,opt,name=outs,proto3" json:"outs,omitempty"`
}

func (x *ImageBuildResponse) Reset() {
	*x = ImageBuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageBuildResponse) ProtoMessage() {}

func (x *ImageBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageBuildResponse.ProtoReflect.Descriptor instead.
func (*ImageBuildResponse) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_image_proto_rawDescGZIP(), []int{1}
}

func (x *ImageBuildResponse) GetStatus() base.RunStatus {
	if x != nil {
		return x.Status
	}
	return base.RunStatus(0)
}

func (x *ImageBuildResponse) GetLog() *base.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *ImageBuildResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ImageBuildResponse) GetOuts() *OutResponse {
	if x != nil {
		return x.Outs
	}
	return nil
}

var File_v1_suite_suite_image_proto protoreflect.FileDescriptor

var file_v1_suite_suite_image_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x12, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70,
	0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x63, 0x68,
	0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70,
	0x75, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x50,
	0x75, 0x73, 0x68, 0x22, 0xb4, 0x01, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x6c,
	0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6f, 0x75,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6f, 0x75, 0x74, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_suite_suite_image_proto_rawDescOnce sync.Once
	file_v1_suite_suite_image_proto_rawDescData = file_v1_suite_suite_image_proto_rawDesc
)

func file_v1_suite_suite_image_proto_rawDescGZIP() []byte {
	file_v1_suite_suite_image_proto_rawDescOnce.Do(func() {
		file_v1_suite_suite_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_suite_suite_image_proto_rawDescData)
	})
	return file_v1_suite_suite_image_proto_rawDescData
}

var file_v1_suite_suite_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_suite_suite_image_proto_goTypes = []interface{}{
	(*ImageBuildParam)(nil),    // 0: meta.v1.suite.ImageBuildParam
	(*ImageBuildResponse)(nil), // 1: meta.v1.suite.ImageBuildResponse
	(*base.JsonString)(nil),    // 2: meta.v1.base.JsonString
	(base.BuildArch)(0),        // 3: meta.v1.base.BuildArch
	(base.RunStatus)(0),        // 4: meta.v1.base.RunStatus
	(*base.Log)(nil),           // 5: meta.v1.base.Log
	(*OutResponse)(nil),        // 6: meta.v1.suite.OutResponse
}
var file_v1_suite_suite_image_proto_depIdxs = []int32{
	2, // 0: meta.v1.suite.ImageBuildParam.params:type_name -> meta.v1.base.JsonString
	3, // 1: meta.v1.suite.ImageBuildParam.arch:type_name -> meta.v1.base.BuildArch
	4, // 2: meta.v1.suite.ImageBuildResponse.status:type_name -> meta.v1.base.RunStatus
	5, // 3: meta.v1.suite.ImageBuildResponse.log:type_name -> meta.v1.base.Log
	6, // 4: meta.v1.suite.ImageBuildResponse.outs:type_name -> meta.v1.suite.OutResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_suite_suite_image_proto_init() }
func file_v1_suite_suite_image_proto_init() {
	if File_v1_suite_suite_image_proto != nil {
		return
	}
	file_v1_suite_meta_running_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_suite_suite_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageBuildParam); i {
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
		file_v1_suite_suite_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageBuildResponse); i {
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
			RawDescriptor: file_v1_suite_suite_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_suite_suite_image_proto_goTypes,
		DependencyIndexes: file_v1_suite_suite_image_proto_depIdxs,
		MessageInfos:      file_v1_suite_suite_image_proto_msgTypes,
	}.Build()
	File_v1_suite_suite_image_proto = out.File
	file_v1_suite_suite_image_proto_rawDesc = nil
	file_v1_suite_suite_image_proto_goTypes = nil
	file_v1_suite_suite_image_proto_depIdxs = nil
}
