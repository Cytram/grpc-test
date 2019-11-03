// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package imagescaler

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

type ScaleImageRequest struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleImageRequest) Reset()         { *m = ScaleImageRequest{} }
func (m *ScaleImageRequest) String() string { return proto.CompactTextString(m) }
func (*ScaleImageRequest) ProtoMessage()    {}
func (*ScaleImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ScaleImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleImageRequest.Unmarshal(m, b)
}
func (m *ScaleImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleImageRequest.Marshal(b, m, deterministic)
}
func (m *ScaleImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleImageRequest.Merge(m, src)
}
func (m *ScaleImageRequest) XXX_Size() int {
	return xxx_messageInfo_ScaleImageRequest.Size(m)
}
func (m *ScaleImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleImageRequest proto.InternalMessageInfo

func (m *ScaleImageRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type ScaleImageReply struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleImageReply) Reset()         { *m = ScaleImageReply{} }
func (m *ScaleImageReply) String() string { return proto.CompactTextString(m) }
func (*ScaleImageReply) ProtoMessage()    {}
func (*ScaleImageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ScaleImageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleImageReply.Unmarshal(m, b)
}
func (m *ScaleImageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleImageReply.Marshal(b, m, deterministic)
}
func (m *ScaleImageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleImageReply.Merge(m, src)
}
func (m *ScaleImageReply) XXX_Size() int {
	return xxx_messageInfo_ScaleImageReply.Size(m)
}
func (m *ScaleImageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleImageReply.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleImageReply proto.InternalMessageInfo

func (m *ScaleImageReply) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Image struct {
	// Image content, represented as a stream of bytes.
	// Note: As with all `bytes` fields, protobuffers use a pure binary
	// representation, whereas JSON representations use base64.
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// If both content and source is present use content.
	Source               *ImageSource `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Image) GetSource() *ImageSource {
	if m != nil {
		return m.Source
	}
	return nil
}

type ImageSource struct {
	// The URI of the source document. Can be either:
	//
	// 1. A publicly-accessible image HTTP/HTTPS URL. When fetching images from
	HttpUri              string   `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageSource) Reset()         { *m = ImageSource{} }
func (m *ImageSource) String() string { return proto.CompactTextString(m) }
func (*ImageSource) ProtoMessage()    {}
func (*ImageSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *ImageSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSource.Unmarshal(m, b)
}
func (m *ImageSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSource.Marshal(b, m, deterministic)
}
func (m *ImageSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSource.Merge(m, src)
}
func (m *ImageSource) XXX_Size() int {
	return xxx_messageInfo_ImageSource.Size(m)
}
func (m *ImageSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSource.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSource proto.InternalMessageInfo

func (m *ImageSource) GetHttpUri() string {
	if m != nil {
		return m.HttpUri
	}
	return ""
}

func init() {
	proto.RegisterType((*ScaleImageRequest)(nil), "imagescaler.ScaleImageRequest")
	proto.RegisterType((*ScaleImageReply)(nil), "imagescaler.ScaleImageReply")
	proto.RegisterType((*Image)(nil), "imagescaler.Image")
	proto.RegisterType((*ImageSource)(nil), "imagescaler.ImageSource")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xcc, 0x4d, 0x4c, 0x4f, 0x2d, 0x4e, 0x4e, 0xcc,
	0x49, 0x2d, 0x52, 0xb2, 0xe5, 0x12, 0x0c, 0x06, 0xb1, 0x3c, 0x41, 0x62, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x1a, 0x5c, 0xac, 0x60, 0x35, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0x42, 0x7a, 0x48, 0x3a, 0xf4, 0x20, 0x2a, 0x21, 0x0a, 0x94, 0xb4, 0xb9, 0xf8, 0x91, 0xb5, 0x17,
	0xe4, 0x54, 0x0a, 0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x80, 0xb5, 0xf3,
	0x04, 0xc1, 0xb8, 0x4a, 0xc1, 0x5c, 0xac, 0x60, 0x75, 0xb8, 0x95, 0x08, 0x19, 0x70, 0xb1, 0x15,
	0xe7, 0x97, 0x16, 0x25, 0xa7, 0x4a, 0x30, 0x81, 0xad, 0x96, 0xc0, 0xb4, 0x3a, 0x18, 0x2c, 0x1f,
	0x04, 0x55, 0xa7, 0xa4, 0xc1, 0xc5, 0x8d, 0x24, 0x2c, 0x24, 0xc9, 0xc5, 0x91, 0x51, 0x52, 0x52,
	0x10, 0x5f, 0x5a, 0x94, 0x09, 0x36, 0x9b, 0x33, 0x88, 0x1d, 0xc4, 0x0f, 0x2d, 0xca, 0x34, 0x8a,
	0x86, 0xa9, 0x04, 0x1b, 0x26, 0xe4, 0xc3, 0xc5, 0x85, 0x70, 0xba, 0x90, 0x1c, 0x8a, 0x45, 0x18,
	0x41, 0x22, 0x25, 0x83, 0x53, 0xbe, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0xb6, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x59, 0x16, 0x12, 0x68, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageScalerClient is the client API for ImageScaler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageScalerClient interface {
	ScaleImage(ctx context.Context, in *ScaleImageRequest, opts ...grpc.CallOption) (*ScaleImageReply, error)
}

type imageScalerClient struct {
	cc *grpc.ClientConn
}

func NewImageScalerClient(cc *grpc.ClientConn) ImageScalerClient {
	return &imageScalerClient{cc}
}

func (c *imageScalerClient) ScaleImage(ctx context.Context, in *ScaleImageRequest, opts ...grpc.CallOption) (*ScaleImageReply, error) {
	out := new(ScaleImageReply)
	err := c.cc.Invoke(ctx, "/imagescaler.ImageScaler/ScaleImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageScalerServer is the server API for ImageScaler service.
type ImageScalerServer interface {
	ScaleImage(context.Context, *ScaleImageRequest) (*ScaleImageReply, error)
}

// UnimplementedImageScalerServer can be embedded to have forward compatible implementations.
type UnimplementedImageScalerServer struct {
}

func (*UnimplementedImageScalerServer) ScaleImage(ctx context.Context, req *ScaleImageRequest) (*ScaleImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleImage not implemented")
}

func RegisterImageScalerServer(s *grpc.Server, srv ImageScalerServer) {
	s.RegisterService(&_ImageScaler_serviceDesc, srv)
}

func _ImageScaler_ScaleImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageScalerServer).ScaleImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imagescaler.ImageScaler/ScaleImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageScalerServer).ScaleImage(ctx, req.(*ScaleImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageScaler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imagescaler.ImageScaler",
	HandlerType: (*ImageScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScaleImage",
			Handler:    _ImageScaler_ScaleImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
