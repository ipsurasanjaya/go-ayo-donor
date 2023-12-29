package handler

import (
	"go-ayo-donor/blood/usecase"
	"go-ayo-donor/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	bloodHandler struct {
		BloodUsecase usecase.BloodUsecase
	}

	GetBloodSupplyByUddResponse struct {
		Product   string            `json:"product"`
		BloodType map[string]string `json:"bloodType"`
	}

	GetBloodSupplies struct {
		BloodType string `json:"bloodType"`
		Amount    string `json:"amount"`
	}
)

func NewHandler(usecase usecase.BloodUsecase) *bloodHandler {
	return &bloodHandler{
		BloodUsecase: usecase,
	}
}

func (bh *bloodHandler) GetBloodSupplyByUdd(c echo.Context) error {
	ctx := c.Request().Context()

	udd := c.Param("udd")
	uddReader := strings.NewReader("udd=" + udd)

	out, err := bh.BloodUsecase.GetBloodSupplyByUdd(ctx, uddReader)
	if err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to get blood supply by udd",
		)
	}

	var supplies []GetBloodSupplyByUddResponse
	for _, v := range out {
		supplies = append(supplies, GetBloodSupplyByUddResponse(v))
	}

	return helper.SuccessResponse(c, http.StatusOK, supplies)
}

func (handler *bloodHandler) GetBloodSupplies(c echo.Context) error {
	ctx := c.Request().Context()
	out, err := handler.BloodUsecase.GetBloodSupplies(ctx)
	if err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"failed to get blood supplies",
		)
	}

	var supplies []GetBloodSupplies
	for _, v := range out {
		supplies = append(supplies, GetBloodSupplies(v))
	}

	return helper.SuccessResponse(c, http.StatusOK, supplies)
}
