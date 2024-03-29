// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/rock.proto

package proto

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

type Transaction_TxType int32

const (
	Transaction_NTX              Transaction_TxType = 0
	Transaction_CTX              Transaction_TxType = 1 // config tx
	Transaction_CROSSCHAINTX     Transaction_TxType = 2 // cross chain tx
	Transaction_ANCHORTX         Transaction_TxType = 3 // anchor tx
	Transaction_CROSSCHAINTXAUTO Transaction_TxType = 4
	Transaction_ANCHORTXAUTO     Transaction_TxType = 5
	Transaction_TIMEOUTTX        Transaction_TxType = 6
)

// Enum value maps for Transaction_TxType.
var (
	Transaction_TxType_name = map[int32]string{
		0: "NTX",
		1: "CTX",
		2: "CROSSCHAINTX",
		3: "ANCHORTX",
		4: "CROSSCHAINTXAUTO",
		5: "ANCHORTXAUTO",
		6: "TIMEOUTTX",
	}
	Transaction_TxType_value = map[string]int32{
		"NTX":              0,
		"CTX":              1,
		"CROSSCHAINTX":     2,
		"ANCHORTX":         3,
		"CROSSCHAINTXAUTO": 4,
		"ANCHORTXAUTO":     5,
		"TIMEOUTTX":        6,
	}
)

func (x Transaction_TxType) Enum() *Transaction_TxType {
	p := new(Transaction_TxType)
	*p = x
	return p
}

func (x Transaction_TxType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Transaction_TxType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_rock_proto_enumTypes[0].Descriptor()
}

func (Transaction_TxType) Type() protoreflect.EnumType {
	return &file_proto_rock_proto_enumTypes[0]
}

func (x Transaction_TxType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Transaction_TxType.Descriptor instead.
func (Transaction_TxType) EnumDescriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{7, 0}
}

type AddTxsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *AddTxsRequest) Reset() {
	*x = AddTxsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTxsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTxsRequest) ProtoMessage() {}

func (x *AddTxsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTxsRequest.ProtoReflect.Descriptor instead.
func (*AddTxsRequest) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{0}
}

func (x *AddTxsRequest) GetTxs() []*Transaction {
	if x != nil {
		return x.Txs
	}
	return nil
}

type SendBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch *Batch `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *SendBatchRequest) Reset() {
	*x = SendBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBatchRequest) ProtoMessage() {}

func (x *SendBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBatchRequest.ProtoReflect.Descriptor instead.
func (*SendBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{1}
}

func (x *SendBatchRequest) GetBatch() *Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

type AddTxsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0: success, 1: fail
	Status uint64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddTxsResponse) Reset() {
	*x = AddTxsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTxsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTxsResponse) ProtoMessage() {}

func (x *AddTxsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTxsResponse.ProtoReflect.Descriptor instead.
func (*AddTxsResponse) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{2}
}

func (x *AddTxsResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type AckBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AckBatchRequest) Reset() {
	*x = AckBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckBatchRequest) ProtoMessage() {}

func (x *AckBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckBatchRequest.ProtoReflect.Descriptor instead.
func (*AckBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{3}
}

func (x *AckBatchRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FetchBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FetchBatchRequest) Reset() {
	*x = FetchBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBatchRequest) ProtoMessage() {}

func (x *FetchBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchBatchRequest.ProtoReflect.Descriptor instead.
func (*FetchBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{4}
}

func (x *FetchBatchRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FetchBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0: success, 1: fail
	Status uint64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Batch  *Batch `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *FetchBatchResponse) Reset() {
	*x = FetchBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBatchResponse) ProtoMessage() {}

