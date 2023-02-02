// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/knit/crosstest/test.proto

package crosstestconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	crosstest "github.com/bufbuild/knit-go/internal/gen/buf/knit/crosstest"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// ParentServiceName is the fully-qualified name of the ParentService service.
	ParentServiceName = "buf.knit.crosstest.ParentService"
	// ChildServiceName is the fully-qualified name of the ChildService service.
	ChildServiceName = "buf.knit.crosstest.ChildService"
	// RelationsServiceName is the fully-qualified name of the RelationsService service.
	RelationsServiceName = "buf.knit.crosstest.RelationsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ParentServiceListParentsProcedure is the fully-qualified name of the ParentService's ListParents
	// RPC.
	ParentServiceListParentsProcedure = "/buf.knit.crosstest.ParentService/ListParents"
	// ParentServiceGetParentProcedure is the fully-qualified name of the ParentService's GetParent RPC.
	ParentServiceGetParentProcedure = "/buf.knit.crosstest.ParentService/GetParent"
	// ParentServiceCreateParentProcedure is the fully-qualified name of the ParentService's
	// CreateParent RPC.
	ParentServiceCreateParentProcedure = "/buf.knit.crosstest.ParentService/CreateParent"
	// ParentServiceUpdateParentProcedure is the fully-qualified name of the ParentService's
	// UpdateParent RPC.
	ParentServiceUpdateParentProcedure = "/buf.knit.crosstest.ParentService/UpdateParent"
	// ParentServiceDeleteParentProcedure is the fully-qualified name of the ParentService's
	// DeleteParent RPC.
	ParentServiceDeleteParentProcedure = "/buf.knit.crosstest.ParentService/DeleteParent"
	// ChildServiceListChildrenProcedure is the fully-qualified name of the ChildService's ListChildren
	// RPC.
	ChildServiceListChildrenProcedure = "/buf.knit.crosstest.ChildService/ListChildren"
	// ChildServiceGetChildProcedure is the fully-qualified name of the ChildService's GetChild RPC.
	ChildServiceGetChildProcedure = "/buf.knit.crosstest.ChildService/GetChild"
	// ChildServiceCreateChildProcedure is the fully-qualified name of the ChildService's CreateChild
	// RPC.
	ChildServiceCreateChildProcedure = "/buf.knit.crosstest.ChildService/CreateChild"
	// ChildServiceUpdateChildProcedure is the fully-qualified name of the ChildService's UpdateChild
	// RPC.
	ChildServiceUpdateChildProcedure = "/buf.knit.crosstest.ChildService/UpdateChild"
	// ChildServiceDeleteChildProcedure is the fully-qualified name of the ChildService's DeleteChild
	// RPC.
	ChildServiceDeleteChildProcedure = "/buf.knit.crosstest.ChildService/DeleteChild"
	// RelationsServiceGetChildParentProcedure is the fully-qualified name of the RelationsService's
	// GetChildParent RPC.
	RelationsServiceGetChildParentProcedure = "/buf.knit.crosstest.RelationsService/GetChildParent"
	// RelationsServiceGetParentChildrenProcedure is the fully-qualified name of the RelationsService's
	// GetParentChildren RPC.
	RelationsServiceGetParentChildrenProcedure = "/buf.knit.crosstest.RelationsService/GetParentChildren"
)

