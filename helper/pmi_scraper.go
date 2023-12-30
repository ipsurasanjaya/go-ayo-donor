package helper

import (
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ClientRequestUrl(params io.Reader, method string, url string) (doc *goquery.Document, err error) {
	req, err := http.NewRequest(method, url, params)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return
	}

	doc, err = documentReaderInit(res.Body)
	if err != nil {
		return
	}

	return
}

func documentReaderInit(body io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	return doc, err
}
