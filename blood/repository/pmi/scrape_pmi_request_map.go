package pmi

import (
	"go-ayo-donor/model/domain"
	"net/http"
)

type (
	linkAndMethod struct {
		link   string
		method string
	}

	pmiScrapperOp map[domain.PmiScrapperRequest]linkAndMethod
)

var (
	pmiRequestMap = pmiScrapperOp{
		domain.GetBloodSupplyByUdd: linkAndMethod{
			link:   "https://ayodonor.pmi.or.id/?page=stok",
			method: http.MethodPost,
		},
		domain.GetBloodSupplies: linkAndMethod{
			link:   "https://ayodonor.pmi.or.id/#",
			method: http.MethodGet,
		},
	}
)
