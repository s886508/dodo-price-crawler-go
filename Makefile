BLD_NAME:=dodo-price-crawler

.PHONY: build 
build: cmd/crawler/main.go 
	go build -o bin/$(BLD_NAME) $<

