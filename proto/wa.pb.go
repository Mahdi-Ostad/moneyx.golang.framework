// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: wa.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	framework "moneyx.golang.framework/proto/ariyana/framework"
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

type SendMessageResultProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDone        *bool                  `protobuf:"varint,1,req,name=IsDone" json:"IsDone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageResultProto) Reset() {
	*x = SendMessageResultProto{}
	mi := &file_wa_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageResultProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResultProto) ProtoMessage() {}

func (x *SendMessageResultProto) ProtoReflect() protoreflect.Message {
	mi := &file_wa_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResultProto.ProtoReflect.Descriptor instead.
func (*SendMessageResultProto) Descriptor() ([]byte, []int) {
	return file_wa_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageResultProto) GetIsDone() bool {
	if x != nil && x.IsDone != nil {
		return *x.IsDone
	}
	return false
}

type StatusQueryProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsValid       *bool                  `protobuf:"varint,1,req,name=IsValid" json:"IsValid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusQueryProto) Reset() {
	*x = StatusQueryProto{}
	mi := &file_wa_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusQueryProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusQueryProto) ProtoMessage() {}

func (x *StatusQueryProto) ProtoReflect() protoreflect.Message {
	mi := &file_wa_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusQueryProto.ProtoReflect.Descriptor instead.
func (*StatusQueryProto) Descriptor() ([]byte, []int) {
	return file_wa_proto_rawDescGZIP(), []int{1}
}

func (x *StatusQueryProto) GetIsValid() bool {
	if x != nil && x.IsValid != nil {
		return *x.IsValid
	}
	return false
}

type InitializeQueryProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	QrCode        *string                `protobuf:"bytes,1,req,name=QrCode" json:"QrCode,omitempty"`
	ExpireTime    *string                `protobuf:"bytes,2,req,name=ExpireTime" json:"ExpireTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitializeQueryProto) Reset() {
	*x = InitializeQueryProto{}
	mi := &file_wa_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitializeQueryProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeQueryProto) ProtoMessage() {}

func (x *InitializeQueryProto) ProtoReflect() protoreflect.Message {
	mi := &file_wa_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeQueryProto.ProtoReflect.Descriptor instead.
func (*InitializeQueryProto) Descriptor() ([]byte, []int) {
	return file_wa_proto_rawDescGZIP(), []int{2}
}

func (x *InitializeQueryProto) GetQrCode() string {
	if x != nil && x.QrCode != nil {
		return *x.QrCode
	}
	return ""
}

func (x *InitializeQueryProto) GetExpireTime() string {
	if x != nil && x.ExpireTime != nil {
		return *x.ExpireTime
	}
	return ""
}

type SendMessageCommandProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PhoneNumber   *string                `protobuf:"bytes,1,req,name=PhoneNumber" json:"PhoneNumber,omitempty"`
	Text          *string                `protobuf:"bytes,2,req,name=Text" json:"Text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageCommandProto) Reset() {
	*x = SendMessageCommandProto{}
	mi := &file_wa_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageCommandProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageCommandProto) ProtoMessage() {}

func (x *SendMessageCommandProto) ProtoReflect() protoreflect.Message {
	mi := &file_wa_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageCommandProto.ProtoReflect.Descriptor instead.
func (*SendMessageCommandProto) Descriptor() ([]byte, []int) {
	return file_wa_proto_rawDescGZIP(), []int{3}
}

func (x *SendMessageCommandProto) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *SendMessageCommandProto) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

type GreetingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      *string                `protobuf:"bytes,1,req,name=Response" json:"Response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GreetingResponse) Reset() {
	*x = GreetingResponse{}
	mi := &file_wa_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResponse) ProtoMessage() {}

func (x *GreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wa_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResponse.ProtoReflect.Descriptor instead.
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return file_wa_proto_rawDescGZIP(), []int{4}
}

func (x *GreetingResponse) GetResponse() string {
	if x != nil && x.Response != nil {
		return *x.Response
	}
	return ""
}

var File_wa_proto protoreflect.FileDescriptor

const file_wa_proto_rawDesc = "" +
	"\n" +
	"\bwa.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1eariyana/framework/common.proto\"0\n" +
	"\x16SendMessageResultProto\x12\x16\n" +
	"\x06IsDone\x18\x01 \x02(\bR\x06IsDone\",\n" +
	"\x10StatusQueryProto\x12\x18\n" +
	"\aIsValid\x18\x01 \x02(\bR\aIsValid\"N\n" +
	"\x14InitializeQueryProto\x12\x16\n" +
	"\x06QrCode\x18\x01 \x02(\tR\x06QrCode\x12\x1e\n" +
	"\n" +
	"ExpireTime\x18\x02 \x02(\tR\n" +
	"ExpireTime\"O\n" +
	"\x17SendMessageCommandProto\x12 \n" +
	"\vPhoneNumber\x18\x01 \x02(\tR\vPhoneNumber\x12\x12\n" +
	"\x04Text\x18\x02 \x02(\tR\x04Text\".\n" +
	"\x10GreetingResponse\x12\x1a\n" +
	"\bResponse\x18\x01 \x02(\tR\bResponse2\xc2\x02\n" +
	"\x0fWhatsappService\x12;\n" +
	"\x0eCheckUserState\x12\x16.google.protobuf.Empty\x1a\x11.StatusQueryProto\x12;\n" +
	"\n" +
	"Initialize\x12\x16.google.protobuf.Empty\x1a\x15.InitializeQueryProto\x12@\n" +
	"\vSendMessage\x12\x18.SendMessageCommandProto\x1a\x17.SendMessageResultProto\x125\n" +
	"\bSayHello\x12\x16.google.protobuf.Empty\x1a\x11.GreetingResponse\x12<\n" +
	"\aSignOut\x12\x1e.ariyana.framework.StringIdArg\x1a\x11.StatusQueryProtoB\x1fZ\x1dmoneyx.golang.framework/proto"

var (
	file_wa_proto_rawDescOnce sync.Once
	file_wa_proto_rawDescData []byte
)

func file_wa_proto_rawDescGZIP() []byte {
	file_wa_proto_rawDescOnce.Do(func() {
		file_wa_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wa_proto_rawDesc), len(file_wa_proto_rawDesc)))
	})
	return file_wa_proto_rawDescData
}

var file_wa_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wa_proto_goTypes = []any{
	(*SendMessageResultProto)(nil),  // 0: SendMessageResultProto
	(*StatusQueryProto)(nil),        // 1: StatusQueryProto
	(*InitializeQueryProto)(nil),    // 2: InitializeQueryProto
	(*SendMessageCommandProto)(nil), // 3: SendMessageCommandProto
	(*GreetingResponse)(nil),        // 4: GreetingResponse
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
	(*framework.StringIdArg)(nil),   // 6: ariyana.framework.StringIdArg
}
var file_wa_proto_depIdxs = []int32{
	5, // 0: WhatsappService.CheckUserState:input_type -> google.protobuf.Empty
	5, // 1: WhatsappService.Initialize:input_type -> google.protobuf.Empty
	3, // 2: WhatsappService.SendMessage:input_type -> SendMessageCommandProto
	5, // 3: WhatsappService.SayHello:input_type -> google.protobuf.Empty
	6, // 4: WhatsappService.SignOut:input_type -> ariyana.framework.StringIdArg
	1, // 5: WhatsappService.CheckUserState:output_type -> StatusQueryProto
	2, // 6: WhatsappService.Initialize:output_type -> InitializeQueryProto
	0, // 7: WhatsappService.SendMessage:output_type -> SendMessageResultProto
	4, // 8: WhatsappService.SayHello:output_type -> GreetingResponse
	1, // 9: WhatsappService.SignOut:output_type -> StatusQueryProto
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wa_proto_init() }
func file_wa_proto_init() {
	if File_wa_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wa_proto_rawDesc), len(file_wa_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wa_proto_goTypes,
		DependencyIndexes: file_wa_proto_depIdxs,
		MessageInfos:      file_wa_proto_msgTypes,
	}.Build()
	File_wa_proto = out.File
	file_wa_proto_goTypes = nil
	file_wa_proto_depIdxs = nil
}
