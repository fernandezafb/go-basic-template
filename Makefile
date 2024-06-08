build:
	@go build -o build/package/go-basic-template cmd/main.go
run:
	@./build/package/go-basic-template
test:
	@go test -v ./...