package helper

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, Response{
		Data:    data,
		Code:    code,
		Message: "OK",
	})
}

func ErrorResponse(
	c echo.Context,
	code int,
	message string,
) error {
	return c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}
