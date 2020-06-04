package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestIsSamplrablePath(t *testing.T) {
	// Empty config
	assert.False(t, isSamplrablePath("config"))
	assert.False(t, isSamplrablePath("config.yml"))

	viper.Set("includes", []string{`\.yml$`})
	assert.False(t, isSamplrablePath("config"))
	assert.True(t, isSamplrablePath("config.yml"))

	viper.Set("includes", []string{`\.yml$`})
	viper.Set("excludes", []string{`config`})
	assert.False(t, isSamplrablePath("config"))
	assert.False(t, isSamplrablePath("config.yml"))
	assert.True(t, isSamplrablePath("settings.yml"))
}
