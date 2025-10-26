.PHONY: build run test clean migrate

# Переменные
APP_NAME=serv_golang_rest_api
DOCKER_COMPOSE=docker compose

# Сборка и запуск
build:
	$(DOCKER_COMPOSE) build

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs -f api

# Разработка
dev:
	$(DOCKER_COMPOSE) up db -d
	go run ./cmd/api

test:
	go test ./...

# Миграции
migrate:
	$(DOCKER_COMPOSE) exec api ./main migrate

# Очистка
clean:
	$(DOCKER_COMPOSE) down -v
	docker system prune -f