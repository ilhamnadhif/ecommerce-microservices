package handler

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	errors2 "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"product/model"
	"product/proto"
	"product/repository"
	"time"
)

func NewProductHandler(db *gorm.DB, productRepository repository.ProductRepository) ProductServiceHandler {
	return ProductServiceHandler{
		DB:                db,
		ProductRepository: productRepository,
	}
}

type ProductServiceHandler struct {
	DB                *gorm.DB
	ProductRepository repository.ProductRepository
}

func (service *ProductServiceHandler) FindOneByID(ctx context.Context, id *proto.ProductID, product *proto.Product) error {
	productResp, err := service.ProductRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	*product = proto.Product{
		ID:          int64(productResp.ID),
		Name:        productResp.Name,
		Description: productResp.Description,
		Price:       int64(productResp.Price),
		CreatedAt:   timestamppb.New(productResp.CreatedAt),
		UpdatedAt:   timestamppb.New(productResp.UpdatedAt),
	}
	return nil
}

func (service *ProductServiceHandler) FindAll(ctx context.Context, empty *emptypb.Empty, stream proto.ProductService_FindAllStream) error {
	_ = empty
	products, err := service.ProductRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, product := range products {
		productsResp := proto.Product{
			ID:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       int64(product.Price),
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		}
		stream.Send(&productsResp)
	}
	return nil
}

func (service *ProductServiceHandler) Create(ctx context.Context, req *proto.ProductCreateReq, product *proto.Product) error {
	productResp, err := service.ProductRepository.Create(ctx, service.DB, model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       int(req.Price),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*product = proto.Product{
		ID:          int64(productResp.ID),
		Name:        productResp.Name,
		Description: productResp.Description,
		Price:       int64(productResp.Price),
		CreatedAt:   timestamppb.New(productResp.CreatedAt),
		UpdatedAt:   timestamppb.New(productResp.UpdatedAt),
	}
	return nil
}

func (service *ProductServiceHandler) Update(ctx context.Context, req *proto.ProductUpdateReq, product *proto.Product) error {
	findProduct, err := service.ProductRepository.FindOneByID(ctx, service.DB, int(req.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	productResp, err := service.ProductRepository.Update(ctx, service.DB, model.Product{
		ID:          int(req.ID),
		Name:        req.Name,
		Description: req.Description,
		Price:       int(req.Price),
		CreatedAt:   findProduct.CreatedAt,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*product = proto.Product{
		ID:          int64(productResp.ID),
		Name:        productResp.Name,
		Description: productResp.Description,
		Price:       int64(productResp.Price),
		CreatedAt:   timestamppb.New(productResp.CreatedAt),
		UpdatedAt:   timestamppb.New(productResp.UpdatedAt),
	}
	return nil
}

func (service *ProductServiceHandler) Delete(ctx context.Context, id *proto.ProductID, empty *emptypb.Empty) error {
	_, err := service.ProductRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	err = service.ProductRepository.Delete(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	return nil
}
