package usecase

import (
	"context"
	"go-ayo-donor/blood/repository/pmi"
	"go-ayo-donor/model/domain"
	"io"
	"regexp"
)

var bloodType = []string{
	"A+", "B+", "O+", "AB+",
}

type BloodUsecase interface {
	GetBloodSupplyByUdd(ctx context.Context, param io.Reader) (out []domain.GetBloodSupplyByUddOut, err error)
	GetBloodSupplies(ctx context.Context) (out []domain.GetBloodSuppliesOut, err error)
}

type bloodUsecase struct {
	ClientPmi pmi.ClientPmiScraper
}

func NewUsecase(clientPmi pmi.ClientPmiScraper) BloodUsecase {
	return &bloodUsecase{
		ClientPmi: clientPmi,
	}
}

func (b *bloodUsecase) GetBloodSupplyByUdd(ctx context.Context, params io.Reader) (out []domain.GetBloodSupplyByUddOut, err error) {
	var (
		stock   domain.GetBloodSupplyByUddOut
		counter int
	)

	selector, err := b.ClientPmi.GetBloodSupplyByUdd(
		ctx,
		params,
		domain.GetBloodSupplyByUdd,
	)
	if err != nil {
		return
	}

	for idx, tds := range selector.Nodes {
		if idx%6 != 0 {
			if idx%6 == 1 {
				stock.Product = tds.FirstChild.Data
				stock.BloodType = map[string]string{}
			} else {
				stock.BloodType[bloodType[counter-1]] = tds.FirstChild.Data
			}
			counter++
		}
		if counter == 5 {
			out = append(out, stock)
			counter = 0
		}
	}

	return
}

func (b *bloodUsecase) GetBloodSupplies(ctx context.Context) (out []domain.GetBloodSuppliesOut, err error) {
	var (
		status      bool
		data        string
		scrapedData []string
	)

	script, err := b.ClientPmi.GetBloodSupplies(ctx, domain.GetBloodSupplies)
	if err != nil {
		return
	}

	regex := regexp.MustCompile(`data: ([ -~]+)`)
	scrapeScript := regex.FindAllString(script, -1)

	for i := 0; i < len(scrapeScript[1]); i++ {
		if string(scrapeScript[1][i]) == " " {
			continue
		}

		if status {
			if string(scrapeScript[1][i]) == "," {
				scrapedData = append(scrapedData, data)
				data = ""
			} else {
				if string(scrapeScript[1][i]) == "]" {
					continue
				}
				data += string(scrapeScript[1][i])
			}
		}
		if string(scrapeScript[1][i]) == "[" {
			status = true
		}
	}

	var blood domain.GetBloodSuppliesOut
	for idx, v := range scrapedData {
		blood.BloodType = bloodType[idx]
		blood.Amount = v
		out = append(out, blood)
	}

	return
}
