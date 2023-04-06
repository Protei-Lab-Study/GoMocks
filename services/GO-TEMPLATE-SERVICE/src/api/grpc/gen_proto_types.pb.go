// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: gen_proto_types.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Результат выполнения запроса
type ResultCode int32

const (
	ResultCode_OK                  ResultCode = 0
	ResultCode_INVALID_CREDENTIALS ResultCode = 1
	ResultCode_BADARG              ResultCode = 2
	ResultCode_PCONNREFUSED        ResultCode = 3
	ResultCode_NOT_FOUND           ResultCode = 4
	ResultCode_DB_CONNECTION_ERROR ResultCode = 5
	ResultCode_SYNTAX_ERROR        ResultCode = 6
	ResultCode_UNKNOWN_ERROR       ResultCode = 7
	ResultCode_ALREADY_EXISTS      ResultCode = 8
	ResultCode_USER_NOT_FOUND      ResultCode = 9
	ResultCode_PARTIAL_SUCCESS     ResultCode = 10
	ResultCode_CHAT_NOT_FOUND      ResultCode = 11
	ResultCode_CREATED             ResultCode = 12
	ResultCode_NOT_IMPLEMENTED     ResultCode = 13
	ResultCode_FORBIDDEN           ResultCode = 14
	ResultCode_NOT_CONFIGURED      ResultCode = 15
	ResultCode_NOT_REGISTERED      ResultCode = 16
	ResultCode_CALL_NOT_FOUND      ResultCode = 17
	ResultCode_REGISTRATION_FAILED ResultCode = 18
	ResultCode_GRPC_ERROR          ResultCode = 19
	ResultCode_USER_DELETED        ResultCode = 20
)

// Enum value maps for ResultCode.
var (
	ResultCode_name = map[int32]string{
		0:  "OK",
		1:  "INVALID_CREDENTIALS",
		2:  "BADARG",
		3:  "PCONNREFUSED",
		4:  "NOT_FOUND",
		5:  "DB_CONNECTION_ERROR",
		6:  "SYNTAX_ERROR",
		7:  "UNKNOWN_ERROR",
		8:  "ALREADY_EXISTS",
		9:  "USER_NOT_FOUND",
		10: "PARTIAL_SUCCESS",
		11: "CHAT_NOT_FOUND",
		12: "CREATED",
		13: "NOT_IMPLEMENTED",
		14: "FORBIDDEN",
		15: "NOT_CONFIGURED",
		16: "NOT_REGISTERED",
		17: "CALL_NOT_FOUND",
		18: "REGISTRATION_FAILED",
		19: "GRPC_ERROR",
		20: "USER_DELETED",
	}
	ResultCode_value = map[string]int32{
		"OK":                  0,
		"INVALID_CREDENTIALS": 1,
		"BADARG":              2,
		"PCONNREFUSED":        3,
		"NOT_FOUND":           4,
		"DB_CONNECTION_ERROR": 5,
		"SYNTAX_ERROR":        6,
		"UNKNOWN_ERROR":       7,
		"ALREADY_EXISTS":      8,
		"USER_NOT_FOUND":      9,
		"PARTIAL_SUCCESS":     10,
		"CHAT_NOT_FOUND":      11,
		"CREATED":             12,
		"NOT_IMPLEMENTED":     13,
		"FORBIDDEN":           14,
		"NOT_CONFIGURED":      15,
		"NOT_REGISTERED":      16,
		"CALL_NOT_FOUND":      17,
		"REGISTRATION_FAILED": 18,
		"GRPC_ERROR":          19,
		"USER_DELETED":        20,
	}
)

func (x ResultCode) Enum() *ResultCode {
	p := new(ResultCode)
	*p = x
	return p
}

func (x ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_gen_proto_types_proto_enumTypes[0].Descriptor()
}

func (ResultCode) Type() protoreflect.EnumType {
	return &file_gen_proto_types_proto_enumTypes[0]
}

func (x ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultCode.Descriptor instead.
func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_gen_proto_types_proto_rawDescGZIP(), []int{0}
}

// Запрос в сервис с передачей токена
type GeneralRequestWithToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Токен пользователя на уровне UC.CORE
}

