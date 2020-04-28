package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Includes []string
	Excludes []string
}

var (
	config Config

	anyKeyCompile, _    = regexp.Compile("#(h|s)?samplr#")
	keyCompile, _       = regexp.Compile("#samplr#")
	hideKeyCompile, _   = regexp.Compile("#hsamplr#")
	secretKeyCompile, _ = regexp.Compile("#ssamplr#")
)

func loadConfig() {
	f, err := os.Open(".samplr.yml")
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
