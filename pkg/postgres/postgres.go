package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"simple_rest_crud/config"
	"simple_rest_crud/pkg/utils"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, pgConf config.PgConfig) (*pgxpool.Pool, error) {
	dns := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", pgConf.Username, pgConf.Password, pgConf.Host, pgConf.Port, pgConf.Database)
	fmt.Println(dns)

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
		// TODO: logger
		return nil, fmt.Errorf("error connecting to DB")
	}

	return pool, nil

}
