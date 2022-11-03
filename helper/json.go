package helper

import (
	"encoding/json"
	"go-ayo-donor/model/web"
	"log"
	"net/http"
)

func WriteToResponseBody(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(body)
	if err != nil {
		log.Fatal(err)
	}
}

func ResponseSuccess(w http.ResponseWriter, message string, data interface{}) {
	apiResponse := web.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: message,
		Data:    data,
	}

	WriteToResponseBody(w, http.StatusOK, apiResponse)
}

func ResponseError(w http.ResponseWriter, message string, code int) {
	apiResponse := web.ApiResponse{
		Code:    code,
		Status:  "Error",
		Message: message,
		Data:    nil,
	}

	WriteToResponseBody(w, code, apiResponse)
}
