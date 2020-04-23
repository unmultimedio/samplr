package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "configuration.yaml"

	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	output, err := os.Create(fileName + ".sample")
	if err != nil {
		log.Fatal(err)
	}

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
			log.Fatal(err)
		}
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
