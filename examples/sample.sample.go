package examples

// Code shouldn't include secrets,
// and that's why is not normal to
// find code files gitignored, but
// anyway is still possible to sample.
//
// Go code does not use # as comments
// but if you want to sample this file
// you can do:

// #samplr#const sampleToken = "REPLACE_ME_TOKEN"
const sampleToken = "REPLACE_ME_TOKEN"

var (
	// #samplr#	sampleOtherToken = "REPLACE_ME_AGAIN"
	sampleOtherToken = "REPLACE_ME_AGAIN"
)
