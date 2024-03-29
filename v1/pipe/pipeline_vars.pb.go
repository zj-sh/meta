// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/pipe/pipeline_vars.proto

package pipe

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

type Var struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"group_id" yaml:"group_id"
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id" yaml:"group_id"`
	// @gotags: json:"variable" yaml:"variable"
	Variable string `protobuf:"bytes,2,opt,name=variable,proto3" json:"variable" yaml:"variable"`
	// @gotags: json:"value" yaml:"value"
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value" yaml:"value"`
	// @gotags: json:"is_running" yaml:"is_running"
	IsRunning bool `protobuf:"varint,4,opt,name=is_running,json=isRunning,proto3" json:"is_running" yaml:"is_running"`
	// @gotags: json:"is_private" yaml:"is_private"
	IsPrivate bool `protobuf:"varint,5,opt,name=is_private,json=isPrivate,proto3" json:"is_private" yaml:"is_private"`
	// @gotags: json:"desc" yaml:"desc"
	Desc string `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc" yaml:"desc"`
	// @gotags: json:"value_real" yaml:"value_real"
	ValueReal string `protobuf:"bytes,7,opt,name=value_real,json=valueReal,proto3" json:"value_real" yaml:"value_real"`
}

func (x *Var) Reset() {
	*x = Var{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pipe_pipeline_vars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Var) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Var) ProtoMessage() {}

func (x *Var) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pipe_pipeline_vars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Var.ProtoReflect.Descriptor instead.
func (*Var) Descriptor() ([]byte, []int) {
	return file_v1_pipe_pipeline_vars_proto_rawDescGZIP(), []int{0}
}

func (x *Var) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Var) GetVariable() string {
	if x != nil {
		return x.Variable
	}
	return ""
}

func (x *Var) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Var) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *Var) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Var) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Var) GetValueReal() string {
	if x != nil {
		return x.ValueReal
	}
	return ""
}

type VarGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"group_id" yaml:"group_id"
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id" yaml:"group_id"`
	// @gotags: json:"group_code" yaml:"group_code"
	GroupCode string `protobuf:"bytes,2,opt,name=group_code,json=groupCode,proto3" json:"group_code" yaml:"group_code"`
	// @gotags: json:"group_name" yaml:"group_name"
	GroupName string `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name" yaml:"group_name"`
	// @gotags: json:"group_desc" yaml:"group_desc"
	GroupDesc string `protobuf:"bytes,4,opt,name=group_desc,json=groupDesc,proto3" json:"group_desc" yaml:"group_desc"`
	// @gotags: json:"variables" yaml:"variables"
	Variables []*Var `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables" yaml:"variables"`
}

func (x *VarGroup) Reset() {
	*x = VarGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pipe_pipeline_vars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VarGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VarGroup) ProtoMessage() {}

func (x *VarGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pipe_pipeline_vars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VarGroup.ProtoReflect.Descriptor instead.
func (*VarGroup) Descriptor() ([]byte, []int) {
	return file_v1_pipe_pipeline_vars_proto_rawDescGZIP(), []int{1}
}

func (x *VarGroup) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *VarGroup) GetGroupCode() string {
	if x != nil {
		return x.GroupCode
	}
	return ""
}

func (x *VarGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *VarGroup) GetGroupDesc() string {
	if x != nil {
		return x.GroupDesc
	}
	return ""
}

func (x *VarGroup) GetVariables() []*Var {
	if x != nil {
		return x.Variables
	}
	return nil
}

var File_v1_pipe_pipeline_vars_proto protoreflect.FileDescriptor

var file_v1_pipe_pipeline_vars_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x03,
	0x56, 0x61, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x61,
	0x6c, 0x22, 0xb3, 0x01, 0x0a, 0x08, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x2f, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x52, 0x09, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_pipe_pipeline_vars_proto_rawDescOnce sync.Once
	file_v1_pipe_pipeline_vars_proto_rawDescData = file_v1_pipe_pipeline_vars_proto_rawDesc
)

func file_v1_pipe_pipeline_vars_proto_rawDescGZIP() []byte {
	file_v1_pipe_pipeline_vars_proto_rawDescOnce.Do(func() {
		file_v1_pipe_pipeline_vars_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_pipe_pipeline_vars_proto_rawDescData)
	})
	return file_v1_pipe_pipeline_vars_proto_rawDescData
}

var file_v1_pipe_pipeline_vars_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_pipe_pipeline_vars_proto_goTypes = []interface{}{
	(*Var)(nil),      // 0: meta.v1.pipe.Var
	(*VarGroup)(nil), // 1: meta.v1.pipe.VarGroup
}
var file_v1_pipe_pipeline_vars_proto_depIdxs = []int32{
	0, // 0: meta.v1.pipe.VarGroup.variables:type_name -> meta.v1.pipe.Var
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_pipe_pipeline_vars_proto_init() }
func file_v1_pipe_pipeline_vars_proto_init() {
	if File_v1_pipe_pipeline_vars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_pipe_pipeline_vars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Var); i {
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
		file_v1_pipe_pipeline_vars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VarGroup); i {
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
			RawDescriptor: file_v1_pipe_pipeline_vars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_pipe_pipeline_vars_proto_goTypes,
		DependencyIndexes: file_v1_pipe_pipeline_vars_proto_depIdxs,
		MessageInfos:      file_v1_pipe_pipeline_vars_proto_msgTypes,
	}.Build()
	File_v1_pipe_pipeline_vars_proto = out.File
	file_v1_pipe_pipeline_vars_proto_rawDesc = nil
	file_v1_pipe_pipeline_vars_proto_goTypes = nil
	file_v1_pipe_pipeline_vars_proto_depIdxs = nil
}
