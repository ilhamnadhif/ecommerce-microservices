package handler

import (
	"context"
	"errors"
	errors2 "go-micro.dev/v4/errors"
	"merchant/helper"
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
	merchantResp, err := service.MerchantRepository.FindOneByID(ctx, service.DB, int(id.GetID()))
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
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		Name:      merchantResp.Name,
		Slug:      merchantResp.Slug,
		Balance:   int64(merchantResp.Balance),
		ImageID:   int64(merchantResp.ImageID),
		Image:     merchantResp.Image,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) FindOneByEmail(ctx context.Context, email *pb.MerchantEmail, merchant *pb.Merchant) error {
	merchantResp, err := service.MerchantRepository.FindOneByEmail(ctx, service.DB, email.GetEmail())
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
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		Name:      merchantResp.Name,
		Slug:      merchantResp.Slug,
		Balance:   int64(merchantResp.Balance),
		ImageID:   int64(merchantResp.ImageID),
		Image:     merchantResp.Image,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) FindOneBySlug(ctx context.Context, slug *pb.MerchantSlug, merchant *pb.Merchant) error {
	merchantResp, err := service.MerchantRepository.FindOneByEmail(ctx, service.DB, slug.GetSlug())
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
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		Name:      merchantResp.Name,
		Slug:      merchantResp.Slug,
		Balance:   int64(merchantResp.Balance),
		ImageID:   int64(merchantResp.ImageID),
		Image:     merchantResp.Image,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) FindAll(ctx context.Context, _ *emptypb.Empty, stream pb.MerchantService_FindAllStream) error {
	merchants, err := service.MerchantRepository.FindAll(ctx, service.DB)
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	for _, merchant := range merchants {
		merchantsResp := pb.Merchant{
			ID:        int64(merchant.ID),
			Email:     merchant.Email,
			Password:  merchant.Password,
			Name:      merchant.Name,
			Slug:      merchant.Slug,
			Balance:   int64(merchant.Balance),
			ImageID:   int64(merchant.ImageID),
			Image:     merchant.Image,
			CreatedAt: timestamppb.New(merchant.CreatedAt),
			UpdatedAt: timestamppb.New(merchant.UpdatedAt),
		}
		stream.Send(&merchantsResp)
	}
	return nil
}

func (service *MerchantServiceHandler) Create(ctx context.Context, req *pb.MerchantCreateReq, merchant *pb.Merchant) error {
	newPassword, err := helper.HashPassword(req.GetPassword())
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	merchantResp, err := service.MerchantRepository.Create(ctx, service.DB, model.Merchant{
		Email:     req.GetEmail(),
		Password:  newPassword,
		Name:      req.Name,
		Slug:      helper.GenerateSlug(req.GetName()),
		Balance:   0,
		ImageID:   int(req.GetImageID()),
		Image:     req.GetImage(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*merchant = pb.Merchant{
		ID:        int64(merchantResp.ID),
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		Name:      merchantResp.Name,
		Slug:      merchantResp.Slug,
		Balance:   int64(merchantResp.Balance),
		ImageID:   int64(merchantResp.ImageID),
		Image:     merchantResp.Image,
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
		ID:        int(req.GetID()),
		Email:     findMerchant.Email,
		Password:  findMerchant.Password,
		Name:      req.GetName(),
		Slug:      helper.GenerateSlug(req.GetName()),
		Balance:   int(req.GetBalance()),
		ImageID:   int(req.GetImageID()),
		Image:     req.GetImage(),
		CreatedAt: findMerchant.CreatedAt,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logrus.Error(err.Error())
		return errors2.BadRequest("", err.Error())
	}
	*merchant = pb.Merchant{
		ID:        int64(merchantResp.ID),
		Email:     merchantResp.Email,
		Password:  merchantResp.Password,
		Name:      merchantResp.Name,
		Slug:      merchantResp.Slug,
		Balance:   int64(merchantResp.Balance),
		ImageID:   int64(merchantResp.ImageID),
		Image:     merchantResp.Image,
		CreatedAt: timestamppb.New(merchantResp.CreatedAt),
		UpdatedAt: timestamppb.New(merchantResp.UpdatedAt),
	}
	return nil
}

func (service *MerchantServiceHandler) Delete(ctx context.Context, id *pb.MerchantID, _ *emptypb.Empty) error {
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
