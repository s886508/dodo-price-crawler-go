package parser

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/s886508/dodo-price-crawler-go/pkg/price"
)

func RetrievePrices(url_ string, page int, stations []string) price.PriceInfos {
	v := url.Values{}
	v.Add("tPage", strconv.Itoa(page))
	v.Add("Send", "搜尋")
	res, err := http.PostForm(url_, v)
	if err != nil {
		log.Printf("Failed to get response from url: %s, err: %v\n", url_, err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("Http request error, url: %s, status code: %d\n", url_, res.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Failed to create document reader, url: %s, err: %v\n", url_, err)
		return nil
	}

	prices := make(price.PriceInfos, 0)
	doc.Find("div.listbox").Each(func(i int, s *goquery.Selection) {
		title := s.Find("p").Text()
		link, _ := s.Find("a").Attr("href")

		for _, station := range stations {
			if strings.Contains(title, station) {
				tempUrl := url_[:strings.LastIndex(url_, "/")+1] + link
				price := retrivePriceDetail(tempUrl, station)
				prices = append(prices, price)
			}
		}
	})

	return prices
}

func retrivePriceDetail(url_ string, station string) *price.PriceInfo {
	res, err := http.Get(url_)
	if err != nil {
		log.Printf("Failed to get response from url: %s, err: %v\n", url_, err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("Http request error, url: %s, status code: %d\n", url_, res.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Failed to create document reader, url: %s, err: %v\n", url_, err)
		return nil
	}

	re := regexp.MustCompile(`\s+`)
	details := make([]string, 0)
	doc.Find(".tx01").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(j int, ss *goquery.Selection) {
			details = append(details, re.ReplaceAllString(ss.Text(), ""))
		})
	})

	info := &price.PriceInfo{
		Station: station,
		Detail:  details,
	}
	return info

}
