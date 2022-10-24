package handler

import (
	"cart/model"
	pb "cart/proto"
	"cart/repository"
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	errors2 "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func NewCartHandler(db *gorm.DB, cartRepository repository.CartRepository) CartServiceHandler {
	return CartServiceHandler{
		DB:             db,
		CartRepository: cartRepository,
	}
}

type CartServiceHandler struct {
	DB             *gorm.DB
	CartRepository repository.CartRepository
}

func (service *CartServiceHandler) FindOneByID(ctx context.Context, id *pb.CartID, cart *pb.Cart) error {
	cartResp, err := service.CartRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	*cart = pb.Cart{
		ID:         int64(cartResp.ID),
		CustomerID: int64(cartResp.CustomerID),
		ProductID:  int64(cartResp.ProductID),
		Quantity:   int64(cartResp.Quantity),
		CreatedAt:  timestamppb.New(cartResp.CreatedAt),
		UpdatedAt:  timestamppb.New(cartResp.UpdatedAt),
	}
	return nil
}

func (service *CartServiceHandler) FindAll(ctx context.Context, _ *emptypb.Empty, stream pb.CartService_FindAllStream) error {
	carts, err := service.CartRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, cart := range carts {
		cartsResp := pb.Cart{
			ID:         int64(cart.ID),
			CustomerID: int64(cart.CustomerID),
			ProductID:  int64(cart.ProductID),
			Quantity:   int64(cart.Quantity),
			CreatedAt:  timestamppb.New(cart.CreatedAt),
			UpdatedAt:  timestamppb.New(cart.UpdatedAt),
		}
		stream.Send(&cartsResp)
	}
	return nil
}

func (service *CartServiceHandler) FindAllByCustomerID(ctx context.Context, id *pb.CustomerID, stream pb.CartService_FindAllByCustomerIDStream) error {
	carts, err := service.CartRepository.FindAllByCustomerID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, cart := range carts {
		cartsResp := pb.Cart{
			ID:         int64(cart.ID),
			CustomerID: int64(cart.CustomerID),
			ProductID:  int64(cart.ProductID),
			Quantity:   int64(cart.Quantity),
			CreatedAt:  timestamppb.New(cart.CreatedAt),
			UpdatedAt:  timestamppb.New(cart.UpdatedAt),
		}
		stream.Send(&cartsResp)
	}
	return nil
}

func (service *CartServiceHandler) Create(ctx context.Context, req *pb.CartCreateReq, cart *pb.Cart) error {
	cartResp, err := service.CartRepository.Create(ctx, service.DB, model.Cart{
		CustomerID: int(req.CustomerID),
		ProductID:  int(req.ProductID),
		Quantity:   int(req.Quantity),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*cart = pb.Cart{
		ID:         int64(cartResp.ID),
		CustomerID: int64(cartResp.CustomerID),
		ProductID:  int64(cartResp.ProductID),
		Quantity:   int64(cartResp.Quantity),
		CreatedAt:  timestamppb.New(cartResp.CreatedAt),
		UpdatedAt:  timestamppb.New(cartResp.UpdatedAt),
	}
	return nil
}

func (service *CartServiceHandler) Update(ctx context.Context, req *pb.CartUpdateReq, cart *pb.Cart) error {
	findCart, err := service.CartRepository.FindOneByID(ctx, service.DB, int(req.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	cartResp, err := service.CartRepository.Update(ctx, service.DB, model.Cart{
		ID:         int(req.ID),
		CustomerID: findCart.CustomerID,
		ProductID:  findCart.ProductID,
		Quantity:   int(req.Quantity),
		CreatedAt:  findCart.CreatedAt,
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*cart = pb.Cart{
		ID:         int64(cartResp.ID),
		CustomerID: int64(cartResp.CustomerID),
		ProductID:  int64(cartResp.ProductID),
		Quantity:   int64(cartResp.Quantity),
		CreatedAt:  timestamppb.New(cartResp.CreatedAt),
		UpdatedAt:  timestamppb.New(cartResp.UpdatedAt),
	}
	return nil
}

func (service *CartServiceHandler) Delete(ctx context.Context, id *pb.CartID, _ *emptypb.Empty) error {
	_, err := service.CartRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	err = service.CartRepository.Delete(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	return nil
}
