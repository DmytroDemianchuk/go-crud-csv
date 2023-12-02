migrate-up:
	migrate -path ./migrations -database postgres://root:test@localhost:5432/store?sslmode=disable up
migrate-down:
	migrate -path ./migrations -database postgres://root:test@localhost:5432/store?sslmode=disable down

migrate-drop:
	migrate -path ./migrations -database postgres://root:test@localhost:5432/store?sslmode=disable drop

run:
	docker-compose up -d --build