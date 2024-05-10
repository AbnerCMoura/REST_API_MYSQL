build:
	@go build -o bin/apirest cmd/main.go
	
test:
	@go test -v ./...
	
run: build
	@./bin/apirest