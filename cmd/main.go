/*
Package main implements the entry point for the application.
*/
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/cli"
	"github.com/anoriqq/pj-tmpl-go/internal/infra/log"
	"github.com/anoriqq/pj-tmpl-go/internal/infra/server"
	"github.com/go-errors/errors"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error(
			"failed to run",
			slog.Any("err", err), log.NewStackTraceSlogAttr(err),
		)
		os.Exit(1)
	}
}

func run(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = cli.ParsePanic(r)
		}
	}()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	env, port, err := cli.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	handler := log.GetLogger(env)
	slog.SetDefault(handler)

	slog.Debug("env", slog.String("env", env.String()))
	slog.Debug("port", slog.String("port", port.String()))

	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, 0)
	}
	slog.Debug("current working directory", slog.String("path", cwd))

	if err := server.Serve(ctx, port); err != nil {
		return err
	}

	return nil
}