func (x *FetchBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchBatchResponse.ProtoReflect.Descriptor instead.
func (*FetchBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{5}
}

func (x *FetchBatchResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FetchBatchResponse) GetBatch() *Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Txs []*Transaction `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{6}
}

func (x *Batch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Batch) GetTxs() []*Transaction {
	if x != nil {
		return x.Txs
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version             []byte             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	From                []byte             `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                  []byte             `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value               []byte             `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp           int64              `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature           []byte             `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Id                  uint64             `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	TransactionHash     []byte             `protobuf:"bytes,8,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	Nonce               int64              `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Other               *NonHash           `protobuf:"bytes,10,opt,name=other,proto3" json:"other,omitempty"`
	TxType              Transaction_TxType `protobuf:"varint,11,opt,name=txType,proto3,enum=rock.Transaction_TxType" json:"txType,omitempty"`
	CName               []byte             `protobuf:"bytes,12,opt,name=cName,proto3" json:"cName,omitempty"`
	ExpirationTimestamp int64              `protobuf:"varint,13,opt,name=expiration_timestamp,json=expirationTimestamp,proto3" json:"expiration_timestamp,omitempty"`
	Participant         *Participant       `protobuf:"bytes,14,opt,name=participant,proto3" json:"participant,omitempty"`
	ExtraData           string             `protobuf:"bytes,15,opt,name=extraData,proto3" json:"extraData,omitempty"`
	Amount              int64              `protobuf:"varint,16,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{7}
}

func (x *Transaction) GetVersion() []byte {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Transaction) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Transaction) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Transaction) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Transaction) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Transaction) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *Transaction) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetOther() *NonHash {
	if x != nil {
		return x.Other
	}
	return nil
}

func (x *Transaction) GetTxType() Transaction_TxType {
	if x != nil {
		return x.TxType
	}
	return Transaction_NTX
}

func (x *Transaction) GetCName() []byte {
	if x != nil {
		return x.CName
	}
	return nil
}

func (x *Transaction) GetExpirationTimestamp() int64 {
	if x != nil {
		return x.ExpirationTimestamp
	}
	return 0
}

func (x *Transaction) GetParticipant() *Participant {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *Transaction) GetExtraData() string {
	if x != nil {
		return x.ExtraData
	}
	return ""
}

func (x *Transaction) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type NonHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeHash      []byte   `protobuf:"bytes,1,opt,name=nodeHash,proto3" json:"nodeHash,omitempty"`
	PrivateTxHash []byte   `protobuf:"bytes,2,opt,name=privateTxHash,proto3" json:"privateTxHash,omitempty"`
	Collection    []string `protobuf:"bytes,3,rep,name=collection,proto3" json:"collection,omitempty"`
	Nonce         uint64   `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *NonHash) Reset() {
	*x = NonHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonHash) ProtoMessage() {}

func (x *NonHash) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonHash.ProtoReflect.Descriptor instead.
func (*NonHash) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{8}
}

func (x *NonHash) GetNodeHash() []byte {
	if x != nil {
		return x.NodeHash
	}
	return nil
}

func (x *NonHash) GetPrivateTxHash() []byte {
	if x != nil {
		return x.PrivateTxHash
	}
	return nil
}

func (x *NonHash) GetCollection() []string {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *NonHash) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Initiator   []byte   `protobuf:"bytes,1,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Withholding [][]byte `protobuf:"bytes,2,rep,name=withholding,proto3" json:"withholding,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_proto_rock_proto_rawDescGZIP(), []int{9}
}

func (x *Participant) GetInitiator() []byte {
	if x != nil {
		return x.Initiator
	}
	return nil
}

func (x *Participant) GetWithholding() [][]byte {
	if x != nil {
		return x.Withholding
	}
	return nil
}

var File_proto_rock_proto protoreflect.FileDescriptor

var file_proto_rock_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x63, 0x6b, 0x22, 0x34, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54,
	0x78, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x74, 0x78, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x35,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x78, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x21, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x3c, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0xeb, 0x04, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x4e, 0x6f, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x74,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x71, 0x0a, 0x06, 0x54, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e,
	0x54, 0x58, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x54, 0x58, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x52, 0x4f, 0x53, 0x53, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x54, 0x58, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x4e, 0x43, 0x48, 0x4f, 0x52, 0x54, 0x58, 0x10, 0x03, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x52, 0x4f, 0x53, 0x53, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x54, 0x58, 0x41, 0x55, 0x54,
	0x4f, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x43, 0x48, 0x4f, 0x52, 0x54, 0x58, 0x41,
	0x55, 0x54, 0x4f, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x54, 0x58, 0x10, 0x06, 0x22, 0x81, 0x01, 0x0a, 0x07, 0x4e, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x68, 0x6f, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68,
	0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x32, 0xfb, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x70,
	0x6f, 0x6f, 0x6c, 0x12, 0x35, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x54, 0x78, 0x73, 0x12, 0x13, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x78, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x78, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x41, 0x63, 0x6b, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x17, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x72, 0x6f, 0x63, 0x6b, 0x2d, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x65, 0x78, 0x74, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rock_proto_rawDescOnce sync.Once
	file_proto_rock_proto_rawDescData = file_proto_rock_proto_rawDesc
)

