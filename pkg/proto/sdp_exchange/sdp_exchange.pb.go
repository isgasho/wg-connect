// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdp_exchange.proto

package sdp_exchange

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OfferMessage struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Body                 *Offer   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfferMessage) Reset()         { *m = OfferMessage{} }
func (m *OfferMessage) String() string { return proto.CompactTextString(m) }
func (*OfferMessage) ProtoMessage()    {}
func (*OfferMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{0}
}

func (m *OfferMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfferMessage.Unmarshal(m, b)
}
func (m *OfferMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfferMessage.Marshal(b, m, deterministic)
}
func (m *OfferMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfferMessage.Merge(m, src)
}
func (m *OfferMessage) XXX_Size() int {
	return xxx_messageInfo_OfferMessage.Size(m)
}
func (m *OfferMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OfferMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OfferMessage proto.InternalMessageInfo

func (m *OfferMessage) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *OfferMessage) GetBody() *Offer {
	if m != nil {
		return m.Body
	}
	return nil
}

type AnswerMessage struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Body                 *Answer  `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerMessage) Reset()         { *m = AnswerMessage{} }
func (m *AnswerMessage) String() string { return proto.CompactTextString(m) }
func (*AnswerMessage) ProtoMessage()    {}
func (*AnswerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{1}
}

func (m *AnswerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerMessage.Unmarshal(m, b)
}
func (m *AnswerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerMessage.Marshal(b, m, deterministic)
}
func (m *AnswerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerMessage.Merge(m, src)
}
func (m *AnswerMessage) XXX_Size() int {
	return xxx_messageInfo_AnswerMessage.Size(m)
}
func (m *AnswerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerMessage proto.InternalMessageInfo

func (m *AnswerMessage) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *AnswerMessage) GetBody() *Answer {
	if m != nil {
		return m.Body
	}
	return nil
}

type OfferResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfferResponse) Reset()         { *m = OfferResponse{} }
func (m *OfferResponse) String() string { return proto.CompactTextString(m) }
func (*OfferResponse) ProtoMessage()    {}
func (*OfferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{2}
}

func (m *OfferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfferResponse.Unmarshal(m, b)
}
func (m *OfferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfferResponse.Marshal(b, m, deterministic)
}
func (m *OfferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfferResponse.Merge(m, src)
}
func (m *OfferResponse) XXX_Size() int {
	return xxx_messageInfo_OfferResponse.Size(m)
}
func (m *OfferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OfferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OfferResponse proto.InternalMessageInfo

func (m *OfferResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AnswerResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerResponse) Reset()         { *m = AnswerResponse{} }
func (m *AnswerResponse) String() string { return proto.CompactTextString(m) }
func (*AnswerResponse) ProtoMessage()    {}
func (*AnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{3}
}

func (m *AnswerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerResponse.Unmarshal(m, b)
}
func (m *AnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerResponse.Marshal(b, m, deterministic)
}
func (m *AnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerResponse.Merge(m, src)
}
func (m *AnswerResponse) XXX_Size() int {
	return xxx_messageInfo_AnswerResponse.Size(m)
}
func (m *AnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerResponse proto.InternalMessageInfo

func (m *AnswerResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PeerMessage struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerMessage) Reset()         { *m = PeerMessage{} }
func (m *PeerMessage) String() string { return proto.CompactTextString(m) }
func (*PeerMessage) ProtoMessage()    {}
func (*PeerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{4}
}

func (m *PeerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerMessage.Unmarshal(m, b)
}
func (m *PeerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerMessage.Marshal(b, m, deterministic)
}
func (m *PeerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerMessage.Merge(m, src)
}
func (m *PeerMessage) XXX_Size() int {
	return xxx_messageInfo_PeerMessage.Size(m)
}
func (m *PeerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PeerMessage proto.InternalMessageInfo

func (m *PeerMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PeerResponse struct {
	BodyJson             string   `protobuf:"bytes,1,opt,name=body_json,json=bodyJson,proto3" json:"body_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerResponse) Reset()         { *m = PeerResponse{} }
func (m *PeerResponse) String() string { return proto.CompactTextString(m) }
func (*PeerResponse) ProtoMessage()    {}
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{5}
}

