package handler

import (
	"context"
	"errors"
	errors2 "go-micro.dev/v4/errors"
	"merchant/model"
	pb "merchant/proto"
	"merchant/repository"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func NewMerchantHandler(db *gorm.DB, merchantRepository repository.MerchantRepository) MerchantServiceHandler {
	return MerchantServiceHandler{
		DB:                 db,
		MerchantRepository: merchantRepository,
	}
}

type MerchantServiceHandler struct {
	DB                 *gorm.DB
	MerchantRepository repository.MerchantRepository
}

func (service *MerchantServiceHandler) FindOneByID(ctx context.Context, id *pb.MerchantID, merchant *pb.Merchant) error {
	merchantResp, err := service.MerchantRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	*merchant = pb.Merchant{
		ID:        int64(merchantResp.ID),
		Name:      merchantResp.Name,
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) FindOneByEmail(ctx context.Context, email *pb.MerchantEmail, merchant *pb.Merchant) error {
	merchantResp, err := service.MerchantRepository.FindOneByEmail(ctx, service.DB, email.Email)
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	*merchant = pb.Merchant{
		ID:        int64(merchantResp.ID),
		Name:      merchantResp.Name,
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) FindAll(ctx context.Context, empty *emptypb.Empty, stream pb.MerchantService_FindAllStream) error {
	_ = empty
	merchants, err := service.MerchantRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, merchant := range merchants {
		merchantsResp := pb.Merchant{
			ID:        int64(merchant.ID),
			Name:      merchant.Name,
			Email:     merchant.Email,
			Password:  merchant.Password,
			CreatedAt: timestamppb.New(merchant.CreatedAt),
			UpdatedAt: timestamppb.New(merchant.UpdatedAt),
		}
		stream.Send(&merchantsResp)
	}
	return nil
}

func (service *MerchantServiceHandler) Create(ctx context.Context, req *pb.MerchantCreateReq, merchant *pb.Merchant) error {
	merchantResp, err := service.MerchantRepository.Create(ctx, service.DB, model.Merchant{
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
	*merchant = pb.Merchant{
		ID:        int64(merchantResp.ID),
		Name:      merchantResp.Name,
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) Update(ctx context.Context, req *pb.MerchantUpdateReq, merchant *pb.Merchant) error {
	findMerchant, err := service.MerchantRepository.FindOneByID(ctx, service.DB, int(req.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	merchantResp, err := service.MerchantRepository.Update(ctx, service.DB, model.Merchant{
		ID:        int(req.ID),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: findMerchant.CreatedAt,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*merchant = pb.Merchant{
		ID:        int64(merchantResp.ID),
		Name:      merchantResp.Name,
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) Delete(ctx context.Context, id *pb.MerchantID, empty *emptypb.Empty) error {
	_, err := service.MerchantRepository.FindOneByID(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors2.NotFound("", err.Error())
		} else {
			return errors2.BadRequest("", err.Error())
		}
	}
	err = service.MerchantRepository.Delete(ctx, service.DB, int(id.ID))
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	return nil
}
