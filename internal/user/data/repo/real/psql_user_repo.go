package real

import (
	"context"
	"fmt"
	"simple_rest_crud/internal/user/domain/entity"
	"simple_rest_crud/internal/user/domain/repo"
	"simple_rest_crud/pkg/postgres"

	"github.com/jackc/pgx/v5/pgconn"
)

type psqlUserRepository struct {
	client postgres.Client
}

// CreateUser implements repo.IUserRepository.
func (repo *psqlUserRepository) CreateUser(ctx context.Context, user entity.UserCreateEntity) (int, error) {
	q := `
		INSERT INTO users (username, email, password) 
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var userId int
	err := repo.client.QueryRow(
		ctx, q,
		user.Username,
		user.Email,
		user.HashedPassword,
	).Scan(&userId)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// TODO: logger
			fmt.Println("DB error on creating user", pgErr.Message, pgErr.Detail, pgErr.Code)
			return -1, pgErr
		}
		return -1, err
	}

	return userId, nil
}

// GetUserByID implements repo.IUserRepository.
func (repo *psqlUserRepository) GetUserByID(ctx context.Context, userId int) (*entity.UserDetailEntity, error) {
	q := `
		SELECT id, username, email FROM users
		WHERE id = $1
	`
	user := &entity.UserDetailEntity{}

	err := repo.client.QueryRow(ctx, q, userId).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
	)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// TODO: logger
			fmt.Println("DB error on getting user by id", pgErr.Message, pgErr.Detail, pgErr.Code, userId)
			return nil, pgErr
		}
		return nil, err
	}

	return user, nil
}

func NewPsqlUserRepository(client postgres.Client) repo.IUserRepository {
	return &psqlUserRepository{client: client}
}
