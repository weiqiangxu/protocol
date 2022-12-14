// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: user_error.proto

package user

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

type ERROR_CODE int32

const (
	ERROR_CODE_SuccessCode        ERROR_CODE = 0 // 成功状态吗
	ERROR_CODE_UndefinedErrorCode ERROR_CODE = 1 // 未定义错误
	ERROR_CODE_AuthFailCode       ERROR_CODE = 2 // 权限验证错误
)

// Enum value maps for ERROR_CODE.
var (
	ERROR_CODE_name = map[int32]string{
		0: "SuccessCode",
		1: "UndefinedErrorCode",
		2: "AuthFailCode",
	}
	ERROR_CODE_value = map[string]int32{
		"SuccessCode":        0,
		"UndefinedErrorCode": 1,
		"AuthFailCode":       2,
	}
)

func (x ERROR_CODE) Enum() *ERROR_CODE {
	p := new(ERROR_CODE)
	*p = x
	return p
}

func (x ERROR_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_user_error_proto_enumTypes[0].Descriptor()
}

func (ERROR_CODE) Type() protoreflect.EnumType {
	return &file_user_error_proto_enumTypes[0]
}

func (x ERROR_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR_CODE.Descriptor instead.
func (ERROR_CODE) EnumDescriptor() ([]byte, []int) {
	return file_user_error_proto_rawDescGZIP(), []int{0}
}

var File_user_error_proto protoreflect.FileDescriptor

var file_user_error_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x2a, 0x47, 0x0a, 0x0a, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x6e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x46, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x10,
	0x02, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x77, 0x65, 0x69, 0x71, 0x69, 0x61, 0x6e, 0x67, 0x78, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_error_proto_rawDescOnce sync.Once
	file_user_error_proto_rawDescData = file_user_error_proto_rawDesc
)

func file_user_error_proto_rawDescGZIP() []byte {
	file_user_error_proto_rawDescOnce.Do(func() {
		file_user_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_error_proto_rawDescData)
	})
	return file_user_error_proto_rawDescData
}

var file_user_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_error_proto_goTypes = []interface{}{
	(ERROR_CODE)(0), // 0: user.ERROR_CODE
}
var file_user_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_error_proto_init() }
func file_user_error_proto_init() {
	if File_user_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_error_proto_goTypes,
		DependencyIndexes: file_user_error_proto_depIdxs,
		EnumInfos:         file_user_error_proto_enumTypes,
	}.Build()
	File_user_error_proto = out.File
	file_user_error_proto_rawDesc = nil
	file_user_error_proto_goTypes = nil
	file_user_error_proto_depIdxs = nil
}
