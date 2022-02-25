// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pair.proto

package pair

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ELeListRsp struct {
	Code                 int64          `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	UserElement          []*UserElement `protobuf:"bytes,3,rep,name=UserElement,proto3" json:"UserElement,omitempty"`
	Page                 int64          `protobuf:"varint,4,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize             int64          `protobuf:"varint,5,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Total                int64          `protobuf:"varint,6,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ELeListRsp) Reset()         { *m = ELeListRsp{} }
func (m *ELeListRsp) String() string { return proto.CompactTextString(m) }
func (*ELeListRsp) ProtoMessage()    {}
func (*ELeListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{0}
}

func (m *ELeListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ELeListRsp.Unmarshal(m, b)
}
func (m *ELeListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ELeListRsp.Marshal(b, m, deterministic)
}
func (m *ELeListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ELeListRsp.Merge(m, src)
}
func (m *ELeListRsp) XXX_Size() int {
	return xxx_messageInfo_ELeListRsp.Size(m)
}
func (m *ELeListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ELeListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ELeListRsp proto.InternalMessageInfo

func (m *ELeListRsp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ELeListRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ELeListRsp) GetUserElement() []*UserElement {
	if m != nil {
		return m.UserElement
	}
	return nil
}

func (m *ELeListRsp) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ELeListRsp) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ELeListRsp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type EleListReq struct {
	Page                 int64    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Keyword              string   `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EleListReq) Reset()         { *m = EleListReq{} }
func (m *EleListReq) String() string { return proto.CompactTextString(m) }
func (*EleListReq) ProtoMessage()    {}
func (*EleListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{1}
}

func (m *EleListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EleListReq.Unmarshal(m, b)
}
func (m *EleListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EleListReq.Marshal(b, m, deterministic)
}
func (m *EleListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EleListReq.Merge(m, src)
}
func (m *EleListReq) XXX_Size() int {
	return xxx_messageInfo_EleListReq.Size(m)
}
func (m *EleListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EleListReq.DiscardUnknown(m)
}

var xxx_messageInfo_EleListReq proto.InternalMessageInfo

func (m *EleListReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *EleListReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *EleListReq) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type ElePairReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize             int64    `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElePairReq) Reset()         { *m = ElePairReq{} }
func (m *ElePairReq) String() string { return proto.CompactTextString(m) }
func (*ElePairReq) ProtoMessage()    {}
func (*ElePairReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{2}
}

func (m *ElePairReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElePairReq.Unmarshal(m, b)
}
func (m *ElePairReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElePairReq.Marshal(b, m, deterministic)
}
func (m *ElePairReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElePairReq.Merge(m, src)
}
func (m *ElePairReq) XXX_Size() int {
	return xxx_messageInfo_ElePairReq.Size(m)
}
func (m *ElePairReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ElePairReq.DiscardUnknown(m)
}

var xxx_messageInfo_ElePairReq proto.InternalMessageInfo

func (m *ElePairReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ElePairReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ElePairReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type EleViewRsp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Element              *Element `protobuf:"bytes,3,opt,name=Element,proto3" json:"Element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EleViewRsp) Reset()         { *m = EleViewRsp{} }
func (m *EleViewRsp) String() string { return proto.CompactTextString(m) }
func (*EleViewRsp) ProtoMessage()    {}
func (*EleViewRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{3}
}

func (m *EleViewRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EleViewRsp.Unmarshal(m, b)
}
func (m *EleViewRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EleViewRsp.Marshal(b, m, deterministic)
}
func (m *EleViewRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EleViewRsp.Merge(m, src)
}
func (m *EleViewRsp) XXX_Size() int {
	return xxx_messageInfo_EleViewRsp.Size(m)
}
func (m *EleViewRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_EleViewRsp.DiscardUnknown(m)
}

var xxx_messageInfo_EleViewRsp proto.InternalMessageInfo

