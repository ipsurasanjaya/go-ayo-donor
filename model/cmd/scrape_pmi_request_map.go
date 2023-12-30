package cmd

import (
	"go-ayo-donor/model/domain"
	"net/http"
)

type (
	LinkAndMethod struct {
		Link   string
		Method string
	}

	pmiScrapperOp map[domain.PmiScrapperRequest]LinkAndMethod
)

var (
	PmiRequestMap = pmiScrapperOp{
		domain.GetBloodSupplyByUdd: LinkAndMethod{
			Link:   "https://ayodonor.pmi.or.id/?page=stok",
			Method: http.MethodPost,
		},
		domain.GetBloodSupplies: LinkAndMethod{
			Link:   "https://ayodonor.pmi.or.id/#",
			Method: http.MethodGet,
		},
		domain.GetMobileDonor: LinkAndMethod{
			Link:   "https://ayodonor.pmi.or.id",
			Method: http.MethodGet,
		},
	}
)
