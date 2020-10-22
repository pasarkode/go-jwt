// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/file_set.proto

package resultstore

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// This resource represents a set of Files and other (nested) FileSets.
// A FileSet is a node in the graph, and the file_sets field represents the
// outgoing edges. A resource may reference various nodes in the graph to
// represent the transitive closure of all files from those nodes.
// The FileSets must form a directed acyclic graph. The Upload API is unable to
// enforce that the graph is acyclic at write time, and if cycles are written,
// it may cause issues at read time.
//
// A FileSet may be referenced by other resources in conjunction with Files. A
// File is preferred for something that can only be ever referenced by one
// resource, and a FileSet is preferred if it can be reference by multiple
// resources.
type FileSet struct {
	// The format of this FileSet resource name must be:
	// invocations/${INVOCATION_ID}/fileSets/${url_encode(FILE_SET_ID)}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the file set. They must match the
	// resource name after proper encoding.
	Id *FileSet_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// List of names of other file sets that are referenced from this one.
	// Each name must point to a file set under the same invocation. The name
	// format must be: invocations/${INVOCATION_ID}/fileSets/${FILE_SET_ID}
	FileSets []string `protobuf:"bytes,3,rep,name=file_sets,json=fileSets,proto3" json:"file_sets,omitempty"`
	// Files that are contained within this file set.
	// The uid field in the file should be unique for the Invocation.
	Files                []*File  `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileSet) Reset()         { *m = FileSet{} }
func (m *FileSet) String() string { return proto.CompactTextString(m) }
func (*FileSet) ProtoMessage()    {}
func (*FileSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8d86e03917a704, []int{0}
}

func (m *FileSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSet.Unmarshal(m, b)
}
func (m *FileSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSet.Marshal(b, m, deterministic)
}
func (m *FileSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSet.Merge(m, src)
}
func (m *FileSet) XXX_Size() int {
	return xxx_messageInfo_FileSet.Size(m)
}
func (m *FileSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSet.DiscardUnknown(m)
}

var xxx_messageInfo_FileSet proto.InternalMessageInfo

func (m *FileSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileSet) GetId() *FileSet_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FileSet) GetFileSets() []string {
	if m != nil {
		return m.FileSets
	}
	return nil
}

func (m *FileSet) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// The resource ID components that identify the FileSet.
type FileSet_Id struct {
	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The FileSet ID.
	FileSetId            string   `protobuf:"bytes,2,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileSet_Id) Reset()         { *m = FileSet_Id{} }
func (m *FileSet_Id) String() string { return proto.CompactTextString(m) }
func (*FileSet_Id) ProtoMessage()    {}
func (*FileSet_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8d86e03917a704, []int{0, 0}
}

func (m *FileSet_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSet_Id.Unmarshal(m, b)
}
func (m *FileSet_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSet_Id.Marshal(b, m, deterministic)
}
func (m *FileSet_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSet_Id.Merge(m, src)
}
func (m *FileSet_Id) XXX_Size() int {
	return xxx_messageInfo_FileSet_Id.Size(m)
}
func (m *FileSet_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSet_Id.DiscardUnknown(m)
}

var xxx_messageInfo_FileSet_Id proto.InternalMessageInfo

func (m *FileSet_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *FileSet_Id) GetFileSetId() string {
	if m != nil {
		return m.FileSetId
	}
	return ""
}

func init() {
	proto.RegisterType((*FileSet)(nil), "google.devtools.resultstore.v2.FileSet")
	proto.RegisterType((*FileSet_Id)(nil), "google.devtools.resultstore.v2.FileSet.Id")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/file_set.proto", fileDescriptor_ca8d86e03917a704)
}

var fileDescriptor_ca8d86e03917a704 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbb, 0x4f, 0xc3, 0x30,
	0x10, 0x87, 0x95, 0xa4, 0x3c, 0x72, 0x81, 0xc5, 0x53, 0x54, 0xa4, 0x2a, 0x2a, 0x0c, 0x01, 0x09,
	0x5b, 0x0a, 0x5b, 0xd9, 0x18, 0x90, 0xbc, 0xa1, 0xb0, 0xb1, 0x44, 0xa1, 0xbe, 0x5a, 0x96, 0xdc,
	0x5c, 0x89, 0x4d, 0xfe, 0x7b, 0x24, 0x44, 0x1e, 0xa2, 0x0b, 0xb4, 0xdb, 0xd9, 0xfe, 0xee, 0x77,
	0x9f, 0x0f, 0xee, 0x35, 0x91, 0xb6, 0x28, 0x14, 0x76, 0x9e, 0xc8, 0x3a, 0xd1, 0xa2, 0xfb, 0xb4,
	0xde, 0x79, 0x6a, 0x51, 0x74, 0x85, 0xd8, 0x18, 0x8b, 0x95, 0x43, 0xcf, 0x77, 0x2d, 0x79, 0x62,
	0x8b, 0x01, 0xe7, 0x13, 0xce, 0xf7, 0x70, 0xde, 0x15, 0xf3, 0xdb, 0x23, 0xe2, 0x86, 0xa8, 0xe5,
	0x57, 0x00, 0x67, 0xcf, 0xc6, 0xe2, 0x2b, 0x7a, 0xc6, 0x60, 0xd6, 0xd4, 0x5b, 0x4c, 0x83, 0x2c,
	0xc8, 0xe3, 0xb2, 0xaf, 0xd9, 0x0a, 0x42, 0xa3, 0xd2, 0x30, 0x0b, 0xf2, 0xa4, 0xb8, 0xe3, 0xff,
	0xcf, 0xe5, 0x63, 0x10, 0x97, 0xaa, 0x0c, 0x8d, 0x62, 0x57, 0x10, 0x4f, 0xe2, 0x2e, 0x8d, 0xb2,
	0x28, 0x8f, 0xcb, 0xf3, 0xcd, 0x80, 0x38, 0xb6, 0x82, 0x93, 0x9f, 0xda, 0xa5, 0xb3, 0x2c, 0xca,
	0x93, 0xe2, 0xe6, 0x98, 0xec, 0x72, 0x68, 0x99, 0x4b, 0x08, 0xa5, 0x62, 0xd7, 0x70, 0x69, 0x9a,
	0x8e, 0xd6, 0xb5, 0x37, 0xd4, 0x54, 0x46, 0x8d, 0xde, 0x17, 0xbf, 0x97, 0x52, 0xb1, 0x05, 0x24,
	0x93, 0x43, 0x35, 0x7e, 0x24, 0x2e, 0xe3, 0xd1, 0x42, 0xaa, 0xa7, 0x0f, 0x58, 0xae, 0x69, 0x7b,
	0x60, 0xf8, 0x4b, 0xf0, 0x26, 0x47, 0x42, 0x93, 0xad, 0x1b, 0xcd, 0xa9, 0xd5, 0x42, 0x63, 0xd3,
	0xef, 0x50, 0x0c, 0x4f, 0xf5, 0xce, 0xb8, 0xbf, 0x36, 0xfe, 0xb8, 0x77, 0x7c, 0x3f, 0xed, 0xbb,
	0x1e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x67, 0xbb, 0x64, 0xf5, 0x01, 0x00, 0x00,
}
