generate-auth-api:
	protoc --proto_path=api/auth/v1 --go_out=pkg/auth_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative api/auth/v1/auth.proto
