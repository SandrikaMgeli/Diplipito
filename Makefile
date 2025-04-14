build:
	@go build -o bin/diplipito

run: build
	 @./bin/diplipito

test:
	@go test ./... -v --race