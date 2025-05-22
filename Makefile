build:
	@go build -o bin/middleware cmd/middleware/main.go

run: build
	@./bin/middleware
