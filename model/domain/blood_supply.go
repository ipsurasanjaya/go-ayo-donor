package domain

type (
	GetBloodSupplyByUddIn struct {
		UnitDonorDarah string
	}

	GetBloodSupplyByUddOut struct {
		Product   string
		BloodType map[string]string
	}

	GetBloodSuppliesOut struct {
		BloodType string
		Amount    string
	}
)
