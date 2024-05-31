//go:build tools
// +build tools

package tools

import (
	// uuid

	// connect
	_ "connectrpc.com/connect"

	// postgres driver
	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)
