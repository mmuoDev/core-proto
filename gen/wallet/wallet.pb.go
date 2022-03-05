// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: wallet.proto

package wallet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionType int32

const (
	TransactionType_CREDIT TransactionType = 0
	TransactionType_DEBIT  TransactionType = 1
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "CREDIT",
		1: "DEBIT",
	}
	TransactionType_value = map[string]int32{
		"CREDIT": 0,
		"DEBIT":  1,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_wallet_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

type CreateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId       int32 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	PreviousBalance int32 `protobuf:"varint,2,opt,name=previous_balance,json=previousBalance,proto3" json:"previous_balance,omitempty"`
	CurrentBalance  int32 `protobuf:"varint,3,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWalletRequest) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateWalletRequest) GetPreviousBalance() int32 {
	if x != nil {
		return x.PreviousBalance
	}
	return 0
}

func (x *CreateWalletRequest) GetCurrentBalance() int32 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

type CreateWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateWalletResponse) Reset() {
	*x = CreateWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletResponse) ProtoMessage() {}

func (x *CreateWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletResponse.ProtoReflect.Descriptor instead.
func (*CreateWalletResponse) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWalletResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RetrieveWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int32 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *RetrieveWalletRequest) Reset() {
	*x = RetrieveWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveWalletRequest) ProtoMessage() {}

func (x *RetrieveWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveWalletRequest.ProtoReflect.Descriptor instead.
func (*RetrieveWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveWalletRequest) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type UpdateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId       int32 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	PreviousBalance int32 `protobuf:"varint,2,opt,name=previous_balance,json=previousBalance,proto3" json:"previous_balance,omitempty"`
	CurrentBalance  int32 `protobuf:"varint,3,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
}

func (x *UpdateWalletRequest) Reset() {
	*x = UpdateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWalletRequest) ProtoMessage() {}

func (x *UpdateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWalletRequest.ProtoReflect.Descriptor instead.
func (*UpdateWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateWalletRequest) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *UpdateWalletRequest) GetPreviousBalance() int32 {
	if x != nil {
		return x.PreviousBalance
	}
	return 0
}

func (x *UpdateWalletRequest) GetCurrentBalance() int32 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

type RetrieveWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId       int32 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	PreviousBalance int32 `protobuf:"varint,2,opt,name=previous_balance,json=previousBalance,proto3" json:"previous_balance,omitempty"`
	CurrentBalance  int32 `protobuf:"varint,3,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
}

func (x *RetrieveWalletResponse) Reset() {
	*x = RetrieveWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveWalletResponse) ProtoMessage() {}

func (x *RetrieveWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveWalletResponse.ProtoReflect.Descriptor instead.
func (*RetrieveWalletResponse) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveWalletResponse) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *RetrieveWalletResponse) GetPreviousBalance() int32 {
	if x != nil {
		return x.PreviousBalance
	}
	return 0
}

func (x *RetrieveWalletResponse) GetCurrentBalance() int32 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x26,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x88,
	0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x16, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2a, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52,
	0x45, 0x44, 0x49, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x49, 0x54, 0x10,
	0x01, 0x32, 0xe9, 0x01, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x49, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x0e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1d,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wallet_proto_goTypes = []interface{}{
	(TransactionType)(0),           // 0: wallet.TransactionType
	(*CreateWalletRequest)(nil),    // 1: wallet.CreateWalletRequest
	(*CreateWalletResponse)(nil),   // 2: wallet.CreateWalletResponse
	(*RetrieveWalletRequest)(nil),  // 3: wallet.RetrieveWalletRequest
	(*UpdateWalletRequest)(nil),    // 4: wallet.UpdateWalletRequest
	(*RetrieveWalletResponse)(nil), // 5: wallet.RetrieveWalletResponse
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_wallet_proto_depIdxs = []int32{
	1, // 0: wallet.Wallet.CreateWallet:input_type -> wallet.CreateWalletRequest
	4, // 1: wallet.Wallet.UpdateWallet:input_type -> wallet.UpdateWalletRequest
	3, // 2: wallet.Wallet.RetrieveWallet:input_type -> wallet.RetrieveWalletRequest
	2, // 3: wallet.Wallet.CreateWallet:output_type -> wallet.CreateWalletResponse
	6, // 4: wallet.Wallet.UpdateWallet:output_type -> google.protobuf.Empty
	5, // 5: wallet.Wallet.RetrieveWallet:output_type -> wallet.RetrieveWalletResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRequest); i {
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
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletResponse); i {
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
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveWalletRequest); i {
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
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWalletRequest); i {
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
		file_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveWalletResponse); i {
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
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		EnumInfos:         file_wallet_proto_enumTypes,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
