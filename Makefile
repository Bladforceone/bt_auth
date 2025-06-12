# Загрузка переменных окружения из .env файла
include .env

# Определение переменной для строки подключения
LOCAL_MIGRATIONS_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) user=$(PG_USER) password=$(PG_PASSWORD) dbname=$(PG_NAME) sslmode=disable"

# Генерация API
generate-user-api:
	protoc \
        --proto_path=api/user/v1 --proto_path=vendor.protogen\
        --go_out=pkg/user_v1 --go_opt=paths=source_relative \
        --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
        --validate_out=lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
        --plugin=protoc-gen-validate=C:/Users/bladf/go/bin/protoc-gen-validate.exe \
        --grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
        --plugin=protoc-gen-grpc-gateway=C:/Users/bladf/go/bin/protoc-gen-grpc-gateway.exe \
        api/user/v1/user.proto


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
vendor-proto:
	@if not exist "vendor.protogen\\validate" ( \
		mkdir "vendor.protogen\\validate" && \
		git clone https://github.com/envoyproxy/protoc-gen-validate "vendor.protogen\\protoc-gen-validate" && \
		move "vendor.protogen\\protoc-gen-validate\\validate\\*.proto" "vendor.protogen\\validate" && \
		rmdir /s /q "vendor.protogen\\protoc-gen-validate" \
	)
	@if not exist "vendor.protogen\\google" ( \
		git clone https://github.com/googleapis/googleapis "vendor.protogen\\googleapis" && \
		mkdir "vendor.protogen\\google" && \
		move "vendor.protogen\\googleapis\\google\\api" "vendor.protogen\\google" && \
		rmdir /s /q "vendor.protogen\\googleapis" \
	)
