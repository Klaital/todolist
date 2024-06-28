
internal/storage/queries/db.go: internal/storage/sqlc.yaml internal/storage/queries/*.sql internal/storage/migrations/*.sql
	@sqlc generate -f ./internal/storage/sqlc.yaml

.PHONY: migrate-dev
migrate-dev:
	migrate -source file://internal/storage/migrations/ -database "postgres://todo_tester:todotester123@localhost:5431/todolist_test?sslmode=disable" up

.PHONY: build
build: internal/storage/db.go
	@go generate ./...
	@go build ./cmd/server

.PHONY: psql
psql:
	PGPASSWORD=todotester123 psql -U todo_tester -h localhost -p 5431 todolist_test
