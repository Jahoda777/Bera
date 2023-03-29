// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polaris/erc20/v1alpha1/tx.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ConvertERC20ToCosmosRequest defines a message to convert ERC20 tokens to Cosmos tokens.
type ConvertERC20ToCosmosRequest struct {
	// token is the address of the ERC20 token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// receipient is the address of the recipient.
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// amount is the amount of tokens to convert.
	Amount *cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount,omitempty"`
}

func (m *ConvertERC20ToCosmosRequest) Reset()         { *m = ConvertERC20ToCosmosRequest{} }
func (m *ConvertERC20ToCosmosRequest) String() string { return proto.CompactTextString(m) }
func (*ConvertERC20ToCosmosRequest) ProtoMessage()    {}
func (*ConvertERC20ToCosmosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba53552aa354681a, []int{0}
}
func (m *ConvertERC20ToCosmosRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConvertERC20ToCosmosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConvertERC20ToCosmosRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConvertERC20ToCosmosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertERC20ToCosmosRequest.Merge(m, src)
}
func (m *ConvertERC20ToCosmosRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConvertERC20ToCosmosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertERC20ToCosmosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertERC20ToCosmosRequest proto.InternalMessageInfo

func (m *ConvertERC20ToCosmosRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ConvertERC20ToCosmosRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

// ConvertERC20ToCosmosResponse defines the response for the ConvertERC20ToCosmosRequest.
type ConvertERC20ToCosmosResponse struct {
	// success indicates if the conversion was successful.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *ConvertERC20ToCosmosResponse) Reset()         { *m = ConvertERC20ToCosmosResponse{} }
func (m *ConvertERC20ToCosmosResponse) String() string { return proto.CompactTextString(m) }
func (*ConvertERC20ToCosmosResponse) ProtoMessage()    {}
func (*ConvertERC20ToCosmosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba53552aa354681a, []int{1}
}
func (m *ConvertERC20ToCosmosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConvertERC20ToCosmosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConvertERC20ToCosmosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConvertERC20ToCosmosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertERC20ToCosmosResponse.Merge(m, src)
}
func (m *ConvertERC20ToCosmosResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConvertERC20ToCosmosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertERC20ToCosmosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertERC20ToCosmosResponse proto.InternalMessageInfo

func (m *ConvertERC20ToCosmosResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// ConvertCosmosToERC20Request defines a message to convert Cosmos tokens to ERC20 tokens.
type ConvertCosmosToERC20Request struct {
	// token is the address of the ERC20 token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// recipient is the address of the recipient.
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// amount is the amount of tokens to convert.
	Amount *cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount,omitempty"`
}

func (m *ConvertCosmosToERC20Request) Reset()         { *m = ConvertCosmosToERC20Request{} }
func (m *ConvertCosmosToERC20Request) String() string { return proto.CompactTextString(m) }
func (*ConvertCosmosToERC20Request) ProtoMessage()    {}
func (*ConvertCosmosToERC20Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba53552aa354681a, []int{2}
}
func (m *ConvertCosmosToERC20Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConvertCosmosToERC20Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConvertCosmosToERC20Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConvertCosmosToERC20Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertCosmosToERC20Request.Merge(m, src)
}
func (m *ConvertCosmosToERC20Request) XXX_Size() int {
	return m.Size()
}
func (m *ConvertCosmosToERC20Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertCosmosToERC20Request.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertCosmosToERC20Request proto.InternalMessageInfo

func (m *ConvertCosmosToERC20Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ConvertCosmosToERC20Request) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

// ConvertCosmosToERC20Response defines the response for the ConvertCosmosToERC20Request.
type ConvertCosmosToERC20Response struct {
	// success indicates if the conversion was successful.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *ConvertCosmosToERC20Response) Reset()         { *m = ConvertCosmosToERC20Response{} }
func (m *ConvertCosmosToERC20Response) String() string { return proto.CompactTextString(m) }
func (*ConvertCosmosToERC20Response) ProtoMessage()    {}
func (*ConvertCosmosToERC20Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba53552aa354681a, []int{3}
}
func (m *ConvertCosmosToERC20Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConvertCosmosToERC20Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConvertCosmosToERC20Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConvertCosmosToERC20Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertCosmosToERC20Response.Merge(m, src)
}
func (m *ConvertCosmosToERC20Response) XXX_Size() int {
	return m.Size()
}
func (m *ConvertCosmosToERC20Response) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertCosmosToERC20Response.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertCosmosToERC20Response proto.InternalMessageInfo

func (m *ConvertCosmosToERC20Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ConvertERC20ToCosmosRequest)(nil), "polaris.erc20.v1alpha1.ConvertERC20ToCosmosRequest")
	proto.RegisterType((*ConvertERC20ToCosmosResponse)(nil), "polaris.erc20.v1alpha1.ConvertERC20ToCosmosResponse")
	proto.RegisterType((*ConvertCosmosToERC20Request)(nil), "polaris.erc20.v1alpha1.ConvertCosmosToERC20Request")
	proto.RegisterType((*ConvertCosmosToERC20Response)(nil), "polaris.erc20.v1alpha1.ConvertCosmosToERC20Response")
}

func init() { proto.RegisterFile("polaris/erc20/v1alpha1/tx.proto", fileDescriptor_ba53552aa354681a) }

var fileDescriptor_ba53552aa354681a = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0x73, 0x41, 0x2d, 0xd4, 0x1b, 0xa7, 0x16, 0xd2, 0x50, 0x5d, 0x51, 0x06, 0x54, 0x81,
	0xb0, 0x9b, 0x84, 0x81, 0xb5, 0x8d, 0x90, 0x60, 0x60, 0x39, 0x3a, 0xb1, 0x44, 0xae, 0xf3, 0x74,
	0xb1, 0x92, 0xf8, 0x19, 0xdb, 0x39, 0x95, 0x0d, 0xf1, 0x09, 0x90, 0xf8, 0x22, 0x80, 0xf8, 0x10,
	0x8c, 0x15, 0x2c, 0x88, 0x01, 0xa1, 0x04, 0x89, 0x0f, 0xc1, 0x82, 0x6a, 0xbb, 0xd0, 0xa2, 0x53,
	0x4f, 0x8c, 0x4c, 0x77, 0xcf, 0xfe, 0xff, 0x9f, 0x7f, 0xef, 0xd9, 0x8f, 0x6c, 0x6b, 0x9c, 0x72,
	0x23, 0x2d, 0x03, 0x23, 0x7a, 0xbb, 0xac, 0xec, 0xf2, 0xa9, 0x1e, 0xf3, 0x2e, 0x73, 0x47, 0x54,
	0x1b, 0x74, 0x98, 0x5e, 0x8b, 0x02, 0xea, 0x05, 0xf4, 0x54, 0xd0, 0xbe, 0x2e, 0xd0, 0xce, 0xd0,
	0xb2, 0x99, 0x2d, 0x58, 0xd9, 0x3d, 0xf9, 0x04, 0x43, 0x7b, 0x33, 0x6c, 0x0c, 0x7d, 0xc4, 0x42,
	0x10, 0xb7, 0xd6, 0x0b, 0x2c, 0x30, 0xac, 0x9f, 0xfc, 0xc5, 0xd5, 0xad, 0x02, 0xb1, 0x98, 0x02,
	0xe3, 0x5a, 0x32, 0xae, 0x14, 0x3a, 0xee, 0x24, 0xaa, 0xe8, 0xe9, 0xbc, 0x4b, 0xc8, 0x8d, 0x01,
	0xaa, 0x12, 0x8c, 0x7b, 0x90, 0x0f, 0x7a, 0xbb, 0x07, 0x38, 0xf0, 0x29, 0x73, 0x78, 0x36, 0x07,
	0xeb, 0xd2, 0x3b, 0x64, 0xc5, 0xe1, 0x04, 0x54, 0x2b, 0xb9, 0x99, 0xec, 0xac, 0xed, 0x6f, 0x7c,
	0x7c, 0x7f, 0xf7, 0x6a, 0x3c, 0x74, 0x4f, 0x88, 0xbd, 0xd1, 0xc8, 0x80, 0xb5, 0x79, 0xd0, 0xa4,
	0x7d, 0xb2, 0x66, 0x40, 0x48, 0x2d, 0x41, 0xb9, 0x56, 0xf3, 0x22, 0xc3, 0x1f, 0x5d, 0xda, 0x25,
	0xab, 0x7c, 0x86, 0x73, 0xe5, 0x5a, 0x97, 0xbc, 0x63, 0xf3, 0xcb, 0xd7, 0xed, 0x8d, 0xe0, 0xb0,
	0xa3, 0x09, 0x95, 0xc8, 0x66, 0xdc, 0x8d, 0xe9, 0x23, 0xe5, 0xf2, 0x28, 0xec, 0xdc, 0x27, 0x5b,
	0xd5, 0xcc, 0x56, 0xa3, 0xb2, 0x90, 0xb6, 0xc8, 0x65, 0x3b, 0x17, 0x02, 0xac, 0xf5, 0xd8, 0x57,
	0xf2, 0xd3, 0xf0, 0x6c, 0xb9, 0xc1, 0x73, 0x80, 0x3e, 0xc5, 0x7f, 0x52, 0xee, 0x5f, 0xcc, 0x75,
	0xe5, 0xf6, 0x7e, 0x36, 0x09, 0x79, 0x6c, 0x8b, 0x27, 0x60, 0x4a, 0x29, 0x20, 0x7d, 0x9b, 0x90,
	0xf5, 0xaa, 0xc6, 0xa5, 0x7d, 0x5a, 0xfd, 0x0c, 0xe9, 0x05, 0x4f, 0xa3, 0x7d, 0xef, 0xdf, 0x4c,
	0x01, 0xb6, 0x43, 0x5f, 0x7e, 0xfa, 0xfe, 0xba, 0xb9, 0xd3, 0xb9, 0xc5, 0xce, 0x8f, 0x86, 0x08,
	0xa6, 0xa1, 0x8f, 0x86, 0x0e, 0x87, 0xa1, 0x27, 0x67, 0x99, 0xcf, 0x55, 0x5f, 0xcb, 0x5c, 0x75,
	0xbf, 0xb5, 0xcc, 0x95, 0x0d, 0xae, 0x65, 0x8e, 0x23, 0xe9, 0x30, 0xd0, 0xb7, 0x57, 0x5e, 0xfc,
	0x78, 0x73, 0x3b, 0xd9, 0x7f, 0xf8, 0x61, 0x91, 0x25, 0xc7, 0x8b, 0x2c, 0xf9, 0xb6, 0xc8, 0x92,
	0x57, 0xcb, 0xac, 0x71, 0xbc, 0xcc, 0x1a, 0x9f, 0x97, 0x59, 0xe3, 0x29, 0xd5, 0x93, 0x82, 0x1e,
	0x82, 0xe1, 0x62, 0xcc, 0xa5, 0xa2, 0x23, 0x28, 0x7f, 0x67, 0x8e, 0x73, 0x7f, 0x14, 0x8f, 0x70,
	0xcf, 0x35, 0xd8, 0xc3, 0x55, 0x3f, 0xac, 0xfd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2c,
	0x36, 0xc6, 0x4f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// ConvertERC20ToCosmos converts ERC20 tokens to Cosmos tokens.
	ConvertERC20ToCosmos(ctx context.Context, in *ConvertERC20ToCosmosRequest, opts ...grpc.CallOption) (*ConvertERC20ToCosmosResponse, error)
	// ConvertCosmosToERC20 converts Cosmos tokens to ERC20 tokens.
	ConvertCosmosToERC20(ctx context.Context, in *ConvertCosmosToERC20Request, opts ...grpc.CallOption) (*ConvertCosmosToERC20Response, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) ConvertERC20ToCosmos(ctx context.Context, in *ConvertERC20ToCosmosRequest, opts ...grpc.CallOption) (*ConvertERC20ToCosmosResponse, error) {
	out := new(ConvertERC20ToCosmosResponse)
	err := c.cc.Invoke(ctx, "/polaris.erc20.v1alpha1.MsgService/ConvertERC20ToCosmos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConvertCosmosToERC20(ctx context.Context, in *ConvertCosmosToERC20Request, opts ...grpc.CallOption) (*ConvertCosmosToERC20Response, error) {
	out := new(ConvertCosmosToERC20Response)
	err := c.cc.Invoke(ctx, "/polaris.erc20.v1alpha1.MsgService/ConvertCosmosToERC20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// ConvertERC20ToCosmos converts ERC20 tokens to Cosmos tokens.
	ConvertERC20ToCosmos(context.Context, *ConvertERC20ToCosmosRequest) (*ConvertERC20ToCosmosResponse, error)
	// ConvertCosmosToERC20 converts Cosmos tokens to ERC20 tokens.
	ConvertCosmosToERC20(context.Context, *ConvertCosmosToERC20Request) (*ConvertCosmosToERC20Response, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) ConvertERC20ToCosmos(ctx context.Context, req *ConvertERC20ToCosmosRequest) (*ConvertERC20ToCosmosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertERC20ToCosmos not implemented")
}
func (*UnimplementedMsgServiceServer) ConvertCosmosToERC20(ctx context.Context, req *ConvertCosmosToERC20Request) (*ConvertCosmosToERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCosmosToERC20 not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_ConvertERC20ToCosmos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertERC20ToCosmosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConvertERC20ToCosmos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polaris.erc20.v1alpha1.MsgService/ConvertERC20ToCosmos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConvertERC20ToCosmos(ctx, req.(*ConvertERC20ToCosmosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConvertCosmosToERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertCosmosToERC20Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConvertCosmosToERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polaris.erc20.v1alpha1.MsgService/ConvertCosmosToERC20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConvertCosmosToERC20(ctx, req.(*ConvertCosmosToERC20Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.erc20.v1alpha1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertERC20ToCosmos",
			Handler:    _MsgService_ConvertERC20ToCosmos_Handler,
		},
		{
			MethodName: "ConvertCosmosToERC20",
			Handler:    _MsgService_ConvertCosmosToERC20_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "polaris/erc20/v1alpha1/tx.proto",
}

func (m *ConvertERC20ToCosmosRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConvertERC20ToCosmosRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConvertERC20ToCosmosRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != nil {
		{
			size := m.Amount.Size()
			i -= size
			if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConvertERC20ToCosmosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConvertERC20ToCosmosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConvertERC20ToCosmosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConvertCosmosToERC20Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConvertCosmosToERC20Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConvertCosmosToERC20Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != nil {
		{
			size := m.Amount.Size()
			i -= size
			if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConvertCosmosToERC20Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConvertCosmosToERC20Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConvertCosmosToERC20Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConvertERC20ToCosmosRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ConvertERC20ToCosmosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func (m *ConvertCosmosToERC20Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ConvertCosmosToERC20Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConvertERC20ToCosmosRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConvertERC20ToCosmosRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConvertERC20ToCosmosRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.Amount = &v
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConvertERC20ToCosmosResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConvertERC20ToCosmosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConvertERC20ToCosmosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConvertCosmosToERC20Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConvertCosmosToERC20Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConvertCosmosToERC20Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.Amount = &v
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConvertCosmosToERC20Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConvertCosmosToERC20Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConvertCosmosToERC20Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
