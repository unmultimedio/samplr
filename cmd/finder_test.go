package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSamplrablePath(t *testing.T) {
	runConfig = config{}
	assert.False(t, isSamplrablePath("config"))
	assert.False(t, isSamplrablePath("config.yml"))

	runConfig = config{includes: []string{`\.yml$`}}
	assert.False(t, isSamplrablePath("config"))
	assert.True(t, isSamplrablePath("config.yml"))

	runConfig = config{
		excludes: []string{`config`},
		includes: []string{`\.yml$`},
	}
	assert.False(t, isSamplrablePath("config"))
	assert.False(t, isSamplrablePath("config.yml"))
	assert.True(t, isSamplrablePath("settings.yml"))
}
