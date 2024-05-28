package config

import (
	"github.com/kefniark/mango/example/codegen/db"
	"github.com/rs/zerolog"
)

type ServerOptions struct {
	Logger *zerolog.Logger
	DB     *db.Queries
}
