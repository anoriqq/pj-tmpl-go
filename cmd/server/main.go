package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/anoriqq/pj-tmpl-go/cmd/server/internal/cli"
	"github.com/anoriqq/pj-tmpl-go/cmd/server/internal/log"
)

func init() {
	slog.SetDefault(slog.New(log.NewPrettyJSONSlogHandler(os.Stdout, nil)))
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("failed to run", slog.Any("err", err), log.NewStackTraceSlogAttr(err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	c := cli.NewCLI(os.Stdout, os.Stderr, os.Stdin, cwd)
	opts := cli.NewOptions()

	slog.Info("start")
	if err := c.Run(ctx, opts); err != nil {
		return err
	}
	slog.Info("end")

	return nil
}
