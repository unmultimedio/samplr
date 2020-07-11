package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputPathFor(t *testing.T) {
	const sampleExt = ".sample"

	assert.Equal(t,
		fmt.Sprintf("config%s.yaml", sampleExt),
		outputPathFor("config.yaml"),
	)
	assert.Equal(t,
		fmt.Sprintf("file.with.many.extensions%s.toml", sampleExt),
		outputPathFor("file.with.many.extensions.toml"),
	)
	assert.Equal(t,
		fmt.Sprintf("no_extension%s", sampleExt),
		outputPathFor("no_extension"),
	)

	assert.Equal(t,
		fmt.Sprintf("./a/b/c/config%s.yaml", sampleExt),
		outputPathFor("./a/b/c/config.yaml"),
	)
	assert.Equal(t,
		fmt.Sprintf("./a/b c/file.with.many.extensions%s.toml", sampleExt),
		outputPathFor("./a/b c/file.with.many.extensions.toml"),
	)
	assert.Equal(t,
		fmt.Sprintf("./a/b/c/some spaces%s", sampleExt),
		outputPathFor("./a/b/c/some spaces"),
	)
}

func TestSamplrNoKeys(t *testing.T) {
	input := "some line with no special keys"
	output, skip := sampleLine(input)
	assert.Equal(t, input+"\n", output)
	assert.Equal(t, false, skip)
}

func TestSamplrKey(t *testing.T) {
	const key = "#samplr#"

	testCases := []struct {
		input              string
		expectedSecondLine string
	}{
		{
			input:              key,
			expectedSecondLine: "",
		},
		{
			input:              key + "some content",
			expectedSecondLine: "some content",
		},
		{
			input:              key + " with lead spaces",
			expectedSecondLine: " with lead spaces",
		},
		{
			// With many keys, it just respects the first one
			input:              key + "many" + key + "keys" + key,
			expectedSecondLine: "many" + key + "keys" + key,
		},
		{
			input:              "  " + key + "space before key",
			expectedSecondLine: "  space before key",
		},
		{
			input:              "  content-" + key + "+before key",
			expectedSecondLine: "  content-+before key",
		},
	}

	for _, tc := range testCases {
		actualOutput, actualSkip := sampleLine(tc.input)
		expectedOutput := tc.input + "\n" + tc.expectedSecondLine + "\n"
		assert.Equal(t, expectedOutput, actualOutput)
		assert.Equal(t, true, actualSkip)
	}
}

func TestHideSamplrKey(t *testing.T) {
	const key = "#hsamplr#"

	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          key,
			expectedOutput: "",
		},
		{
			input:          key + "some content",
			expectedOutput: "some content",
		},
		{
			input:          key + " with lead spaces",
			expectedOutput: " with lead spaces",
		},
		{
			// With many keys, it just respects the first one
			input:          key + "many" + key + "keys" + key,
			expectedOutput: "many" + key + "keys" + key,
		},
		{
			input:          "  " + key + "space before key",
			expectedOutput: "  space before key",
		},
		{
			input:          "  content-" + key + "+before key",
			expectedOutput: "  content-+before key",
		},
	}

	for _, tc := range testCases {
		actualOutput, actualSkip := sampleLine(tc.input)
		assert.Equal(t, tc.expectedOutput+"\n", actualOutput)
		assert.Equal(t, true, actualSkip)
	}
}
