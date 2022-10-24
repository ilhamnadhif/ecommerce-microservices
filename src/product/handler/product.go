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
	pb "product/proto"
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

func (service *ProductServiceHandler) FindOneByID(ctx context.Context, id *pb.ProductID, product *pb.Product) error {
	productResp, err := service.ProductRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	*product = pb.Product{
		ID:          int64(productResp.ID),
		MerchantID:  int64(productResp.MerchantID),
		Name:        productResp.Name,
		Description: productResp.Description,
		Price:       int64(productResp.Price),
		CreatedAt:   timestamppb.New(productResp.CreatedAt),
		UpdatedAt:   timestamppb.New(productResp.UpdatedAt),
	}
	return nil
}

func (service *ProductServiceHandler) FindAll(ctx context.Context, _ *emptypb.Empty, stream pb.ProductService_FindAllStream) error {
	products, err := service.ProductRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, product := range products {
		productsResp := pb.Product{
			ID:          int64(product.ID),
			MerchantID:  int64(product.MerchantID),
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

func (service *ProductServiceHandler) FindAllByMerchantID(ctx context.Context, id *pb.MerchantID, stream pb.ProductService_FindAllByMerchantIDStream) error {
	products, err := service.ProductRepository.FindAllByMerchantID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, product := range products {
		productsResp := pb.Product{
			ID:          int64(product.ID),
			MerchantID:  int64(product.MerchantID),
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

func (service *ProductServiceHandler) Create(ctx context.Context, req *pb.ProductCreateReq, product *pb.Product) error {
	productResp, err := service.ProductRepository.Create(ctx, service.DB, model.Product{
		MerchantID:  int(req.MerchantID),
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
	*product = pb.Product{
		ID:          int64(productResp.ID),
		MerchantID:  int64(productResp.MerchantID),
		Name:        productResp.Name,
		Description: productResp.Description,
		Price:       int64(productResp.Price),
		CreatedAt:   timestamppb.New(productResp.CreatedAt),
		UpdatedAt:   timestamppb.New(productResp.UpdatedAt),
	}
	return nil
}

func (service *ProductServiceHandler) Update(ctx context.Context, req *pb.ProductUpdateReq, product *pb.Product) error {
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
		MerchantID:  findProduct.MerchantID,
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
	*product = pb.Product{
		ID:          int64(productResp.ID),
		MerchantID:  int64(productResp.MerchantID),
		Name:        productResp.Name,
		Description: productResp.Description,
		Price:       int64(productResp.Price),
		CreatedAt:   timestamppb.New(productResp.CreatedAt),
		UpdatedAt:   timestamppb.New(productResp.UpdatedAt),
	}
	return nil
}

func (service *ProductServiceHandler) Delete(ctx context.Context, id *pb.ProductID, _ *emptypb.Empty) error {
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
