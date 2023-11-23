postgres:
	docker run --name learn-go-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d --rm postgres:12

createdb:
	docker exec -it learn-go-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it learn-go-postgres dropdb simple_bank

migrate-force-version:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" force 1

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -count=1 -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc migrate-force-version
