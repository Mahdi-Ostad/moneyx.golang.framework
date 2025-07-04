// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.2
// source: RemoveBankTransaction.proto

package RemoveBankTransaction

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

type RollbackRemovedBankTransactionFiles struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []string               `protobuf:"bytes,1,rep,name=Ids,proto3" json:"Ids,omitempty"`
	ManagerId     string                 `protobuf:"bytes,2,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RollbackRemovedBankTransactionFiles) Reset() {
	*x = RollbackRemovedBankTransactionFiles{}
	mi := &file_RemoveBankTransaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackRemovedBankTransactionFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackRemovedBankTransactionFiles) ProtoMessage() {}

func (x *RollbackRemovedBankTransactionFiles) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveBankTransaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackRemovedBankTransactionFiles.ProtoReflect.Descriptor instead.
func (*RollbackRemovedBankTransactionFiles) Descriptor() ([]byte, []int) {
	return file_RemoveBankTransaction_proto_rawDescGZIP(), []int{0}
}

func (x *RollbackRemovedBankTransactionFiles) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *RollbackRemovedBankTransactionFiles) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *RollbackRemovedBankTransactionFiles) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RollbackRemovedBankTransactionFiles) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveBankTransactionFilesResponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDone        bool                   `protobuf:"varint,1,opt,name=IsDone,proto3" json:"IsDone,omitempty"`
	CorrelationId string                 `protobuf:"bytes,2,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveBankTransactionFilesResponseIntegratedCommand) Reset() {
	*x = RemoveBankTransactionFilesResponseIntegratedCommand{}
	mi := &file_RemoveBankTransaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveBankTransactionFilesResponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBankTransactionFilesResponseIntegratedCommand) ProtoMessage() {}

func (x *RemoveBankTransactionFilesResponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveBankTransaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBankTransactionFilesResponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RemoveBankTransactionFilesResponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_RemoveBankTransaction_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveBankTransactionFilesResponseIntegratedCommand) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *RemoveBankTransactionFilesResponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RemoveBankTransactionFilesResponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveBankTransactionFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileIds       []string               `protobuf:"bytes,1,rep,name=FileIds,proto3" json:"FileIds,omitempty"`
	ManagerId     string                 `protobuf:"bytes,2,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveBankTransactionFilesIntegratedCommand) Reset() {
	*x = RemoveBankTransactionFilesIntegratedCommand{}
	mi := &file_RemoveBankTransaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveBankTransactionFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBankTransactionFilesIntegratedCommand) ProtoMessage() {}

func (x *RemoveBankTransactionFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveBankTransaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBankTransactionFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RemoveBankTransactionFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_RemoveBankTransaction_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveBankTransactionFilesIntegratedCommand) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

func (x *RemoveBankTransactionFilesIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *RemoveBankTransactionFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RemoveBankTransactionFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveAccountingBankTransactionDocumentResponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDone        bool                   `protobuf:"varint,1,opt,name=IsDone,proto3" json:"IsDone,omitempty"`
	Failure       *common.Failure        `protobuf:"bytes,2,opt,name=Failure,proto3" json:"Failure,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) Reset() {
	*x = RemoveAccountingBankTransactionDocumentResponseIntegratedCommand{}
	mi := &file_RemoveBankTransaction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) ProtoMessage() {}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveBankTransaction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAccountingBankTransactionDocumentResponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_RemoveBankTransaction_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) GetFailure() *common.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RemoveAccountingBankTransactionDocumentResponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveAccountingBankTransactionDocumentIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountIds    []string               `protobuf:"bytes,1,rep,name=AccountIds,proto3" json:"AccountIds,omitempty"`
	TransactionId string                 `protobuf:"bytes,2,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
	IsKnown       bool                   `protobuf:"varint,3,opt,name=IsKnown,proto3" json:"IsKnown,omitempty"`
	ManagerId     string                 `protobuf:"bytes,4,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,5,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,6,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) Reset() {
	*x = RemoveAccountingBankTransactionDocumentIntegratedCommand{}
	mi := &file_RemoveBankTransaction_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAccountingBankTransactionDocumentIntegratedCommand) ProtoMessage() {}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveBankTransaction_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAccountingBankTransactionDocumentIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RemoveAccountingBankTransactionDocumentIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_RemoveBankTransaction_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) GetAccountIds() []string {
	if x != nil {
		return x.AccountIds
	}
	return nil
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) GetIsKnown() bool {
	if x != nil {
		return x.IsKnown
	}
	return false
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RemoveAccountingBankTransactionDocumentIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	IsCommited    bool                             `protobuf:"varint,1,opt,name=IsCommited,proto3" json:"IsCommited,omitempty"`
	WhatsappInfo  []*common.AccountingWhatsappInfo `protobuf:"bytes,2,rep,name=WhatsappInfo,proto3" json:"WhatsappInfo,omitempty"`
	CorrelationId string                           `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                           `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) Reset() {
	*x = CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand{}
	mi := &file_RemoveBankTransaction_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) ProtoMessage() {}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveBankTransaction_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand.ProtoReflect.Descriptor instead.
func (*CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_RemoveBankTransaction_proto_rawDescGZIP(), []int{5}
}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) GetIsCommited() bool {
	if x != nil {
		return x.IsCommited
	}
	return false
}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) GetWhatsappInfo() []*common.AccountingWhatsappInfo {
	if x != nil {
		return x.WhatsappInfo
	}
	return nil
}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_RemoveBankTransaction_proto protoreflect.FileDescriptor

