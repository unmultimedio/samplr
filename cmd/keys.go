package cmd

import "regexp"

var (
	anyKeyCompile, _    = regexp.Compile("#(h|s)?samplr#")
	keyCompile, _       = regexp.Compile("#samplr#")
	hideKeyCompile, _   = regexp.Compile("#hsamplr#")
	secretKeyCompile, _ = regexp.Compile("#ssamplr#")
)
