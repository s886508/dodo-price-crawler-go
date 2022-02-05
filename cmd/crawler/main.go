package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
	"github.com/s886508/dodo-price-crawler-go/pkg/config"
	"github.com/s886508/dodo-price-crawler-go/pkg/parser"
	"github.com/s886508/dodo-price-crawler-go/pkg/price"
)

var args struct {
	ConfigFile string `default:"configs/config.json"`
}

func main() {
	arg.MustParse(&args)
	cfg := config.LoadConfig(args.ConfigFile)
	prices := make(price.PriceInfos, 0)
	for page := 1; page <= 3; page++ {
		prices = append(prices, parser.RetrievePrices(cfg.Url, page, cfg.Stations)...)
	}

	bytes, err := json.Marshal(prices)
	if err != nil {
		log.Fatalf("Failed to encoding prices to json, err: %v\n", err)
	}
	fmt.Println(string(bytes))
}
