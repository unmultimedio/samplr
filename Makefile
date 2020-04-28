deps:
	go mod tidy && go mod vendor

lint:
	golangci-lint run -v

release:
	go build -o ./build/samplr

test:
	go test ./... -v -race -cover
