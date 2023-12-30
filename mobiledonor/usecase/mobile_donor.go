package usecase

import (
	"context"
	"go-ayo-donor/mobiledonor/repository"
	"go-ayo-donor/model/domain"
	"strconv"
	"strings"
)

type MobileDonorUsecase interface {
	Get(ctx context.Context) ([]domain.GetMobileDonorOut, error)
}

type mobileDonorUsecase struct {
	client repository.ClientMobileScraper
}

func NewUsecase(client repository.ClientMobileScraper) *mobileDonorUsecase {
	return &mobileDonorUsecase{
		client: client,
	}
}

func (m *mobileDonorUsecase) Get(ctx context.Context) ([]domain.GetMobileDonorOut, error) {
	selector, err := m.client.Get(ctx, domain.GetMobileDonor)
	if err != nil {
		return nil, err
	}

	out := []domain.GetMobileDonorOut{}
	md := domain.GetMobileDonorOut{}
	for i, v := range selector.Nodes {
		if i%2 == 0 {
			prov := v.FirstChild.FirstChild.FirstChild.Data
			md.Province = strings.Replace(prov, "\u00a0", "", -1)
			continue
		}

		strAmt := v.FirstChild.FirstChild.Data
		strAmt = strings.Trim(strAmt, " Lokasi")
		amt, err := strconv.Atoi(strAmt)
		if err != nil {
			return nil, err
		}

		md.Amount = amt
		out = append(out, md)
	}
	return out, nil
}