func file_proto_rock_proto_rawDescGZIP() []byte {
	file_proto_rock_proto_rawDescOnce.Do(func() {
		file_proto_rock_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rock_proto_rawDescData)
	})
	return file_proto_rock_proto_rawDescData
}

var file_proto_rock_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_rock_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_rock_proto_goTypes = []interface{}{
	(Transaction_TxType)(0),    // 0: rock.Transaction.TxType
	(*AddTxsRequest)(nil),      // 1: rock.AddTxsRequest
	(*SendBatchRequest)(nil),   // 2: rock.SendBatchRequest
	(*AddTxsResponse)(nil),     // 3: rock.AddTxsResponse
	(*AckBatchRequest)(nil),    // 4: rock.AckBatchRequest
	(*FetchBatchRequest)(nil),  // 5: rock.FetchBatchRequest
	(*FetchBatchResponse)(nil), // 6: rock.FetchBatchResponse
	(*Batch)(nil),              // 7: rock.Batch
	(*Transaction)(nil),        // 8: rock.Transaction
	(*NonHash)(nil),            // 9: rock.NonHash
	(*Participant)(nil),        // 10: rock.Participant
}
var file_proto_rock_proto_depIdxs = []int32{
	8,  // 0: rock.AddTxsRequest.txs:type_name -> rock.Transaction
	7,  // 1: rock.SendBatchRequest.batch:type_name -> rock.Batch
	7,  // 2: rock.FetchBatchResponse.batch:type_name -> rock.Batch
	8,  // 3: rock.Batch.txs:type_name -> rock.Transaction
	9,  // 4: rock.Transaction.other:type_name -> rock.NonHash
	0,  // 5: rock.Transaction.txType:type_name -> rock.Transaction.TxType
	10, // 6: rock.Transaction.participant:type_name -> rock.Participant
	1,  // 7: rock.Mempool.AddTxs:input_type -> rock.AddTxsRequest
	2,  // 8: rock.Mempool.SendBatch:input_type -> rock.SendBatchRequest
	4,  // 9: rock.Mempool.AckBatch:input_type -> rock.AckBatchRequest
	5,  // 10: rock.Mempool.FetchBatch:input_type -> rock.FetchBatchRequest
	3,  // 11: rock.Mempool.AddTxs:output_type -> rock.AddTxsResponse
	3,  // 12: rock.Mempool.SendBatch:output_type -> rock.AddTxsResponse
	3,  // 13: rock.Mempool.AckBatch:output_type -> rock.AddTxsResponse
	6,  // 14: rock.Mempool.FetchBatch:output_type -> rock.FetchBatchResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_rock_proto_init() }
func file_proto_rock_proto_init() {
	if File_proto_rock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTxsRequest); i {
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
		file_proto_rock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBatchRequest); i {
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
		file_proto_rock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTxsResponse); i {
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
		file_proto_rock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckBatchRequest); i {
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
		file_proto_rock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchBatchRequest); i {
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
		file_proto_rock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchBatchResponse); i {
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
		file_proto_rock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
		file_proto_rock_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_proto_rock_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonHash); i {
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
		file_proto_rock_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
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
			RawDescriptor: file_proto_rock_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rock_proto_goTypes,
		DependencyIndexes: file_proto_rock_proto_depIdxs,
		EnumInfos:         file_proto_rock_proto_enumTypes,
		MessageInfos:      file_proto_rock_proto_msgTypes,
	}.Build()
	File_proto_rock_proto = out.File
	file_proto_rock_proto_rawDesc = nil
	file_proto_rock_proto_goTypes = nil
	file_proto_rock_proto_depIdxs = nil
}
