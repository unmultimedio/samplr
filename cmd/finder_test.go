package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSamplrablePath(t *testing.T) {
	// Empty config
	assert.False(t, isSamplrablePath("config"))
	assert.False(t, isSamplrablePath("config.yml"))

	config.Includes = []string{`\.yml$`}
	assert.False(t, isSamplrablePath("config"))
	assert.True(t, isSamplrablePath("config.yml"))

	config.Includes = []string{`\.yml$`}
	config.Excludes = []string{`config`}
	assert.False(t, isSamplrablePath("config"))
	assert.False(t, isSamplrablePath("config.yml"))
	assert.True(t, isSamplrablePath("settings.yml"))
}
