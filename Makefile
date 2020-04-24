build:
	go build -o build/samplr

checks:
	golangci-lint run -v
