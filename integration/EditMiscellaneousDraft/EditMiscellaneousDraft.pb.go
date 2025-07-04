// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.2
// source: EditMiscellaneousDraft.proto

package EditMiscellaneousDraft

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

type RollbackEditedMiscellaneousDraftFilesIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CorrelationId string                 `protobuf:"bytes,1,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RollbackEditedMiscellaneousDraftFilesIntegratedCommand) Reset() {
	*x = RollbackEditedMiscellaneousDraftFilesIntegratedCommand{}
	mi := &file_EditMiscellaneousDraft_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackEditedMiscellaneousDraftFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackEditedMiscellaneousDraftFilesIntegratedCommand) ProtoMessage() {}

func (x *RollbackEditedMiscellaneousDraftFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_EditMiscellaneousDraft_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackEditedMiscellaneousDraftFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*RollbackEditedMiscellaneousDraftFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_EditMiscellaneousDraft_proto_rawDescGZIP(), []int{0}
}

func (x *RollbackEditedMiscellaneousDraftFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *RollbackEditedMiscellaneousDraftFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EditMiscellaneousDraftFilesResponseIntegratedCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	FileIds       []string               `protobuf:"bytes,2,rep,name=FileIds,proto3" json:"FileIds,omitempty"`
	CorrelationId string                 `protobuf:"bytes,3,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) Reset() {
	*x = EditMiscellaneousDraftFilesResponseIntegratedCommand{}
	mi := &file_EditMiscellaneousDraft_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMiscellaneousDraftFilesResponseIntegratedCommand) ProtoMessage() {}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_EditMiscellaneousDraft_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMiscellaneousDraftFilesResponseIntegratedCommand.ProtoReflect.Descriptor instead.
func (*EditMiscellaneousDraftFilesResponseIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_EditMiscellaneousDraft_proto_rawDescGZIP(), []int{1}
}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *EditMiscellaneousDraftFilesResponseIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EditMiscellaneousDraftFilesIntegratedCommand struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Files          []*common.FileItem     `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty"`
	DeletedFileIds []string               `protobuf:"bytes,2,rep,name=DeletedFileIds,proto3" json:"DeletedFileIds,omitempty"`
	ManagerId      string                 `protobuf:"bytes,3,opt,name=ManagerId,proto3" json:"ManagerId,omitempty"`
	CorrelationId  string                 `protobuf:"bytes,4,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
	Id             string                 `protobuf:"bytes,5,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) Reset() {
	*x = EditMiscellaneousDraftFilesIntegratedCommand{}
	mi := &file_EditMiscellaneousDraft_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMiscellaneousDraftFilesIntegratedCommand) ProtoMessage() {}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_EditMiscellaneousDraft_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMiscellaneousDraftFilesIntegratedCommand.ProtoReflect.Descriptor instead.
func (*EditMiscellaneousDraftFilesIntegratedCommand) Descriptor() ([]byte, []int) {
	return file_EditMiscellaneousDraft_proto_rawDescGZIP(), []int{2}
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) GetFiles() []*common.FileItem {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) GetDeletedFileIds() []string {
	if x != nil {
		return x.DeletedFileIds
	}
	return nil
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *EditMiscellaneousDraftFilesIntegratedCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_EditMiscellaneousDraft_proto protoreflect.FileDescriptor

const file_EditMiscellaneousDraft_proto_rawDesc = "" +
	"\n" +
	"\x1cEditMiscellaneousDraft.proto\x12\x16EditMiscellaneousDraft\x1a\fCommon.proto\"n\n" +
	"6RollbackEditedMiscellaneousDraftFilesIntegratedCommand\x12$\n" +
	"\rCorrelationId\x18\x01 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x02 \x01(\tR\x02Id\"\x9e\x01\n" +
	"4EditMiscellaneousDraftFilesResponseIntegratedCommand\x12\x16\n" +
	"\x06Status\x18\x01 \x01(\bR\x06Status\x12\x18\n" +
	"\aFileIds\x18\x02 \x03(\tR\aFileIds\x12$\n" +
	"\rCorrelationId\x18\x03 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x04 \x01(\tR\x02Id\"\xd2\x01\n" +
	",EditMiscellaneousDraftFilesIntegratedCommand\x12&\n" +
	"\x05Files\x18\x01 \x03(\v2\x10.Common.FileItemR\x05Files\x12&\n" +
	"\x0eDeletedFileIds\x18\x02 \x03(\tR\x0eDeletedFileIds\x12\x1c\n" +
	"\tManagerId\x18\x03 \x01(\tR\tManagerId\x12$\n" +
	"\rCorrelationId\x18\x04 \x01(\tR\rCorrelationId\x12\x0e\n" +
	"\x02Id\x18\x05 \x01(\tR\x02IdBnZ:moneyx.golang.framework/integration/EditMiscellaneousDraft\xaa\x02/Ariyana.Framework.Schema.EditMiscellaneousDraftb\x06proto3"

var (
	file_EditMiscellaneousDraft_proto_rawDescOnce sync.Once
	file_EditMiscellaneousDraft_proto_rawDescData []byte
)

func file_EditMiscellaneousDraft_proto_rawDescGZIP() []byte {
	file_EditMiscellaneousDraft_proto_rawDescOnce.Do(func() {
		file_EditMiscellaneousDraft_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_EditMiscellaneousDraft_proto_rawDesc), len(file_EditMiscellaneousDraft_proto_rawDesc)))
	})
	return file_EditMiscellaneousDraft_proto_rawDescData
}

var file_EditMiscellaneousDraft_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_EditMiscellaneousDraft_proto_goTypes = []any{
	(*RollbackEditedMiscellaneousDraftFilesIntegratedCommand)(nil), // 0: EditMiscellaneousDraft.RollbackEditedMiscellaneousDraftFilesIntegratedCommand
	(*EditMiscellaneousDraftFilesResponseIntegratedCommand)(nil),   // 1: EditMiscellaneousDraft.EditMiscellaneousDraftFilesResponseIntegratedCommand
	(*EditMiscellaneousDraftFilesIntegratedCommand)(nil),           // 2: EditMiscellaneousDraft.EditMiscellaneousDraftFilesIntegratedCommand
	(*common.FileItem)(nil), // 3: Common.FileItem
}
var file_EditMiscellaneousDraft_proto_depIdxs = []int32{
	3, // 0: EditMiscellaneousDraft.EditMiscellaneousDraftFilesIntegratedCommand.Files:type_name -> Common.FileItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EditMiscellaneousDraft_proto_init() }
func file_EditMiscellaneousDraft_proto_init() {
	if File_EditMiscellaneousDraft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_EditMiscellaneousDraft_proto_rawDesc), len(file_EditMiscellaneousDraft_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EditMiscellaneousDraft_proto_goTypes,
		DependencyIndexes: file_EditMiscellaneousDraft_proto_depIdxs,
		MessageInfos:      file_EditMiscellaneousDraft_proto_msgTypes,
	}.Build()
	File_EditMiscellaneousDraft_proto = out.File
	file_EditMiscellaneousDraft_proto_goTypes = nil
	file_EditMiscellaneousDraft_proto_depIdxs = nil
}
