// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/reqs/cron.proto

package reqs

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

type CronMode int32

const (
	CronMode_CRON_MODE_UNKNOWN CronMode = 0
	CronMode_CRON_MODE_ONCE    CronMode = 1
	CronMode_CRON_MODE_LOOP    CronMode = 2
)

// Enum value maps for CronMode.
var (
	CronMode_name = map[int32]string{
		0: "CRON_MODE_UNKNOWN",
		1: "CRON_MODE_ONCE",
		2: "CRON_MODE_LOOP",
	}
	CronMode_value = map[string]int32{
		"CRON_MODE_UNKNOWN": 0,
		"CRON_MODE_ONCE":    1,
		"CRON_MODE_LOOP":    2,
	}
)

func (x CronMode) Enum() *CronMode {
	p := new(CronMode)
	*p = x
	return p
}

func (x CronMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CronMode) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_reqs_cron_proto_enumTypes[0].Descriptor()
}

func (CronMode) Type() protoreflect.EnumType {
	return &file_v1_reqs_cron_proto_enumTypes[0]
}

func (x CronMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CronMode.Descriptor instead.
func (CronMode) EnumDescriptor() ([]byte, []int) {
	return file_v1_reqs_cron_proto_rawDescGZIP(), []int{0}
}

type ParamCronStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"cron_name" binding:"required"
	CronName string `protobuf:"bytes,1,opt,name=cron_name,json=cronName,proto3" json:"cron_name" binding:"required"`
	// @gotags: json:"cron_spec" binding:"required"
	CronSpec string `protobuf:"bytes,2,opt,name=cron_spec,json=cronSpec,proto3" json:"cron_spec" binding:"required"`
	// @gotags: json:"cron_desc" binding:"required"
	CronDesc string `protobuf:"bytes,3,opt,name=cron_desc,json=cronDesc,proto3" json:"cron_desc" binding:"required"`
	// @gotags: json:"cron_serve" binding:"required"
	CronServe string `protobuf:"bytes,4,opt,name=cron_serve,json=cronServe,proto3" json:"cron_serve" binding:"required"`
	// @gotags: json:"cron_url"
	CronUrl string `protobuf:"bytes,5,opt,name=cron_url,json=cronUrl,proto3" json:"cron_url"`
	// @gotags: json:"expires_time"
	ExpiresTime int64 `protobuf:"varint,6,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time"`
	// @gotags: json:"cron_param"
	CronParam string `protobuf:"bytes,7,opt,name=cron_param,json=cronParam,proto3" json:"cron_param"`
	// @gotags: json:"cron_mode"
	CronMode CronMode `protobuf:"varint,8,opt,name=cron_mode,json=cronMode,proto3,enum=meta.v1.reqs.CronMode" json:"cron_mode"`
}

func (x *ParamCronStart) Reset() {
	*x = ParamCronStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_reqs_cron_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamCronStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamCronStart) ProtoMessage() {}

func (x *ParamCronStart) ProtoReflect() protoreflect.Message {
	mi := &file_v1_reqs_cron_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamCronStart.ProtoReflect.Descriptor instead.
func (*ParamCronStart) Descriptor() ([]byte, []int) {
	return file_v1_reqs_cron_proto_rawDescGZIP(), []int{0}
}

func (x *ParamCronStart) GetCronName() string {
	if x != nil {
		return x.CronName
	}
	return ""
}

func (x *ParamCronStart) GetCronSpec() string {
	if x != nil {
		return x.CronSpec
	}
	return ""
}

func (x *ParamCronStart) GetCronDesc() string {
	if x != nil {
		return x.CronDesc
	}
	return ""
}

func (x *ParamCronStart) GetCronServe() string {
	if x != nil {
		return x.CronServe
	}
	return ""
}

func (x *ParamCronStart) GetCronUrl() string {
	if x != nil {
		return x.CronUrl
	}
	return ""
}

func (x *ParamCronStart) GetExpiresTime() int64 {
	if x != nil {
		return x.ExpiresTime
	}
	return 0
}

func (x *ParamCronStart) GetCronParam() string {
	if x != nil {
		return x.CronParam
	}
	return ""
}

func (x *ParamCronStart) GetCronMode() CronMode {
	if x != nil {
		return x.CronMode
	}
	return CronMode_CRON_MODE_UNKNOWN
}

type ParamCronStop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"cron_name" binding:"required"
	CronName string `protobuf:"bytes,1,opt,name=cron_name,json=cronName,proto3" json:"cron_name" binding:"required"`
}

func (x *ParamCronStop) Reset() {
	*x = ParamCronStop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_reqs_cron_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamCronStop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamCronStop) ProtoMessage() {}

func (x *ParamCronStop) ProtoReflect() protoreflect.Message {
	mi := &file_v1_reqs_cron_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamCronStop.ProtoReflect.Descriptor instead.
func (*ParamCronStop) Descriptor() ([]byte, []int) {
	return file_v1_reqs_cron_proto_rawDescGZIP(), []int{1}
}

func (x *ParamCronStop) GetCronName() string {
	if x != nil {
		return x.CronName
	}
	return ""
}

var File_v1_reqs_cron_proto protoreflect.FileDescriptor

var file_v1_reqs_cron_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65,
	0x71, 0x73, 0x22, 0x98, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43, 0x72, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x72, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x72, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x73, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a,
	0x0d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43, 0x72, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x49, 0x0a, 0x08, 0x43,
	0x72, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x4f, 0x4e, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x52, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4e, 0x43, 0x45,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x4c, 0x4f, 0x4f, 0x50, 0x10, 0x02, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_reqs_cron_proto_rawDescOnce sync.Once
	file_v1_reqs_cron_proto_rawDescData = file_v1_reqs_cron_proto_rawDesc
)

func file_v1_reqs_cron_proto_rawDescGZIP() []byte {
	file_v1_reqs_cron_proto_rawDescOnce.Do(func() {
		file_v1_reqs_cron_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_reqs_cron_proto_rawDescData)
	})
	return file_v1_reqs_cron_proto_rawDescData
}

var file_v1_reqs_cron_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_reqs_cron_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_reqs_cron_proto_goTypes = []interface{}{
	(CronMode)(0),          // 0: meta.v1.reqs.CronMode
	(*ParamCronStart)(nil), // 1: meta.v1.reqs.ParamCronStart
	(*ParamCronStop)(nil),  // 2: meta.v1.reqs.ParamCronStop
}
var file_v1_reqs_cron_proto_depIdxs = []int32{
	0, // 0: meta.v1.reqs.ParamCronStart.cron_mode:type_name -> meta.v1.reqs.CronMode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_reqs_cron_proto_init() }
func file_v1_reqs_cron_proto_init() {
	if File_v1_reqs_cron_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_reqs_cron_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamCronStart); i {
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
		file_v1_reqs_cron_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamCronStop); i {
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
			RawDescriptor: file_v1_reqs_cron_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_reqs_cron_proto_goTypes,
		DependencyIndexes: file_v1_reqs_cron_proto_depIdxs,
		EnumInfos:         file_v1_reqs_cron_proto_enumTypes,
		MessageInfos:      file_v1_reqs_cron_proto_msgTypes,
	}.Build()
	File_v1_reqs_cron_proto = out.File
	file_v1_reqs_cron_proto_rawDesc = nil
	file_v1_reqs_cron_proto_goTypes = nil
	file_v1_reqs_cron_proto_depIdxs = nil
}
