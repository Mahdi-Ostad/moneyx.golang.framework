// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.2
// source: ExecutePayReceipt.proto

package ExecutePayReceipt

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

type RollbackExecutedPayReceiptFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CorrelationId string                 `protobuf:"bytes,1,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RollbackExecutedPayReceiptFilesIntegratedCommand) Reset() {
	*x = RollbackExecutedPayReceiptFilesIntegratedCommand{}
	mi := &file_ExecutePayReceipt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackExecutedPayReceiptFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackExecutedPayReceiptFilesIntegratedCommand) ProtoMessage() {}

func (x *RollbackExecutedPayReceiptFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutePayReceipt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackExecutedPayReceiptFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RollbackExecutedPayReceiptFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_ExecutePayReceipt_proto_rawDescGZIP(), []int{0}
}

func (x *RollbackExecutedPayReceiptFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RollbackExecutedPayReceiptFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ExecutePayReceiptReponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDone        bool                   `protobuf:"varint,1,opt,name=IsDone,proto3" json:"IsDone,omitempty"`
	Failure       *common.Failure        `protobuf:"bytes,2,opt,name=Failure,proto3" json:"Failure,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutePayReceiptReponseIntegratedCommand) Reset() {
	*x = ExecutePayReceiptReponseIntegratedCommand{}
	mi := &file_ExecutePayReceipt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutePayReceiptReponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePayReceiptReponseIntegratedCommand) ProtoMessage() {}

func (x *ExecutePayReceiptReponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutePayReceipt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePayReceiptReponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*ExecutePayReceiptReponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_ExecutePayReceipt_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutePayReceiptReponseIntegratedCommand) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *ExecutePayReceiptReponseIntegratedCommand) GetFailure() *common.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *ExecutePayReceiptReponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ExecutePayReceiptReponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ExecutePayReceiptIntegratedCommand struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	DebtorAccountId    string                 `protobuf:"bytes,1,opt,name=DebtorAccountId,proto3" json:"DebtorAccountId,omitempty"`
	CreditorAccountId  string                 `protobuf:"bytes,2,opt,name=CreditorAccountId,proto3" json:"CreditorAccountId,omitempty"`
	IsPrivate          bool                   `protobuf:"varint,3,opt,name=IsPrivate,proto3" json:"IsPrivate,omitempty"`
	Date               int64                  `protobuf:"varint,4,opt,name=Date,proto3" json:"Date,omitempty"`
	DebtorCommission   *common.MoneyDocument  `protobuf:"bytes,5,opt,name=DebtorCommission,proto3" json:"DebtorCommission,omitempty"`
	CreditorCommission *common.MoneyDocument  `protobuf:"bytes,6,opt,name=CreditorCommission,proto3" json:"CreditorCommission,omitempty"`
	UpdateCreditLimit  bool                   `protobuf:"varint,7,opt,name=UpdateCreditLimit,proto3" json:"UpdateCreditLimit,omitempty"`
	FileIds            []string               `protobuf:"bytes,8,rep,name=FileIds,proto3" json:"FileIds,omitempty"`
	Money              *common.Money          `protobuf:"bytes,9,opt,name=Money,proto3" json:"Money,omitempty"`
	Description        string                 `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	ReceiptId          int32                  `protobuf:"varint,11,opt,name=ReceiptId,proto3" json:"ReceiptId,omitempty"`
	MangerId           string                 `protobuf:"bytes,12,opt,name=MangerId,proto3" json:"MangerId,omitempty"`
	CorrelationId      string                 `protobuf:"bytes,13,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id                 string                 `protobuf:"bytes,14,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ExecutePayReceiptIntegratedCommand) Reset() {
	*x = ExecutePayReceiptIntegratedCommand{}
	mi := &file_ExecutePayReceipt_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutePayReceiptIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePayReceiptIntegratedCommand) ProtoMessage() {}

func (x *ExecutePayReceiptIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutePayReceipt_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePayReceiptIntegratedCommand.ProtoReflect.Descriptor instead.
func (*ExecutePayReceiptIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_ExecutePayReceipt_proto_rawDescGZIP(), []int{2}
}

func (x *ExecutePayReceiptIntegratedCommand) GetDebtorAccountId() string {
	if x != nil {
		return x.DebtorAccountId
	}
	return ""
}

func (x *ExecutePayReceiptIntegratedCommand) GetCreditorAccountId() string {
	if x != nil {
		return x.CreditorAccountId
	}
	return ""
}

func (x *ExecutePayReceiptIntegratedCommand) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *ExecutePayReceiptIntegratedCommand) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *ExecutePayReceiptIntegratedCommand) GetDebtorCommission() *common.MoneyDocument {
	if x != nil {
		return x.DebtorCommission
	}
	return nil
}

