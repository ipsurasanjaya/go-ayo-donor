package http

import (
	"go-ayo-donor/helper"
	"go-ayo-donor/mobiledonor/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	GetMobileDonorRes struct {
		Province string `json:"province"`
		Amount   int    `json:"amount"`
	}
)

type mobileDonorHandler struct {
	muc usecase.MobileDonorUsecase
}

func NewHandler(muc usecase.MobileDonorUsecase) *mobileDonorHandler {
	return &mobileDonorHandler{
		muc: muc,
	}
}

func (mh *mobileDonorHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	out, err := mh.muc.Get(ctx)
	if err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to get mobile donor",
		)
	}

	res := []GetMobileDonorRes{}

	for _, v := range out {
		res = append(res, GetMobileDonorRes(v))
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		res,
	)
}
