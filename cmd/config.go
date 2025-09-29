package main

import (
	"log/slog"
	"os"
	"strconv"
	"sync"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
	"github.com/anoriqq/pj-tmpl-go/internal/domain/port"
)

type config struct {
	env  env.Env
	port port.Port
}

// LogValue implements the [slog.LogValuer] interface.
func (c config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("env", c.env.String()),
		slog.String("port", c.port.String()),
	)
}

func loadConfig() *config {
	return sync.OnceValue(func() *config {
		var cfg config

		if v, ok := os.LookupEnv("ENV"); ok {
			cfg.env = env.FromStringZero(v)
		}

		if v, ok := os.LookupEnv("PORT"); ok {
			portValue, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil
			}

			if portValue > port.MaxPortValue {
				return nil
			}

			cfg.port = port.New(uint16(portValue))
		}

		return &cfg
	})()
}
