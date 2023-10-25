test:
	go test -v ./... -cover

run:
	go run cmd/api/main.go serve-http

install:
	@go get cmd/api
