dev:
	docker compose up -d
	air

db-up:
	docker compose up -d

db-down:
	docker compose down

sqlc:
	sqlc generate
