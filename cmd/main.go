package main

import (
	userAPI "bt_auth/internal/api/user"
	"bt_auth/internal/config"
	"bt_auth/internal/config/env"
	userRepository "bt_auth/internal/repository/user"
	userService "bt_auth/internal/service/user"
	desc "bt_auth/pkg/user_v1"
	"context"
	"flag"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	dbConfig, err := env.NewDBConfig()
	if err != nil {
		log.Fatalf("failed to get db config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, dbConfig.DSN())

	userRepo := userRepository.NewRepository(pool)
	userServ := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, userAPI.NewServer(userServ))

	log.Printf("server listening at %v", grpcConfig.Address())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
