// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.2
// source: DetermineDocumentWithBankTransaction.proto

package DetermineDocumentWithBankTransaction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "moneyx.golang.framework/integration/common"
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

type RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CorrelationId string                 `protobuf:"bytes,1,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) Reset() {
	*x = RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand{}
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) ProtoMessage() {}

func (x *RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP(), []int{0}
}

func (x *RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DetermineDocumentWithBankTransactionResponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDone        bool                   `protobuf:"varint,1,opt,name=IsDone,proto3" json:"IsDone,omitempty"`
	Failure       *common.Failure        `protobuf:"bytes,2,opt,name=Failure,proto3" json:"Failure,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) Reset() {
	*x = DetermineDocumentWithBankTransactionResponseIntegratedCommand{}
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineDocumentWithBankTransactionResponseIntegratedCommand) ProtoMessage() {}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineDocumentWithBankTransactionResponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*DetermineDocumentWithBankTransactionResponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP(), []int{1}
}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) GetFailure() *common.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionResponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DetermineDocumentWithBankTransactionIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransactionId string                 `protobuf:"bytes,1,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
	FileIds       []string               `protobuf:"bytes,2,rep,name=FileIds,proto3" json:"FileIds,omitempty"`
	Date          int64                  `protobuf:"varint,3,opt,name=Date,proto3" json:"Date,omitempty"`
	IsPayment     bool                   `protobuf:"varint,4,opt,name=IsPayment,proto3" json:"IsPayment,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,5,opt,name=IsPrivate,proto3" json:"IsPrivate,omitempty"`
	BranchName    string                 `protobuf:"bytes,6,opt,name=BranchName,proto3" json:"BranchName,omitempty"`
	OpponentAcc   *common.AccountingInfo `protobuf:"bytes,7,opt,name=OpponentAcc,proto3" json:"OpponentAcc,omitempty"`
	BankAccountId string                 `protobuf:"bytes,8,opt,name=BankAccountId,proto3" json:"BankAccountId,omitempty"`
	ReceiptNumber string                 `protobuf:"bytes,9,opt,name=ReceiptNumber,proto3" json:"ReceiptNumber,omitempty"`
	Description   string                 `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	UserId        string                 `protobuf:"bytes,11,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ManagerId     string                 `protobuf:"bytes,12,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,13,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,14,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) Reset() {
	*x = DetermineDocumentWithBankTransactionIntegratedCommand{}
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineDocumentWithBankTransactionIntegratedCommand) ProtoMessage() {}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineDocumentWithBankTransactionIntegratedCommand.ProtoReflect.Descriptor instead.
func (*DetermineDocumentWithBankTransactionIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP(), []int{2}
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetIsPayment() bool {
	if x != nil {
		return x.IsPayment
	}
	return false
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetOpponentAcc() *common.AccountingInfo {
	if x != nil {
		return x.OpponentAcc
	}
	return nil
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetBankAccountId() string {
	if x != nil {
		return x.BankAccountId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetReceiptNumber() string {
	if x != nil {
		return x.ReceiptNumber
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	NewFileIds    []string               `protobuf:"bytes,2,rep,name=NewFileIds,proto3" json:"NewFileIds,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) Reset() {
	*x = DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand{}
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) ProtoMessage() {}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP(), []int{3}
}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) GetNewFileIds() []string {
	if x != nil {
		return x.NewFileIds
	}
	return nil
}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DetermineDocumentWithBankTransactionFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*common.FileItem     `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty"`
	ManagerId     string                 `protobuf:"bytes,2,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) Reset() {
	*x = DetermineDocumentWithBankTransactionFilesIntegratedCommand{}
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineDocumentWithBankTransactionFilesIntegratedCommand) ProtoMessage() {}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineDocumentWithBankTransactionFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*DetermineDocumentWithBankTransactionFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP(), []int{4}
}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) GetFiles() []*common.FileItem {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DetermineDocumentWithBankTransactionFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsCommit      bool                   `protobuf:"varint,1,opt,name=IsCommit,proto3" json:"IsCommit,omitempty"`
	CorrelationId string                 `protobuf:"bytes,2,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) Reset() {
	*x = CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand{}
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) ProtoMessage() {}

func (x *CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_DetermineDocumentWithBankTransaction_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand.ProtoReflect.Descriptor instead.
func (*CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP(), []int{5}
}

func (x *CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) GetIsCommit() bool {
	if x != nil {
		return x.IsCommit
	}
	return false
}

func (x *CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_DetermineDocumentWithBankTransaction_proto protoreflect.FileDescriptor

const file_DetermineDocumentWithBankTransaction_proto_rawDesc = "" +
	"\n" +
	"*DetermineDocumentWithBankTransaction.proto\x12$DetermineDocumentWithBankTransaction\x1a\fCommon.proto\"{\n" +
	"CRollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand\x12$\n" +
	"\rCorrelationId\x18\x01 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x02 \x01(\tR\x02Id\"\xb8\x01\n" +
	"=DetermineDocumentWithBankTransactionResponseIntegratedCommand\x12\x16\n" +
	"\x06IsDone\x18\x01 \x01(\bR\x06IsDone\x12)\n" +
	"\aFailure\x18\x02 \x01(\v2\x0f.Common.FailureR\aFailure\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xfb\x03\n" +
	"5DetermineDocumentWithBankTransactionIntegratedCommand\x12$\n" +
	"\rTransactionId\x18\x01 \x01(\tR\rTransactionId\x12\x18\n" +
	"\aFileIds\x18\x02 \x03(\tR\aFileIds\x12\x12\n" +
	"\x04Date\x18\x03 \x01(\x03R\x04Date\x12\x1c\n" +
	"\tIsPayment\x18\x04 \x01(\bR\tIsPayment\x12\x1c\n" +
	"\tIsPrivate\x18\x05 \x01(\bR\tIsPrivate\x12\x1e\n" +
	"\n" +
	"BranchName\x18\x06 \x01(\tR\n" +
	"BranchName\x128\n" +
	"\vOpponentAcc\x18\a \x01(\v2\x16.Common.AccountingInfoR\vOpponentAcc\x12$\n" +
	"\rBankAccountId\x18\b \x01(\tR\rBankAccountId\x12$\n" +
	"\rReceiptNumber\x18\t \x01(\tR\rReceiptNumber\x12 \n" +
	"\vDescription\x18\n" +
	" \x01(\tR\vDescription\x12\x16\n" +
	"\x06UserId\x18\v \x01(\tR\x06UserId\x12\x1c\n" +
	"\tManagerId\x18\f \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\r \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x0e \x01(\tR\x02Id\"\xb2\x01\n" +
	"BDetermineDocumentWithBankTransactionFilesResponseIntegratedCommand\x12\x16\n" +
	"\x06Status\x18\x01 \x01(\bR\x06Status\x12\x1e\n" +
	"\n" +
	"NewFileIds\x18\x02 \x03(\tR\n" +
	"NewFileIds\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xb8\x01\n" +
	":DetermineDocumentWithBankTransactionFilesIntegratedCommand\x12&\n" +
	"\x05Files\x18\x01 \x03(\v2\x10.Common.FileItemR\x05Files\x12\x1c\n" +
	"\tManagerId\x18\x02 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\x9a\x01\n" +
	"FCommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand\x12\x1a\n" +
	"\bIsCommit\x18\x01 \x01(\bR\bIsCommit\x12$\n" +
	"\rCorrelationId\x18\x02 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x03 \x01(\tR\x02IdB\x8a\x01ZHmoneyx.golang.framework/integration/DetermineDocumentWithBankTransaction\xaa\x02=Ariyana.Framework.Schema.DetermineDocumentWithBankTransactionb\x06proto3"

var (
	file_DetermineDocumentWithBankTransaction_proto_rawDescOnce sync.Once
	file_DetermineDocumentWithBankTransaction_proto_rawDescData []byte
)

func file_DetermineDocumentWithBankTransaction_proto_rawDescGZIP() []byte {
	file_DetermineDocumentWithBankTransaction_proto_rawDescOnce.Do(func() {
		file_DetermineDocumentWithBankTransaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_DetermineDocumentWithBankTransaction_proto_rawDesc), len(file_DetermineDocumentWithBankTransaction_proto_rawDesc)))
	})
	return file_DetermineDocumentWithBankTransaction_proto_rawDescData
}

var file_DetermineDocumentWithBankTransaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_DetermineDocumentWithBankTransaction_proto_goTypes = []any{
	(*RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand)(nil),    // 0: DetermineDocumentWithBankTransaction.RollbackDeterminedDocumentWithBankTransactionFilesIntegratedCommand
	(*DetermineDocumentWithBankTransactionResponseIntegratedCommand)(nil),          // 1: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionResponseIntegratedCommand
	(*DetermineDocumentWithBankTransactionIntegratedCommand)(nil),                  // 2: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionIntegratedCommand
	(*DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand)(nil),     // 3: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionFilesResponseIntegratedCommand
	(*DetermineDocumentWithBankTransactionFilesIntegratedCommand)(nil),             // 4: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionFilesIntegratedCommand
	(*CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand)(nil), // 5: DetermineDocumentWithBankTransaction.CommitOrRollbackDeterminedDocumentWithBankTransactionIntegratedCommand
	(*common.Failure)(nil),        // 6: Common.Failure
	(*common.AccountingInfo)(nil), // 7: Common.AccountingInfo
	(*common.FileItem)(nil),       // 8: Common.FileItem
}
var file_DetermineDocumentWithBankTransaction_proto_depIdxs = []int32{
	6, // 0: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionResponseIntegratedCommand.Failure:type_name -> Common.Failure
	7, // 1: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionIntegratedCommand.OpponentAcc:type_name -> Common.AccountingInfo
	8, // 2: DetermineDocumentWithBankTransaction.DetermineDocumentWithBankTransactionFilesIntegratedCommand.Files:type_name -> Common.FileItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_DetermineDocumentWithBankTransaction_proto_init() }
func file_DetermineDocumentWithBankTransaction_proto_init() {
	if File_DetermineDocumentWithBankTransaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_DetermineDocumentWithBankTransaction_proto_rawDesc), len(file_DetermineDocumentWithBankTransaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DetermineDocumentWithBankTransaction_proto_goTypes,
		DependencyIndexes: file_DetermineDocumentWithBankTransaction_proto_depIdxs,
		MessageInfos:      file_DetermineDocumentWithBankTransaction_proto_msgTypes,
	}.Build()
	File_DetermineDocumentWithBankTransaction_proto = out.File
	file_DetermineDocumentWithBankTransaction_proto_goTypes = nil
	file_DetermineDocumentWithBankTransaction_proto_depIdxs = nil
}
