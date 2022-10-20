// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ProductService service

func NewProductServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductService service

type ProductService interface {
	FindOneByID(ctx context.Context, in *ProductID, opts ...client.CallOption) (*Product, error)
	FindAll(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (ProductService_FindAllService, error)
	Create(ctx context.Context, in *ProductCreateReq, opts ...client.CallOption) (*Product, error)
	Update(ctx context.Context, in *ProductUpdateReq, opts ...client.CallOption) (*Product, error)
	Delete(ctx context.Context, in *ProductID, opts ...client.CallOption) (*emptypb.Empty, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) FindOneByID(ctx context.Context, in *ProductID, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductService.FindOneByID", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindAll(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (ProductService_FindAllService, error) {
	req := c.c.NewRequest(c.name, "ProductService.FindAll", &emptypb.Empty{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &productServiceFindAll{stream}, nil
}

type ProductService_FindAllService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*Product, error)
}

type productServiceFindAll struct {
	stream client.Stream
}

func (x *productServiceFindAll) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *productServiceFindAll) Close() error {
	return x.stream.Close()
}

func (x *productServiceFindAll) Context() context.Context {
	return x.stream.Context()
}

func (x *productServiceFindAll) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *productServiceFindAll) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *productServiceFindAll) Recv() (*Product, error) {
	m := new(Product)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productService) Create(ctx context.Context, in *ProductCreateReq, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductService.Create", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) Update(ctx context.Context, in *ProductUpdateReq, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductService.Update", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) Delete(ctx context.Context, in *ProductID, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "ProductService.Delete", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	FindOneByID(context.Context, *ProductID, *Product) error
	FindAll(context.Context, *emptypb.Empty, ProductService_FindAllStream) error
	Create(context.Context, *ProductCreateReq, *Product) error
	Update(context.Context, *ProductUpdateReq, *Product) error
	Delete(context.Context, *ProductID, *emptypb.Empty) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		FindOneByID(ctx context.Context, in *ProductID, out *Product) error
		FindAll(ctx context.Context, stream server.Stream) error
		Create(ctx context.Context, in *ProductCreateReq, out *Product) error
		Update(ctx context.Context, in *ProductUpdateReq, out *Product) error
		Delete(ctx context.Context, in *ProductID, out *emptypb.Empty) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) FindOneByID(ctx context.Context, in *ProductID, out *Product) error {
	return h.ProductServiceHandler.FindOneByID(ctx, in, out)
}

func (h *productServiceHandler) FindAll(ctx context.Context, stream server.Stream) error {
	m := new(emptypb.Empty)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ProductServiceHandler.FindAll(ctx, m, &productServiceFindAllStream{stream})
}

type ProductService_FindAllStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Product) error
}

type productServiceFindAllStream struct {
	stream server.Stream
}

func (x *productServiceFindAllStream) Close() error {
	return x.stream.Close()
}

func (x *productServiceFindAllStream) Context() context.Context {
	return x.stream.Context()
}

func (x *productServiceFindAllStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *productServiceFindAllStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *productServiceFindAllStream) Send(m *Product) error {
	return x.stream.Send(m)
}

func (h *productServiceHandler) Create(ctx context.Context, in *ProductCreateReq, out *Product) error {
	return h.ProductServiceHandler.Create(ctx, in, out)
}

func (h *productServiceHandler) Update(ctx context.Context, in *ProductUpdateReq, out *Product) error {
	return h.ProductServiceHandler.Update(ctx, in, out)
}

func (h *productServiceHandler) Delete(ctx context.Context, in *ProductID, out *emptypb.Empty) error {
	return h.ProductServiceHandler.Delete(ctx, in, out)
}