func (x *GeneralRequestWithToken) Reset() {
	*x = GeneralRequestWithToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gen_proto_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralRequestWithToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralRequestWithToken) ProtoMessage() {}

func (x *GeneralRequestWithToken) ProtoReflect() protoreflect.Message {
	mi := &file_gen_proto_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralRequestWithToken.ProtoReflect.Descriptor instead.
func (*GeneralRequestWithToken) Descriptor() ([]byte, []int) {
	return file_gen_proto_types_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralRequestWithToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Ответ от сервиса
type GeneralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode ResultCode `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3,enum=protei.uc.api.ResultCode" json:"result_code,omitempty"` // Код возврата на выполнения запроса
	Error      *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                                                            // Поле для передачи ошибки, при выполнении запроса (заполняется если resultCode != OK)
}

func (x *GeneralResponse) Reset() {
	*x = GeneralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gen_proto_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResponse) ProtoMessage() {}

func (x *GeneralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gen_proto_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResponse.ProtoReflect.Descriptor instead.
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return file_gen_proto_types_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralResponse) GetResultCode() ResultCode {
	if x != nil {
		return x.ResultCode
	}
	return ResultCode_OK
}

func (x *GeneralResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason      string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`           // Строковое описание причины ошибки
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"` // Описание ошибки
	Command     string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`         // Команда в которой возникла ошибка (может быть строкой в модуле или названием метода)
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gen_proto_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_gen_proto_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_gen_proto_types_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Error) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Error) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

var File_gen_proto_types_proto protoreflect.FileDescriptor

var file_gen_proto_types_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x2e,
	0x75, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x2f, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x2e, 0x75, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x2e, 0x75,
	0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x5b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2a,
	0x91, 0x03, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x42, 0x41, 0x44, 0x41, 0x52, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x43, 0x4f, 0x4e, 0x4e, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13,
	0x44, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x08, 0x12, 0x12,
	0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x54, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f,
	0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x0d, 0x0a,
	0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x45, 0x44, 0x10, 0x0f,
	0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52,
	0x45, 0x44, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x47, 0x49,
	0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x12, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x52, 0x50, 0x43, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x13, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x14, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gen_proto_types_proto_rawDescOnce sync.Once
	file_gen_proto_types_proto_rawDescData = file_gen_proto_types_proto_rawDesc
)

func file_gen_proto_types_proto_rawDescGZIP() []byte {
	file_gen_proto_types_proto_rawDescOnce.Do(func() {
		file_gen_proto_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_gen_proto_types_proto_rawDescData)
	})
	return file_gen_proto_types_proto_rawDescData
}

var file_gen_proto_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gen_proto_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gen_proto_types_proto_goTypes = []interface{}{
	(ResultCode)(0),                 // 0: protei.uc.api.ResultCode
	(*GeneralRequestWithToken)(nil), // 1: protei.uc.api.GeneralRequestWithToken
	(*GeneralResponse)(nil),         // 2: protei.uc.api.GeneralResponse
	(*Error)(nil),                   // 3: protei.uc.api.Error
}
var file_gen_proto_types_proto_depIdxs = []int32{
	0, // 0: protei.uc.api.GeneralResponse.result_code:type_name -> protei.uc.api.ResultCode
	3, // 1: protei.uc.api.GeneralResponse.error:type_name -> protei.uc.api.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gen_proto_types_proto_init() }
func file_gen_proto_types_proto_init() {
	if File_gen_proto_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gen_proto_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralRequestWithToken); i {
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
		file_gen_proto_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResponse); i {
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
		file_gen_proto_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_gen_proto_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gen_proto_types_proto_goTypes,
		DependencyIndexes: file_gen_proto_types_proto_depIdxs,
		EnumInfos:         file_gen_proto_types_proto_enumTypes,
		MessageInfos:      file_gen_proto_types_proto_msgTypes,
	}.Build()
	File_gen_proto_types_proto = out.File
	file_gen_proto_types_proto_rawDesc = nil
	file_gen_proto_types_proto_goTypes = nil
	file_gen_proto_types_proto_depIdxs = nil
}
