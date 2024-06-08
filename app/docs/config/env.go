package config

import (
	"fmt"
	"os"
)

type AppModeEnum int // defined type
const (
	PROD AppModeEnum = iota
	DEV
)

func AppMode() AppModeEnum {
	env := ""
	if val, ok := os.LookupEnv("ENV"); ok {
		env = val
	}

	switch env {
	case "DEV":
		return DEV
	default:
		return PROD
	}
}

func IsDev() bool {
	return AppMode() == DEV
}

func AppAddr() string {
	addr := ":5600"
	if port, ok := os.LookupEnv("PORT"); ok {
		addr = fmt.Sprintf(":%s", port)
	}
	return addr
}

func AppPublicAddr() string {
	app := ""
	proxy := ""

	if port, ok := os.LookupEnv("PORT"); ok {
		app = port
	}
	if port, ok := os.LookupEnv("PROXY_PORT"); ok {
		proxy = port
	}

	if proxy != "" {
		return fmt.Sprintf("%s -> %s (Proxied)", proxy, app)
	}

	return app
}
