include .env

.SILENT:

build: 
	go build -o bin/expense-x-api cmd/api/main.go

run:build
	./bin/expense-x-api

mgcreate:
	migrate create -ext sql -dir migrations/ -seq $(name)

mgup:
	migrate -path migrations/ -database "${DB_URL}" -verbose up

mgdown:
	migrate -path migrations/ -database "${DB_URL}" -verbose down

mgfix:
	migrate -path migrations/ -database "${DB_URL}" force $(v)

