build:
	@go build -o bin/restAPI cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/restAPI