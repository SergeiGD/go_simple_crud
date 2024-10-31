package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"simple_rest_crud/config"
	"simple_rest_crud/internal/user/domain/entity"
	"simple_rest_crud/internal/user/domain/repo"
	"simple_rest_crud/pkg/logging"
	"simple_rest_crud/pkg/utils"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepo repo.IUserRepository
	config   *config.Config
	logger   *logging.Logger
}

func NewUserService(userRepo repo.IUserRepository, config *config.Config) *UserService {
	return &UserService{
		userRepo: userRepo,
		config:   config,
	}
}

func (service *UserService) GetUserByID(ctx context.Context, userId int) (*entity.UserDetailEntity, error) {
	user, err := service.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) CreateUser(ctx context.Context, userData entity.UserCreateRawEntity) (int, error) {
	saltBytes := make([]byte, service.config.Auth.SaltLenth)
	_, err := rand.Read(saltBytes)
	if err != nil {
		service.logger.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err.Error(),
		}).Error("error generating salt")
		return -1, err
	}

	hashedPassword := utils.HashValue(
		[]byte(userData.RawPassword),
		saltBytes,
	)

	userCreateEntity := entity.UserCreateEntity{
		Username:       userData.Username,
		Email:          userData.Email,
		HashedPassword: hashedPassword,
		PasswordSalt:   fmt.Sprintf("%x", saltBytes),
	}

	userId, err := service.userRepo.CreateUser(ctx, userCreateEntity)

	if err != nil {
		return -1, err
	}
	return userId, nil

}
