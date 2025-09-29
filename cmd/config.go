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
		var c config

		if v, ok := os.LookupEnv("ENV"); ok {
			c.env = env.FromStringZero(v)
		}
		if v, ok := os.LookupEnv("PORT"); ok {
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil
			}
			if i > port.MaxPortValue {
				return nil
			}
			c.port = port.New(uint16(i))
		}

		return &c
	})()
}
