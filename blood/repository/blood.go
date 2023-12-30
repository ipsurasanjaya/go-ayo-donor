package repository

import (
	"context"
	"errors"
	"go-ayo-donor/helper"
	"go-ayo-donor/model/cmd"
	"go-ayo-donor/model/domain"
	"io"

	"github.com/PuerkitoBio/goquery"
)

type (
	clientBloodScraper struct {
	}

	ClientBloodScraper interface {
		Get(ctx context.Context, pmiScrapperOp domain.PmiScrapperRequest) (script string, err error)
		GetByUdd(ctx context.Context, in io.Reader, pmiScrapperOp domain.PmiScrapperRequest) (selector *goquery.Selection, err error)
	}
)

func NewClient() *clientBloodScraper {
	return &clientBloodScraper{}
}

func (client *clientBloodScraper) GetByUdd(ctx context.Context, in io.Reader, pmiScraperOp domain.PmiScrapperRequest) (selector *goquery.Selection, err error) {
	pmiReq, ok := cmd.PmiRequestMap[pmiScraperOp]
	if !ok {
		return selector, errors.New("PMI operation not found")
	}
	curl, err := helper.ClientRequestUrl(in, pmiReq.Method, pmiReq.Url)
	if err != nil {
		return
	}

	curl.Find(".table-striped").Children().Each(func(i int, s *goquery.Selection) {
		selector = s.Find("td")
	})

	return
}

func (client *clientBloodScraper) Get(ctx context.Context, pmiScraperOp domain.PmiScrapperRequest) (script string, err error) {
	pmiReq, ok := cmd.PmiRequestMap[pmiScraperOp]
	if !ok {
		return script, errors.New("PMI operation not found")
	}

	curl, err := helper.ClientRequestUrl(nil, pmiReq.Method, pmiReq.Url)
	if err != nil {
		return
	}

	strScript := curl.Find("script").Text()

	return strScript, nil
}
