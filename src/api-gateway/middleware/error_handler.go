package middleware

import (
	"api-gateway/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	data, ok := err.(*echo.HTTPError)
	if !ok {
		data = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := dto.WebResponse{
		Code:   data.Code,
		Status: http.StatusText(data.Code),
		Data:   data.Message,
	}
	c.JSON(data.Code, response)
}
