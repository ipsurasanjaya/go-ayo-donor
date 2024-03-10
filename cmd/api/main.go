package main

import (
	"fmt"
	bHandler "go-ayo-donor/blood/delivery/http"
	bRepo "go-ayo-donor/blood/repository"
	bUsecase "go-ayo-donor/blood/usecase"
	mHandler "go-ayo-donor/mobiledonor/delivery/http"
	mRepo "go-ayo-donor/mobiledonor/repository"
	mUsecase "go-ayo-donor/mobiledonor/usecase"
	"go-ayo-donor/model/domain"
	"go-ayo-donor/pql"
	pHandler "go-ayo-donor/provinces/delivery/http"
	pRepo "go-ayo-donor/provinces/repository"
	pUsecase "go-ayo-donor/provinces/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	webPort = 8080
)

func main() {
	e := echo.New()
	cfg := domain.Config{
		Host:     "localhost",
		Port:     "5432",
		Schema:   "public",
		DBName:   "go_ayo_donor",
		User:     "suras",
		SSLMode:  "disable",
		TimeZone: "Asia/Jakarta",
	}
	db, err := pql.CreateSQLDB(cfg)
	if err != nil {
		log.Fatalf("error %v when creating psql DB", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bc := bRepo.NewClient()
	buc := bUsecase.NewUsecase(bc)

	pc := mRepo.NewClient()
	muc := mUsecase.NewUsecase(pc)

	pr := pRepo.NewProvincesRepo(db)
	puc := pUsecase.NewUsecase(pr)

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

			ph := pHandler.NewHandler(puc)
			provinceGroup := api.Group("/provinces")
			provinceGroup.GET("", ph.Get)
		}
	}

	if err := e.Start(fmt.Sprintf(":%d", webPort)); err != nil {
		log.Fatal("failed to start the server")
	}
}