func (m *PeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerResponse.Unmarshal(m, b)
}
func (m *PeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerResponse.Marshal(b, m, deterministic)
}
func (m *PeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerResponse.Merge(m, src)
}
func (m *PeerResponse) XXX_Size() int {
	return xxx_messageInfo_PeerResponse.Size(m)
}
func (m *PeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerResponse proto.InternalMessageInfo

func (m *PeerResponse) GetBodyJson() string {
	if m != nil {
		return m.BodyJson
	}
	return ""
}

type Offer struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Offer) Reset()         { *m = Offer{} }
func (m *Offer) String() string { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()    {}
func (*Offer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{6}
}

func (m *Offer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Offer.Unmarshal(m, b)
}
func (m *Offer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Offer.Marshal(b, m, deterministic)
}
func (m *Offer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offer.Merge(m, src)
}
func (m *Offer) XXX_Size() int {
	return xxx_messageInfo_Offer.Size(m)
}
func (m *Offer) XXX_DiscardUnknown() {
	xxx_messageInfo_Offer.DiscardUnknown(m)
}

var xxx_messageInfo_Offer proto.InternalMessageInfo

func (m *Offer) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Answer struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{7}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Token struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{8}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*OfferMessage)(nil), "OfferMessage")
	proto.RegisterType((*AnswerMessage)(nil), "AnswerMessage")
	proto.RegisterType((*OfferResponse)(nil), "OfferResponse")
	proto.RegisterType((*AnswerResponse)(nil), "AnswerResponse")
	proto.RegisterType((*PeerMessage)(nil), "PeerMessage")
	proto.RegisterType((*PeerResponse)(nil), "PeerResponse")
	proto.RegisterType((*Offer)(nil), "Offer")
	proto.RegisterType((*Answer)(nil), "Answer")
	proto.RegisterType((*Token)(nil), "Token")
}

func init() { proto.RegisterFile("sdp_exchange.proto", fileDescriptor_c88d59dbcf37a6c3) }

