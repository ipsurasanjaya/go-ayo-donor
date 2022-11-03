package handler

import (
	"go-ayo-donor/blood/usecase"
	"go-ayo-donor/helper"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type (
	bloodHandler struct {
		BloodUsecase usecase.BloodUsecase
	}

	BloodHandler interface {
		GetBloodSupplyByUdd(w http.ResponseWriter, r *http.Request, params httprouter.Params)
		GetBloodSupplies(w http.ResponseWriter, r *http.Request, params httprouter.Params)
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

func NewHandler(usecase usecase.BloodUsecase) BloodHandler {
	return &bloodHandler{
		BloodUsecase: usecase,
	}
}

func (handler *bloodHandler) GetBloodSupplyByUdd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	udd := params.ByName("udd")
	uddReader := strings.NewReader("udd=" + udd)

	out, err := handler.BloodUsecase.GetBloodSupplyByUdd(uddReader)
	if err != nil {
		helper.ResponseError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var supplies []GetBloodSupplyByUddResponse
	for _, v := range out {
		supplies = append(supplies, GetBloodSupplyByUddResponse(v))
	}

	helper.ResponseSuccess(w, "Success get supply by UDD", supplies)
}

func (handler *bloodHandler) GetBloodSupplies(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	out, err := handler.BloodUsecase.GetBloodSupplies()
	if err != nil {
		helper.ResponseError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var supplies []GetBloodSupplies
	for _, v := range out {
		supplies = append(supplies, GetBloodSupplies(v))
	}

	helper.ResponseSuccess(w, "Success get blood supplies", supplies)
}
