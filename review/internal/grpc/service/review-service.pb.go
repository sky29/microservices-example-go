// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microservices-example-go/review/internal/proto-files/service/review-service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	domain "microservices-example-go/review/internal/grpc/domain"
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

type GetReviewResponse struct {
	Review               *domain.Review `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetReviewResponse) Reset()         { *m = GetReviewResponse{} }
func (m *GetReviewResponse) String() string { return proto.CompactTextString(m) }
func (*GetReviewResponse) ProtoMessage()    {}
func (*GetReviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe594616099c2767, []int{0}
}

func (m *GetReviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReviewResponse.Unmarshal(m, b)
}
func (m *GetReviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReviewResponse.Marshal(b, m, deterministic)
}
func (m *GetReviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReviewResponse.Merge(m, src)
}
func (m *GetReviewResponse) XXX_Size() int {
	return xxx_messageInfo_GetReviewResponse.Size(m)
}
func (m *GetReviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetReviewResponse proto.InternalMessageInfo

func (m *GetReviewResponse) GetReview() *domain.Review {
	if m != nil {
		return m.Review
	}
	return nil
}

func init() {
	proto.RegisterType((*GetReviewResponse)(nil), "service.GetReviewResponse")
}

func init() {
	proto.RegisterFile("microservices-example-go/review/internal/proto-files/service/review-service.proto", fileDescriptor_fe594616099c2767)
}

var fileDescriptor_fe594616099c2767 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0xcc, 0xcd, 0x4c, 0x2e,
	0xca, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0xd5, 0x4d, 0xcf, 0xd7, 0x2f, 0x4a, 0x2d, 0xcb, 0x4c, 0x2d, 0xd7, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcb, 0xcc, 0x49, 0x2d,
	0xd6, 0x87, 0xaa, 0x87, 0xaa, 0xd1, 0x85, 0x72, 0xf5, 0xc0, 0x4a, 0x84, 0xd8, 0xa1, 0x5c, 0x29,
	0x0f, 0xb2, 0xcc, 0x4e, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83, 0x2a, 0x81, 0x18, 0xa9, 0x64, 0xcd,
	0x25, 0xe8, 0x9e, 0x5a, 0x12, 0x04, 0x16, 0x0a, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x52, 0xe3, 0x62, 0x83, 0x28, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd3, 0x83, 0x68,
	0xd5, 0x83, 0x88, 0x06, 0x41, 0x65, 0x8d, 0x9c, 0xb8, 0x78, 0x21, 0x3a, 0x83, 0x21, 0x2e, 0x11,
	0x32, 0xe4, 0x62, 0x4e, 0x4f, 0x2d, 0x11, 0x42, 0x53, 0x2f, 0x25, 0xa5, 0x07, 0xf3, 0x07, 0x86,
	0x5d, 0x4e, 0xe6, 0x51, 0xa6, 0x44, 0x7b, 0x26, 0xbd, 0xa8, 0x20, 0x19, 0x16, 0x42, 0x49, 0x6c,
	0x60, 0x0f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x1e, 0x6e, 0xa3, 0x68, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewServiceClient interface {
	Get(ctx context.Context, in *domain.Review, opts ...grpc.CallOption) (*GetReviewResponse, error)
}

type reviewServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewServiceClient(cc *grpc.ClientConn) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) Get(ctx context.Context, in *domain.Review, opts ...grpc.CallOption) (*GetReviewResponse, error) {
	out := new(GetReviewResponse)
	err := c.cc.Invoke(ctx, "/service.ReviewService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
type ReviewServiceServer interface {
	Get(context.Context, *domain.Review) (*GetReviewResponse, error)
}

// UnimplementedReviewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (*UnimplementedReviewServiceServer) Get(ctx context.Context, req *domain.Review) (*GetReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ReviewService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Get(ctx, req.(*domain.Review))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _ReviewService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservices-example-go/review/internal/proto-files/service/review-service.proto",
}
