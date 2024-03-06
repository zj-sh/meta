// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/base/meta.proto

package base

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

// 通用执行状态
type RunStatus int32

const (
	// 未知
	RunStatus_RUN_STATUS_UNKNOWN RunStatus = 0
	// 进行中
	RunStatus_RUN_STATUS_PROCESSING RunStatus = 1
	// 成功
	RunStatus_RUN_STATUS_SUCCESS RunStatus = 2
	// 失败
	RunStatus_RUN_STATUS_FAILED RunStatus = 3
	// 完成
	RunStatus_RUN_STATUS_COMPLETED RunStatus = 4
)

// Enum value maps for RunStatus.
var (
	RunStatus_name = map[int32]string{
		0: "RUN_STATUS_UNKNOWN",
		1: "RUN_STATUS_PROCESSING",
		2: "RUN_STATUS_SUCCESS",
		3: "RUN_STATUS_FAILED",
		4: "RUN_STATUS_COMPLETED",
	}
	RunStatus_value = map[string]int32{
		"RUN_STATUS_UNKNOWN":    0,
		"RUN_STATUS_PROCESSING": 1,
		"RUN_STATUS_SUCCESS":    2,
		"RUN_STATUS_FAILED":     3,
		"RUN_STATUS_COMPLETED":  4,
	}
)

func (x RunStatus) Enum() *RunStatus {
	p := new(RunStatus)
	*p = x
	return p
}

func (x RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_base_meta_proto_enumTypes[0].Descriptor()
}

func (RunStatus) Type() protoreflect.EnumType {
	return &file_v1_base_meta_proto_enumTypes[0]
}

func (x RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStatus.Descriptor instead.
func (RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_base_meta_proto_rawDescGZIP(), []int{0}
}

// 通用状态
type Status int32

const (
	Status_STATUS_UNKNOWN Status = 0
	// 正常
	Status_STATUS_NORMAL Status = 1
	// 锁定
	Status_STATUS_LOCKED Status = 2
	// 禁用
	Status_STATUS_DISABLE Status = 3
	// 删除
	Status_STATUS_DELETED Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "STATUS_NORMAL",
		2: "STATUS_LOCKED",
		3: "STATUS_DISABLE",
		4: "STATUS_DELETED",
	}
	Status_value = map[string]int32{
		"STATUS_UNKNOWN": 0,
		"STATUS_NORMAL":  1,
		"STATUS_LOCKED":  2,
		"STATUS_DISABLE": 3,
		"STATUS_DELETED": 4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_base_meta_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_v1_base_meta_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_v1_base_meta_proto_rawDescGZIP(), []int{1}
}

// String的别名，通用JSON字符串，用以区分是否普通字符串
type JsonString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JsonString) Reset() {
	*x = JsonString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonString) ProtoMessage() {}

func (x *JsonString) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonString.ProtoReflect.Descriptor instead.
func (*JsonString) Descriptor() ([]byte, []int) {
	return file_v1_base_meta_proto_rawDescGZIP(), []int{0}
}

func (x *JsonString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 通用选项
type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"label" gorm:"-"
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label" gorm:"-"`
	// @gotags: json:"value" gorm:"-"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value" gorm:"-"`
	// @gotags: json:"children" gorm:"-"
	Children []*Option `protobuf:"bytes,4,rep,name=children,proto3" json:"children" gorm:"-"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_v1_base_meta_proto_rawDescGZIP(), []int{1}
}

func (x *Option) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Option) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Option) GetChildren() []*Option {
	if x != nil {
		return x.Children
	}
	return nil
}

// 通用分页
type Pages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"` // 从1开始
	Rows int64 `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	// Types that are assignable to Cursor:
	//
	//	*Pages_Next
	//	*Pages_Limit
	Cursor isPages_Cursor `protobuf_oneof:"cursor"`
}

func (x *Pages) Reset() {
	*x = Pages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pages) ProtoMessage() {}

func (x *Pages) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pages.ProtoReflect.Descriptor instead.
func (*Pages) Descriptor() ([]byte, []int) {
	return file_v1_base_meta_proto_rawDescGZIP(), []int{2}
}

func (x *Pages) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pages) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (m *Pages) GetCursor() isPages_Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (x *Pages) GetNext() string {
	if x, ok := x.GetCursor().(*Pages_Next); ok {
		return x.Next
	}
	return ""
}

func (x *Pages) GetLimit() int64 {
	if x, ok := x.GetCursor().(*Pages_Limit); ok {
		return x.Limit
	}
	return 0
}

type isPages_Cursor interface {
	isPages_Cursor()
}

type Pages_Next struct {
	Next string `protobuf:"bytes,3,opt,name=next,proto3,oneof"`
}

type Pages_Limit struct {
	Limit int64 `protobuf:"varint,4,opt,name=limit,proto3,oneof"`
}

func (*Pages_Next) isPages_Cursor() {}

func (*Pages_Limit) isPages_Cursor() {}

var File_v1_base_meta_proto protoreflect.FileDescriptor

var file_v1_base_meta_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x66, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x67,
	0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12,
	0x14, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x08, 0x0a,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2a, 0x87, 0x01, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x55, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x04, 0x2a, 0x6a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0x1f, 0x5a,
	0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73,
	0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_base_meta_proto_rawDescOnce sync.Once
	file_v1_base_meta_proto_rawDescData = file_v1_base_meta_proto_rawDesc
)

func file_v1_base_meta_proto_rawDescGZIP() []byte {
	file_v1_base_meta_proto_rawDescOnce.Do(func() {
		file_v1_base_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_base_meta_proto_rawDescData)
	})
	return file_v1_base_meta_proto_rawDescData
}

var file_v1_base_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_base_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_base_meta_proto_goTypes = []interface{}{
	(RunStatus)(0),     // 0: meta.v1.base.RunStatus
	(Status)(0),        // 1: meta.v1.base.Status
	(*JsonString)(nil), // 2: meta.v1.base.JsonString
	(*Option)(nil),     // 3: meta.v1.base.Option
	(*Pages)(nil),      // 4: meta.v1.base.Pages
}
var file_v1_base_meta_proto_depIdxs = []int32{
	3, // 0: meta.v1.base.Option.children:type_name -> meta.v1.base.Option
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_base_meta_proto_init() }
func file_v1_base_meta_proto_init() {
	if File_v1_base_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_base_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonString); i {
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
		file_v1_base_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_v1_base_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pages); i {
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
	file_v1_base_meta_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Pages_Next)(nil),
		(*Pages_Limit)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_base_meta_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_base_meta_proto_goTypes,
		DependencyIndexes: file_v1_base_meta_proto_depIdxs,
		EnumInfos:         file_v1_base_meta_proto_enumTypes,
		MessageInfos:      file_v1_base_meta_proto_msgTypes,
	}.Build()
	File_v1_base_meta_proto = out.File
	file_v1_base_meta_proto_rawDesc = nil
	file_v1_base_meta_proto_goTypes = nil
	file_v1_base_meta_proto_depIdxs = nil
}
