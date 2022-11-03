package pmi

import (
	"fmt"
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

	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusOK {
		return
	}

	doc, err = DocumentReaderInit(res.Body)
	if err != nil {
		return
	}

	return
}

func DocumentReaderInit(body io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	return doc, err
}
