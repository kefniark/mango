package db

import (
	"context"
	"database/sql"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/kefniark/mango/example/codegen/database"

	// embed database schema.
	_ "embed"
)

//go:embed schema.sql
var dbSchema string

type DB struct {
	db *sql.DB
}

func (db *DB) Migrate() error {
	_, err := db.db.ExecContext(context.Background(), dbSchema)
	return err
}

func (db *DB) Client() *database.Queries {
	return database.New(db.db)
}

func New() (*DB, error) {
	url := "postgres://postgres:password@localhost:5432/mangodb"
	if databaseURL, ok := os.LookupEnv("DATABASE_URL"); ok {
		url = databaseURL
	}

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: stdlib.OpenDBFromPool(pool),
	}, nil
}
