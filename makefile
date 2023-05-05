local:
	go run cmd/main.go
db_unix:
	docker run -d --name frog --hostname plandB -p 5430:5432 -e POSTGRES_PASSWORD=14881488 -v /var/lib/pgsql/data:/var/lib/postgresql/data --net net_postgres postgres:latest
create_migration:
# make create_migration name=name_your_migration_without_spaces 
	migrate create -ext sql -dir db/migrations -seq ${name}
migrate:
	migrate -database 'postgres://user:123456789@localhost:5430/test1?sslmode=disable' -path ./db/migrations up
migrate_down:
	migrate -database 'postgres://user:123456789@localhost:5430/test1?sslmode=disable' -path ./db/migrations down

db_test:
	docker run --name test-pg -p 5430:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=123456789 -d postgres:latest