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

	ScholarId string `protobuf:"bytes,1,opt,name=scholarId,proto3" json:"scholarId,omitempty"`
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

func (x *CheckScholarReq) GetScholarId() string {
	if x != nil {
		return x.ScholarId
	}
	return ""
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

type ListCheckScholarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScholarId []string `protobuf:"bytes,1,rep,name=scholarId,proto3" json:"scholarId,omitempty"`
}

func (x *ListCheckScholarReq) Reset() {
	*x = ListCheckScholarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCheckScholarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCheckScholarReq) ProtoMessage() {}

func (x *ListCheckScholarReq) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCheckScholarReq.ProtoReflect.Descriptor instead.
func (*ListCheckScholarReq) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{2}
}

func (x *ListCheckScholarReq) GetScholarId() []string {
	if x != nil {
		return x.ScholarId
	}
	return nil
}

type ListCreateScholarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scholars []*CreateScholarReply `protobuf:"bytes,1,rep,name=scholars,proto3" json:"scholars,omitempty"`
}

func (x *ListCreateScholarReply) Reset() {
	*x = ListCreateScholarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCreateScholarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCreateScholarReply) ProtoMessage() {}

func (x *ListCreateScholarReply) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCreateScholarReply.ProtoReflect.Descriptor instead.
func (*ListCreateScholarReply) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{3}
}

func (x *ListCreateScholarReply) GetScholars() []*CreateScholarReply {
	if x != nil {
		return x.Scholars
	}
	return nil
}

type MovePaperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperId  string `protobuf:"bytes,1,opt,name=paperId,proto3" json:"paperId,omitempty"`
	OwnerId  string `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	TargetId string `protobuf:"bytes,3,opt,name=targetId,proto3" json:"targetId,omitempty"`
}

func (x *MovePaperReq) Reset() {
	*x = MovePaperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePaperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePaperReq) ProtoMessage() {}

func (x *MovePaperReq) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePaperReq.ProtoReflect.Descriptor instead.
func (*MovePaperReq) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{4}
}

func (x *MovePaperReq) GetPaperId() string {
	if x != nil {
		return x.PaperId
	}
	return ""
}

func (x *MovePaperReq) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *MovePaperReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

type MovePaperReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MovePaperReply) Reset() {
	*x = MovePaperReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePaperReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePaperReply) ProtoMessage() {}

func (x *MovePaperReply) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePaperReply.ProtoReflect.Descriptor instead.
func (*MovePaperReply) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{5}
}

type GetPaperNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperId string `protobuf:"bytes,1,opt,name=paperId,proto3" json:"paperId,omitempty"`
}

func (x *GetPaperNameReq) Reset() {
	*x = GetPaperNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaperNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaperNameReq) ProtoMessage() {}

func (x *GetPaperNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaperNameReq.ProtoReflect.Descriptor instead.
func (*GetPaperNameReq) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{6}
}

func (x *GetPaperNameReq) GetPaperId() string {
	if x != nil {
		return x.PaperId
	}
	return ""
}

type GetPaperNameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperName string `protobuf:"bytes,1,opt,name=paperName,proto3" json:"paperName,omitempty"`
}

func (x *GetPaperNameReply) Reset() {
	*x = GetPaperNameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaperNameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaperNameReply) ProtoMessage() {}

func (x *GetPaperNameReply) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaperNameReply.ProtoReflect.Descriptor instead.
func (*GetPaperNameReply) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{7}
}

func (x *GetPaperNameReply) GetPaperName() string {
	if x != nil {
		return x.PaperName
	}
	return ""
}

var File_paper_proto protoreflect.FileDescriptor

var file_paper_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x68,
	0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c,
	0x61, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x6f,
	0x6c, 0x61, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x33, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x68, 0x6f,
	0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6c,
	0x61, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35,
	0x0a, 0x08, 0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x08, 0x73, 0x63, 0x68,
	0x6f, 0x6c, 0x61, 0x72, 0x73, 0x22, 0x5e, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x9c, 0x02, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72,
	0x12, 0x1a, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x4d,
	0x6f, 0x76, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_paper_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_paper_proto_goTypes = []interface{}{
	(*CheckScholarReq)(nil),        // 0: paper.CheckScholarReq
	(*CreateScholarReply)(nil),     // 1: paper.CreateScholarReply
	(*ListCheckScholarReq)(nil),    // 2: paper.ListCheckScholarReq
	(*ListCreateScholarReply)(nil), // 3: paper.ListCreateScholarReply
	(*MovePaperReq)(nil),           // 4: paper.MovePaperReq
	(*MovePaperReply)(nil),         // 5: paper.MovePaperReply
	(*GetPaperNameReq)(nil),        // 6: paper.GetPaperNameReq
	(*GetPaperNameReply)(nil),      // 7: paper.GetPaperNameReply
}
var file_paper_proto_depIdxs = []int32{
	1, // 0: paper.ListCreateScholarReply.scholars:type_name -> paper.CreateScholarReply
	0, // 1: paper.StreamGreeter.CheckScholar:input_type -> paper.CheckScholarReq
	2, // 2: paper.StreamGreeter.ListCheckScholar:input_type -> paper.ListCheckScholarReq
	4, // 3: paper.StreamGreeter.MovePaper:input_type -> paper.MovePaperReq
	6, // 4: paper.StreamGreeter.GetPaperName:input_type -> paper.GetPaperNameReq
	1, // 5: paper.StreamGreeter.CheckScholar:output_type -> paper.CreateScholarReply
	3, // 6: paper.StreamGreeter.ListCheckScholar:output_type -> paper.ListCreateScholarReply
	5, // 7: paper.StreamGreeter.MovePaper:output_type -> paper.MovePaperReply
	7, // 8: paper.StreamGreeter.GetPaperName:output_type -> paper.GetPaperNameReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_paper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCheckScholarReq); i {
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
		file_paper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCreateScholarReply); i {
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
		file_paper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePaperReq); i {
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
		file_paper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePaperReply); i {
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
		file_paper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaperNameReq); i {
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
		file_paper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaperNameReply); i {
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
			NumMessages:   8,
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
