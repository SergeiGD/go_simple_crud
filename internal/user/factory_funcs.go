package user

import (
	"context"
	"simple_rest_crud/config"
	realRepo "simple_rest_crud/internal/user/data/repo/real"
	iRepo "simple_rest_crud/internal/user/domain/repo"
	"simple_rest_crud/internal/user/domain/service"
	"simple_rest_crud/pkg/logging"
	"simple_rest_crud/pkg/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// TODO: dependency injection + use mock envs

func getConfig(logger *logging.Logger) *config.Config {
	cfg, err := config.GetConfig(logger)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("error connecting to DB")
		return nil
	}
	return cfg
}

func getPsqlClient(logger *logging.Logger) *pgxpool.Pool {
	postgreClient, err := postgres.NewClient(
		context.Background(),
		logger,
		getConfig(logger).Database,
	)
	if err != nil {
		return nil
	}
	return postgreClient
}

func GetUserRepo(logger *logging.Logger) iRepo.IUserRepository {
	return realRepo.NewPsqlUserRepository(getPsqlClient(logger), logger)
}

func GetUserService(logger *logging.Logger) *service.UserService {
	return service.NewUserService(GetUserRepo(logger), getConfig(logger))
}
