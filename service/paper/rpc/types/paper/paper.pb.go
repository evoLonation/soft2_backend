// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: paper.proto

package paper

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

type CheckScholarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScholarId int64 `protobuf:"varint,1,opt,name=scholarId,proto3" json:"scholarId,omitempty"`
}

func (x *CheckScholarReq) Reset() {
	*x = CheckScholarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckScholarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckScholarReq) ProtoMessage() {}

func (x *CheckScholarReq) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckScholarReq.ProtoReflect.Descriptor instead.
func (*CheckScholarReq) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{0}
}

func (x *CheckScholarReq) GetScholarId() int64 {
	if x != nil {
		return x.ScholarId
	}
	return 0
}

type CreateScholarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScholarName string `protobuf:"bytes,1,opt,name=scholarName,proto3" json:"scholarName,omitempty"`
	Institution string `protobuf:"bytes,2,opt,name=institution,proto3" json:"institution,omitempty"`
}

func (x *CreateScholarReply) Reset() {
	*x = CreateScholarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScholarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScholarReply) ProtoMessage() {}

func (x *CreateScholarReply) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScholarReply.ProtoReflect.Descriptor instead.
func (*CreateScholarReply) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScholarReply) GetScholarName() string {
	if x != nil {
		return x.ScholarName
	}
	return ""
}

func (x *CreateScholarReply) GetInstitution() string {
	if x != nil {
		return x.Institution
	}
	return ""
}

var File_paper_proto protoreflect.FileDescriptor

var file_paper_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x68,
	0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c,
	0x61, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x63, 0x68, 0x6f,
	0x6c, 0x61, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x52, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x41, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72,
	0x12, 0x16, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63,
	0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paper_proto_rawDescOnce sync.Once
	file_paper_proto_rawDescData = file_paper_proto_rawDesc
)

func file_paper_proto_rawDescGZIP() []byte {
	file_paper_proto_rawDescOnce.Do(func() {
		file_paper_proto_rawDescData = protoimpl.X.CompressGZIP(file_paper_proto_rawDescData)
	})
	return file_paper_proto_rawDescData
}

var file_paper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_paper_proto_goTypes = []interface{}{
	(*CheckScholarReq)(nil),    // 0: paper.CheckScholarReq
	(*CreateScholarReply)(nil), // 1: paper.CreateScholarReply
}
var file_paper_proto_depIdxs = []int32{
	0, // 0: paper.StreamGreeter.CheckScholar:input_type -> paper.CheckScholarReq
	1, // 1: paper.StreamGreeter.CheckScholar:output_type -> paper.CreateScholarReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paper_proto_init() }
func file_paper_proto_init() {
	if File_paper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckScholarReq); i {
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
		file_paper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScholarReply); i {
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
			RawDescriptor: file_paper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paper_proto_goTypes,
		DependencyIndexes: file_paper_proto_depIdxs,
		MessageInfos:      file_paper_proto_msgTypes,
	}.Build()
	File_paper_proto = out.File
	file_paper_proto_rawDesc = nil
	file_paper_proto_goTypes = nil
	file_paper_proto_depIdxs = nil
}