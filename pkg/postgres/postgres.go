package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"simple_rest_crud/config"
	"simple_rest_crud/pkg/logging"
	"simple_rest_crud/pkg/utils"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, logger *logging.Logger, pgConf config.PgConfig) (*pgxpool.Pool, error) {
	dns := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", pgConf.Username, pgConf.Password, pgConf.Host, pgConf.Port, pgConf.Database)
	logger.WithFields(logrus.Fields{
		"dns": dns,
	}).Info("trying to connect to psql")

	var pool *pgxpool.Pool
	var err error

	err = utils.DoWithAttemps(func() error {

		ctx, cancel := context.WithTimeout(ctx, time.Duration(pgConf.Timeount)*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dns)
		if err != nil {
			return err
		}

		return nil

	}, pgConf.MaxAttemps, time.Duration(pgConf.ConnDelay)*time.Second)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("error connecting to DB")
		return nil, err
	}

	return pool, nil

}
