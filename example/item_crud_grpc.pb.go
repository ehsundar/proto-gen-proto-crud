// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ItemStorageClient is the client API for ItemStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemStorageClient interface {
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	GetAllItems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllItemsResponse, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type itemStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewItemStorageClient(cc grpc.ClientConnInterface) ItemStorageClient {
	return &itemStorageClient{cc}
}

func (c *itemStorageClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/example.ItemStorage/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemStorageClient) GetAllItems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllItemsResponse, error) {
	out := new(GetAllItemsResponse)
	err := c.cc.Invoke(ctx, "/example.ItemStorage/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemStorageClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, "/example.ItemStorage/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemStorageClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, "/example.ItemStorage/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemStorageClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/example.ItemStorage/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemStorageServer is the server API for ItemStorage service.
// All implementations must embed UnimplementedItemStorageServer
// for forward compatibility
type ItemStorageServer interface {
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	GetAllItems(context.Context, *emptypb.Empty) (*GetAllItemsResponse, error)
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedItemStorageServer()
}

// UnimplementedItemStorageServer must be embedded to have forward compatible implementations.
type UnimplementedItemStorageServer struct {
}

func (UnimplementedItemStorageServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedItemStorageServer) GetAllItems(context.Context, *emptypb.Empty) (*GetAllItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (UnimplementedItemStorageServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedItemStorageServer) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedItemStorageServer) DeleteItem(context.Context, *DeleteItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedItemStorageServer) mustEmbedUnimplementedItemStorageServer() {}

// UnsafeItemStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemStorageServer will
// result in compilation errors.
type UnsafeItemStorageServer interface {
	mustEmbedUnimplementedItemStorageServer()
}

func RegisterItemStorageServer(s grpc.ServiceRegistrar, srv ItemStorageServer) {
	s.RegisterService(&ItemStorage_ServiceDesc, srv)
}

func _ItemStorage_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemStorageServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ItemStorage/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemStorageServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemStorage_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemStorageServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ItemStorage/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemStorageServer).GetAllItems(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemStorage_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemStorageServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ItemStorage/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemStorageServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemStorage_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemStorageServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ItemStorage/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemStorageServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemStorage_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemStorageServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ItemStorage/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemStorageServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemStorage_ServiceDesc is the grpc.ServiceDesc for ItemStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ItemStorage",
	HandlerType: (*ItemStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItem",
			Handler:    _ItemStorage_GetItem_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _ItemStorage_GetAllItems_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _ItemStorage_CreateItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ItemStorage_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ItemStorage_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/item_crud.proto",
}
