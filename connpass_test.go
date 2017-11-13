package connpass

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

const (
	proxyURL = ""
	apiURL   = "https://connpass.com/api/v1/event/"
)

var (
	count       = 20
	keywordsAnd = []string{"東京"}
	keywordsOr  = []string{"elasticsearch", "golang", "docker", "kubernetes"}
)

func TestDefault(t *testing.T) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyURL)
	}

	query := Query{
		KeywordAnd: keywordsAnd,
		KeywordOr:  keywordsOr,
		Count:      count,
	}

	requestURL := apiURL + query.getQueryString()
	log.Println(requestURL)

	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	res, err := client.Get(requestURL)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var record Response
	if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	for _, event := range record.Events {
		fmt.Println(event.Address, event.Title, event.EventURL)
	}
}
