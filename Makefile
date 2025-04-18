build:
	@go1.24.2 build -o bin/diplipito

run: build
	 @./bin/diplipito

test:
	@go1.24.2 test ./... -v --race