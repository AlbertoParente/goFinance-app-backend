createDb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

migrationup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop

.PHONY: createDb postgres


# postgres:
# 	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

# dropDb:
# 	dropdb --username=postgres --owner=postgres go_finance

# createDb:
# 	createdb --username=postgres --owner=postgres go_finance

# migrationDrop:
# 	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop

# migrationUp:
# 	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

# test:
# 	go test -v -cover ./...

# server:
# 	go run main.go

# .PHONY: createDb postgres dropdb migrateup migrationdrop test server