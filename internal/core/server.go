package core

import (
	dbClient "github.com/kefniark/go-web-server/gen/db"
	"github.com/rs/zerolog"
)

type ServerOptions struct {
	Logger *zerolog.Logger
	DB     *dbClient.Queries
}
