include .env

postgres:
	docker run --name postgres_container --network dog-network -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:alpine

createdb:
	docker exec -it postgres_container createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres_container dropdb ${DB_NAME}

stopdb: 
	docker stop postgres_container

rmdb:
	docker rm postgres_container

psql:
	docker exec -ti postgres_container psql -U ${DB_USER} -d ${DB_NAME}

migrateup:
	migrate -path db/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

build:
	go build -v ./...

test:
	go test -v -cover ./...

server:
	go mod tidy && go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go dog-recommend/db/sqlc Store

dockerbuild:
	docker build -t dog-recommend:latest .

dockerrun:
	docker run --name dogrecommend --network dog-network -p 8080:8080 -e GIN_MODE=release dog-recommend

networkdocker:
	docker network create dog-network

.PHONY:*