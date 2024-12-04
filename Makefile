sqlc:
	sqlc generate -f ./database/sqlc.yaml
new-migration:
	migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up:
	migrate -database $(url) -path database/migrations up
migrate-down:
	migrate -database $(url) -path database/migrations down
