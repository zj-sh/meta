// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/suite/suite_test.proto

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

type TestRunParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params    *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Cmd       string           `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Timeout   int64            `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	ReportDir string           `protobuf:"bytes,4,opt,name=report_dir,json=reportDir,proto3" json:"report_dir,omitempty"`
}

func (x *TestRunParam) Reset() {
	*x = TestRunParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunParam) ProtoMessage() {}

func (x *TestRunParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunParam.ProtoReflect.Descriptor instead.
func (*TestRunParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestRunParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *TestRunParam) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *TestRunParam) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *TestRunParam) GetReportDir() string {
	if x != nil {
		return x.ReportDir
	}
	return ""
}

type TestRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  base.RunStatus `protobuf:"varint,1,opt,name=status,proto3,enum=meta.v1.base.RunStatus" json:"status,omitempty"`
	Log     *base.Log      `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	Message string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Id      string         `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Outs    *OutResponse   `protobuf:"bytes,5,opt,name=outs,proto3" json:"outs,omitempty"`
}

func (x *TestRunResponse) Reset() {
	*x = TestRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunResponse) ProtoMessage() {}

func (x *TestRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunResponse.ProtoReflect.Descriptor instead.
func (*TestRunResponse) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestRunResponse) GetStatus() base.RunStatus {
	if x != nil {
		return x.Status
	}
	return base.RunStatus(0)
}

func (x *TestRunResponse) GetLog() *base.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *TestRunResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TestRunResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TestRunResponse) GetOuts() *OutResponse {
	if x != nil {
		return x.Outs
	}
	return nil
}

type TestCancelParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *base.JsonString `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Id     string           `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TestCancelParam) Reset() {
	*x = TestCancelParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_suite_suite_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCancelParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCancelParam) ProtoMessage() {}

func (x *TestCancelParam) ProtoReflect() protoreflect.Message {
	mi := &file_v1_suite_suite_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCancelParam.ProtoReflect.Descriptor instead.
func (*TestCancelParam) Descriptor() ([]byte, []int) {
	return file_v1_suite_suite_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestCancelParam) GetParams() *base.JsonString {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *TestCancelParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_suite_suite_test_proto protoreflect.FileDescriptor

var file_v1_suite_suite_test_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x12, 0x76, 0x31, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x69,
	0x72, 0x22, 0xc1, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x75,
	0x69, 0x74, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x6f, 0x75, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_suite_suite_test_proto_rawDescOnce sync.Once
	file_v1_suite_suite_test_proto_rawDescData = file_v1_suite_suite_test_proto_rawDesc
)

func file_v1_suite_suite_test_proto_rawDescGZIP() []byte {
	file_v1_suite_suite_test_proto_rawDescOnce.Do(func() {
		file_v1_suite_suite_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_suite_suite_test_proto_rawDescData)
	})
	return file_v1_suite_suite_test_proto_rawDescData
}

var file_v1_suite_suite_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_suite_suite_test_proto_goTypes = []interface{}{
	(*TestRunParam)(nil),    // 0: meta.v1.suite.TestRunParam
	(*TestRunResponse)(nil), // 1: meta.v1.suite.TestRunResponse
	(*TestCancelParam)(nil), // 2: meta.v1.suite.TestCancelParam
	(*base.JsonString)(nil), // 3: meta.v1.base.JsonString
	(base.RunStatus)(0),     // 4: meta.v1.base.RunStatus
	(*base.Log)(nil),        // 5: meta.v1.base.Log
	(*OutResponse)(nil),     // 6: meta.v1.suite.OutResponse
}
var file_v1_suite_suite_test_proto_depIdxs = []int32{
	3, // 0: meta.v1.suite.TestRunParam.params:type_name -> meta.v1.base.JsonString
	4, // 1: meta.v1.suite.TestRunResponse.status:type_name -> meta.v1.base.RunStatus
	5, // 2: meta.v1.suite.TestRunResponse.log:type_name -> meta.v1.base.Log
	6, // 3: meta.v1.suite.TestRunResponse.outs:type_name -> meta.v1.suite.OutResponse
	3, // 4: meta.v1.suite.TestCancelParam.params:type_name -> meta.v1.base.JsonString
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_suite_suite_test_proto_init() }
func file_v1_suite_suite_test_proto_init() {
	if File_v1_suite_suite_test_proto != nil {
		return
	}
	file_v1_suite_meta_running_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_suite_suite_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunParam); i {
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
		file_v1_suite_suite_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunResponse); i {
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
		file_v1_suite_suite_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCancelParam); i {
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
			RawDescriptor: file_v1_suite_suite_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_suite_suite_test_proto_goTypes,
		DependencyIndexes: file_v1_suite_suite_test_proto_depIdxs,
		MessageInfos:      file_v1_suite_suite_test_proto_msgTypes,
	}.Build()
	File_v1_suite_suite_test_proto = out.File
	file_v1_suite_suite_test_proto_rawDesc = nil
	file_v1_suite_suite_test_proto_goTypes = nil
	file_v1_suite_suite_test_proto_depIdxs = nil
}
