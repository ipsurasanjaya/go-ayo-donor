package domain

type (
	GetMobileDonorOut struct {
		Province string
		Amount   int
	}

	GetMobileDonorByProvinceOut struct {
		InstanceName string
		GoogleMapURL string
		Hour         string
		DonorPlan    int
	}
)
