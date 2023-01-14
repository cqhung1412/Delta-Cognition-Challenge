include .env

networkdocker:
	docker network create dog-network

postgres:
	docker run --name dog_container --network dog-network -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:alpine

createdb:
	docker exec -it dog_container createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

migrateup:
	migrate -path db/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

dropdb:
	docker exec -it dog_container dropdb ${DB_NAME}

stopdb: 
	docker stop dog_container

rmdb:
	docker rm dog_container

psql:
	docker exec -ti dog_container psql -U ${DB_USER} -d ${DB_NAME}

migratedown:
	migrate -path db/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go dog-recommend/db/sqlc Store

build:
	go build -v ./...

test:
	go test -v -cover ./...

server:
	go mod tidy && go run main.go

dev:
	nodemon --exec go run main.go --signal SIGTERM

swagger:
	swag init --parseDependency --parseInternal
	go run main.go

dockerbuild:
	docker build -t dog-recommend:latest .

dockerrun:
	docker run --name dogrecommend --network dog-network -p 8080:8080 -e GIN_MODE=release dog-recommend

.PHONY:*