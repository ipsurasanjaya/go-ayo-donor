package cmd

import (
	"go-ayo-donor/model/domain"
	"net/http"
)

type (
	LinkAndMethod struct {
		Url    string
		Method string
	}

	pmiScrapperOp map[domain.PmiScrapperRequest]LinkAndMethod
)

var (
	PmiRequestMap = pmiScrapperOp{
		domain.GetBloodSupplyByUdd: LinkAndMethod{
			Url:    "https://ayodonor.pmi.or.id/?page=stok",
			Method: http.MethodPost,
		},
		domain.GetBloodSupplies: LinkAndMethod{
			Url:    "https://ayodonor.pmi.or.id/#",
			Method: http.MethodGet,
		},
		domain.GetMobileDonor: LinkAndMethod{
			Url:    "https://ayodonor.pmi.or.id",
			Method: http.MethodGet,
		},
		domain.GetMobileDonorByProvince: LinkAndMethod{
			Url:    "https://ayodonor.pmi.or.id/?page=mobile&module=MjAyMy0xMi0zMA==&prov=",
			Method: http.MethodGet,
		},
	}
)
