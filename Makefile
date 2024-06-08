include .env.example

build:
	@go build -o build/package/go-basic-template cmd/main.go
run:
	@PORT=$(PORT) \
	 VERSION=$(VERSION) \
	./build/package/go-basic-template
test:
	@go test -v ./...