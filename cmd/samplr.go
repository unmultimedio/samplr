// Package cmd contains the samplr commands
package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Sample samples the files
func Sample() {
	fileName := "configuration.yaml"

	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = input.Close() }()

	output, err := os.Create(fileName + ".sample")
	if err != nil {
		log.Fatal(err)
	}

	if err := replicateFile(input, output); err != nil {
		log.Fatal(err)
	}
}

func replicateFile(input *os.File, output *os.File) error {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	var l string
	var skip bool
	for scanner.Scan() {
		if skip {
			skip = false
			continue
		}
		l, skip = sampleLine(scanner.Text())
		if _, err := writer.WriteString(l); err != nil {
			return err
		}
	}
	defer func() { _ = writer.Flush() }()

	return scanner.Err()
}

// sampleLine returns sampled line and flag to skip next one
func sampleLine(l string) (string, bool) {
	const key = "#samplr#"
	i := strings.LastIndex(l, key)
	if i == -1 {
		return l + "\n", false
	}

	return l[i+len(key):] + "\n", true
}
