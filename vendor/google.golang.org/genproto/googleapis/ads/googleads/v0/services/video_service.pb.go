// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/video_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [VideoService.GetVideo][google.ads.googleads.v0.services.VideoService.GetVideo].
type GetVideoRequest struct {
	// The resource name of the video to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVideoRequest) Reset()         { *m = GetVideoRequest{} }
func (m *GetVideoRequest) String() string { return proto.CompactTextString(m) }
func (*GetVideoRequest) ProtoMessage()    {}
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_service_8b8505c516feb5c1, []int{0}
}
func (m *GetVideoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoRequest.Unmarshal(m, b)
}
func (m *GetVideoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoRequest.Marshal(b, m, deterministic)
}
func (dst *GetVideoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoRequest.Merge(dst, src)
}
func (m *GetVideoRequest) XXX_Size() int {
	return xxx_messageInfo_GetVideoRequest.Size(m)
}
func (m *GetVideoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoRequest proto.InternalMessageInfo

func (m *GetVideoRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetVideoRequest)(nil), "google.ads.googleads.v0.services.GetVideoRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoServiceClient interface {
	// Returns the requested video in full detail.
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*resources.Video, error)
}

type videoServiceClient struct {
	cc *grpc.ClientConn
}

func NewVideoServiceClient(cc *grpc.ClientConn) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*resources.Video, error) {
	out := new(resources.Video)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.VideoService/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
type VideoServiceServer interface {
	// Returns the requested video in full detail.
	GetVideo(context.Context, *GetVideoRequest) (*resources.Video, error)
}

func RegisterVideoServiceServer(s *grpc.Server, srv VideoServiceServer) {
	s.RegisterService(&_VideoService_serviceDesc, srv)
}

func _VideoService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.VideoService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/video_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/video_service.proto", fileDescriptor_video_service_8b8505c516feb5c1)
}

var fileDescriptor_video_service_8b8505c516feb5c1 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0x3e, 0xf8, 0xd0, 0x50, 0x11, 0xb3, 0x92, 0xe2, 0xa2, 0xd4, 0x2e, 0x4a, 0xc1,
	0x99, 0xf8, 0x07, 0x17, 0x23, 0x2e, 0xd2, 0x4d, 0x5d, 0x49, 0xa9, 0x90, 0x85, 0x04, 0xca, 0xd8,
	0x5c, 0x42, 0xa0, 0xc9, 0xad, 0xb9, 0xd3, 0x6c, 0xc4, 0x8d, 0xaf, 0xe0, 0xca, 0xad, 0xcb, 0x3e,
	0x8a, 0x5b, 0x5f, 0xc1, 0x95, 0x4f, 0xe0, 0x52, 0x92, 0xe9, 0x04, 0x15, 0x42, 0x77, 0x27, 0x93,
	0xdf, 0x39, 0x73, 0xef, 0x19, 0xe7, 0x2c, 0x46, 0x8c, 0xe7, 0xc0, 0x65, 0x44, 0x5c, 0xcb, 0x52,
	0x15, 0x1e, 0x27, 0xc8, 0x8b, 0x64, 0x06, 0xc4, 0x8b, 0x24, 0x02, 0x9c, 0xae, 0x3f, 0xd9, 0x22,
	0x47, 0x85, 0x6e, 0x47, 0xa3, 0x4c, 0x46, 0xc4, 0x6a, 0x17, 0x2b, 0x3c, 0x66, 0x5c, 0xed, 0xa3,
	0xa6, 0xdc, 0x1c, 0x08, 0x97, 0x79, 0x1d, 0xac, 0x03, 0xdb, 0x07, 0x06, 0x5f, 0x24, 0x5c, 0x66,
	0x19, 0x2a, 0xa9, 0x12, 0xcc, 0x48, 0xff, 0xed, 0x9e, 0x3b, 0xbb, 0x23, 0x50, 0x41, 0xc9, 0x4f,
	0xe0, 0x7e, 0x09, 0xa4, 0xdc, 0x43, 0x67, 0xc7, 0x24, 0x4d, 0x33, 0x99, 0xc2, 0xbe, 0xd5, 0xb1,
	0xfa, 0xdb, 0x93, 0x96, 0x39, 0xbc, 0x96, 0x29, 0x9c, 0xac, 0x2c, 0xa7, 0x55, 0xb9, 0x6e, 0xf4,
	0x58, 0xee, 0x8b, 0xe5, 0x6c, 0x99, 0x24, 0xf7, 0x98, 0x6d, 0xda, 0x82, 0xfd, 0xb9, 0xb5, 0xdd,
	0x6f, 0xb4, 0xd4, 0x6b, 0xb1, 0xca, 0xd0, 0xf5, 0x9e, 0xde, 0x3f, 0x9e, 0xed, 0x81, 0xdb, 0x2f,
	0x77, 0x7e, 0xf8, 0x35, 0xea, 0xe5, 0x6c, 0x49, 0x0a, 0x53, 0xc8, 0x89, 0x0f, 0x74, 0x09, 0xc4,
	0x07, 0x8f, 0xc3, 0x2f, 0xcb, 0xe9, 0xcd, 0x30, 0xdd, 0x38, 0xd4, 0x70, 0xef, 0xe7, 0x4a, 0xe3,
	0xb2, 0xa0, 0xb1, 0x75, 0x7b, 0xb5, 0xb6, 0xc5, 0x38, 0x97, 0x59, 0xcc, 0x30, 0x8f, 0x79, 0x0c,
	0x59, 0x55, 0x9f, 0xe9, 0x7f, 0x91, 0x50, 0xf3, 0x33, 0x5f, 0x18, 0xf1, 0x6a, 0xff, 0x1b, 0xf9,
	0xfe, 0xca, 0xee, 0x8c, 0x74, 0xa0, 0x1f, 0x11, 0xd3, 0xb2, 0x54, 0x81, 0xc7, 0xd6, 0x17, 0xd3,
	0x9b, 0x41, 0x42, 0x3f, 0xa2, 0xb0, 0x46, 0xc2, 0xc0, 0x0b, 0x0d, 0xf2, 0x69, 0xf7, 0xf4, 0xb9,
	0x10, 0x7e, 0x44, 0x42, 0xd4, 0x90, 0x10, 0x81, 0x27, 0x84, 0xc1, 0xee, 0xfe, 0x57, 0x73, 0x9e,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x69, 0x6b, 0x45, 0x12, 0x8d, 0x02, 0x00, 0x00,
}
