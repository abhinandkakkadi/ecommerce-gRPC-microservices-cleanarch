// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: pkg/cart/pb/cart.proto

package pb

import (
	context "context"
	"fmt"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Productid int64 `protobuf:"varint,1,opt,name=productid,proto3" json:"productid,omitempty"`
	Userid    int64 `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetProductid() int64 {
	if x != nil {
		return x.Productid
	}
	return 0
}

func (x *AddToCartRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DisplayCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *DisplayCartRequest) Reset() {
	*x = DisplayCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplayCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplayCartRequest) ProtoMessage() {}

func (x *DisplayCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplayCartRequest.ProtoReflect.Descriptor instead.
func (*DisplayCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *DisplayCartRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prdoductid int64   `protobuf:"varint,1,opt,name=prdoductid,proto3" json:"prdoductid,omitempty"`
	Moviename  string  `protobuf:"bytes,2,opt,name=moviename,proto3" json:"moviename,omitempty"`
	Quantity   int64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Totalprice float32 `protobuf:"fixed32,4,opt,name=totalprice,proto3" json:"totalprice,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *Cart) GetPrdoductid() int64 {
	if x != nil {
		return x.Prdoductid
	}
	return 0
}

func (x *Cart) GetMoviename() string {
	if x != nil {
		return x.Moviename
	}
	return ""
}

func (x *Cart) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Cart) GetTotalprice() float32 {
	if x != nil {
		return x.Totalprice
	}
	return 0
}

type DisplayCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Totalprice   float32 `protobuf:"fixed32,2,opt,name=totalprice,proto3" json:"totalprice,omitempty"`
	Cartproducts []*Cart `protobuf:"bytes,3,rep,name=cartproducts,proto3" json:"cartproducts,omitempty"`
}

func (x *DisplayCartResponse) Reset() {
	*x = DisplayCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplayCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplayCartResponse) ProtoMessage() {}

func (x *DisplayCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplayCartResponse.ProtoReflect.Descriptor instead.
func (*DisplayCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{4}
}

func (x *DisplayCartResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DisplayCartResponse) GetTotalprice() float32 {
	if x != nil {
		return x.Totalprice
	}
	return 0
}

func (x *DisplayCartResponse) GetCartproducts() []*Cart {
	if x != nil {
		return x.Cartproducts
	}
	return nil
}

var File_pkg_cart_pb_cart_proto protoreflect.FileDescriptor

var file_pkg_cart_pb_cart_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x48,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x72, 0x64, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x70, 0x72, 0x64, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x63, 0x61,
	0x72, 0x74, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x0c, 0x63, 0x61,
	0x72, 0x74, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x93, 0x01, 0x0a, 0x0b, 0x43,
	0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_cart_pb_cart_proto_rawDescOnce sync.Once
	file_pkg_cart_pb_cart_proto_rawDescData = file_pkg_cart_pb_cart_proto_rawDesc
)

func file_pkg_cart_pb_cart_proto_rawDescGZIP() []byte {
	file_pkg_cart_pb_cart_proto_rawDescOnce.Do(func() {
		file_pkg_cart_pb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_cart_pb_cart_proto_rawDescData)
	})
	return file_pkg_cart_pb_cart_proto_rawDescData
}

var file_pkg_cart_pb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_cart_pb_cart_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),    // 0: cart.AddToCartRequest
	(*AddToCartResponse)(nil),   // 1: cart.AddToCartResponse
	(*DisplayCartRequest)(nil),  // 2: cart.DisplayCartRequest
	(*Cart)(nil),                // 3: cart.Cart
	(*DisplayCartResponse)(nil), // 4: cart.DisplayCartResponse
}
var file_pkg_cart_pb_cart_proto_depIdxs = []int32{
	3, // 0: cart.DisplayCartResponse.cartproducts:type_name -> cart.Cart
	0, // 1: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	2, // 2: cart.CartService.DisplayCart:input_type -> cart.DisplayCartRequest
	1, // 3: cart.CartService.AddToCart:output_type -> cart.AddToCartResponse
	4, // 4: cart.CartService.DisplayCart:output_type -> cart.DisplayCartResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_cart_pb_cart_proto_init() }
func file_pkg_cart_pb_cart_proto_init() {
	if File_pkg_cart_pb_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_cart_pb_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponse); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplayCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplayCartResponse); i {
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
			RawDescriptor: file_pkg_cart_pb_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_cart_pb_cart_proto_goTypes,
		DependencyIndexes: file_pkg_cart_pb_cart_proto_depIdxs,
		MessageInfos:      file_pkg_cart_pb_cart_proto_msgTypes,
	}.Build()
	File_pkg_cart_pb_cart_proto = out.File
	file_pkg_cart_pb_cart_proto_rawDesc = nil
	file_pkg_cart_pb_cart_proto_goTypes = nil
	file_pkg_cart_pb_cart_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error)
	DisplayCart(ctx context.Context, in *DisplayCartRequest, opts ...grpc.CallOption) (*DisplayCartResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error) {
	out := new(AddToCartResponse)
	fmt.Println("cart request userID := ",in.Productid,"productID := ",in.Productid)
	err := c.cc.Invoke(ctx, "/cart.CartService/AddToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DisplayCart(ctx context.Context, in *DisplayCartRequest, opts ...grpc.CallOption) (*DisplayCartResponse, error) {
	out := new(DisplayCartResponse)
	err := c.cc.Invoke(ctx, "/cart.CartService/DisplayCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error)
	DisplayCart(context.Context, *DisplayCartRequest) (*DisplayCartResponse, error)
}

// UnimplementedCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (*UnimplementedCartServiceServer) AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (*UnimplementedCartServiceServer) DisplayCart(context.Context, *DisplayCartRequest) (*DisplayCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayCart not implemented")
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/AddToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DisplayCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DisplayCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/DisplayCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DisplayCart(ctx, req.(*DisplayCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToCart",
			Handler:    _CartService_AddToCart_Handler,
		},
		{
			MethodName: "DisplayCart",
			Handler:    _CartService_DisplayCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/cart/pb/cart.proto",
}
