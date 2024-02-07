.SILENT:

build: 
	go build -o bin/expense-x-api cmd/api/main.go

run:build
	./bin/expense-x-api
