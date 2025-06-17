package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/anoriqq/pj-tmpl-go/cmd/exp/internal"
)

func init() {
	slog.SetDefault(slog.New(internal.NewPrettyJSONSlogHandler(os.Stdout, nil)))
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("failed to run", slog.Any("err", err), internal.NewStackTraceSlogAttr(err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	cli := internal.NewCLI(os.Stdout, os.Stderr, os.Stdin, cwd)
	opts := internal.NewCLIOptions()

	slog.Info("start")
	if err := cli.Run(ctx, opts); err != nil {
		return err
	}
	slog.Info("end")

	return nil
}
