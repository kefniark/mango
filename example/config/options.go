package config

import (
	"github.com/kefniark/go-web-server/example/codegen/db"
	"github.com/rs/zerolog"
)

type ServerOptions struct {
	Logger *zerolog.Logger
	DB     *db.Queries
}
