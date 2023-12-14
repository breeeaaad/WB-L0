package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	context.Context
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Database {
	return &Database{
		conn: conn,
	}
}
