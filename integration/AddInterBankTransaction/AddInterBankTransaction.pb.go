// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.2
// source: AddInterBankTransaction.proto

package AddInterBankTransaction

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

type RollbackStoredBankInterTransactionFiles struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []string               `protobuf:"bytes,1,rep,name=Ids,proto3" json:"Ids,omitempty"`
	ManagerId     string                 `protobuf:"bytes,2,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RollbackStoredBankInterTransactionFiles) Reset() {
	*x = RollbackStoredBankInterTransactionFiles{}
	mi := &file_AddInterBankTransaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackStoredBankInterTransactionFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackStoredBankInterTransactionFiles) ProtoMessage() {}

func (x *RollbackStoredBankInterTransactionFiles) ProtoReflect() protoreflect.Message {
	mi := &file_AddInterBankTransaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackStoredBankInterTransactionFiles.ProtoReflect.Descriptor instead.
func (*RollbackStoredBankInterTransactionFiles) Descriptor() ([]byte, []int) {
	return file_AddInterBankTransaction_proto_rawDescGZIP(), []int{0}
}

func (x *RollbackStoredBankInterTransactionFiles) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *RollbackStoredBankInterTransactionFiles) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *RollbackStoredBankInterTransactionFiles) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RollbackStoredBankInterTransactionFiles) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StoreInterBankTransactionFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*common.FileItem     `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty"`
	ManagerId     string                 `protobuf:"bytes,2,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreInterBankTransactionFilesIntegratedCommand) Reset() {
	*x = StoreInterBankTransactionFilesIntegratedCommand{}
	mi := &file_AddInterBankTransaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreInterBankTransactionFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInterBankTransactionFilesIntegratedCommand) ProtoMessage() {}

func (x *StoreInterBankTransactionFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_AddInterBankTransaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInterBankTransactionFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*StoreInterBankTransactionFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_AddInterBankTransaction_proto_rawDescGZIP(), []int{1}
}

func (x *StoreInterBankTransactionFilesIntegratedCommand) GetFiles() []*common.FileItem {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *StoreInterBankTransactionFilesIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *StoreInterBankTransactionFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *StoreInterBankTransactionFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StoreInterBankTransactionFilesResponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDone        bool                   `protobuf:"varint,1,opt,name=IsDone,proto3" json:"IsDone,omitempty"`
	Files         []string               `protobuf:"bytes,2,rep,name=Files,proto3" json:"Files,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) Reset() {
	*x = StoreInterBankTransactionFilesResponseIntegratedCommand{}
	mi := &file_AddInterBankTransaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInterBankTransactionFilesResponseIntegratedCommand) ProtoMessage() {}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_AddInterBankTransaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInterBankTransactionFilesResponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*StoreInterBankTransactionFilesResponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_AddInterBankTransaction_proto_rawDescGZIP(), []int{2}
}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *StoreInterBankTransactionFilesResponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_AddInterBankTransaction_proto protoreflect.FileDescriptor

const file_AddInterBankTransaction_proto_rawDesc = "" +
	"\n" +
	"\x1dAddInterBankTransaction.proto\x12\x17AddInterBankTransaction\x1a\fCommon.proto\"\x8f\x01\n" +
	"'RollbackStoredBankInterTransactionFiles\x12\x10\n" +
	"\x03Ids\x18\x01 \x03(\tR\x03Ids\x12\x1c\n" +
	"\tManagerId\x18\x02 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xad\x01\n" +
	"/StoreInterBankTransactionFilesIntegratedCommand\x12&\n" +
	"\x05Files\x18\x01 \x03(\v2\x10.Common.FileItemR\x05Files\x12\x1c\n" +
	"\tManagerId\x18\x02 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\x9d\x01\n" +
	"7StoreInterBankTransactionFilesResponseIntegratedCommand\x12\x16\n" +
	"\x06IsDone\x18\x01 \x01(\bR\x06IsDone\x12\x14\n" +
	"\x05Files\x18\x02 \x03(\tR\x05Files\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02IdBpZ;moneyx.golang.framework/integration/AddInterBankTransaction\xaa\x020Ariyana.Framework.Schema.AddInterBankTransactionb\x06proto3"

var (
	file_AddInterBankTransaction_proto_rawDescOnce sync.Once
	file_AddInterBankTransaction_proto_rawDescData []byte
)

func file_AddInterBankTransaction_proto_rawDescGZIP() []byte {
	file_AddInterBankTransaction_proto_rawDescOnce.Do(func() {
		file_AddInterBankTransaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_AddInterBankTransaction_proto_rawDesc), len(file_AddInterBankTransaction_proto_rawDesc)))
	})
	return file_AddInterBankTransaction_proto_rawDescData
}

var file_AddInterBankTransaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_AddInterBankTransaction_proto_goTypes = []any{
	(*RollbackStoredBankInterTransactionFiles)(nil),                 // 0: AddInterBankTransaction.RollbackStoredBankInterTransactionFiles
	(*StoreInterBankTransactionFilesIntegratedCommand)(nil),         // 1: AddInterBankTransaction.StoreInterBankTransactionFilesIntegratedCommand
	(*StoreInterBankTransactionFilesResponseIntegratedCommand)(nil), // 2: AddInterBankTransaction.StoreInterBankTransactionFilesResponseIntegratedCommand
	(*common.FileItem)(nil),                                         // 3: Common.FileItem
}
var file_AddInterBankTransaction_proto_depIdxs = []int32{
	3, // 0: AddInterBankTransaction.StoreInterBankTransactionFilesIntegratedCommand.Files:type_name -> Common.FileItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AddInterBankTransaction_proto_init() }
func file_AddInterBankTransaction_proto_init() {
	if File_AddInterBankTransaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_AddInterBankTransaction_proto_rawDesc), len(file_AddInterBankTransaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddInterBankTransaction_proto_goTypes,
		DependencyIndexes: file_AddInterBankTransaction_proto_depIdxs,
		MessageInfos:      file_AddInterBankTransaction_proto_msgTypes,
	}.Build()
	File_AddInterBankTransaction_proto = out.File
	file_AddInterBankTransaction_proto_goTypes = nil
	file_AddInterBankTransaction_proto_depIdxs = nil
}
