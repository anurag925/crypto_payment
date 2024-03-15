all:
	go run cmd/server/main.go 

build:
	go build cmd/server/main.go
run:
	./main
worker:
	go run cmd/worker/main.go
asynqmon:
	./cmd/worker/asynqmon --redis-url redis://127.0.0.1:6379/1
migrate:
	go run cmd/migrate/main.go 
script:
	go run cmd/script/main.go 