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
	GetByProvince(ctx context.Context, province string) ([]domain.GetMobileDonorByProvinceOut, error)
}

type mobileDonorUsecase struct {
	client repository.ClientMobileScraper
}

var _ MobileDonorUsecase = (*mobileDonorUsecase)(nil)

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

func (m *mobileDonorUsecase) GetByProvince(
	ctx context.Context,
	province string,
) ([]domain.GetMobileDonorByProvinceOut, error) {
	selector, err := m.client.GetByProvince(
		ctx,
		domain.GetMobileDonorByProvince,
		province,
	)
	if err != nil {
		return nil, err
	}

	out := []domain.GetMobileDonorByProvinceOut{}
	var i int = 1
	node := selector.Nodes

	if len(node) == 0 {
		return nil, domain.ErrDataNotFound
	}

	for i < len(node) {
		attr := node[i].LastChild.Attr
		var url string
		if len(attr) > 0 {
			url = attr[0].Val
		}

		strDP := node[i+3].FirstChild.Data
		dp, err := strconv.Atoi(strDP)
		if err != nil {
			return nil, err
		}

		md := domain.GetMobileDonorByProvinceOut{
			InstanceName: node[i].FirstChild.Data,
			GoogleMapURL: url,
			Hour:         node[i+2].FirstChild.Data,
			DonorPlan:    dp,
		}

		out = append(out, md)
		i += 5
	}
	return out, nil
}
