run_api:
	go run ./cmd/api/main.go
	
create_migration:
	migrate create -ext sql -dir db/migrations $(FILENAME)

run_migrations:
	migrate -database "postgresql://suras@localhost:5432/go_ayo_donor?sslmode=disable" -path db/migrations up