// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: authSAS.proto

package sasv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_authSAS_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponce) Reset() {
	*x = LoginResponce{}
	mi := &file_authSAS_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponce) ProtoMessage() {}

func (x *LoginResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponce.ProtoReflect.Descriptor instead.
func (*LoginResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponce) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponce) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_authSAS_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponce) Reset() {
	*x = LogoutResponce{}
	mi := &file_authSAS_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponce) ProtoMessage() {}

func (x *LogoutResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponce.ProtoReflect.Descriptor instead.
func (*LogoutResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutResponce) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LoginWith2FACodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code          int32                  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginWith2FACodeRequest) Reset() {
	*x = LoginWith2FACodeRequest{}
	mi := &file_authSAS_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginWith2FACodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWith2FACodeRequest) ProtoMessage() {}

func (x *LoginWith2FACodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWith2FACodeRequest.ProtoReflect.Descriptor instead.
func (*LoginWith2FACodeRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{4}
}

func (x *LoginWith2FACodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginWith2FACodeRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type LoginWith2FACodeResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginWith2FACodeResponce) Reset() {
	*x = LoginWith2FACodeResponce{}
	mi := &file_authSAS_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginWith2FACodeResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWith2FACodeResponce) ProtoMessage() {}

func (x *LoginWith2FACodeResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWith2FACodeResponce.ProtoReflect.Descriptor instead.
func (*LoginWith2FACodeResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{5}
}

func (x *LoginWith2FACodeResponce) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_authSAS_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponce) Reset() {
	*x = RegisterResponce{}
	mi := &file_authSAS_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponce) ProtoMessage() {}

func (x *RegisterResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponce.ProtoReflect.Descriptor instead.
func (*RegisterResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterResponce) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type EmailVerifySendCodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailVerifySendCodeRequest) Reset() {
	*x = EmailVerifySendCodeRequest{}
	mi := &file_authSAS_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailVerifySendCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailVerifySendCodeRequest) ProtoMessage() {}

func (x *EmailVerifySendCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailVerifySendCodeRequest.ProtoReflect.Descriptor instead.
func (*EmailVerifySendCodeRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{8}
}

func (x *EmailVerifySendCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type EmailVerifySendCodeResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailVerifySendCodeResponce) Reset() {
	*x = EmailVerifySendCodeResponce{}
	mi := &file_authSAS_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailVerifySendCodeResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailVerifySendCodeResponce) ProtoMessage() {}

func (x *EmailVerifySendCodeResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailVerifySendCodeResponce.ProtoReflect.Descriptor instead.
func (*EmailVerifySendCodeResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{9}
}

func (x *EmailVerifySendCodeResponce) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type EmailVerifyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code          int32                  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailVerifyRequest) Reset() {
	*x = EmailVerifyRequest{}
	mi := &file_authSAS_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailVerifyRequest) ProtoMessage() {}

func (x *EmailVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailVerifyRequest.ProtoReflect.Descriptor instead.
func (*EmailVerifyRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{10}
}

func (x *EmailVerifyRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailVerifyRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type EmailVerifyResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailVerifyResponce) Reset() {
	*x = EmailVerifyResponce{}
	mi := &file_authSAS_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailVerifyResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailVerifyResponce) ProtoMessage() {}

func (x *EmailVerifyResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailVerifyResponce.ProtoReflect.Descriptor instead.
func (*EmailVerifyResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{11}
}

func (x *EmailVerifyResponce) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PasswordRecoverSendCodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordRecoverSendCodeRequest) Reset() {
	*x = PasswordRecoverSendCodeRequest{}
	mi := &file_authSAS_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordRecoverSendCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRecoverSendCodeRequest) ProtoMessage() {}

func (x *PasswordRecoverSendCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRecoverSendCodeRequest.ProtoReflect.Descriptor instead.
func (*PasswordRecoverSendCodeRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{12}
}

func (x *PasswordRecoverSendCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PasswordRecoverSendCodeResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordRecoverSendCodeResponce) Reset() {
	*x = PasswordRecoverSendCodeResponce{}
	mi := &file_authSAS_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordRecoverSendCodeResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRecoverSendCodeResponce) ProtoMessage() {}

func (x *PasswordRecoverSendCodeResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRecoverSendCodeResponce.ProtoReflect.Descriptor instead.
func (*PasswordRecoverSendCodeResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{13}
}

func (x *PasswordRecoverSendCodeResponce) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PasswordRecoverRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	NewPassword   string                 `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	Code          int32                  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordRecoverRequest) Reset() {
	*x = PasswordRecoverRequest{}
	mi := &file_authSAS_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordRecoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRecoverRequest) ProtoMessage() {}

func (x *PasswordRecoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRecoverRequest.ProtoReflect.Descriptor instead.
func (*PasswordRecoverRequest) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{14}
}

func (x *PasswordRecoverRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PasswordRecoverRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *PasswordRecoverRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type PasswordRecoverResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordRecoverResponce) Reset() {
	*x = PasswordRecoverResponce{}
	mi := &file_authSAS_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordRecoverResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRecoverResponce) ProtoMessage() {}

func (x *PasswordRecoverResponce) ProtoReflect() protoreflect.Message {
	mi := &file_authSAS_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRecoverResponce.ProtoReflect.Descriptor instead.
func (*PasswordRecoverResponce) Descriptor() ([]byte, []int) {
	return file_authSAS_proto_rawDescGZIP(), []int{15}
}

func (x *PasswordRecoverResponce) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_authSAS_proto protoreflect.FileDescriptor

var file_authSAS_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0e, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x43,
	0x0a, 0x17, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x32, 0x46, 0x41, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x30, 0x0a, 0x18, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x32, 0x46, 0x41, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x10, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x1a, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x1b, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3e, 0x0a, 0x12,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x13,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x36, 0x0a, 0x1e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x33, 0x0a,
	0x1f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x65, 0x0a, 0x16, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2b, 0x0a, 0x17, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x83, 0x05, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53,
	0x41, 0x53, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x53, 0x41, 0x53, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x32,
	0x46, 0x41, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x32, 0x46, 0x41, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53,
	0x41, 0x53, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x32, 0x46, 0x41, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41,
	0x53, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x13,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53,
	0x41, 0x53, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x53, 0x41, 0x53, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x17, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x53, 0x41, 0x53, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x53, 0x41, 0x53, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x53, 0x41, 0x53, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x16, 0x5a, 0x14,
	0x62, 0x65, 0x67, 0x75, 0x6e, 0x6f, 0x76, 0x2e, 0x73, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x73,
	0x61, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_authSAS_proto_rawDescOnce sync.Once
	file_authSAS_proto_rawDescData []byte
)

func file_authSAS_proto_rawDescGZIP() []byte {
	file_authSAS_proto_rawDescOnce.Do(func() {
		file_authSAS_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_authSAS_proto_rawDesc), len(file_authSAS_proto_rawDesc)))
	})
	return file_authSAS_proto_rawDescData
}

var file_authSAS_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_authSAS_proto_goTypes = []any{
	(*LoginRequest)(nil),                    // 0: authSAS.LoginRequest
	(*LoginResponce)(nil),                   // 1: authSAS.LoginResponce
	(*LogoutRequest)(nil),                   // 2: authSAS.LogoutRequest
	(*LogoutResponce)(nil),                  // 3: authSAS.LogoutResponce
	(*LoginWith2FACodeRequest)(nil),         // 4: authSAS.LoginWith2FACodeRequest
	(*LoginWith2FACodeResponce)(nil),        // 5: authSAS.LoginWith2FACodeResponce
	(*RegisterRequest)(nil),                 // 6: authSAS.RegisterRequest
	(*RegisterResponce)(nil),                // 7: authSAS.RegisterResponce
	(*EmailVerifySendCodeRequest)(nil),      // 8: authSAS.EmailVerifySendCodeRequest
	(*EmailVerifySendCodeResponce)(nil),     // 9: authSAS.EmailVerifySendCodeResponce
	(*EmailVerifyRequest)(nil),              // 10: authSAS.EmailVerifyRequest
	(*EmailVerifyResponce)(nil),             // 11: authSAS.EmailVerifyResponce
	(*PasswordRecoverSendCodeRequest)(nil),  // 12: authSAS.PasswordRecoverSendCodeRequest
	(*PasswordRecoverSendCodeResponce)(nil), // 13: authSAS.PasswordRecoverSendCodeResponce
	(*PasswordRecoverRequest)(nil),          // 14: authSAS.PasswordRecoverRequest
	(*PasswordRecoverResponce)(nil),         // 15: authSAS.PasswordRecoverResponce
}
var file_authSAS_proto_depIdxs = []int32{
	0,  // 0: authSAS.Auth.Login:input_type -> authSAS.LoginRequest
	2,  // 1: authSAS.Auth.Logout:input_type -> authSAS.LogoutRequest
	4,  // 2: authSAS.Auth.LoginWith2FACode:input_type -> authSAS.LoginWith2FACodeRequest
	6,  // 3: authSAS.Auth.Register:input_type -> authSAS.RegisterRequest
	8,  // 4: authSAS.Auth.EmailVerifySendCode:input_type -> authSAS.EmailVerifySendCodeRequest
	10, // 5: authSAS.Auth.EmailVerify:input_type -> authSAS.EmailVerifyRequest
	12, // 6: authSAS.Auth.PasswordRecoverSendCode:input_type -> authSAS.PasswordRecoverSendCodeRequest
	14, // 7: authSAS.Auth.PasswordRecover:input_type -> authSAS.PasswordRecoverRequest
	1,  // 8: authSAS.Auth.Login:output_type -> authSAS.LoginResponce
	3,  // 9: authSAS.Auth.Logout:output_type -> authSAS.LogoutResponce
	5,  // 10: authSAS.Auth.LoginWith2FACode:output_type -> authSAS.LoginWith2FACodeResponce
	7,  // 11: authSAS.Auth.Register:output_type -> authSAS.RegisterResponce
	9,  // 12: authSAS.Auth.EmailVerifySendCode:output_type -> authSAS.EmailVerifySendCodeResponce
	11, // 13: authSAS.Auth.EmailVerify:output_type -> authSAS.EmailVerifyResponce
	13, // 14: authSAS.Auth.PasswordRecoverSendCode:output_type -> authSAS.PasswordRecoverSendCodeResponce
	15, // 15: authSAS.Auth.PasswordRecover:output_type -> authSAS.PasswordRecoverResponce
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_authSAS_proto_init() }
func file_authSAS_proto_init() {
	if File_authSAS_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authSAS_proto_rawDesc), len(file_authSAS_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authSAS_proto_goTypes,
		DependencyIndexes: file_authSAS_proto_depIdxs,
		MessageInfos:      file_authSAS_proto_msgTypes,
	}.Build()
	File_authSAS_proto = out.File
	file_authSAS_proto_goTypes = nil
	file_authSAS_proto_depIdxs = nil
}
