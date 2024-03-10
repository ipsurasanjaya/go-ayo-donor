package http

import (
	"go-ayo-donor/helper"
	"go-ayo-donor/model/domain"
	"go-ayo-donor/provinces/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	GetProvinceReq struct {
		Limit  int64  `query:"limit"`
		Search string `query:"search"`
	}
	GetProvinceRes struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

type handler struct {
	puc usecase.ProvinceUsecase
}

func NewHandler(puc usecase.ProvinceUsecase) *handler {
	return &handler{
		puc: puc,
	}
}

func (ph *handler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	req := GetProvinceReq{}

	if err := c.Bind(&req); err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"bad request, check your json request",
		)
	}

	out, err := ph.puc.Get(ctx, domain.GetProvinceIn(req))
	if err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	res := []GetProvinceRes{}
	for _, v := range out {
		res = append(res, GetProvinceRes(v))
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		res,
	)
}
