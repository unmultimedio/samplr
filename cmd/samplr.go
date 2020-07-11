// Package cmd contains the samplr commands
package cmd

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Sample samples the project
func Sample() {
	for _, path := range samplrableFiles() {
		if err := sampleFile(path); err != nil {
			logger.Error(err)
		}
	}
}

func sampleFile(filePath string) error {
	input, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() { _ = input.Close() }()

	output, err := os.Create(outputPathFor(filePath))
	if err != nil {
		return err
	}

	return replicateFile(input, output)
}

// outputPathFor returns a sampe file path for a given path:
// `myconfig.yaml` => `myconfig.sample.yaml`
// `Dockerfile` => `Dockerfile.sample`
func outputPathFor(filePath string) string {
	ext := filepath.Ext(filePath)
	if len(ext) > 0 {
		filePathNoExt := strings.TrimSuffix(filePath, ext)
		return filePathNoExt + ".sample" + ext
	}
	return filePath + ".sample"
}

func replicateFile(input *os.File, output *os.File) error {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	if viper.GetBool("autogencomments") {
		fileExt := filepath.Ext(output.Name())
		h := CommentFor(fileExt)
		if _, err := writer.WriteString(h); err != nil {
			return err
		}
	}

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
	logger.Info("sample generated: " + output.Name())

	return scanner.Err()
}

// sampleLine returns sampled line and flag to skip next one
func sampleLine(l string) (string, bool) {
	if secretKeyCompile.Match([]byte(l)) {
		return "", false
	}

	if m := hideKeyCompile.FindAllIndex([]byte(l), 1); m != nil {
		return l[m[0][1]:] + "\n", true
	}

	if m := keyCompile.FindAllIndex([]byte(l), 1); m != nil {
		return l + "\n" + l[m[0][1]:] + "\n", true
	}

	return l + "\n", false
}
