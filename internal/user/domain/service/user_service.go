package service

import (
	"context"
	"fmt"
	"simple_rest_crud/internal/user/domain/entity"
	"simple_rest_crud/internal/user/domain/repo"
)

type UserService struct {
	AuthorRepo repo.IUserRepository
}

func NewUserService(authorRepo repo.IUserRepository) *UserService {
	return &UserService{
		AuthorRepo: authorRepo,
	}
}

func (service *UserService) GetUserByID(ctx context.Context, userId int) (*entity.UserDetailEntity, error) {
	user, err := service.AuthorRepo.GetUserByID(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
