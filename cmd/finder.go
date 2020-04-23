package cmd

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// samplrableFiles looks the complete directory and returns filepaths for samplrable files
// - matches in the regexp patterns
// - contains somewhere the samplr key
func samplrableFiles() []string {
	var files []string
	filepath.Walk(".", func(filePath string, file os.FileInfo, err error) error {
		if !file.IsDir() && isSamplrable(filePath) {
			files = append(files, filePath)
		}
		return nil
	})
	return files
}

// isSamplrable checks that a file is included in patterns, and includes the keywords
func isSamplrable(filePath string) bool {
	// TODO support in configuration file
	var matchPatterns = []string{
		`\.yaml$`,
	}

	var matches bool
	for _, pattern := range matchPatterns {
		m, err := regexp.Match(pattern, []byte(filePath))
		if err != nil {
			log.Fatal(err)
			continue
		}
		if m {
			matches = true
			break
		}
	}
	if !matches {
		return false
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return false
	}

	return false
}
