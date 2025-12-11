run:
	go run cmd/main.go

build:
	mkdir build
	go build -o build/app ./cmd/main.go

migrate:
	go run db/migrate.go