// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: common.proto

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

type DataBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      DATA_STATUS `protobuf:"varint,1,opt,name=status,proto3,enum=user.DATA_STATUS" json:"status,omitempty"`        // 状态
	CreatedTime int64       `protobuf:"varint,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"` // 创建时间
	UpdatedTime int64       `protobuf:"varint,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"` // 修改时间
	Creator     int64       `protobuf:"varint,4,opt,name=creator,proto3" json:"creator,omitempty"`                            // 创建人员id
	Modifier    int64       `protobuf:"varint,5,opt,name=modifier,proto3" json:"modifier,omitempty"`                          // 修改人员id
}

func (x *DataBase) Reset() {
	*x = DataBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBase) ProtoMessage() {}

func (x *DataBase) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBase.ProtoReflect.Descriptor instead.
func (*DataBase) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *DataBase) GetStatus() DATA_STATUS {
	if x != nil {
		return x.Status
	}
	return DATA_STATUS_UNIVERSAL_DATA
}

func (x *DataBase) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *DataBase) GetUpdatedTime() int64 {
	if x != nil {
		return x.UpdatedTime
	}
	return 0
}

func (x *DataBase) GetCreator() int64 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *DataBase) GetModifier() int64 {
	if x != nil {
		return x.Modifier
	}
	return 0
}

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ERROR_CODE `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=user.ERROR_CODE" json:"error_code,omitempty"`
	DataBase  *DataBase  `protobuf:"bytes,2,opt,name=data_base,json=dataBase,proto3" json:"data_base,omitempty"`
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *CommonReply) GetErrorCode() ERROR_CODE {
	if x != nil {
		return x.ErrorCode
	}
	return ERROR_CODE_SuccessCode
}

func (x *CommonReply) GetDataBase() *DataBase {
	if x != nil {
		return x.DataBase
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0x6b, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2f, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x69,
	0x71, 0x69, 0x61, 0x6e, 0x67, 0x78, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(*DataBase)(nil),    // 0: user.DataBase
	(*CommonReply)(nil), // 1: user.CommonReply
	(DATA_STATUS)(0),    // 2: user.DATA_STATUS
	(ERROR_CODE)(0),     // 3: user.ERROR_CODE
}
var file_common_proto_depIdxs = []int32{
	2, // 0: user.DataBase.status:type_name -> user.DATA_STATUS
	3, // 1: user.CommonReply.error_code:type_name -> user.ERROR_CODE
	0, // 2: user.CommonReply.data_base:type_name -> user.DataBase
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	file_enum_proto_init()
	file_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBase); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
