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

	walkFunc := func(filePath string, file os.FileInfo, err error) error {
		if !file.IsDir() && isSamplrable(filePath) {
			files = append(files, filePath)
		}
		return nil
	}

	err := filepath.Walk(".", walkFunc)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

// isSamplrable checks that a file is included in patterns, and includes the keywords
func isSamplrable(filePath string) bool {
	if !isSamplrablePath(filePath) {
		// log.Print("ignored: " + filePath)
		return false
	}
	log.Print("matches: " + filePath)

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			log.Print("key in: " + filePath)
			return true
		}
	}
	log.Print("key not in: " + filePath)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return false
	}

	return false
}

func isSamplrablePath(filePath string) bool {
	included := pathMatches(filePath, true)
	excluded := pathMatches(filePath, false)
	if included && excluded {
		log.Print("excluded: " + filePath)
	}

	// Matches if included and not excluded
	return included && !excluded
}

// pathMatches checks if a filePath matches with includes (true) or excludes (false)
func pathMatches(filePath string, includes bool) bool {
	var patterns []string
	if includes {
		patterns = config.Includes
	} else {
		patterns = config.Excludes
	}

	var matches bool
	for _, pattern := range patterns {
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

	return matches
}
