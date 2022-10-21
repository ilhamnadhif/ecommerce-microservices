package handler

import (
	"api-gateway/dto"
	"api-gateway/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		ProductService: productService,
	}
}

type productHandler struct {
	ProductService service.ProductService
}

func (handler *productHandler) FindOneByID(c echo.Context) error {
	ctx := c.Request().Context()
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	product, err := handler.ProductService.FindOneByID(ctx, productID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(product))
}

func (handler *productHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	var merchantID int
	echo.QueryParamsBinder(c).Int("merchant_id", &merchantID)
	products, err := handler.ProductService.FindAll(ctx, merchantID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(products))
}

func (handler *productHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.ProductCreateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	product, err := handler.ProductService.Create(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(product))
}
func (handler *productHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var req dto.ProductUpdateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = productID
	product, err := handler.ProductService.Update(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(product))
}

func (handler *productHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = handler.ProductService.Delete(ctx, productID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(nil))
}
