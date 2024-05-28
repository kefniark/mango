package config

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

func init() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.InfoLevel)
	Logger = &log
}
