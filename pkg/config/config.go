package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Url      string   `json:"url"`
	Stations []string `json:"stations"`
}

func LoadConfig(filePath string) *Config {
	if len(filePath) == 0 {
		log.Fatal("empty path to load config")
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Fail to read config file")
	}

	config := &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		log.Fatalf("Fail to parse config file")
	}
	return config

}
