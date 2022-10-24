package handler

import (
	"api-gateway/dto"
	"api-gateway/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func NewMerchantHandler(merchantService service.MerchantService) merchantHandler {
	return merchantHandler{
		MerchantService: merchantService,
	}
}

type merchantHandler struct {
	MerchantService service.MerchantService
}

func (handler *merchantHandler) FindOneByID(c echo.Context) error {
	ctx := c.Request().Context()
	merchantID, err := strconv.Atoi(c.Param("merchantID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	merchant, err := handler.MerchantService.FindOneByID(ctx, merchantID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchant))
}

func (handler *merchantHandler) FindOneByCommon(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	if claims.Role != dto.MERCHANT_ROLE {
		return echo.NewHTTPError(http.StatusForbidden, "customer: access denied for this role")
	}
	merchant, err := handler.MerchantService.FindOneByID(ctx, claims.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchant))
}

func (handler *merchantHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	merchants, err := handler.MerchantService.FindAll(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchants))
}

func (handler *merchantHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.MerchantCreateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	merchant, err := handler.MerchantService.Create(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchant))
}
func (handler *merchantHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	merchantID, err := strconv.Atoi(c.Param("merchantID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var req dto.MerchantUpdateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = merchantID
	req.QueryData.ID = claims.ID
	req.QueryData.Role = claims.Role
	merchant, err := handler.MerchantService.Update(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchant))
}

func (handler *merchantHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	merchantID, err := strconv.Atoi(c.Param("merchantID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = handler.MerchantService.Delete(ctx, dto.MerchantDeleteReq{
		ID: merchantID,
		QueryData: dto.QueryData{
			ID:   claims.ID,
			Role: claims.Role,
		},
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(nil))
}
