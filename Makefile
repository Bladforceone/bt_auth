# Загрузка переменных окружения из .env файла
include .env

# Определение переменной для строки подключения
LOCAL_MIGRATIONS_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) user=$(PG_USER) password=$(PG_PASSWORD) dbname=$(PG_NAME) sslmode=disable"

# Генерация API
generate-user-api:
	protoc --proto_path=api/user/v1 --go_out=pkg/user_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative api/user/v1/user.proto

# Сборка для Linux
build-linux:
	set GOOS=linux&& set GOARCH=amd64&& go build -o auth cmd/main.go

# Построение и пуш Docker-образа
docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/bladforceone/test-server:v0.0.1 .
	docker push cr.selcloud.ru/bladforceone/test-server:v0.0.1

# Установка зависимостей
install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Проверка статуса миграций
local-migration-status:
	goose -dir ${LOCAL_MIGRATIONS_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

# Применение миграций
local-migration-up:
	goose -dir ${LOCAL_MIGRATIONS_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

# Откат миграций
local-migration-down:
	goose -dir ${LOCAL_MIGRATIONS_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
