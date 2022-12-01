// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: apply.proto

package apply

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

// 创建认证
type CreateIdentifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ScholarId string `protobuf:"bytes,2,opt,name=scholarId,proto3" json:"scholarId,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateIdentifyReq) Reset() {
	*x = CreateIdentifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdentifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentifyReq) ProtoMessage() {}

func (x *CreateIdentifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_apply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentifyReq.ProtoReflect.Descriptor instead.
func (*CreateIdentifyReq) Descriptor() ([]byte, []int) {
	return file_apply_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIdentifyReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateIdentifyReq) GetScholarId() string {
	if x != nil {
		return x.ScholarId
	}
	return ""
}

func (x *CreateIdentifyReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateIdentifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateIdentifyReply) Reset() {
	*x = CreateIdentifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apply_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdentifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentifyReply) ProtoMessage() {}

func (x *CreateIdentifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_apply_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentifyReply.ProtoReflect.Descriptor instead.
func (*CreateIdentifyReply) Descriptor() ([]byte, []int) {
	return file_apply_proto_rawDescGZIP(), []int{1}
}

// 查看用户认证的学者
type CheckIdentifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CheckIdentifyReq) Reset() {
	*x = CheckIdentifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apply_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIdentifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIdentifyReq) ProtoMessage() {}

func (x *CheckIdentifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_apply_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIdentifyReq.ProtoReflect.Descriptor instead.
func (*CheckIdentifyReq) Descriptor() ([]byte, []int) {
	return file_apply_proto_rawDescGZIP(), []int{2}
}

func (x *CheckIdentifyReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckIdentifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsScholar bool   `protobuf:"varint,1,opt,name=isScholar,proto3" json:"isScholar,omitempty"`
	ScholarId string `protobuf:"bytes,2,opt,name=scholarId,proto3" json:"scholarId,omitempty"`
}

func (x *CheckIdentifyReply) Reset() {
	*x = CheckIdentifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apply_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIdentifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIdentifyReply) ProtoMessage() {}

func (x *CheckIdentifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_apply_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIdentifyReply.ProtoReflect.Descriptor instead.
func (*CheckIdentifyReply) Descriptor() ([]byte, []int) {
	return file_apply_proto_rawDescGZIP(), []int{3}
}

func (x *CheckIdentifyReply) GetIsScholar() bool {
	if x != nil {
		return x.IsScholar
	}
	return false
}

func (x *CheckIdentifyReply) GetScholarId() string {
	if x != nil {
		return x.ScholarId
	}
	return ""
}

// 查看学者的认证用户
type CheckUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScholarId string `protobuf:"bytes,1,opt,name=scholarId,proto3" json:"scholarId,omitempty"`
}

func (x *CheckUserReq) Reset() {
	*x = CheckUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apply_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserReq) ProtoMessage() {}

func (x *CheckUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_apply_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserReq.ProtoReflect.Descriptor instead.
func (*CheckUserReq) Descriptor() ([]byte, []int) {
	return file_apply_proto_rawDescGZIP(), []int{4}
}

func (x *CheckUserReq) GetScholarId() string {
	if x != nil {
		return x.ScholarId
	}
	return ""
}

type CheckUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsVerified bool  `protobuf:"varint,1,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	UserId     int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CheckUserReply) Reset() {
	*x = CheckUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apply_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserReply) ProtoMessage() {}

func (x *CheckUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_apply_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserReply.ProtoReflect.Descriptor instead.
func (*CheckUserReply) Descriptor() ([]byte, []int) {
	return file_apply_proto_rawDescGZIP(), []int{5}
}

func (x *CheckUserReply) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *CheckUserReply) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_apply_proto protoreflect.FileDescriptor

var file_apply_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x22, 0x5b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f,
	0x6c, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68,
	0x6f, 0x6c, 0x61, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c,
	0x61, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xcd,
	0x01, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x43, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apply_proto_rawDescOnce sync.Once
	file_apply_proto_rawDescData = file_apply_proto_rawDesc
)

func file_apply_proto_rawDescGZIP() []byte {
	file_apply_proto_rawDescOnce.Do(func() {
		file_apply_proto_rawDescData = protoimpl.X.CompressGZIP(file_apply_proto_rawDescData)
	})
	return file_apply_proto_rawDescData
}

var file_apply_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apply_proto_goTypes = []interface{}{
	(*CreateIdentifyReq)(nil),   // 0: apply.CreateIdentifyReq
	(*CreateIdentifyReply)(nil), // 1: apply.CreateIdentifyReply
	(*CheckIdentifyReq)(nil),    // 2: apply.CheckIdentifyReq
	(*CheckIdentifyReply)(nil),  // 3: apply.CheckIdentifyReply
	(*CheckUserReq)(nil),        // 4: apply.CheckUserReq
	(*CheckUserReply)(nil),      // 5: apply.CheckUserReply
}
var file_apply_proto_depIdxs = []int32{
	0, // 0: apply.Apply.CreateIdentify:input_type -> apply.CreateIdentifyReq
	2, // 1: apply.Apply.CheckIdentify:input_type -> apply.CheckIdentifyReq
	4, // 2: apply.Apply.CheckUser:input_type -> apply.CheckUserReq
	1, // 3: apply.Apply.CreateIdentify:output_type -> apply.CreateIdentifyReply
	3, // 4: apply.Apply.CheckIdentify:output_type -> apply.CheckIdentifyReply
	5, // 5: apply.Apply.CheckUser:output_type -> apply.CheckUserReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apply_proto_init() }
func file_apply_proto_init() {
	if File_apply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdentifyReq); i {
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
		file_apply_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdentifyReply); i {
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
		file_apply_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIdentifyReq); i {
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
		file_apply_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIdentifyReply); i {
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
		file_apply_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserReq); i {
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
		file_apply_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserReply); i {
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
			RawDescriptor: file_apply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apply_proto_goTypes,
		DependencyIndexes: file_apply_proto_depIdxs,
		MessageInfos:      file_apply_proto_msgTypes,
	}.Build()
	File_apply_proto = out.File
	file_apply_proto_rawDesc = nil
	file_apply_proto_goTypes = nil
	file_apply_proto_depIdxs = nil
}
