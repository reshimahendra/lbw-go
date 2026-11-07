test:
	go test ./... -failfast -cover -short
tests:
	go test ./... -v -failfast -cover -short
test-proof:
	go test ./... -coverprofile=proof.out -cover -short && go tool cover -html=proof.out 
show-proof:
	go tool cover -html=proof.out 
run:
	go run ./cmd/app/main.go
build:
	mv -f dist/server dist/server_$$(date +"%Y%m%d_%H%M%S") 2>/dev/null || true
	go build -o ./dist/server -ldflags '-s -w' ./cmd/app/main.go
tidy:
	go mod tidy
vendor:
	go mod vendor