func (m *EleViewRsp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EleViewRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EleViewRsp) GetElement() *Element {
	if m != nil {
		return m.Element
	}
	return nil
}

type EleViewReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EleViewReq) Reset()         { *m = EleViewReq{} }
func (m *EleViewReq) String() string { return proto.CompactTextString(m) }
func (*EleViewReq) ProtoMessage()    {}
func (*EleViewReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{4}
}

func (m *EleViewReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EleViewReq.Unmarshal(m, b)
}
func (m *EleViewReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EleViewReq.Marshal(b, m, deterministic)
}
func (m *EleViewReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EleViewReq.Merge(m, src)
}
func (m *EleViewReq) XXX_Size() int {
	return xxx_messageInfo_EleViewReq.Size(m)
}
func (m *EleViewReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EleViewReq.DiscardUnknown(m)
}

var xxx_messageInfo_EleViewReq proto.InternalMessageInfo

func (m *EleViewReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type EleSaveReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Element              *Element `protobuf:"bytes,2,opt,name=Element,proto3" json:"Element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EleSaveReq) Reset()         { *m = EleSaveReq{} }
func (m *EleSaveReq) String() string { return proto.CompactTextString(m) }
func (*EleSaveReq) ProtoMessage()    {}
func (*EleSaveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{5}
}

func (m *EleSaveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EleSaveReq.Unmarshal(m, b)
}
func (m *EleSaveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EleSaveReq.Marshal(b, m, deterministic)
}
func (m *EleSaveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EleSaveReq.Merge(m, src)
}
func (m *EleSaveReq) XXX_Size() int {
	return xxx_messageInfo_EleSaveReq.Size(m)
}
func (m *EleSaveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EleSaveReq.DiscardUnknown(m)
}

var xxx_messageInfo_EleSaveReq proto.InternalMessageInfo

func (m *EleSaveReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *EleSaveReq) GetElement() *Element {
	if m != nil {
		return m.Element
	}
	return nil
}

type Response struct {
	Code                 int64    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserElement struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Account              string   `protobuf:"bytes,4,opt,name=Account,proto3" json:"Account,omitempty"`
	Element              *Element `protobuf:"bytes,5,opt,name=Element,proto3" json:"Element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserElement) Reset()         { *m = UserElement{} }
func (m *UserElement) String() string { return proto.CompactTextString(m) }
func (*UserElement) ProtoMessage()    {}
func (*UserElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{7}
}

func (m *UserElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserElement.Unmarshal(m, b)
}
func (m *UserElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserElement.Marshal(b, m, deterministic)
}
func (m *UserElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserElement.Merge(m, src)
}
func (m *UserElement) XXX_Size() int {
	return xxx_messageInfo_UserElement.Size(m)
}
func (m *UserElement) XXX_DiscardUnknown() {
	xxx_messageInfo_UserElement.DiscardUnknown(m)
}

var xxx_messageInfo_UserElement proto.InternalMessageInfo

func (m *UserElement) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserElement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserElement) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserElement) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserElement) GetElement() *Element {
	if m != nil {
		return m.Element
	}
	return nil
}

