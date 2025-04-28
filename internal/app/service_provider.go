package app

import (
	"bt_auth/internal/api/userAPI"
	"bt_auth/internal/closer"
	"bt_auth/internal/config"
	"bt_auth/internal/config/env"
	"bt_auth/internal/repository"
	"bt_auth/internal/repository/userRepository"
	"bt_auth/internal/service"
	"bt_auth/internal/service/userService"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type ServiceProvider struct {
	dbConfig   config.DBConfig
	grpcConfig config.GRPCConfig

	dbPool         *pgxpool.Pool
	userRepository repository.UserRepository

	userService service.UserService
	userServer  *userAPI.Server
}

func newServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) DBConfig() config.DBConfig {
	if sp.dbConfig == nil {
		cfg, err := env.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		sp.dbConfig = cfg
	}

	return sp.dbConfig
}

func (sp *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		sp.grpcConfig = cfg
	}

	return sp.grpcConfig
}

func (sp *ServiceProvider) DBPool(ctx context.Context) *pgxpool.Pool {
	if sp.dbPool == nil {
		pool, err := pgxpool.New(ctx, sp.DBConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db pool: %s", err.Error())
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %s", err.Error())
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		sp.dbPool = pool
	}

	return sp.dbPool
}

func (sp *ServiceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepository.NewRepository(sp.DBPool(ctx))
	}

	return sp.userRepository
}

func (sp *ServiceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewService(sp.UserRepository(ctx))
	}

	return sp.userService
}

func (sp *ServiceProvider) UserServer(ctx context.Context) *userAPI.Server {
	if sp.userServer == nil {
		sp.userServer = userAPI.NewServer(sp.UserService(ctx))
	}

	return sp.userServer
}
