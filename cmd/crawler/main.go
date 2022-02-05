package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
	"github.com/s886508/dodo-price-crawler-go/pkg/config"
	"github.com/s886508/dodo-price-crawler-go/pkg/parser"
)

var args struct {
	ConfigFile string `default:"configs/config.json"`
}

func main() {
	arg.MustParse(&args)
	cfg := config.LoadConfig(args.ConfigFile)
	prices := parser.RetrievePrices(cfg.Url, cfg.Stations)

	bytes, err := json.Marshal(prices)
	if err != nil {
		log.Fatalf("Failed to encoding prices to json, err: %v\n", err)
	}
	fmt.Println(string(bytes))
}
