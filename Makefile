deps:
	go mod tidy && go mod vendor

checks:
	golangci-lint run -v

test:
	go test ./... -v -race
