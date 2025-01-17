// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/donor.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DonorService_PrepareDonor_FullMethodName  = "/donor.DonorService/PrepareDonor"
	DonorService_CommitDonor_FullMethodName   = "/donor.DonorService/CommitDonor"
	DonorService_RollbackDonor_FullMethodName = "/donor.DonorService/RollbackDonor"
)

// DonorServiceClient is the client API for DonorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DonorServiceClient interface {
	PrepareDonor(ctx context.Context, in *PrepareDonorRequest, opts ...grpc.CallOption) (*PrepareDonorResponse, error)
	CommitDonor(ctx context.Context, in *CommitDonorRequest, opts ...grpc.CallOption) (*CommitDonorResponse, error)
	RollbackDonor(ctx context.Context, in *RollbackDonorRequest, opts ...grpc.CallOption) (*RollbackDonorResponse, error)
}

type donorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDonorServiceClient(cc grpc.ClientConnInterface) DonorServiceClient {
	return &donorServiceClient{cc}
}

func (c *donorServiceClient) PrepareDonor(ctx context.Context, in *PrepareDonorRequest, opts ...grpc.CallOption) (*PrepareDonorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrepareDonorResponse)
	err := c.cc.Invoke(ctx, DonorService_PrepareDonor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donorServiceClient) CommitDonor(ctx context.Context, in *CommitDonorRequest, opts ...grpc.CallOption) (*CommitDonorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommitDonorResponse)
	err := c.cc.Invoke(ctx, DonorService_CommitDonor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donorServiceClient) RollbackDonor(ctx context.Context, in *RollbackDonorRequest, opts ...grpc.CallOption) (*RollbackDonorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RollbackDonorResponse)
	err := c.cc.Invoke(ctx, DonorService_RollbackDonor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DonorServiceServer is the server API for DonorService service.
// All implementations must embed UnimplementedDonorServiceServer
// for forward compatibility.
type DonorServiceServer interface {
	PrepareDonor(context.Context, *PrepareDonorRequest) (*PrepareDonorResponse, error)
	CommitDonor(context.Context, *CommitDonorRequest) (*CommitDonorResponse, error)
	RollbackDonor(context.Context, *RollbackDonorRequest) (*RollbackDonorResponse, error)
	mustEmbedUnimplementedDonorServiceServer()
}

// UnimplementedDonorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDonorServiceServer struct{}

func (UnimplementedDonorServiceServer) PrepareDonor(context.Context, *PrepareDonorRequest) (*PrepareDonorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDonor not implemented")
}
func (UnimplementedDonorServiceServer) CommitDonor(context.Context, *CommitDonorRequest) (*CommitDonorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitDonor not implemented")
}
func (UnimplementedDonorServiceServer) RollbackDonor(context.Context, *RollbackDonorRequest) (*RollbackDonorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackDonor not implemented")
}
func (UnimplementedDonorServiceServer) mustEmbedUnimplementedDonorServiceServer() {}
func (UnimplementedDonorServiceServer) testEmbeddedByValue()                      {}

// UnsafeDonorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DonorServiceServer will
// result in compilation errors.
type UnsafeDonorServiceServer interface {
	mustEmbedUnimplementedDonorServiceServer()
}

func RegisterDonorServiceServer(s grpc.ServiceRegistrar, srv DonorServiceServer) {
	// If the following call pancis, it indicates UnimplementedDonorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DonorService_ServiceDesc, srv)
}

func _DonorService_PrepareDonor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDonorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonorServiceServer).PrepareDonor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DonorService_PrepareDonor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonorServiceServer).PrepareDonor(ctx, req.(*PrepareDonorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DonorService_CommitDonor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitDonorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonorServiceServer).CommitDonor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DonorService_CommitDonor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonorServiceServer).CommitDonor(ctx, req.(*CommitDonorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DonorService_RollbackDonor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackDonorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonorServiceServer).RollbackDonor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DonorService_RollbackDonor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonorServiceServer).RollbackDonor(ctx, req.(*RollbackDonorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DonorService_ServiceDesc is the grpc.ServiceDesc for DonorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DonorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "donor.DonorService",
	HandlerType: (*DonorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareDonor",
			Handler:    _DonorService_PrepareDonor_Handler,
		},
		{
			MethodName: "CommitDonor",
			Handler:    _DonorService_CommitDonor_Handler,
		},
		{
			MethodName: "RollbackDonor",
			Handler:    _DonorService_RollbackDonor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/donor.proto",
}
