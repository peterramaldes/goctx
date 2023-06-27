build:
	@go build -o ./bin/goctx main.go

run: build
	@./bin/goctx


