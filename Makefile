NAME=costCalculator

## Running current project in the terminal.
run:
	go run cmd/main.go

## DockerRun
docker-run:
	CGO_ENABLED=0 go build -o main ./cmd
	docker build -t costcalculatorimage -f ./Dockerfile .
	docker run -d --network host --name costcalculator costcalculatorimage

## Build
build:
	CGO_ENABLED=0 go build -o main ./cmd


## Format all code in the project.
format:
	@echo formating is running...
	go vet ./...
	go fmt ./...

##------------------------------------------------ DB migration commands -----------------------------------------------

## "postgres://user:password@host:port/name_db?sslmode=disable"
database=postgres://efim:25121997@localhost:6543/postgres?sslmode=disable

## Creates migrations files with name wich should be specified with flag 'n=some_name'
create-migrations:
	migrate create -ext sql -dir migrations -seq $(n)

## Roll migrations
migrate-up:
	migrate -path ./migrations -database $(database) up

## Rollback all migrations
## If you specify flag 's=i' this will rollback 'i' migrations.
migrate-down:
	migrate -path ./migrations -database $(database) down $(s)