const file_RemoveBankTransaction_proto_rawDesc = "" +
	"\n" +
	"\x1bRemoveBankTransaction.proto\x12\x15RemoveBankTransaction\x1a\fCommon.proto\"\x8b\x01\n" +
	"#RollbackRemovedBankTransactionFiles\x12\x10\n" +
	"\x03Ids\x18\x01 \x03(\tR\x03Ids\x12\x1c\n" +
	"\tManagerId\x18\x02 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\x83\x01\n" +
	"3RemoveBankTransactionFilesResponseIntegratedCommand\x12\x16\n" +
	"\x06IsDone\x18\x01 \x01(\bR\x06IsDone\x12$\n" +
	"\rCorrelationId\x18\x02 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x03 \x01(\tR\x02Id\"\x9b\x01\n" +
	"+RemoveBankTransactionFilesIntegratedCommand\x12\x18\n" +
	"\aFileIds\x18\x01 \x03(\tR\aFileIds\x12\x1c\n" +
	"\tManagerId\x18\x02 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xbb\x01\n" +
	"@RemoveAccountingBankTransactionDocumentResponseIntegratedCommand\x12\x16\n" +
	"\x06IsDone\x18\x01 \x01(\bR\x06IsDone\x12)\n" +
	"\aFailure\x18\x02 \x01(\v2\x0f.Common.FailureR\aFailure\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xee\x01\n" +
	"8RemoveAccountingBankTransactionDocumentIntegratedCommand\x12\x1e\n" +
	"\n" +
	"AccountIds\x18\x01 \x03(\tR\n" +
	"AccountIds\x12$\n" +
	"\rTransactionId\x18\x02 \x01(\tR\rTransactionId\x12\x18\n" +
	"\aIsKnown\x18\x03 \x01(\bR\aIsKnown\x12\x1c\n" +
	"\tManagerId\x18\x04 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x05 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x06 \x01(\tR\x02Id\"\xd7\x01\n" +
	";CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand\x12\x1e\n" +
	"\n" +
	"IsCommited\x18\x01 \x01(\bR\n" +
	"IsCommited\x12B\n" +
	"\fWhatsappInfo\x18\x02 \x03(\v2\x1e.Common.AccountingWhatsappInfoR\fWhatsappInfo\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02IdBlZ9moneyx.golang.framework/integration/RemoveBankTransaction\xaa\x02.Ariyana.Framework.Schema.RemoveBankTransactionb\x06proto3"

var (
	file_RemoveBankTransaction_proto_rawDescOnce sync.Once
	file_RemoveBankTransaction_proto_rawDescData []byte
)

func file_RemoveBankTransaction_proto_rawDescGZIP() []byte {
	file_RemoveBankTransaction_proto_rawDescOnce.Do(func() {
		file_RemoveBankTransaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_RemoveBankTransaction_proto_rawDesc), len(file_RemoveBankTransaction_proto_rawDesc)))
	})
	return file_RemoveBankTransaction_proto_rawDescData
}

var file_RemoveBankTransaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_RemoveBankTransaction_proto_goTypes = []any{
	(*RollbackRemovedBankTransactionFiles)(nil),                              // 0: RemoveBankTransaction.RollbackRemovedBankTransactionFiles
	(*RemoveBankTransactionFilesResponseIntegratedCommand)(nil),              // 1: RemoveBankTransaction.RemoveBankTransactionFilesResponseIntegratedCommand
	(*RemoveBankTransactionFilesIntegratedCommand)(nil),                      // 2: RemoveBankTransaction.RemoveBankTransactionFilesIntegratedCommand
	(*RemoveAccountingBankTransactionDocumentResponseIntegratedCommand)(nil), // 3: RemoveBankTransaction.RemoveAccountingBankTransactionDocumentResponseIntegratedCommand
	(*RemoveAccountingBankTransactionDocumentIntegratedCommand)(nil),         // 4: RemoveBankTransaction.RemoveAccountingBankTransactionDocumentIntegratedCommand
	(*CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand)(nil),      // 5: RemoveBankTransaction.CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand
	(*common.Failure)(nil),                // 6: Common.Failure
	(*common.AccountingWhatsappInfo)(nil), // 7: Common.AccountingWhatsappInfo
}
var file_RemoveBankTransaction_proto_depIdxs = []int32{
	6, // 0: RemoveBankTransaction.RemoveAccountingBankTransactionDocumentResponseIntegratedCommand.Failure:type_name -> Common.Failure
	7, // 1: RemoveBankTransaction.CommitOrRollbackRemovedAccountingDocumentsIntegratedCommand.WhatsappInfo:type_name -> Common.AccountingWhatsappInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RemoveBankTransaction_proto_init() }
func file_RemoveBankTransaction_proto_init() {
	if File_RemoveBankTransaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_RemoveBankTransaction_proto_rawDesc), len(file_RemoveBankTransaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RemoveBankTransaction_proto_goTypes,
		DependencyIndexes: file_RemoveBankTransaction_proto_depIdxs,
		MessageInfos:      file_RemoveBankTransaction_proto_msgTypes,
	}.Build()
	File_RemoveBankTransaction_proto = out.File
	file_RemoveBankTransaction_proto_goTypes = nil
	file_RemoveBankTransaction_proto_depIdxs = nil
}
