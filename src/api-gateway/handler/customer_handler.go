package handler

import (
	"api-gateway/dto"
	"api-gateway/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func NewCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{
		CustomerService: customerService,
	}
}

type customerHandler struct {
	CustomerService service.CustomerService
}

func (handler *customerHandler) FindOneByID(c echo.Context) error {
	ctx := c.Request().Context()
	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	customer, err := handler.CustomerService.FindOneByID(ctx, customerID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(customer))
}

func (handler *customerHandler) FindOneByCommon(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	if claims.Role != dto.CUSTOMER_ROLE {
		return echo.NewHTTPError(http.StatusForbidden, "customer: access denied for this role")
	}
	customer, err := handler.CustomerService.FindOneByID(ctx, claims.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(customer))
}

func (handler *customerHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	customers, err := handler.CustomerService.FindAll(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(customers))
}

func (handler *customerHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.CustomerCreateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	customer, err := handler.CustomerService.Create(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(customer))
}
func (handler *customerHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var req dto.CustomerUpdateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = customerID
	req.QueryData.ID = claims.ID
	req.QueryData.Role = claims.Role
	customer, err := handler.CustomerService.Update(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(customer))
}

func (handler *customerHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.JWTCustomClaims)
	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = handler.CustomerService.Delete(ctx, dto.CustomerDeleteReq{
		ID: customerID,
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
