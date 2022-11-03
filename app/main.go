package main

import (
	handler "go-ayo-donor/blood/delivery/http"
	"go-ayo-donor/blood/repository/pmi"
	"go-ayo-donor/blood/usecase"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	webPort = "8080"
)

func main() {
	pmiClient := pmi.NewClient()
	bloodUsecase := usecase.NewUsecase(pmiClient)
	bloodHandler := handler.NewHandler(bloodUsecase)

	router := httprouter.New()
	router.GET("/api/v1/supplies", bloodHandler.GetBloodSupplies)
	router.GET("/api/v1/supplies/:udd", bloodHandler.GetBloodSupplyByUdd)

	server := http.Server{
		Addr:    ":" + webPort,
		Handler: router,
	}

	log.Println("Start listening to server!")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
