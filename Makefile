run:
	go run cmd/backend/main.go

run-bin:
	./bin/backend

build:
	go build -v -p 10 -o bin/backend cmd/backend/main.go

build-run: build run-bin

test:
	go get gotest.tools/gotestsum
	gotestsum --format testname -- ./... -v -coverprofile=coverage.out
