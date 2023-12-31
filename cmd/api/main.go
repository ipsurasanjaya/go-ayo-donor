package main

import (
	"fmt"
	bHandler "go-ayo-donor/blood/delivery/http"
	bRepo "go-ayo-donor/blood/repository"
	bUsecase "go-ayo-donor/blood/usecase"
	mHandler "go-ayo-donor/mobiledonor/delivery/http"
	mRepo "go-ayo-donor/mobiledonor/repository"
	mUsecase "go-ayo-donor/mobiledonor/usecase"
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

	bc := bRepo.NewClient()
	buc := bUsecase.NewUsecase(bc)

	pc := mRepo.NewClient()
	muc := mUsecase.NewUsecase(pc)

	v1 := e.Group("/v1")
	{
		api := v1.Group("/api")
		{
			bh := bHandler.NewHandler(buc)
			bloodGroup := api.Group("/bloods")

			bloodGroup.GET("/supplies", bh.GetBloodSupplies)
			bloodGroup.GET("/supplies/:udd", bh.GetBloodSupplyByUdd)

			mh := mHandler.NewHandler(muc)
			mobileGroup := api.Group("/mobiles")
			mobileGroup.GET("", mh.Get)
			mobileGroup.GET("/:province", mh.GetByProvince)
		}
	}

	if err := e.Start(fmt.Sprintf(":%d", 8080)); err != nil {
		log.Fatal("failed to start the server")
	}
}
