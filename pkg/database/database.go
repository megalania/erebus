package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Config pgxpool.Config
	Pool   *pgxpool.Pool
}

func NewPool(str string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(
		ctx,
		str,
	)
	if err != nil {
		return nil, err
	}
	return pool, err
}
