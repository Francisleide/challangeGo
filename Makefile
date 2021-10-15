test:
	go test -v -race ./...

build:
	go build cmd/challenge-api/main.go

run: 
	go run cmd/challenge-api/main.go
swagger:
	swag init -g docs/docs.go


