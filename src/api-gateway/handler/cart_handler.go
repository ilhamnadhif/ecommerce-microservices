package handler

import (
	"api-gateway/dto"
	"api-gateway/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func NewCartHandler(cartService service.CartService) cartHandler {
	return cartHandler{
		CartService: cartService,
	}
}

type cartHandler struct {
	CartService service.CartService
}

func (handler *cartHandler) FindOneByID(c echo.Context) error {
	ctx := c.Request().Context()
	cartID, err := strconv.Atoi(c.Param("cartID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cart, err := handler.CartService.FindOneByID(ctx, cartID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(cart))
}

func (handler *cartHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	carts, err := handler.CartService.FindAll(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(carts))
}

func (handler *cartHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	var req dto.CartCreateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.QueryData.ID = claims.ID
	req.QueryData.Role = claims.Role
	cart, err := handler.CartService.Create(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(cart))
}
func (handler *cartHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	cartID, err := strconv.Atoi(c.Param("cartID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var req dto.CartUpdateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = cartID
	req.QueryData.ID = claims.ID
	req.QueryData.Role = claims.Role
	cart, err := handler.CartService.Update(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(cart))
}

func (handler *cartHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	cartID, err := strconv.Atoi(c.Param("cartID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = handler.CartService.Delete(ctx, dto.CartDeleteReq{
		ID: cartID,
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
