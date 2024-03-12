package http

import (
	"errors"
	"go-ayo-donor/helper"
	"go-ayo-donor/model/domain"
	"go-ayo-donor/transfusionunits/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	GetByProvinceIDRes struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}
)

type handler struct {
	tuc usecase.TransfusionUnitUsecase
}

func NewHandler(tuc usecase.TransfusionUnitUsecase) *handler {
	return &handler{
		tuc: tuc,
	}
}

func (th *handler) GetByProvinceID(c echo.Context) error {
	ctx := c.Request().Context()

	provinceID := c.Param("id")
	log.Println(provinceID)

	provinceIDInt, err := strconv.Atoi(provinceID)
	if err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"bad request, check your json request",
		)
	}

	out, err := th.tuc.GetByProvinceID(ctx, int64(provinceIDInt))
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			return helper.ErrorResponse(
				c,
				http.StatusNotFound,
				"not found, the data is not found",
			)
		}

		return helper.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"internal server error, failed get province by id",
		)
	}

	res := []GetByProvinceIDRes{}

	for _, v := range out {
		res = append(res, GetByProvinceIDRes(v))
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		res,
	)
}
