package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Includes []string
	Excludes []string
}

var config Config

func loadConfig() {
	f, err := os.Open(".samplr.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() { _ = f.Close() }()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config loaded: %+v\n", config)
}
