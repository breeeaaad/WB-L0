package database

import (
	"context"

	"github.com/breeeaaad/WB-L0/internal/repository/cache"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	context context.Context
	conn    *pgx.Conn
	C       *cache.Cache
}

func New(conn *pgx.Conn) *Database {
	return &Database{
		conn:    conn,
		context: context.Background(),
		C:       cache.New(),
	}
}
