dbml2sql:
	dbml2sql database.dbml -o database.sql --postgres
dbdocs:
	dbdocs build database.dbml

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root soybean_admin

dropdb:
	docker exec -it postgres12 dropdb soybean_admin

migrateup:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5432/soybean_admin?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5432/soybean_admin?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: dbml2sql dbdocs postgres createdb dropdb migrateup migratedown sqlc