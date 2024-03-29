// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: v1/reqs/logs.proto

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

type ParamLogsLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Time      int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Delay     int64  `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	Server    string `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
	Auth      string `protobuf:"bytes,5,opt,name=auth,proto3" json:"auth,omitempty"`
	Head      string `protobuf:"bytes,6,opt,name=head,proto3" json:"head,omitempty"`
	Param     string `protobuf:"bytes,7,opt,name=param,proto3" json:"param,omitempty"`
	Data      string `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Error     string `protobuf:"bytes,9,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ParamLogsLink) Reset() {
	*x = ParamLogsLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_reqs_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamLogsLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamLogsLink) ProtoMessage() {}

func (x *ParamLogsLink) ProtoReflect() protoreflect.Message {
	mi := &file_v1_reqs_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamLogsLink.ProtoReflect.Descriptor instead.
func (*ParamLogsLink) Descriptor() ([]byte, []int) {
	return file_v1_reqs_logs_proto_rawDescGZIP(), []int{0}
}

func (x *ParamLogsLink) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ParamLogsLink) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ParamLogsLink) GetDelay() int64 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *ParamLogsLink) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ParamLogsLink) GetAuth() string {
	if x != nil {
		return x.Auth
	}
	return ""
}

func (x *ParamLogsLink) GetHead() string {
	if x != nil {
		return x.Head
	}
	return ""
}

