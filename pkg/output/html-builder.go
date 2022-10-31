package output

import (
	"html/template"
	"log"
	"os"
	"path"

	"github.com/s886508/dodo-price-crawler-go/pkg/price"
)

func PriceInfoToHtml(prices price.PriceInfos, fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Failed to open file for writing html content, err: %v\n", err)
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(path.Join("template", "index.html.template"))
	if err != nil {
		log.Fatalf("Failed to parse html template, err: %v\n", err)

	}

	data := struct {
		Items price.PriceInfos
	}{
		Items: prices,
	}
	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Failed to write html content, err: %v\n", err)

	}
}
