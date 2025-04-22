generate-auth-api:
	protoc --proto_path=api/auth/v1 --go_out=pkg/auth_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative api/auth/v1/auth.proto

build:
	set GOOS=linux&& set GOARCH=amd64&& go build -o auth cmd/main.go

copy-to-server:
    powershell -Command "scp auth root@87.228.80.229:"

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/bladforceone/test-server:v0.0.1 .
	docker push cr.selcloud.ru/bladforceone/test-server:v0.0.1