package handler

import (
	"context"
	"customer/model"
	pb "customer/proto"
	"customer/repository"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

func NewCustomerHandler(db *gorm.DB, customerRepository repository.CustomerRepository) CustomerServiceHandler {
	return CustomerServiceHandler{
		DB:                 db,
		CustomerRepository: customerRepository,
	}
}

type CustomerServiceHandler struct {
	DB                 *gorm.DB
	CustomerRepository repository.CustomerRepository
}

func (service *CustomerServiceHandler) FindOneByID(ctx context.Context, id *pb.CustomerID, customer *pb.Customer) error {
	customerResp, err := service.CustomerRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.InvalidArgument, err.Error())
	}
	*customer = pb.Customer{
		ID:        int64(customerResp.ID),
		Name:      customerResp.Name,
		Email:     customerResp.Email,
		Password:  customerResp.Password,
		CreatedAt: timestamppb.New(customerResp.CreatedAt),
		UpdatedAt: timestamppb.New(customerResp.UpdatedAt),
	}
	return nil
}

func (service *CustomerServiceHandler) FindOneByEmail(ctx context.Context, email *pb.CustomerEmail, customer *pb.Customer) error {
	customerResp, err := service.CustomerRepository.FindOneByEmail(ctx, service.DB, email.Email)
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.InvalidArgument, err.Error())
	}
	*customer = pb.Customer{
		ID:        int64(customerResp.ID),
		Name:      customerResp.Name,
		Email:     customerResp.Email,
		Password:  customerResp.Password,
		CreatedAt: timestamppb.New(customerResp.CreatedAt),
		UpdatedAt: timestamppb.New(customerResp.UpdatedAt),
	}
	return nil
}

func (service *CustomerServiceHandler) FindAll(ctx context.Context, empty *emptypb.Empty, stream pb.CustomerService_FindAllStream) error {
	_ = empty
	customers, err := service.CustomerRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}
	for _, customer := range customers {
		customersResp := pb.Customer{
			ID:        int64(customer.ID),
			Name:      customer.Name,
			Email:     customer.Email,
			Password:  customer.Password,
			CreatedAt: timestamppb.New(customer.CreatedAt),
			UpdatedAt: timestamppb.New(customer.UpdatedAt),
		}
		stream.Send(&customersResp)
	}
	return nil
}

func (service *CustomerServiceHandler) Create(ctx context.Context, req *pb.CustomerCreateReq, customer *pb.Customer) error {
	customerResp, err := service.CustomerRepository.Create(ctx, service.DB, model.Customer{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}
	*customer = pb.Customer{
		ID:        int64(customerResp.ID),
		Name:      customerResp.Name,
		Email:     customerResp.Email,
		Password:  customerResp.Password,
		CreatedAt: timestamppb.New(customerResp.CreatedAt),
		UpdatedAt: timestamppb.New(customerResp.UpdatedAt),
	}
	return nil
}

func (service *CustomerServiceHandler) Update(ctx context.Context, req *pb.CustomerUpdateReq, customer *pb.Customer) error {
	findCustomer, err := service.CustomerRepository.FindOneByID(ctx, service.DB, int(req.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.InvalidArgument, err.Error())
	}
	customerResp, err := service.CustomerRepository.Update(ctx, service.DB, model.Customer{
		ID:        int(req.ID),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: findCustomer.CreatedAt,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}
	*customer = pb.Customer{
		ID:        int64(customerResp.ID),
		Name:      customerResp.Name,
		Email:     customerResp.Email,
		Password:  customerResp.Password,
		CreatedAt: timestamppb.New(customerResp.CreatedAt),
		UpdatedAt: timestamppb.New(customerResp.UpdatedAt),
	}
	return nil
}

func (service *CustomerServiceHandler) Delete(ctx context.Context, id *pb.CustomerID, empty *emptypb.Empty) error {
	_, err := service.CustomerRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.InvalidArgument, err.Error())
	}
	err = service.CustomerRepository.Delete(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}