// ParentServiceClient is a client for the buf.knit.crosstest.ParentService service.
type ParentServiceClient interface {
	ListParents(context.Context, *connect_go.Request[crosstest.ListParentsRequest]) (*connect_go.Response[crosstest.ListParentsResponse], error)
	GetParent(context.Context, *connect_go.Request[crosstest.GetParentRequest]) (*connect_go.Response[crosstest.Parent], error)
	CreateParent(context.Context, *connect_go.Request[crosstest.CreateParentRequest]) (*connect_go.Response[crosstest.Parent], error)
	UpdateParent(context.Context, *connect_go.Request[crosstest.UpdateParentRequest]) (*connect_go.Response[crosstest.Parent], error)
	DeleteParent(context.Context, *connect_go.Request[crosstest.DeleteParentRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewParentServiceClient constructs a client for the buf.knit.crosstest.ParentService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewParentServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ParentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &parentServiceClient{
		listParents: connect_go.NewClient[crosstest.ListParentsRequest, crosstest.ListParentsResponse](
			httpClient,
			baseURL+ParentServiceListParentsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getParent: connect_go.NewClient[crosstest.GetParentRequest, crosstest.Parent](
			httpClient,
			baseURL+ParentServiceGetParentProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		createParent: connect_go.NewClient[crosstest.CreateParentRequest, crosstest.Parent](
			httpClient,
			baseURL+ParentServiceCreateParentProcedure,
			opts...,
		),
		updateParent: connect_go.NewClient[crosstest.UpdateParentRequest, crosstest.Parent](
			httpClient,
			baseURL+ParentServiceUpdateParentProcedure,
			opts...,
		),
		deleteParent: connect_go.NewClient[crosstest.DeleteParentRequest, emptypb.Empty](
			httpClient,
			baseURL+ParentServiceDeleteParentProcedure,
			opts...,
		),
	}
}

// parentServiceClient implements ParentServiceClient.
type parentServiceClient struct {
	listParents  *connect_go.Client[crosstest.ListParentsRequest, crosstest.ListParentsResponse]
	getParent    *connect_go.Client[crosstest.GetParentRequest, crosstest.Parent]
	createParent *connect_go.Client[crosstest.CreateParentRequest, crosstest.Parent]
	updateParent *connect_go.Client[crosstest.UpdateParentRequest, crosstest.Parent]
	deleteParent *connect_go.Client[crosstest.DeleteParentRequest, emptypb.Empty]
}

// ListParents calls buf.knit.crosstest.ParentService.ListParents.
func (c *parentServiceClient) ListParents(ctx context.Context, req *connect_go.Request[crosstest.ListParentsRequest]) (*connect_go.Response[crosstest.ListParentsResponse], error) {
	return c.listParents.CallUnary(ctx, req)
}

// GetParent calls buf.knit.crosstest.ParentService.GetParent.
func (c *parentServiceClient) GetParent(ctx context.Context, req *connect_go.Request[crosstest.GetParentRequest]) (*connect_go.Response[crosstest.Parent], error) {
	return c.getParent.CallUnary(ctx, req)
}

// CreateParent calls buf.knit.crosstest.ParentService.CreateParent.
func (c *parentServiceClient) CreateParent(ctx context.Context, req *connect_go.Request[crosstest.CreateParentRequest]) (*connect_go.Response[crosstest.Parent], error) {
	return c.createParent.CallUnary(ctx, req)
}

// UpdateParent calls buf.knit.crosstest.ParentService.UpdateParent.
func (c *parentServiceClient) UpdateParent(ctx context.Context, req *connect_go.Request[crosstest.UpdateParentRequest]) (*connect_go.Response[crosstest.Parent], error) {
	return c.updateParent.CallUnary(ctx, req)
}

// DeleteParent calls buf.knit.crosstest.ParentService.DeleteParent.
func (c *parentServiceClient) DeleteParent(ctx context.Context, req *connect_go.Request[crosstest.DeleteParentRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteParent.CallUnary(ctx, req)
}

// ParentServiceHandler is an implementation of the buf.knit.crosstest.ParentService service.
type ParentServiceHandler interface {
	ListParents(context.Context, *connect_go.Request[crosstest.ListParentsRequest]) (*connect_go.Response[crosstest.ListParentsResponse], error)
	GetParent(context.Context, *connect_go.Request[crosstest.GetParentRequest]) (*connect_go.Response[crosstest.Parent], error)
	CreateParent(context.Context, *connect_go.Request[crosstest.CreateParentRequest]) (*connect_go.Response[crosstest.Parent], error)
	UpdateParent(context.Context, *connect_go.Request[crosstest.UpdateParentRequest]) (*connect_go.Response[crosstest.Parent], error)
	DeleteParent(context.Context, *connect_go.Request[crosstest.DeleteParentRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewParentServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewParentServiceHandler(svc ParentServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ParentServiceListParentsProcedure, connect_go.NewUnaryHandler(
		ParentServiceListParentsProcedure,
		svc.ListParents,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(ParentServiceGetParentProcedure, connect_go.NewUnaryHandler(
		ParentServiceGetParentProcedure,
		svc.GetParent,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(ParentServiceCreateParentProcedure, connect_go.NewUnaryHandler(
		ParentServiceCreateParentProcedure,
		svc.CreateParent,
		opts...,
	))
	mux.Handle(ParentServiceUpdateParentProcedure, connect_go.NewUnaryHandler(
		ParentServiceUpdateParentProcedure,
		svc.UpdateParent,
		opts...,
	))
	mux.Handle(ParentServiceDeleteParentProcedure, connect_go.NewUnaryHandler(
		ParentServiceDeleteParentProcedure,
		svc.DeleteParent,
		opts...,
	))
	return "/buf.knit.crosstest.ParentService/", mux
}

// UnimplementedParentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedParentServiceHandler struct{}

func (UnimplementedParentServiceHandler) ListParents(context.Context, *connect_go.Request[crosstest.ListParentsRequest]) (*connect_go.Response[crosstest.ListParentsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ParentService.ListParents is not implemented"))
}

func (UnimplementedParentServiceHandler) GetParent(context.Context, *connect_go.Request[crosstest.GetParentRequest]) (*connect_go.Response[crosstest.Parent], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ParentService.GetParent is not implemented"))
}

func (UnimplementedParentServiceHandler) CreateParent(context.Context, *connect_go.Request[crosstest.CreateParentRequest]) (*connect_go.Response[crosstest.Parent], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ParentService.CreateParent is not implemented"))
}

func (UnimplementedParentServiceHandler) UpdateParent(context.Context, *connect_go.Request[crosstest.UpdateParentRequest]) (*connect_go.Response[crosstest.Parent], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ParentService.UpdateParent is not implemented"))
}

func (UnimplementedParentServiceHandler) DeleteParent(context.Context, *connect_go.Request[crosstest.DeleteParentRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ParentService.DeleteParent is not implemented"))
}

// ChildServiceClient is a client for the buf.knit.crosstest.ChildService service.
type ChildServiceClient interface {
	ListChildren(context.Context, *connect_go.Request[crosstest.ListChildrenRequest]) (*connect_go.Response[crosstest.ListChildrenResponse], error)
	GetChild(context.Context, *connect_go.Request[crosstest.GetChildRequest]) (*connect_go.Response[crosstest.Child], error)
	CreateChild(context.Context, *connect_go.Request[crosstest.CreateChildRequest]) (*connect_go.Response[crosstest.Child], error)
	UpdateChild(context.Context, *connect_go.Request[crosstest.UpdateChildRequest]) (*connect_go.Response[crosstest.Child], error)
	DeleteChild(context.Context, *connect_go.Request[crosstest.DeleteChildRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewChildServiceClient constructs a client for the buf.knit.crosstest.ChildService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChildServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ChildServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &childServiceClient{
		listChildren: connect_go.NewClient[crosstest.ListChildrenRequest, crosstest.ListChildrenResponse](
			httpClient,
			baseURL+ChildServiceListChildrenProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getChild: connect_go.NewClient[crosstest.GetChildRequest, crosstest.Child](
			httpClient,
			baseURL+ChildServiceGetChildProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		createChild: connect_go.NewClient[crosstest.CreateChildRequest, crosstest.Child](
			httpClient,
			baseURL+ChildServiceCreateChildProcedure,
			opts...,
		),
		updateChild: connect_go.NewClient[crosstest.UpdateChildRequest, crosstest.Child](
			httpClient,
			baseURL+ChildServiceUpdateChildProcedure,
			opts...,
		),
		deleteChild: connect_go.NewClient[crosstest.DeleteChildRequest, emptypb.Empty](
			httpClient,
			baseURL+ChildServiceDeleteChildProcedure,
			opts...,
		),
	}
}

// childServiceClient implements ChildServiceClient.
type childServiceClient struct {
	listChildren *connect_go.Client[crosstest.ListChildrenRequest, crosstest.ListChildrenResponse]
	getChild     *connect_go.Client[crosstest.GetChildRequest, crosstest.Child]
	createChild  *connect_go.Client[crosstest.CreateChildRequest, crosstest.Child]
	updateChild  *connect_go.Client[crosstest.UpdateChildRequest, crosstest.Child]
	deleteChild  *connect_go.Client[crosstest.DeleteChildRequest, emptypb.Empty]
}

// ListChildren calls buf.knit.crosstest.ChildService.ListChildren.
func (c *childServiceClient) ListChildren(ctx context.Context, req *connect_go.Request[crosstest.ListChildrenRequest]) (*connect_go.Response[crosstest.ListChildrenResponse], error) {
	return c.listChildren.CallUnary(ctx, req)
}

// GetChild calls buf.knit.crosstest.ChildService.GetChild.
func (c *childServiceClient) GetChild(ctx context.Context, req *connect_go.Request[crosstest.GetChildRequest]) (*connect_go.Response[crosstest.Child], error) {
	return c.getChild.CallUnary(ctx, req)
}

// CreateChild calls buf.knit.crosstest.ChildService.CreateChild.
func (c *childServiceClient) CreateChild(ctx context.Context, req *connect_go.Request[crosstest.CreateChildRequest]) (*connect_go.Response[crosstest.Child], error) {
	return c.createChild.CallUnary(ctx, req)
}

// UpdateChild calls buf.knit.crosstest.ChildService.UpdateChild.
func (c *childServiceClient) UpdateChild(ctx context.Context, req *connect_go.Request[crosstest.UpdateChildRequest]) (*connect_go.Response[crosstest.Child], error) {
	return c.updateChild.CallUnary(ctx, req)
}

// DeleteChild calls buf.knit.crosstest.ChildService.DeleteChild.
func (c *childServiceClient) DeleteChild(ctx context.Context, req *connect_go.Request[crosstest.DeleteChildRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteChild.CallUnary(ctx, req)
}

// ChildServiceHandler is an implementation of the buf.knit.crosstest.ChildService service.
type ChildServiceHandler interface {
	ListChildren(context.Context, *connect_go.Request[crosstest.ListChildrenRequest]) (*connect_go.Response[crosstest.ListChildrenResponse], error)
	GetChild(context.Context, *connect_go.Request[crosstest.GetChildRequest]) (*connect_go.Response[crosstest.Child], error)
	CreateChild(context.Context, *connect_go.Request[crosstest.CreateChildRequest]) (*connect_go.Response[crosstest.Child], error)
	UpdateChild(context.Context, *connect_go.Request[crosstest.UpdateChildRequest]) (*connect_go.Response[crosstest.Child], error)
	DeleteChild(context.Context, *connect_go.Request[crosstest.DeleteChildRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewChildServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChildServiceHandler(svc ChildServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ChildServiceListChildrenProcedure, connect_go.NewUnaryHandler(
		ChildServiceListChildrenProcedure,
		svc.ListChildren,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(ChildServiceGetChildProcedure, connect_go.NewUnaryHandler(
		ChildServiceGetChildProcedure,
		svc.GetChild,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(ChildServiceCreateChildProcedure, connect_go.NewUnaryHandler(
		ChildServiceCreateChildProcedure,
		svc.CreateChild,
		opts...,
	))
	mux.Handle(ChildServiceUpdateChildProcedure, connect_go.NewUnaryHandler(
		ChildServiceUpdateChildProcedure,
		svc.UpdateChild,
		opts...,
	))
	mux.Handle(ChildServiceDeleteChildProcedure, connect_go.NewUnaryHandler(
		ChildServiceDeleteChildProcedure,
		svc.DeleteChild,
		opts...,
	))
	return "/buf.knit.crosstest.ChildService/", mux
}

// UnimplementedChildServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChildServiceHandler struct{}

func (UnimplementedChildServiceHandler) ListChildren(context.Context, *connect_go.Request[crosstest.ListChildrenRequest]) (*connect_go.Response[crosstest.ListChildrenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ChildService.ListChildren is not implemented"))
}

func (UnimplementedChildServiceHandler) GetChild(context.Context, *connect_go.Request[crosstest.GetChildRequest]) (*connect_go.Response[crosstest.Child], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ChildService.GetChild is not implemented"))
}

func (UnimplementedChildServiceHandler) CreateChild(context.Context, *connect_go.Request[crosstest.CreateChildRequest]) (*connect_go.Response[crosstest.Child], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ChildService.CreateChild is not implemented"))
}

func (UnimplementedChildServiceHandler) UpdateChild(context.Context, *connect_go.Request[crosstest.UpdateChildRequest]) (*connect_go.Response[crosstest.Child], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ChildService.UpdateChild is not implemented"))
}

func (UnimplementedChildServiceHandler) DeleteChild(context.Context, *connect_go.Request[crosstest.DeleteChildRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.ChildService.DeleteChild is not implemented"))
}

// RelationsServiceClient is a client for the buf.knit.crosstest.RelationsService service.
type RelationsServiceClient interface {
	GetChildParent(context.Context, *connect_go.Request[crosstest.GetChildParentRequest]) (*connect_go.Response[crosstest.GetChildParentResponse], error)
	GetParentChildren(context.Context, *connect_go.Request[crosstest.GetParentChildrenRequest]) (*connect_go.Response[crosstest.GetParentChildrenResponse], error)
}

// NewRelationsServiceClient constructs a client for the buf.knit.crosstest.RelationsService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRelationsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RelationsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &relationsServiceClient{
		getChildParent: connect_go.NewClient[crosstest.GetChildParentRequest, crosstest.GetChildParentResponse](
			httpClient,
			baseURL+RelationsServiceGetChildParentProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getParentChildren: connect_go.NewClient[crosstest.GetParentChildrenRequest, crosstest.GetParentChildrenResponse](
			httpClient,
			baseURL+RelationsServiceGetParentChildrenProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// relationsServiceClient implements RelationsServiceClient.
type relationsServiceClient struct {
	getChildParent    *connect_go.Client[crosstest.GetChildParentRequest, crosstest.GetChildParentResponse]
	getParentChildren *connect_go.Client[crosstest.GetParentChildrenRequest, crosstest.GetParentChildrenResponse]
}

// GetChildParent calls buf.knit.crosstest.RelationsService.GetChildParent.
func (c *relationsServiceClient) GetChildParent(ctx context.Context, req *connect_go.Request[crosstest.GetChildParentRequest]) (*connect_go.Response[crosstest.GetChildParentResponse], error) {
	return c.getChildParent.CallUnary(ctx, req)
}

// GetParentChildren calls buf.knit.crosstest.RelationsService.GetParentChildren.
func (c *relationsServiceClient) GetParentChildren(ctx context.Context, req *connect_go.Request[crosstest.GetParentChildrenRequest]) (*connect_go.Response[crosstest.GetParentChildrenResponse], error) {
	return c.getParentChildren.CallUnary(ctx, req)
}

// RelationsServiceHandler is an implementation of the buf.knit.crosstest.RelationsService service.
type RelationsServiceHandler interface {
	GetChildParent(context.Context, *connect_go.Request[crosstest.GetChildParentRequest]) (*connect_go.Response[crosstest.GetChildParentResponse], error)
	GetParentChildren(context.Context, *connect_go.Request[crosstest.GetParentChildrenRequest]) (*connect_go.Response[crosstest.GetParentChildrenResponse], error)
}

// NewRelationsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRelationsServiceHandler(svc RelationsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(RelationsServiceGetChildParentProcedure, connect_go.NewUnaryHandler(
		RelationsServiceGetChildParentProcedure,
		svc.GetChildParent,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(RelationsServiceGetParentChildrenProcedure, connect_go.NewUnaryHandler(
		RelationsServiceGetParentChildrenProcedure,
		svc.GetParentChildren,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	return "/buf.knit.crosstest.RelationsService/", mux
}

// UnimplementedRelationsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRelationsServiceHandler struct{}

func (UnimplementedRelationsServiceHandler) GetChildParent(context.Context, *connect_go.Request[crosstest.GetChildParentRequest]) (*connect_go.Response[crosstest.GetChildParentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.RelationsService.GetChildParent is not implemented"))
}

func (UnimplementedRelationsServiceHandler) GetParentChildren(context.Context, *connect_go.Request[crosstest.GetParentChildrenRequest]) (*connect_go.Response[crosstest.GetParentChildrenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.knit.crosstest.RelationsService.GetParentChildren is not implemented"))
}