func (x *ParamLogsLink) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *ParamLogsLink) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ParamLogsLink) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ParamLogsOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Corpid     string `protobuf:"bytes,2,opt,name=corpid,proto3" json:"corpid,omitempty"`
	Userid     string `protobuf:"bytes,3,opt,name=userid,proto3" json:"userid,omitempty"`
	UserName   string `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	ClientIp   string `protobuf:"bytes,5,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ServerName string `protobuf:"bytes,6,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	Operate    string `protobuf:"bytes,7,opt,name=operate,proto3" json:"operate,omitempty"`
	Func       string `protobuf:"bytes,8,opt,name=func,proto3" json:"func,omitempty"`
	Param      string `protobuf:"bytes,9,opt,name=param,proto3" json:"param,omitempty"`
	Time       int64  `protobuf:"varint,10,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ParamLogsOperation) Reset() {
	*x = ParamLogsOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_reqs_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamLogsOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamLogsOperation) ProtoMessage() {}

func (x *ParamLogsOperation) ProtoReflect() protoreflect.Message {
	mi := &file_v1_reqs_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamLogsOperation.ProtoReflect.Descriptor instead.
func (*ParamLogsOperation) Descriptor() ([]byte, []int) {
	return file_v1_reqs_logs_proto_rawDescGZIP(), []int{1}
}

func (x *ParamLogsOperation) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ParamLogsOperation) GetCorpid() string {
	if x != nil {
		return x.Corpid
	}
	return ""
}

func (x *ParamLogsOperation) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *ParamLogsOperation) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ParamLogsOperation) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *ParamLogsOperation) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *ParamLogsOperation) GetOperate() string {
	if x != nil {
		return x.Operate
	}
	return ""
}

func (x *ParamLogsOperation) GetFunc() string {
	if x != nil {
		return x.Func
	}
	return ""
}

func (x *ParamLogsOperation) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *ParamLogsOperation) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type ParamLogsLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Corpid       string `protobuf:"bytes,1,opt,name=corpid,proto3" json:"corpid,omitempty"`
	Userid       string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`
	UserName     string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	LoginMode    string `protobuf:"bytes,4,opt,name=login_mode,json=loginMode,proto3" json:"login_mode,omitempty"`
	LoginClient  string `protobuf:"bytes,5,opt,name=login_client,json=loginClient,proto3" json:"login_client,omitempty"`
	LoginAccount string `protobuf:"bytes,6,opt,name=login_account,json=loginAccount,proto3" json:"login_account,omitempty"`
	LoginIp      string `protobuf:"bytes,7,opt,name=login_ip,json=loginIp,proto3" json:"login_ip,omitempty"`
	LoginStatus  int32  `protobuf:"varint,8,opt,name=login_status,json=loginStatus,proto3" json:"login_status,omitempty"`
	Time         int64  `protobuf:"varint,9,opt,name=time,proto3" json:"time,omitempty"`
	Desc         string `protobuf:"bytes,10,opt,name=desc,proto3" json:"desc,omitempty"`
	Token        string `protobuf:"bytes,11,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ParamLogsLogin) Reset() {
	*x = ParamLogsLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_reqs_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamLogsLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamLogsLogin) ProtoMessage() {}

func (x *ParamLogsLogin) ProtoReflect() protoreflect.Message {
	mi := &file_v1_reqs_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamLogsLogin.ProtoReflect.Descriptor instead.
func (*ParamLogsLogin) Descriptor() ([]byte, []int) {
	return file_v1_reqs_logs_proto_rawDescGZIP(), []int{2}
}

func (x *ParamLogsLogin) GetCorpid() string {
	if x != nil {
		return x.Corpid
	}
	return ""
}

func (x *ParamLogsLogin) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *ParamLogsLogin) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ParamLogsLogin) GetLoginMode() string {
	if x != nil {
		return x.LoginMode
	}
	return ""
}

func (x *ParamLogsLogin) GetLoginClient() string {
	if x != nil {
		return x.LoginClient
	}
	return ""
}

func (x *ParamLogsLogin) GetLoginAccount() string {
	if x != nil {
		return x.LoginAccount
	}
	return ""
}

func (x *ParamLogsLogin) GetLoginIp() string {
	if x != nil {
		return x.LoginIp
	}
	return ""
}

func (x *ParamLogsLogin) GetLoginStatus() int32 {
	if x != nil {
		return x.LoginStatus
	}
	return 0
}

func (x *ParamLogsLogin) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ParamLogsLogin) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ParamLogsLogin) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_v1_reqs_logs_proto protoreflect.FileDescriptor

var file_v1_reqs_logs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65,
	0x71, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x96, 0x02,
	0x0a, 0x12, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x63,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x2d, 0x73, 0x68, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_reqs_logs_proto_rawDescOnce sync.Once
	file_v1_reqs_logs_proto_rawDescData = file_v1_reqs_logs_proto_rawDesc
)

func file_v1_reqs_logs_proto_rawDescGZIP() []byte {
	file_v1_reqs_logs_proto_rawDescOnce.Do(func() {
		file_v1_reqs_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_reqs_logs_proto_rawDescData)
	})
	return file_v1_reqs_logs_proto_rawDescData
}

var file_v1_reqs_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_reqs_logs_proto_goTypes = []interface{}{
	(*ParamLogsLink)(nil),      // 0: meta.v1.reqs.ParamLogsLink
	(*ParamLogsOperation)(nil), // 1: meta.v1.reqs.ParamLogsOperation
	(*ParamLogsLogin)(nil),     // 2: meta.v1.reqs.ParamLogsLogin
}
var file_v1_reqs_logs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_reqs_logs_proto_init() }
func file_v1_reqs_logs_proto_init() {
	if File_v1_reqs_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_reqs_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamLogsLink); i {
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
		file_v1_reqs_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamLogsOperation); i {
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
		file_v1_reqs_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamLogsLogin); i {
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
			RawDescriptor: file_v1_reqs_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_reqs_logs_proto_goTypes,
		DependencyIndexes: file_v1_reqs_logs_proto_depIdxs,
		MessageInfos:      file_v1_reqs_logs_proto_msgTypes,
	}.Build()
	File_v1_reqs_logs_proto = out.File
	file_v1_reqs_logs_proto_rawDesc = nil
	file_v1_reqs_logs_proto_goTypes = nil
	file_v1_reqs_logs_proto_depIdxs = nil
}
