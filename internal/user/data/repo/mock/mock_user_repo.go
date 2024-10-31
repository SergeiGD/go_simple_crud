package mock

import (
	"context"
	"simple_rest_crud/internal/user/domain/entity"
	"simple_rest_crud/internal/user/domain/repo"
)

type mockUserRepository struct {
}

// CreateUser implements repo.IUserRepository.
func (m *mockUserRepository) CreateUser(ctx context.Context, user entity.UserCreateEntity) (int, error) {
	return 1, nil
}

// GetUserByID implements repo.IUserRepository.
func (m *mockUserRepository) GetUserByID(ctx context.Context, userId int) (*entity.UserDetailEntity, error) {
	user := &entity.UserDetailEntity{
		Email:    "mock@gmail.com",
		Username: "mockusername",
		ID:       1,
	}
	return user, nil
}

func NewMockUserRepository() repo.IUserRepository {
	return &mockUserRepository{}
}
