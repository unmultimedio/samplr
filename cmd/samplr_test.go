package cmd

import (
	"fmt"
	"strings"
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

var testCases = []struct {
	input              string
	expectedSecondLine string
}{
	{
		input:              "%s",
		expectedSecondLine: "",
	},
	{
		input:              "%ssome content",
		expectedSecondLine: "some content",
	},
	{
		input:              "%s with lead spaces",
		expectedSecondLine: " with lead spaces",
	},
	{
		// With many keys, it just respects the first one
		input:              "%s many %s keys %s",
		expectedSecondLine: " many %s keys %s",
	},
	{
		// space before key is removed
		input:              "  %sspace before key",
		expectedSecondLine: "space before key",
	},
	{
		// any content before key is removed
		input:              "before-%s+after",
		expectedSecondLine: "+after",
	},
}

func TestSamplrRegularKey(t *testing.T) {
	const key = "#samplr#"

	for _, tc := range testCases {
		tc.input = strings.ReplaceAll(tc.input, "%s", key)
		tc.expectedSecondLine = strings.ReplaceAll(tc.expectedSecondLine, "%s", key)
		expectedOutput := tc.input + "\n" + tc.expectedSecondLine + "\n"

		actualOutput, actualSkip := sampleLine(tc.input)
		assert.Equal(t, expectedOutput, actualOutput)
		assert.Equal(t, true, actualSkip)
	}
}

func TestSamplrHideKey(t *testing.T) {
	const hkey = "#hsamplr#"

	for _, tc := range testCases {
		tc.input = strings.ReplaceAll(tc.input, "%s", hkey)
		tc.expectedSecondLine = strings.ReplaceAll(tc.expectedSecondLine, "%s", hkey)
		expectedOutput := tc.expectedSecondLine + "\n"

		actualOutput, actualSkip := sampleLine(tc.input)
		assert.Equal(t, expectedOutput, actualOutput)
		assert.Equal(t, true, actualSkip)
	}
}

func TestSamplrSecretKey(t *testing.T) {
	const skey = "#ssamplr#"

	for _, tc := range testCases {
		tc.input = strings.ReplaceAll(tc.input, "%s", skey)

		actualOutput, actualSkip := sampleLine(tc.input)
		assert.Empty(t, actualOutput)
		assert.Equal(t, false, actualSkip)
	}
}
