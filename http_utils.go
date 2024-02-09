package main

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func queryHtmlFromUrl(url string, query string, attr string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	attr, exists := doc.Find(query).Attr(attr)
	if !exists {
		return "", errors.New("The specified tag doesn't exist in the retrieved html")
	}
	return attr, nil
}
