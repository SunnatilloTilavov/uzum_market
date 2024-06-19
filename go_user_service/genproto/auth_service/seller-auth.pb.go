// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: seller-auth.proto

package auth_service

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

type SellerEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SellerEmpty) Reset() {
	*x = SellerEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerEmpty) ProtoMessage() {}

func (x *SellerEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerEmpty.ProtoReflect.Descriptor instead.
func (*SellerEmpty) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{0}
}

type SellerCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail    string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SellerCreateRequest) Reset() {
	*x = SellerCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerCreateRequest) ProtoMessage() {}

func (x *SellerCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerCreateRequest.ProtoReflect.Descriptor instead.
func (*SellerCreateRequest) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SellerCreateRequest) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *SellerCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SellerRConfirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp      string `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	Gmail    string `protobuf:"bytes,2,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender   string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Language string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *SellerRConfirm) Reset() {
	*x = SellerRConfirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerRConfirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerRConfirm) ProtoMessage() {}

func (x *SellerRConfirm) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerRConfirm.ProtoReflect.Descriptor instead.
func (*SellerRConfirm) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SellerRConfirm) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *SellerRConfirm) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *SellerRConfirm) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SellerRConfirm) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *SellerRConfirm) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type SellerLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail    string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SellerLoginRequest) Reset() {
	*x = SellerLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerLoginRequest) ProtoMessage() {}

func (x *SellerLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerLoginRequest.ProtoReflect.Descriptor instead.
func (*SellerLoginRequest) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SellerLoginRequest) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *SellerLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SellerLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accesstoken  string `protobuf:"bytes,1,opt,name=accesstoken,proto3" json:"accesstoken,omitempty"`
	Refreshtoken string `protobuf:"bytes,2,opt,name=refreshtoken,proto3" json:"refreshtoken,omitempty"`
}

func (x *SellerLoginResponse) Reset() {
	*x = SellerLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerLoginResponse) ProtoMessage() {}

func (x *SellerLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerLoginResponse.ProtoReflect.Descriptor instead.
func (*SellerLoginResponse) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SellerLoginResponse) GetAccesstoken() string {
	if x != nil {
		return x.Accesstoken
	}
	return ""
}

func (x *SellerLoginResponse) GetRefreshtoken() string {
	if x != nil {
		return x.Refreshtoken
	}
	return ""
}

type SellerGmailCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *SellerGmailCheckRequest) Reset() {
	*x = SellerGmailCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerGmailCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerGmailCheckRequest) ProtoMessage() {}

func (x *SellerGmailCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerGmailCheckRequest.ProtoReflect.Descriptor instead.
func (*SellerGmailCheckRequest) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{5}
}

func (x *SellerGmailCheckRequest) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

type SellerGmailCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SellerGmailCheckResponse) Reset() {
	*x = SellerGmailCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerGmailCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerGmailCheckResponse) ProtoMessage() {}

func (x *SellerGmailCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerGmailCheckResponse.ProtoReflect.Descriptor instead.
func (*SellerGmailCheckResponse) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{6}
}

func (x *SellerGmailCheckResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SellerLoginByGmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp   string `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	Gmail string `protobuf:"bytes,2,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *SellerLoginByGmailRequest) Reset() {
	*x = SellerLoginByGmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerLoginByGmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerLoginByGmailRequest) ProtoMessage() {}

func (x *SellerLoginByGmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seller_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerLoginByGmailRequest.ProtoReflect.Descriptor instead.
func (*SellerLoginByGmailRequest) Descriptor() ([]byte, []int) {
	return file_seller_auth_proto_rawDescGZIP(), []int{7}
}

func (x *SellerLoginByGmailRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *SellerLoginByGmailRequest) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

var File_seller_auth_proto protoreflect.FileDescriptor

var file_seller_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x47, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x52, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5b, 0x0a, 0x13,
	0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2f, 0x0a, 0x17, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x36, 0x0a, 0x18, 0x53, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x43, 0x0a, 0x19, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xa8, 0x07, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x5e, 0x0a, 0x15, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x10, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x47, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x47, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x14, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x79, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x1b, 0x53, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x12, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61,
	0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x19, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6d, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x59,
	0x0a, 0x13, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x1a, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_seller_auth_proto_rawDescOnce sync.Once
	file_seller_auth_proto_rawDescData = file_seller_auth_proto_rawDesc
)

func file_seller_auth_proto_rawDescGZIP() []byte {
	file_seller_auth_proto_rawDescOnce.Do(func() {
		file_seller_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_seller_auth_proto_rawDescData)
	})
	return file_seller_auth_proto_rawDescData
}

