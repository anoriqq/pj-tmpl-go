/*
Package main implements the entry point for the application.
*/
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/server"
	"github.com/go-errors/errors"
)

func main() {
	ctx := context.Background()

	if err := eval(ctx, run); err != nil {
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

	if err := server.Serve(ctx, cfg.port); err != nil {
		return err
	}

	return nil
}
