package main

import (
	"bufio"
	"log"
	"os"
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
	// for scanner.Scan() {
	// 	// fmt.Println(scanner.Text())
	// 	if _, err := output.WriteString(scanner.Text() + "\n"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// err = output.Sync()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	writer := bufio.NewWriter(output)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		if _, err := writer.WriteString(scanner.Text() + "\n"); err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
