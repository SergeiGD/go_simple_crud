package user

import (
	"context"
	"fmt"
	"simple_rest_crud/config"
	realRepo "simple_rest_crud/internal/user/data/repo/real"
	iRepo "simple_rest_crud/internal/user/domain/repo"
	"simple_rest_crud/internal/user/domain/service"
	"simple_rest_crud/pkg/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: dependency injection + use mock envs

func getConfig() *config.Config {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println("error reading conf")
		fmt.Println(err)
		return nil
	}
	return cfg
}

func getPsqlClient() *pgxpool.Pool {
	postgreClient, err := postgres.NewClient(context.TODO(), getConfig().Database)
	if err != nil {
		// TODO: logger
		fmt.Println("Error connecting to db")
		fmt.Println(err)
		return nil
	}
	return postgreClient
}

func GetUserRepo() iRepo.IUserRepository {
	return realRepo.NewPsqlUserRepository(getPsqlClient())
}

func GetUserService() *service.UserService {
	return service.NewUserService(GetUserRepo(), getConfig())
}
