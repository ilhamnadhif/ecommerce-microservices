package service

import (
	"api-gateway/dto"
	pb "api-gateway/proto"
	"context"
	"github.com/labstack/echo/v4"
	"go-micro.dev/v4/errors"
	"io"
	"net/http"
)

type MerchantService interface {
	FindOneByID(ctx context.Context, merchantID int) (dto.MerchantResponse, error)
	FindAll(ctx context.Context) ([]dto.MerchantResponse, error)
	Create(ctx context.Context, request dto.MerchantCreateReq) (dto.MerchantResponse, error)
	Update(ctx context.Context, request dto.MerchantUpdateReq) (dto.MerchantResponse, error)
	Delete(ctx context.Context, merchantID int) error
}

func NewMerchantService(service pb.MerchantService) MerchantService {
	return &merchantServiceImpl{
		MerchantRPC: service,
	}
}

type merchantServiceImpl struct {
	MerchantRPC pb.MerchantService
}

func (service *merchantServiceImpl) FindOneByID(ctx context.Context, merchantID int) (dto.MerchantResponse, error) {
	merchant, err := service.MerchantRPC.FindOneByID(ctx, &pb.MerchantID{
		ID: int64(merchantID),
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.MerchantResponse{
		ID:        int(merchant.ID),
		Name:      merchant.Name,
		Email:     merchant.Email,
		Password:  merchant.Password,
		CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
	}, nil
}

func (service *merchantServiceImpl) FindAll(ctx context.Context) ([]dto.MerchantResponse, error) {
	merchantsResponse := make([]dto.MerchantResponse, 0)
	stream, err := service.MerchantRPC.FindAll(ctx, nil)
	if err != nil {
		e := errors.FromError(err)
		return nil, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		merchantsResponse = append(merchantsResponse, dto.MerchantResponse{
			ID:        int(msg.ID),
			Name:      msg.Name,
			Email:     msg.Email,
			Password:  msg.Password,
			CreatedAt: dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt: dto.DateTime(msg.UpdatedAt.AsTime()),
		})
	}
	return merchantsResponse, nil
}

func (service *merchantServiceImpl) Create(ctx context.Context, request dto.MerchantCreateReq) (dto.MerchantResponse, error) {
	merchant, err := service.MerchantRPC.Create(ctx, &pb.MerchantCreateReq{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.MerchantResponse{
		ID:        int(merchant.ID),
		Name:      merchant.Name,
		Email:     merchant.Email,
		Password:  merchant.Password,
		CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
	}, nil
}

func (service *merchantServiceImpl) Update(ctx context.Context, request dto.MerchantUpdateReq) (dto.MerchantResponse, error) {
	merchant, err := service.MerchantRPC.Update(ctx, &pb.MerchantUpdateReq{
		ID:       int64(request.ID),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.MerchantResponse{
		ID:        int(merchant.ID),
		Name:      merchant.Name,
		Email:     merchant.Email,
		Password:  merchant.Password,
		CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
	}, nil
}

func (service *merchantServiceImpl) Delete(ctx context.Context, merchantID int) error {
	_, err := service.MerchantRPC.Delete(ctx, &pb.MerchantID{ID: int64(merchantID)})
	if err != nil {
		e := errors.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return nil
}
