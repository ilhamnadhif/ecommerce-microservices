package handler

import (
	"context"
	"customer/model"
	pb "customer/proto"
	"customer/repository"
	"errors"
	"time"

	errors2 "go-micro.dev/v4/errors"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
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
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
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
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
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

func (service *CustomerServiceHandler) FindAll(ctx context.Context, _ *emptypb.Empty, stream pb.CustomerService_FindAllStream) error {
	customers, err := service.CustomerRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
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
		return errors2.BadRequest("", err.Error())
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
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
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
		return errors2.BadRequest("", err.Error())
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

func (service *CustomerServiceHandler) Delete(ctx context.Context, id *pb.CustomerID, _ *emptypb.Empty) error {
	_, err := service.CustomerRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	err = service.CustomerRepository.Delete(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	return nil
}