var fileDescriptor_c88d59dbcf37a6c3 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0x53, 0xe9, 0x4f, 0x32, 0x4d, 0x22, 0x0c, 0x5e, 0x94, 0x24, 0x88, 0x6c, 0xbd, 0x28,
	0x28, 0x2b, 0xd4, 0x27, 0xf0, 0x42, 0x22, 0x05, 0x51, 0xa2, 0xf7, 0xa5, 0x6d, 0xa6, 0x15, 0x85,
	0xdd, 0xd0, 0x0d, 0xa8, 0x8f, 0xe9, 0x1b, 0xc9, 0xfe, 0x44, 0x13, 0x28, 0xe2, 0x5d, 0xb2, 0x67,
	0xe6, 0xcc, 0xb7, 0x73, 0x16, 0x50, 0x95, 0xd5, 0x92, 0x3e, 0x36, 0x2f, 0x2b, 0xb1, 0x23, 0x5e,
	0xed, 0x65, 0x2d, 0xd9, 0x1d, 0x84, 0x0f, 0xdb, 0x2d, 0xed, 0xef, 0x49, 0xa9, 0xd5, 0x8e, 0x30,
	0x83, 0x41, 0x2d, 0xdf, 0x48, 0x4c, 0x7a, 0x67, 0xbd, 0xd9, 0x78, 0x3e, 0xe4, 0xcf, 0xfa, 0xaf,
	0xb0, 0x87, 0x98, 0x40, 0x7f, 0x2d, 0xcb, 0xcf, 0xc9, 0x91, 0x13, 0x4d, 0x6b, 0x61, 0xce, 0xd8,
	0x02, 0xa2, 0x1b, 0xa1, 0xde, 0xff, 0x6b, 0x95, 0x76, 0xac, 0x46, 0xdc, 0xf6, 0x3a, 0xaf, 0x29,
	0x44, 0xd6, 0x9a, 0x54, 0x25, 0x85, 0x22, 0x44, 0xe8, 0x6f, 0x64, 0x49, 0xc6, 0x2a, 0x28, 0xcc,
	0x37, 0x3b, 0x87, 0xd8, 0x35, 0xfd, 0x55, 0x35, 0x85, 0xf1, 0x23, 0xfd, 0x42, 0x9d, 0xb4, 0xa1,
	0x02, 0x07, 0xc3, 0x2e, 0x20, 0xd4, 0x45, 0x3f, 0x46, 0x29, 0x04, 0x9a, 0x63, 0xf9, 0xaa, 0x64,
	0x53, 0xe9, 0xeb, 0x83, 0x85, 0x92, 0x82, 0xa5, 0x30, 0x30, 0x70, 0x7a, 0x9c, 0xb9, 0x82, 0x1b,
	0x67, 0xc8, 0x33, 0x18, 0x5a, 0xa8, 0x83, 0x6a, 0x0a, 0x03, 0xb3, 0x84, 0x43, 0xe2, 0xfc, 0xab,
	0x07, 0xfe, 0xad, 0x4b, 0x07, 0x67, 0x30, 0xca, 0xa9, 0xd6, 0x50, 0x18, 0xf2, 0xd6, 0x05, 0x92,
	0x88, 0xb7, 0x49, 0x99, 0x87, 0xa7, 0x10, 0xe4, 0x54, 0xbb, 0xa1, 0x6e, 0xc9, 0x49, 0xb3, 0x4f,
	0xe6, 0x61, 0x06, 0x7e, 0x4e, 0xb5, 0x25, 0x6e, 0x64, 0x97, 0x1c, 0xf3, 0xf0, 0x12, 0x82, 0x27,
	0x12, 0xa5, 0x95, 0x23, 0xde, 0x7e, 0x0b, 0x49, 0xcc, 0x3b, 0x21, 0x30, 0x0f, 0xaf, 0x00, 0x74,
	0xb5, 0x1b, 0x16, 0xf3, 0x4e, 0xe0, 0xc9, 0x31, 0xef, 0xe6, 0xc1, 0xbc, 0xf5, 0xd0, 0xbc, 0xb2,
	0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x9a, 0x1d, 0x76, 0x7b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	GetPeer(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerResponse, error)
	GetAnswer(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Answer, error)
	GetOffer(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Offer, error)
	SendOffer(ctx context.Context, in *OfferMessage, opts ...grpc.CallOption) (*OfferResponse, error)
	SendAnswer(ctx context.Context, in *AnswerMessage, opts ...grpc.CallOption) (*AnswerResponse, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) GetPeer(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerResponse, error) {
	out := new(PeerResponse)
	err := c.cc.Invoke(ctx, "/Exchange/GetPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) GetAnswer(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/Exchange/GetAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) GetOffer(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Offer, error) {
	out := new(Offer)
	err := c.cc.Invoke(ctx, "/Exchange/GetOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) SendOffer(ctx context.Context, in *OfferMessage, opts ...grpc.CallOption) (*OfferResponse, error) {
	out := new(OfferResponse)
	err := c.cc.Invoke(ctx, "/Exchange/SendOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) SendAnswer(ctx context.Context, in *AnswerMessage, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, "/Exchange/SendAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	GetPeer(context.Context, *PeerMessage) (*PeerResponse, error)
	GetAnswer(context.Context, *Token) (*Answer, error)
	GetOffer(context.Context, *Token) (*Offer, error)
	SendOffer(context.Context, *OfferMessage) (*OfferResponse, error)
	SendAnswer(context.Context, *AnswerMessage) (*AnswerResponse, error)
}

// UnimplementedExchangeServer can be embedded to have forward compatible implementations.
type UnimplementedExchangeServer struct {
}

func (*UnimplementedExchangeServer) GetPeer(ctx context.Context, req *PeerMessage) (*PeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeer not implemented")
}
func (*UnimplementedExchangeServer) GetAnswer(ctx context.Context, req *Token) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}
func (*UnimplementedExchangeServer) GetOffer(ctx context.Context, req *Token) (*Offer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffer not implemented")
}
func (*UnimplementedExchangeServer) SendOffer(ctx context.Context, req *OfferMessage) (*OfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOffer not implemented")
}
func (*UnimplementedExchangeServer) SendAnswer(ctx context.Context, req *AnswerMessage) (*AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAnswer not implemented")
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_GetPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).GetPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/GetPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).GetPeer(ctx, req.(*PeerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_GetAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).GetAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/GetAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).GetAnswer(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_GetOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).GetOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/GetOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).GetOffer(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_SendOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfferMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).SendOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/SendOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).SendOffer(ctx, req.(*OfferMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_SendAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).SendAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/SendAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).SendAnswer(ctx, req.(*AnswerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeer",
			Handler:    _Exchange_GetPeer_Handler,
		},
		{
			MethodName: "GetAnswer",
			Handler:    _Exchange_GetAnswer_Handler,
		},
		{
			MethodName: "GetOffer",
			Handler:    _Exchange_GetOffer_Handler,
		},
		{
			MethodName: "SendOffer",
			Handler:    _Exchange_SendOffer_Handler,
		},
		{
			MethodName: "SendAnswer",
			Handler:    _Exchange_SendAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdp_exchange.proto",
}
