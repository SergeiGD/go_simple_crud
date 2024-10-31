package real

import (
	"context"
	"errors"
	"simple_rest_crud/internal/user/domain/entity"
	"simple_rest_crud/internal/user/domain/repo"
	"simple_rest_crud/pkg/logging"
	"simple_rest_crud/pkg/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type psqlUserRepository struct {
	client postgres.Client
	logger *logging.Logger
}

// CreateUser implements repo.IUserRepository.
func (repo *psqlUserRepository) CreateUser(ctx context.Context, user entity.UserCreateEntity) (int, error) {
	q := `
		INSERT INTO users (username, email, password, salt) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var userId int
	err := repo.client.QueryRow(
		ctx, q,
		user.Username,
		user.Email,
		user.HashedPassword,
		user.PasswordSalt,
	).Scan(&userId)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			repo.logger.WithFields(logrus.Fields{
				"message": pgErr.Message,
				"detail":  pgErr.Detail,
				"code":    pgErr.Code,
			}).Error("db error on creating user")
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			repo.logger.WithFields(logrus.Fields{
				"message": pgErr.Message,
				"detail":  pgErr.Detail,
				"code":    pgErr.Code,
			}).Error("db error on getting user by id")
			return nil, err
		}

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func NewPsqlUserRepository(client postgres.Client, logger *logging.Logger) repo.IUserRepository {
	return &psqlUserRepository{client: client, logger: logger}
}
