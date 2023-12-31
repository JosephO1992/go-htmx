build:
	go build -o bin/main.go cmd/main.go

run: 
	templ generate
	go run cmd/main.go
