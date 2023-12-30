package repository

import (
	"context"
	"errors"
	"go-ayo-donor/helper"
	"go-ayo-donor/model/cmd"
	"go-ayo-donor/model/domain"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type (
	clientMobileScraper struct{}

	ClientMobileScraper interface {
		Get(ctx context.Context, pmiScraperOp domain.PmiScrapperRequest) (*goquery.Selection, error)
		GetByProvince(
			ctx context.Context,
			pmiScraperOp domain.PmiScrapperRequest,
			province string,
		) (*goquery.Selection, error)
	}
)

var _ ClientMobileScraper = (*clientMobileScraper)(nil)

func NewClient() *clientMobileScraper {
	return &clientMobileScraper{}
}

func (c *clientMobileScraper) Get(ctx context.Context, pmiScraperOp domain.PmiScrapperRequest) (*goquery.Selection, error) {
	pmiReq, ok := cmd.PmiRequestMap[pmiScraperOp]
	if !ok {
		return nil, errors.New("PMI operation not found")
	}

	curl, err := helper.ClientRequestUrl(nil, pmiReq.Method, pmiReq.Url)
	if err != nil {
		return nil, err
	}

	var selector *goquery.Selection
	curl.Find(".table-striped").Children().Each(func(i int, s *goquery.Selection) {
		selector = s.Find("td")
	})

	return selector, nil
}

func (c *clientMobileScraper) GetByProvince(
	ctx context.Context,
	pmiScraperOp domain.PmiScrapperRequest,
	province string,
) (*goquery.Selection, error) {
	pmiReq, ok := cmd.PmiRequestMap[pmiScraperOp]
	if !ok {
		return nil, errors.New("PMI operation not found")
	}

	p := cases.Title(language.Und, cases.NoLower).String(province)
	prov := url.QueryEscape(p)
	url := pmiReq.Url + prov

	curl, err := helper.ClientRequestUrl(nil, pmiReq.Method, url)
	if err != nil {
		return nil, err
	}

	var selector *goquery.Selection
	curl.Find(".table-striped").Children().Each(func(i int, s *goquery.Selection) {
		selector = s.Find("td")
	})

	return selector, nil
}
