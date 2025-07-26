package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
	"github.com/anoriqq/pj-tmpl-go/internal/infra/cli"
	"github.com/anoriqq/pj-tmpl-go/internal/infra/log"
	"github.com/go-errors/errors"
)

func init() {
	// 初期値はLCL
	setupLogger(env.LCL)
}

func main() {
	ctx := context.Background()

	err := run(ctx)
	if err != nil {
		slog.Error("failed to run", slog.Any("err", err), log.NewStackTraceSlogAttr(err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	c := cli.NewCLI(os.Stdout, os.Stderr, os.Stdin, cwd)

	args := os.Args[1:]

	opts, err := cli.NewOptions(args)
	if err != nil {
		return err
	}

	if opts.Help() {
		return nil
	}

	setupLogger(opts.Env())

	slog.Info("start")

	if err := c.Main(ctx, opts); err != nil {
		return err
	}

	slog.Info("end")

	return nil
}

func setupLogger(e env.Env) {
	switch e {
	case env.LCL:
		handler := log.NewPrettyJSONSlogHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		logger := slog.New(handler)
		slog.SetDefault(logger)
	case env.DEV:
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		logger := slog.New(handler)
		slog.SetDefault(logger)
	default:
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		logger := slog.New(handler)
		slog.SetDefault(logger)
	}
}
