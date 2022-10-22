// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/customer.proto

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

// Api Endpoints for CustomerService service

func NewCustomerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CustomerService service

type CustomerService interface {
	FindOneByID(ctx context.Context, in *CustomerID, opts ...client.CallOption) (*Customer, error)
	FindOneByEmail(ctx context.Context, in *CustomerEmail, opts ...client.CallOption) (*Customer, error)
	FindAll(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (CustomerService_FindAllService, error)
	Create(ctx context.Context, in *CustomerCreateReq, opts ...client.CallOption) (*Customer, error)
	Update(ctx context.Context, in *CustomerUpdateReq, opts ...client.CallOption) (*Customer, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*emptypb.Empty, error)
}

type customerService struct {
	c    client.Client
	name string
}

func NewCustomerService(name string, c client.Client) CustomerService {
	return &customerService{
		c:    c,
		name: name,
	}
}

func (c *customerService) FindOneByID(ctx context.Context, in *CustomerID, opts ...client.CallOption) (*Customer, error) {
	req := c.c.NewRequest(c.name, "CustomerService.FindOneByID", in)
	out := new(Customer)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) FindOneByEmail(ctx context.Context, in *CustomerEmail, opts ...client.CallOption) (*Customer, error) {
	req := c.c.NewRequest(c.name, "CustomerService.FindOneByEmail", in)
	out := new(Customer)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) FindAll(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (CustomerService_FindAllService, error) {
	req := c.c.NewRequest(c.name, "CustomerService.FindAll", &emptypb.Empty{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &customerServiceFindAll{stream}, nil
}

type CustomerService_FindAllService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*Customer, error)
}

type customerServiceFindAll struct {
	stream client.Stream
}

func (x *customerServiceFindAll) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *customerServiceFindAll) Close() error {
	return x.stream.Close()
}

func (x *customerServiceFindAll) Context() context.Context {
	return x.stream.Context()
}

func (x *customerServiceFindAll) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *customerServiceFindAll) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *customerServiceFindAll) Recv() (*Customer, error) {
	m := new(Customer)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerService) Create(ctx context.Context, in *CustomerCreateReq, opts ...client.CallOption) (*Customer, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Create", in)
	out := new(Customer)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Update(ctx context.Context, in *CustomerUpdateReq, opts ...client.CallOption) (*Customer, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Update", in)
	out := new(Customer)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "CustomerService.Delete", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerService service

type CustomerServiceHandler interface {
	FindOneByID(context.Context, *CustomerID, *Customer) error
	FindOneByEmail(context.Context, *CustomerEmail, *Customer) error
	FindAll(context.Context, *emptypb.Empty, CustomerService_FindAllStream) error
	Create(context.Context, *CustomerCreateReq, *Customer) error
	Update(context.Context, *CustomerUpdateReq, *Customer) error
	Delete(context.Context, *DeleteReq, *emptypb.Empty) error
}

func RegisterCustomerServiceHandler(s server.Server, hdlr CustomerServiceHandler, opts ...server.HandlerOption) error {
	type customerService interface {
		FindOneByID(ctx context.Context, in *CustomerID, out *Customer) error
		FindOneByEmail(ctx context.Context, in *CustomerEmail, out *Customer) error
		FindAll(ctx context.Context, stream server.Stream) error
		Create(ctx context.Context, in *CustomerCreateReq, out *Customer) error
		Update(ctx context.Context, in *CustomerUpdateReq, out *Customer) error
		Delete(ctx context.Context, in *DeleteReq, out *emptypb.Empty) error
	}
	type CustomerService struct {
		customerService
	}
	h := &customerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CustomerService{h}, opts...))
}

type customerServiceHandler struct {
	CustomerServiceHandler
}

func (h *customerServiceHandler) FindOneByID(ctx context.Context, in *CustomerID, out *Customer) error {
	return h.CustomerServiceHandler.FindOneByID(ctx, in, out)
}

func (h *customerServiceHandler) FindOneByEmail(ctx context.Context, in *CustomerEmail, out *Customer) error {
	return h.CustomerServiceHandler.FindOneByEmail(ctx, in, out)
}

func (h *customerServiceHandler) FindAll(ctx context.Context, stream server.Stream) error {
	m := new(emptypb.Empty)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.CustomerServiceHandler.FindAll(ctx, m, &customerServiceFindAllStream{stream})
}

type CustomerService_FindAllStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Customer) error
}

type customerServiceFindAllStream struct {
	stream server.Stream
}

func (x *customerServiceFindAllStream) Close() error {
	return x.stream.Close()
}

func (x *customerServiceFindAllStream) Context() context.Context {
	return x.stream.Context()
}

func (x *customerServiceFindAllStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *customerServiceFindAllStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *customerServiceFindAllStream) Send(m *Customer) error {
	return x.stream.Send(m)
}

func (h *customerServiceHandler) Create(ctx context.Context, in *CustomerCreateReq, out *Customer) error {
	return h.CustomerServiceHandler.Create(ctx, in, out)
}

func (h *customerServiceHandler) Update(ctx context.Context, in *CustomerUpdateReq, out *Customer) error {
	return h.CustomerServiceHandler.Update(ctx, in, out)
}

func (h *customerServiceHandler) Delete(ctx context.Context, in *DeleteReq, out *emptypb.Empty) error {
	return h.CustomerServiceHandler.Delete(ctx, in, out)
}
