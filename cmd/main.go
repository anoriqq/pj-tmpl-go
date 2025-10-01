/*
Package main implements the entry point for the application.
*/
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-errors/errors"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/server"
)

func main() {
	ctx := context.Background()

	err := eval(ctx, run)
	if err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	slog.Debug("current working directory", slog.String("path", cwd))

	cfg := loadConfig()

	err = server.Serve(ctx, cfg.port)
	if err != nil {
		return err
	}

	return nil
}
