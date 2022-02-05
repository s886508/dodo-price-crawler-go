package parser

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/s886508/dodo-price-crawler-go/pkg/price"
)

func RetrievePrices(url string, stations []string) price.PriceInfos {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get response from url: %s, err: %v\n", url, err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("Http request error, url: %s, status code: %d\n", url, res.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Failed to create document reader, url: %s, err: %v\n", url, err)
		return nil
	}

	prices := make(price.PriceInfos, 0)
	doc.Find("div.listbox").Each(func(i int, s *goquery.Selection) {
		title := s.Find("p").Text()
		link, _ := s.Find("a").Attr("href")

		for _, station := range stations {
			if strings.Contains(title, station) {
				tempUrl := url[:strings.LastIndex(url, "/")+1] + link
				price := retrivePriceDetail(tempUrl, station)
				prices = append(prices, price)
			}
		}
	})

	return prices
}

func retrivePriceDetail(url string, station string) *price.PriceInfo {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get response from url: %s, err: %v\n", url, err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("Http request error, url: %s, status code: %d\n", url, res.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Failed to create document reader, url: %s, err: %v\n", url, err)
		return nil
	}

	details := make([]string, 0)
	doc.Find(".tx01").Each(func(i int, s *goquery.Selection) {
		details = append(details, s.Find("p").Text())
	})

	info := &price.PriceInfo{
		Station: station,
		Detail:  strings.Join(details, ";"),
	}
	return info

}
