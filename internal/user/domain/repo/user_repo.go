package repo

import (
	"context"
	"simple_rest_crud/internal/user/domain/entity"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.UserCreateEntity) (int, error)
	GetUserByID(ctx context.Context, userId int) (*entity.UserDetailEntity, error)
}
