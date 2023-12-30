package repository

import (
	"context"
	"errors"
	"go-ayo-donor/helper"
	"go-ayo-donor/model/cmd"
	"go-ayo-donor/model/domain"

	"github.com/PuerkitoBio/goquery"
)

type (
	clientMobileScraper struct{}

	ClientMobileScraper interface {
		Get(ctx context.Context, pmiScraperOp domain.PmiScrapperRequest) (*goquery.Selection, error)
	}
)

func NewClient() *clientMobileScraper {
	return &clientMobileScraper{}
}

func (c *clientMobileScraper) Get(ctx context.Context, pmiScraperOp domain.PmiScrapperRequest) (*goquery.Selection, error) {
	pmiReq, ok := cmd.PmiRequestMap[pmiScraperOp]
	if !ok {
		return nil, errors.New("PMI operation not found")
	}

	curl, err := helper.ClientRequestUrl(nil, pmiReq.Method, pmiReq.Link)
	if err != nil {
		return nil, err
	}

	var selector *goquery.Selection
	curl.Find(".table-striped").Children().Each(func(i int, s *goquery.Selection) {
		selector = s.Find("td")
	})

	return selector, nil
}
