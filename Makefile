test:
	go test ./... -failfast -cover -short
test-proof:
	go test ./... -coverprofile=proof.out -cover -short && go tool cover -html=proof.out 
show-proof:
	go tool cover -html=proof.out 
run:
	go run ./cmd/app/main.go
build:
	go build -o ./dist/server -ldflags '-s -w' ./cmd/app/main.go
tidy:
	go mod tidy
vendor:
	go mod vendor