type Element struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Skill                string   `protobuf:"bytes,2,opt,name=Skill,proto3" json:"Skill,omitempty"`
	SkillNeed            string   `protobuf:"bytes,3,opt,name=SkillNeed,proto3" json:"SkillNeed,omitempty"`
	Star                 bool     `protobuf:"varint,4,opt,name=Star,proto3" json:"Star,omitempty"`
	Boost                int64    `protobuf:"varint,5,opt,name=Boost,proto3" json:"Boost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Element) Reset()         { *m = Element{} }
func (m *Element) String() string { return proto.CompactTextString(m) }
func (*Element) ProtoMessage()    {}
func (*Element) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{8}
}

func (m *Element) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Element.Unmarshal(m, b)
}
func (m *Element) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Element.Marshal(b, m, deterministic)
}
func (m *Element) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Element.Merge(m, src)
}
func (m *Element) XXX_Size() int {
	return xxx_messageInfo_Element.Size(m)
}
func (m *Element) XXX_DiscardUnknown() {
	xxx_messageInfo_Element.DiscardUnknown(m)
}

var xxx_messageInfo_Element proto.InternalMessageInfo

func (m *Element) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Element) GetSkill() string {
	if m != nil {
		return m.Skill
	}
	return ""
}

func (m *Element) GetSkillNeed() string {
	if m != nil {
		return m.SkillNeed
	}
	return ""
}

func (m *Element) GetStar() bool {
	if m != nil {
		return m.Star
	}
	return false
}

func (m *Element) GetBoost() int64 {
	if m != nil {
		return m.Boost
	}
	return 0
}

func init() {
	proto.RegisterType((*ELeListRsp)(nil), "pair.ELeListRsp")
	proto.RegisterType((*EleListReq)(nil), "pair.EleListReq")
	proto.RegisterType((*ElePairReq)(nil), "pair.ElePairReq")
	proto.RegisterType((*EleViewRsp)(nil), "pair.EleViewRsp")
	proto.RegisterType((*EleViewReq)(nil), "pair.EleViewReq")
	proto.RegisterType((*EleSaveReq)(nil), "pair.EleSaveReq")
	proto.RegisterType((*Response)(nil), "pair.Response")
	proto.RegisterType((*UserElement)(nil), "pair.UserElement")
	proto.RegisterType((*Element)(nil), "pair.Element")
}

func init() { proto.RegisterFile("pair.proto", fileDescriptor_b6c646fab57af36d) }

var fileDescriptor_b6c646fab57af36d = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6a, 0x1b, 0x31,
	0x10, 0x46, 0xbb, 0xfe, 0x1d, 0xd3, 0x90, 0x8a, 0x50, 0x84, 0x29, 0xc5, 0xec, 0xa5, 0x3e, 0x39,
	0xd4, 0xb9, 0xf4, 0x9a, 0x14, 0x53, 0x02, 0xa9, 0x09, 0x72, 0x93, 0x43, 0x6f, 0xaa, 0x3d, 0x14,
	0x91, 0x8d, 0xb5, 0x5d, 0x6d, 0x13, 0xd2, 0x67, 0xe8, 0xeb, 0xf4, 0x3d, 0xfa, 0x48, 0x65, 0x24,
	0xad, 0x76, 0x93, 0x38, 0x81, 0x9c, 0x76, 0x66, 0xb4, 0xf3, 0x7d, 0xdf, 0xfc, 0x48, 0x00, 0x85,
	0xd2, 0xe5, 0xac, 0x28, 0x4d, 0x65, 0x78, 0x87, 0xec, 0xec, 0x2f, 0x03, 0x58, 0x9c, 0xe1, 0x99,
	0xb6, 0x95, 0xb4, 0x05, 0xe7, 0xd0, 0xf9, 0x64, 0x36, 0x28, 0xd8, 0x84, 0x4d, 0x53, 0xe9, 0x6c,
	0x2e, 0xa0, 0xff, 0x05, 0xad, 0x55, 0x3f, 0x50, 0x24, 0x13, 0x36, 0x1d, 0xca, 0xda, 0xe5, 0x47,
	0x30, 0xba, 0xb0, 0x58, 0x2e, 0x72, 0xbc, 0xc6, 0x6d, 0x25, 0xd2, 0x49, 0x3a, 0x1d, 0xcd, 0x5f,
	0xcf, 0x1c, 0x49, 0xeb, 0x40, 0xb6, 0xff, 0x22, 0x8a, 0x73, 0xc2, 0xea, 0x78, 0x0a, 0xb2, 0xf9,
	0x18, 0x06, 0xf4, 0x5d, 0xe9, 0xdf, 0x28, 0xba, 0x2e, 0x1e, 0x7d, 0x7e, 0x00, 0xdd, 0xaf, 0xa6,
	0x52, 0xb9, 0xe8, 0xb9, 0x03, 0xef, 0x64, 0x97, 0x00, 0x8b, 0xdc, 0xcb, 0xc6, 0x9f, 0x11, 0x93,
	0x3d, 0x81, 0x99, 0x3c, 0xc0, 0x14, 0xd0, 0xbf, 0xc2, 0xbb, 0x5b, 0x53, 0x6e, 0x44, 0xea, 0x4b,
	0x0a, 0x6e, 0xb6, 0x74, 0xb8, 0xe7, 0x4a, 0x97, 0x84, 0xbb, 0x0f, 0xe9, 0x85, 0xde, 0x38, 0xd8,
	0x8e, 0x24, 0x33, 0x32, 0x25, 0x4f, 0x30, 0xa5, 0xf7, 0x99, 0xb2, 0xb5, 0xc3, 0xbb, 0xd4, 0x78,
	0xfb, 0xf2, 0xf6, 0xbe, 0x87, 0x7e, 0xd3, 0x5a, 0x36, 0x1d, 0xcd, 0x5f, 0xf9, 0xd6, 0xd6, 0x6d,
	0xad, 0x4f, 0xb3, 0x77, 0x0d, 0xc9, 0x2e, 0xd1, 0xd9, 0x67, 0x77, 0xbe, 0x52, 0x37, 0xb8, 0xbb,
	0xa8, 0x16, 0x51, 0xf2, 0x2c, 0xd1, 0x47, 0x18, 0x48, 0xb4, 0x85, 0xd9, 0x5a, 0x7c, 0x59, 0x2d,
	0xd9, 0x1f, 0x76, 0x6f, 0x57, 0xf8, 0x1e, 0x24, 0xa7, 0xb5, 0x86, 0xe4, 0xd4, 0xf5, 0x75, 0xa9,
	0xae, 0xeb, 0x34, 0x67, 0xf3, 0x37, 0xd0, 0x3b, 0xbe, 0x51, 0x95, 0x2a, 0xc3, 0x90, 0x82, 0x47,
	0x2c, 0xc7, 0xeb, 0xb5, 0xf9, 0xb5, 0xad, 0xdc, 0x12, 0x0d, 0x65, 0xed, 0xb6, 0x0b, 0xe9, 0x3e,
	0x5b, 0xc8, 0x5d, 0xfc, 0x71, 0x47, 0x3b, 0x0e, 0xa0, 0xbb, 0xba, 0xd2, 0x79, 0x1e, 0xc4, 0x78,
	0x87, 0xbf, 0x85, 0xa1, 0x33, 0x96, 0x88, 0xf5, 0xd6, 0x34, 0x01, 0xd2, 0xbf, 0x22, 0xa5, 0x24,
	0x68, 0x20, 0x9d, 0x4d, 0x38, 0x27, 0xc6, 0xd8, 0x2a, 0xac, 0xb4, 0x77, 0xe6, 0xff, 0x18, 0xad,
	0x90, 0x2e, 0xf9, 0x21, 0x8c, 0x82, 0x06, 0x9a, 0x0c, 0xdf, 0x8f, 0x52, 0xc3, 0xa0, 0xc6, 0x7b,
	0x3e, 0x12, 0x3b, 0xfe, 0x21, 0x26, 0xd0, 0xa8, 0x5b, 0x09, 0x61, 0xf2, 0xe3, 0x07, 0x11, 0x5b,
	0xb4, 0x52, 0xe8, 0xaa, 0xb4, 0x52, 0xc2, 0xcd, 0x89, 0x29, 0xcd, 0x13, 0xd0, 0xa4, 0x38, 0x95,
	0x4d, 0x4a, 0xb8, 0x14, 0x8f, 0x53, 0x4e, 0x06, 0xdf, 0x7a, 0xb3, 0x43, 0x0a, 0x7e, 0xef, 0xb9,
	0xb7, 0xe5, 0xe8, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x84, 0x83, 0x4b, 0x69, 0x04, 0x00,
	0x00,
}