func (x *ExecutePayReceiptIntegratedCommand) GetCreditorCommission() *common.MoneyDocument {
	if x != nil {
		return x.CreditorCommission
	}
	return nil
}

func (x *ExecutePayReceiptIntegratedCommand) GetUpdateCreditLimit() bool {
	if x != nil {
		return x.UpdateCreditLimit
	}
	return false
}

func (x *ExecutePayReceiptIntegratedCommand) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

func (x *ExecutePayReceiptIntegratedCommand) GetMoney() *common.Money {
	if x != nil {
		return x.Money
	}
	return nil
}

func (x *ExecutePayReceiptIntegratedCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ExecutePayReceiptIntegratedCommand) GetReceiptId() int32 {
	if x != nil {
		return x.ReceiptId
	}
	return 0
}

func (x *ExecutePayReceiptIntegratedCommand) GetMangerId() string {
	if x != nil {
		return x.MangerId
	}
	return ""
}

func (x *ExecutePayReceiptIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ExecutePayReceiptIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ExecutePayReceiptFilesReponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	FileIds       []string               `protobuf:"bytes,2,rep,name=FileIds,proto3" json:"FileIds,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) Reset() {
	*x = ExecutePayReceiptFilesReponseIntegratedCommand{}
	mi := &file_ExecutePayReceipt_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePayReceiptFilesReponseIntegratedCommand) ProtoMessage() {}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutePayReceipt_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePayReceiptFilesReponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*ExecutePayReceiptFilesReponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_ExecutePayReceipt_proto_rawDescGZIP(), []int{3}
}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ExecutePayReceiptFilesReponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ExecutePayReceiptFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*common.FileItem     `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty"`
	ManagerId     string                 `protobuf:"bytes,2,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecutePayReceiptFilesIntegratedCommand) Reset() {
	*x = ExecutePayReceiptFilesIntegratedCommand{}
	mi := &file_ExecutePayReceipt_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutePayReceiptFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePayReceiptFilesIntegratedCommand) ProtoMessage() {}

func (x *ExecutePayReceiptFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutePayReceipt_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePayReceiptFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*ExecutePayReceiptFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_ExecutePayReceipt_proto_rawDescGZIP(), []int{4}
}

func (x *ExecutePayReceiptFilesIntegratedCommand) GetFiles() []*common.FileItem {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ExecutePayReceiptFilesIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *ExecutePayReceiptFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ExecutePayReceiptFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_ExecutePayReceipt_proto protoreflect.FileDescriptor

const file_ExecutePayReceipt_proto_rawDesc = "" +
	"\n" +
	"\x17ExecutePayReceipt.proto\x12\x11ExecutePayReceipt\x1a\fCommon.proto\"h\n" +
	"0RollbackExecutedPayReceiptFilesIntegratedCommand\x12$\n" +
	"\rCorrelationId\x18\x01 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x02 \x01(\tR\x02Id\"\xa4\x01\n" +
	")ExecutePayReceiptReponseIntegratedCommand\x12\x16\n" +
	"\x06IsDone\x18\x01 \x01(\bR\x06IsDone\x12)\n" +
	"\aFailure\x18\x02 \x01(\v2\x0f.Common.FailureR\aFailure\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xb7\x04\n" +
	"\"ExecutePayReceiptIntegratedCommand\x12(\n" +
	"\x0fDebtorAccountId\x18\x01 \x01(\tR\x0fDebtorAccountId\x12,\n" +
	"\x11CreditorAccountId\x18\x02 \x01(\tR\x11CreditorAccountId\x12\x1c\n" +
	"\tIsPrivate\x18\x03 \x01(\bR\tIsPrivate\x12\x12\n" +
	"\x04Date\x18\x04 \x01(\x03R\x04Date\x12A\n" +
	"\x10DebtorCommission\x18\x05 \x01(\v2\x15.Common.MoneyDocumentR\x10DebtorCommission\x12E\n" +
	"\x12CreditorCommission\x18\x06 \x01(\v2\x15.Common.MoneyDocumentR\x12CreditorCommission\x12,\n" +
	"\x11UpdateCreditLimit\x18\a \x01(\bR\x11UpdateCreditLimit\x12\x18\n" +
	"\aFileIds\x18\b \x03(\tR\aFileIds\x12#\n" +
	"\x05Money\x18\t \x01(\v2\r.Common.MoneyR\x05Money\x12 \n" +
	"\vDescription\x18\n" +
	" \x01(\tR\vDescription\x12\x1c\n" +
	"\tReceiptId\x18\v \x01(\x05R\tReceiptId\x12\x1a\n" +
	"\bMangerId\x18\f \x01(\tR\bMangerId\x12$\n" +
	"\rCorrelationId\x18\r \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x0e \x01(\tR\x02Id\"\x98\x01\n" +
	".ExecutePayReceiptFilesReponseIntegratedCommand\x12\x16\n" +
	"\x06Status\x18\x01 \x01(\bR\x06Status\x12\x18\n" +
	"\aFileIds\x18\x02 \x03(\tR\aFileIds\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xa5\x01\n" +
	"'ExecutePayReceiptFilesIntegratedCommand\x12&\n" +
	"\x05Files\x18\x01 \x03(\v2\x10.Common.FileItemR\x05Files\x12\x1c\n" +
	"\tManagerId\x18\x02 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02IdBdZ5moneyx.golang.framework/integration/ExecutePayReceipt\xaa\x02*Ariyana.Framework.Schema.ExecutePayReceiptb\x06proto3"

var (
	file_ExecutePayReceipt_proto_rawDescOnce sync.Once
	file_ExecutePayReceipt_proto_rawDescData []byte
)

func file_ExecutePayReceipt_proto_rawDescGZIP() []byte {
	file_ExecutePayReceipt_proto_rawDescOnce.Do(func() {
		file_ExecutePayReceipt_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ExecutePayReceipt_proto_rawDesc), len(file_ExecutePayReceipt_proto_rawDesc)))
	})
	return file_ExecutePayReceipt_proto_rawDescData
}

var file_ExecutePayReceipt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ExecutePayReceipt_proto_goTypes = []any{
	(*RollbackExecutedPayReceiptFilesIntegratedCommand)(nil), // 0: ExecutePayReceipt.RollbackExecutedPayReceiptFilesIntegratedCommand
	(*ExecutePayReceiptReponseIntegratedCommand)(nil),        // 1: ExecutePayReceipt.ExecutePayReceiptReponseIntegratedCommand
	(*ExecutePayReceiptIntegratedCommand)(nil),               // 2: ExecutePayReceipt.ExecutePayReceiptIntegratedCommand
	(*ExecutePayReceiptFilesReponseIntegratedCommand)(nil),   // 3: ExecutePayReceipt.ExecutePayReceiptFilesReponseIntegratedCommand
	(*ExecutePayReceiptFilesIntegratedCommand)(nil),          // 4: ExecutePayReceipt.ExecutePayReceiptFilesIntegratedCommand
	(*common.Failure)(nil),                                   // 5: Common.Failure
	(*common.MoneyDocument)(nil),                             // 6: Common.MoneyDocument
	(*common.Money)(nil),                                     // 7: Common.Money
	(*common.FileItem)(nil),                                  // 8: Common.FileItem
}
var file_ExecutePayReceipt_proto_depIdxs = []int32{
	5, // 0: ExecutePayReceipt.ExecutePayReceiptReponseIntegratedCommand.Failure:type_name -> Common.Failure
	6, // 1: ExecutePayReceipt.ExecutePayReceiptIntegratedCommand.DebtorCommission:type_name -> Common.MoneyDocument
	6, // 2: ExecutePayReceipt.ExecutePayReceiptIntegratedCommand.CreditorCommission:type_name -> Common.MoneyDocument
	7, // 3: ExecutePayReceipt.ExecutePayReceiptIntegratedCommand.Money:type_name -> Common.Money
	8, // 4: ExecutePayReceipt.ExecutePayReceiptFilesIntegratedCommand.Files:type_name -> Common.FileItem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ExecutePayReceipt_proto_init() }
func file_ExecutePayReceipt_proto_init() {
	if File_ExecutePayReceipt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ExecutePayReceipt_proto_rawDesc), len(file_ExecutePayReceipt_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExecutePayReceipt_proto_goTypes,
		DependencyIndexes: file_ExecutePayReceipt_proto_depIdxs,
		MessageInfos:      file_ExecutePayReceipt_proto_msgTypes,
	}.Build()
	File_ExecutePayReceipt_proto = out.File
	file_ExecutePayReceipt_proto_goTypes = nil
	file_ExecutePayReceipt_proto_depIdxs = nil
}