var file_seller_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_seller_auth_proto_goTypes = []interface{}{
	(*SellerEmpty)(nil),               // 0: auth_service.SellerEmpty
	(*SellerCreateRequest)(nil),       // 1: auth_service.SellerCreateRequest
	(*SellerRConfirm)(nil),            // 2: auth_service.SellerRConfirm
	(*SellerLoginRequest)(nil),        // 3: auth_service.SellerLoginRequest
	(*SellerLoginResponse)(nil),       // 4: auth_service.SellerLoginResponse
	(*SellerGmailCheckRequest)(nil),   // 5: auth_service.SellerGmailCheckRequest
	(*SellerGmailCheckResponse)(nil),  // 6: auth_service.SellerGmailCheckResponse
	(*SellerLoginByGmailRequest)(nil), // 7: auth_service.SellerLoginByGmailRequest
}
var file_seller_auth_proto_depIdxs = []int32{
	3,  // 0: auth_service.SellerAuth.SellerLoginByPassword:input_type -> auth_service.SellerLoginRequest
	5,  // 1: auth_service.SellerAuth.SellerGmailCheck:input_type -> auth_service.SellerGmailCheckRequest
	5,  // 2: auth_service.SellerAuth.SellerRegisterByMail:input_type -> auth_service.SellerGmailCheckRequest
	2,  // 3: auth_service.SellerAuth.SellerRegisterByMailConfirm:input_type -> auth_service.SellerRConfirm
	1,  // 4: auth_service.SellerAuth.SellerCreate:input_type -> auth_service.SellerCreateRequest
	5,  // 5: auth_service.SellerAuth.SellerLoginByGmail:input_type -> auth_service.SellerGmailCheckRequest
	7,  // 6: auth_service.SellerAuth.SellerLoginByGmailComfirm:input_type -> auth_service.SellerLoginByGmailRequest
	1,  // 7: auth_service.SellerAuth.SellerUpdatePassword:input_type -> auth_service.SellerCreateRequest
	5,  // 8: auth_service.SellerAuth.SellerResetPassword:input_type -> auth_service.SellerGmailCheckRequest
	2,  // 9: auth_service.SellerAuth.SellerResetPasswordConfirm:input_type -> auth_service.SellerRConfirm
	4,  // 10: auth_service.SellerAuth.SellerLoginByPassword:output_type -> auth_service.SellerLoginResponse
	6,  // 11: auth_service.SellerAuth.SellerGmailCheck:output_type -> auth_service.SellerGmailCheckResponse
	0,  // 12: auth_service.SellerAuth.SellerRegisterByMail:output_type -> auth_service.SellerEmpty
	0,  // 13: auth_service.SellerAuth.SellerRegisterByMailConfirm:output_type -> auth_service.SellerEmpty
	0,  // 14: auth_service.SellerAuth.SellerCreate:output_type -> auth_service.SellerEmpty
	0,  // 15: auth_service.SellerAuth.SellerLoginByGmail:output_type -> auth_service.SellerEmpty
	4,  // 16: auth_service.SellerAuth.SellerLoginByGmailComfirm:output_type -> auth_service.SellerLoginResponse
	0,  // 17: auth_service.SellerAuth.SellerUpdatePassword:output_type -> auth_service.SellerEmpty
	0,  // 18: auth_service.SellerAuth.SellerResetPassword:output_type -> auth_service.SellerEmpty
	0,  // 19: auth_service.SellerAuth.SellerResetPasswordConfirm:output_type -> auth_service.SellerEmpty
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_seller_auth_proto_init() }
func file_seller_auth_proto_init() {
	if File_seller_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seller_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerEmpty); i {
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
		file_seller_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerCreateRequest); i {
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
		file_seller_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerRConfirm); i {
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
		file_seller_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerLoginRequest); i {
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
		file_seller_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerLoginResponse); i {
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
		file_seller_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerGmailCheckRequest); i {
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
		file_seller_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerGmailCheckResponse); i {
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
		file_seller_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerLoginByGmailRequest); i {
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
			RawDescriptor: file_seller_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seller_auth_proto_goTypes,
		DependencyIndexes: file_seller_auth_proto_depIdxs,
		MessageInfos:      file_seller_auth_proto_msgTypes,
	}.Build()
	File_seller_auth_proto = out.File
	file_seller_auth_proto_rawDesc = nil
	file_seller_auth_proto_goTypes = nil
	file_seller_auth_proto_depIdxs = nil
}
