// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: sso/permissions.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetUserPermissionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserPermissionsRequest) Reset() {
	*x = GetUserPermissionsRequest{}
	mi := &file_sso_permissions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPermissionsRequest) ProtoMessage() {}

func (x *GetUserPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetUserPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserPermissionsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserPermissionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permissions   map[string]bool        `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserPermissionsResponse) Reset() {
	*x = UserPermissionsResponse{}
	mi := &file_sso_permissions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPermissionsResponse) ProtoMessage() {}

func (x *UserPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPermissionsResponse.ProtoReflect.Descriptor instead.
func (*UserPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *UserPermissionsResponse) GetPermissions() map[string]bool {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_sso_permissions_proto protoreflect.FileDescriptor

var file_sso_permissions_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x73, 0x73, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb5, 0x01,
	0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x77, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x68, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17,
	0x5a, 0x15, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_sso_permissions_proto_rawDescOnce sync.Once
	file_sso_permissions_proto_rawDescData []byte
)

func file_sso_permissions_proto_rawDescGZIP() []byte {
	file_sso_permissions_proto_rawDescOnce.Do(func() {
		file_sso_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sso_permissions_proto_rawDesc), len(file_sso_permissions_proto_rawDesc)))
	})
	return file_sso_permissions_proto_rawDescData
}

var file_sso_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sso_permissions_proto_goTypes = []any{
	(*GetUserPermissionsRequest)(nil), // 0: permissions.v1.GetUserPermissionsRequest
	(*UserPermissionsResponse)(nil),   // 1: permissions.v1.UserPermissionsResponse
	nil,                               // 2: permissions.v1.UserPermissionsResponse.PermissionsEntry
}
var file_sso_permissions_proto_depIdxs = []int32{
	2, // 0: permissions.v1.UserPermissionsResponse.permissions:type_name -> permissions.v1.UserPermissionsResponse.PermissionsEntry
	0, // 1: permissions.v1.Permissions.GetUserPermissions:input_type -> permissions.v1.GetUserPermissionsRequest
	1, // 2: permissions.v1.Permissions.GetUserPermissions:output_type -> permissions.v1.UserPermissionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sso_permissions_proto_init() }
func file_sso_permissions_proto_init() {
	if File_sso_permissions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sso_permissions_proto_rawDesc), len(file_sso_permissions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_permissions_proto_goTypes,
		DependencyIndexes: file_sso_permissions_proto_depIdxs,
		MessageInfos:      file_sso_permissions_proto_msgTypes,
	}.Build()
	File_sso_permissions_proto = out.File
	file_sso_permissions_proto_goTypes = nil
	file_sso_permissions_proto_depIdxs = nil
}
