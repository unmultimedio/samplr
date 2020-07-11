package cmd

import "github.com/spf13/viper"

type config struct {
	autogencomments bool
	excludes        []string
	includes        []string
}

var runConfig config

func InitConfig() {
	agc := viper.GetBool("autogencomments")
	exc := viper.GetStringSlice("excludes")
	inc := viper.GetStringSlice("includes")

	runConfig = config{
		autogencomments: agc,
		excludes:        exc,
		includes:        inc,
	}
}
