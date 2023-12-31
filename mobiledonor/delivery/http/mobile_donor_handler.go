package http

import (
	"errors"
	"go-ayo-donor/helper"
	"go-ayo-donor/mobiledonor/usecase"
	"go-ayo-donor/model/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	GetMobileDonorRes struct {
		Province string `json:"province"`
		Amount   int    `json:"amount"`
	}

	GetMobileByProvinceRes struct {
		InstanceName string `json:"instanceName"`
		GoogleMapURL string `json:"googleMapURL"`
		Hour         string `json:"hour"`
		DonorPlan    int    `json:"donorPlan"`
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

func (mh *mobileDonorHandler) GetByProvince(c echo.Context) error {
	ctx := c.Request().Context()

	province := c.Param("province")

	out, err := mh.muc.GetByProvince(ctx, province)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			return helper.ErrorResponse(
				c,
				http.StatusBadRequest,
				"mobiles is not found",
			)
		}
		return helper.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to get mobiles by province",
		)
	}

	res := []GetMobileByProvinceRes{}

	for _, v := range out {
		res = append(res, GetMobileByProvinceRes(v))
	}

	return helper.SuccessResponse(c, http.StatusOK, res)
}
