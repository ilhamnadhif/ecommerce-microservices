package handler

import (
	"api-gateway/dto"
	"api-gateway/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewAuthHandler(authService service.AuthService) authHandler {
	return authHandler{
		AuthService: authService,
	}
}

type authHandler struct {
	AuthService service.AuthService
}

func (handler *authHandler) LoginMerchant(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.LoginReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	merchant, err := handler.AuthService.LoginMerchant(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchant))
}

func (handler *authHandler) LoginCustomer(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.LoginReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	merchant, err := handler.AuthService.LoginCustomer(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(merchant))
}
