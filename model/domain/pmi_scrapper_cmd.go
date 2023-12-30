package domain

type PmiScrapperRequest int32

const (
	UndifinedPmiScrapperRequest PmiScrapperRequest = iota
	GetBloodSupplyByUdd
	GetBloodSupplies
	GetMobileDonor
)
