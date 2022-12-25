

DB_CONTAINER = postgres12-auth
DB_USERNAME = root
DB_PASSWORD = root
DB_DBNAME = auth
DB_PORT = 5432
DB_HOST = localhost


startdb:
	docker start $(DB_CONTAINER)

stopdb:
	docker stop $(DB_CONTAINER)

postgres:
	docker run --name $(DB_CONTAINER) -p $(DB_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:12-alpine 

createdb:
	docker exec -it $(DB_CONTAINER) createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) $(DB_DBNAME)

dropdb:
	docker exec -it $(DB_CONTAINER) dropdb $(DB_DBNAME)

migrateup:
	migrate -path db/migration -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DBNAME)?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DBNAME)?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run .

run:
	CompileDaemon -command="make server"

mock:
	mockgen -package mockdb --destination=db/mock/store.go github.com/brkss/go-auth/db/sqlc Store

gen: sqlc mock

.PHONY: startdb stopdb postgres createdb dropdb migrateup migratedown sqlc test server run mock
