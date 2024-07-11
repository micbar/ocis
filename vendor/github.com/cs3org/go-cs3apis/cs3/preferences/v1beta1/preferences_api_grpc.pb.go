// Copyright 2018-2019 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cs3/preferences/v1beta1/preferences_api.proto

package preferencesv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PreferencesAPI_SetKey_FullMethodName = "/cs3.preferences.v1beta1.PreferencesAPI/SetKey"
	PreferencesAPI_GetKey_FullMethodName = "/cs3.preferences.v1beta1.PreferencesAPI/GetKey"
)

// PreferencesAPIClient is the client API for PreferencesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreferencesAPIClient interface {
	// Maps the key-value pair.
	SetKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*SetKeyResponse, error)
	// Returns the value associated with the
	// requested key.
	GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResponse, error)
}

type preferencesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPreferencesAPIClient(cc grpc.ClientConnInterface) PreferencesAPIClient {
	return &preferencesAPIClient{cc}
}

func (c *preferencesAPIClient) SetKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*SetKeyResponse, error) {
	out := new(SetKeyResponse)
	err := c.cc.Invoke(ctx, PreferencesAPI_SetKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *preferencesAPIClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResponse, error) {
	out := new(GetKeyResponse)
	err := c.cc.Invoke(ctx, PreferencesAPI_GetKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreferencesAPIServer is the server API for PreferencesAPI service.
// All implementations should embed UnimplementedPreferencesAPIServer
// for forward compatibility
type PreferencesAPIServer interface {
	// Maps the key-value pair.
	SetKey(context.Context, *SetKeyRequest) (*SetKeyResponse, error)
	// Returns the value associated with the
	// requested key.
	GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error)
}

// UnimplementedPreferencesAPIServer should be embedded to have forward compatible implementations.
type UnimplementedPreferencesAPIServer struct {
}

func (UnimplementedPreferencesAPIServer) SetKey(context.Context, *SetKeyRequest) (*SetKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKey not implemented")
}
func (UnimplementedPreferencesAPIServer) GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}

// UnsafePreferencesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreferencesAPIServer will
// result in compilation errors.
type UnsafePreferencesAPIServer interface {
	mustEmbedUnimplementedPreferencesAPIServer()
}

func RegisterPreferencesAPIServer(s grpc.ServiceRegistrar, srv PreferencesAPIServer) {
	s.RegisterService(&PreferencesAPI_ServiceDesc, srv)
}

func _PreferencesAPI_SetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesAPIServer).SetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PreferencesAPI_SetKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesAPIServer).SetKey(ctx, req.(*SetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PreferencesAPI_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesAPIServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PreferencesAPI_GetKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesAPIServer).GetKey(ctx, req.(*GetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PreferencesAPI_ServiceDesc is the grpc.ServiceDesc for PreferencesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PreferencesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.preferences.v1beta1.PreferencesAPI",
	HandlerType: (*PreferencesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetKey",
			Handler:    _PreferencesAPI_SetKey_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _PreferencesAPI_GetKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/preferences/v1beta1/preferences_api.proto",
}
