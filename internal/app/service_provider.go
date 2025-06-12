package app

import (
	"bt_auth/internal/api/userAPI"
	"bt_auth/internal/client/db"
	"bt_auth/internal/client/db/pg"
	"bt_auth/internal/client/db/transaction"
	"bt_auth/internal/closer"
	"bt_auth/internal/config"
	"bt_auth/internal/config/env"
	"bt_auth/internal/repository"
	"bt_auth/internal/repository/userRepository"
	"bt_auth/internal/service"
	"bt_auth/internal/service/userService"
	"context"
	"log"
)

type ServiceProvider struct {
	dbConfig   config.DBConfig
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	dbClient       db.Client
	txManager      db.TxManager
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

func (sp *ServiceProvider) HTTPConfig() config.HTTPConfig {
	if sp.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		sp.httpConfig = cfg
	}

	return sp.httpConfig
}

func (sp *ServiceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		cl, err := pg.New(ctx, sp.DBConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db pool: %s", err.Error())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %s", err.Error())
		}
		closer.Add(cl.Close)

		sp.dbClient = cl
	}

	return sp.dbClient
}

func (sp *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DBClient(ctx).DB())
	}

	return sp.txManager
}

func (sp *ServiceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.userRepository
}

func (sp *ServiceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewService(
			sp.UserRepository(ctx),
			sp.TxManager(ctx),
		)
	}

	return sp.userService
}

func (sp *ServiceProvider) UserServer(ctx context.Context) *userAPI.Server {
	if sp.userServer == nil {
		sp.userServer = userAPI.NewServer(sp.UserService(ctx))
	}

	return sp.userServer
}
