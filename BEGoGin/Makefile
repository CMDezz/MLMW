postgres :
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb: 
	docker exec -it postgres16 createdb --username=root --owner=root MLMW

dropdb: 
	docker exec -it postgres16 dropdb MLMW

migrateInit:
	migrate create -ext sql -dir dto/migration -seq MLMW

migrateInitUp:
	migrate -path dto/migration -database "postgresql://root:secret@localhost:5432/MLMW?sslmode=disable" -verbose up

migrateInitDown:
	migrate -path dto/migration -database "postgresql://root:secret@localhost:5432/MLMW?sslmode=disable" -verbose down

start:
	go run main.go

.PHONY: postgres createdb dropdb migrateInitUp migrateInitDown migrateInit start 