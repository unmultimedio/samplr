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

func TestSampleLineKey(t *testing.T) {
	const key = "#samplr#"
	var l string
	var skip bool

	l, skip = sampleLine(fmt.Sprintf("%scontent", key))
	assert.Equal(t, fmt.Sprintf("%scontent\ncontent\n", key), l)
	assert.True(t, skip)

	l, skip = sampleLine(fmt.Sprintf("   %scontent", key))
	assert.Equal(t, fmt.Sprintf("   %scontent\ncontent\n", key), l)
	assert.True(t, skip)

	// Does not keep relative space for content
	l, skip = sampleLine(fmt.Sprintf("   %scontent", key))
	assert.NotEqual(t, fmt.Sprintf("   %scontent\n   content\n", key), l)
	assert.True(t, skip)
}

func TestSampleLineHideKey(t *testing.T) {
	const hideKey = "#hsamplr#"
	var l string
	var skip bool

	l, skip = sampleLine(fmt.Sprintf("%scontent", hideKey))
	assert.Equal(t, "content\n", l)
	assert.True(t, skip)

	l, skip = sampleLine(fmt.Sprintf("   %scontent", hideKey))
	assert.Equal(t, "content\n", l)
	assert.True(t, skip)

	// Does not keep relative space for content
	l, skip = sampleLine(fmt.Sprintf("   %scontent", hideKey))
	assert.NotEqual(t, "   content\n", l)
	assert.True(t, skip)
}

func TestSampleLineSecretKey(t *testing.T) {
	const secretKey = "#ssamplr#"
	var l string
	var skip bool

	l, skip = sampleLine(fmt.Sprintf("%scontent", secretKey))
	assert.Equal(t, "", l)
	assert.False(t, skip)

	l, skip = sampleLine(fmt.Sprintf("   %scontent", secretKey))
	assert.Equal(t, "", l)
	assert.False(t, skip)
}
