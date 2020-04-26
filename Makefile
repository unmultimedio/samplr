deps:
	go mod tidy && go mod vendor

lint:
	golangci-lint run -v

test:
	go test ./... -v -race
