package main

import (
	"fmt"
	handler "go-ayo-donor/blood/delivery/http"
	"go-ayo-donor/blood/repository/pmi"
	"go-ayo-donor/blood/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	webPort = "8080"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pc := pmi.NewClient()
	buc := usecase.NewUsecase(pc)

	v1 := e.Group("/v1")
	{
		api := v1.Group("/api")
		{
			bh := handler.NewHandler(buc)
			bloodGroup := api.Group("/bloods")

			bloodGroup.GET("/supplies", bh.GetBloodSupplies)
			bloodGroup.GET("/supplies/:udd", bh.GetBloodSupplyByUdd)
		}
	}

	if err := e.Start(fmt.Sprintf(":%d", 8080)); err != nil {
		log.Fatal("failed to start the server")
	}
}
