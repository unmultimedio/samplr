// Package cmd contains the samplr commands
package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const key = "#samplr#"

// Sample samples the project
func Sample() {
	for _, path := range samplrableFiles() {
		if err := sampleFile(path); err != nil {
			log.Fatal(err)
		}
	}
}

func sampleFile(filePath string) error {
	input, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() { _ = input.Close() }()

	output, err := os.Create(filePath + ".sample")
	if err != nil {
		return err
	}

	return replicateFile(input, output)
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
	i := strings.LastIndex(l, key)
	if i == -1 {
		return l + "\n", false
	}

	return l[i+len(key):] + "\n", true
}
