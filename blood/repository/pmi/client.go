package pmi

import (
	"context"
	"errors"
	"go-ayo-donor/model/domain"
	"io"

	"github.com/PuerkitoBio/goquery"
)

type (
	clientPmiScraper struct {
	}

	ClientPmiScraper interface {
		GetBloodSupplies(ctx context.Context, pmiScrapperOp domain.PmiScrapperRequest) (script string, err error)
		GetBloodSupplyByUdd(ctx context.Context, in io.Reader, pmiScrapperOp domain.PmiScrapperRequest) (selector *goquery.Selection, err error)
	}
)

func NewClient() ClientPmiScraper {
	return &clientPmiScraper{}
}

func (client *clientPmiScraper) GetBloodSupplyByUdd(ctx context.Context, in io.Reader, pmiScrapperOp domain.PmiScrapperRequest) (selector *goquery.Selection, err error) {
	pmiReq, ok := pmiRequestMap[pmiScrapperOp]
	if !ok {
		return selector, errors.New("PMI operation not found")
	}
	curl, err := ClientRequestUrl(in, pmiReq.method, pmiReq.link)
	if err != nil {
		return
	}

	curl.Find(".table-striped").Children().Each(func(i int, s *goquery.Selection) {
		selector = s.Find("td")
	})

	return
}

func (client *clientPmiScraper) GetBloodSupplies(ctx context.Context, pmiScrapperOp domain.PmiScrapperRequest) (script string, err error) {
	pmiReq, ok := pmiRequestMap[pmiScrapperOp]
	if !ok {
		return script, errors.New("PMI operation not found")
	}

	curl, err := ClientRequestUrl(nil, pmiReq.method, pmiReq.link)
	if err != nil {
		return
	}

	strScript := curl.Find("script").Text()

	return strScript, nil
}
