DB_CONTAINER_NAME=postgres15
DB_PORT=5433
DB_USER=root
DB_PASSWORD=A112233a
DB_NAME=go-chat
POSTGRES_IMAGE=postgres:15-alpine
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

psqlinit:
	docker run --name $(DB_CONTAINER_NAME) \
		-p $(DB_PORT):5432 \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-d $(POSTGRES_IMAGE)

psql:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_USER)

createdb:
	docker exec -it $(DB_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DB_CONTAINER_NAME) dropdb --username=$(DB_USER) --if-exists $(DB_NAME)
	
mup:
	migrate -path db/migrations -database "$(DB_URL)" up

mdown:
	migrate -path db/migrations -database "$(DB_URL)" down

.PHONY: psqlinit psql createdb dropdb mup mdown
