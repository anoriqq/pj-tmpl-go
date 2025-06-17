package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/anoriqq/pj-tmpl-go/cmd/exp/internal"
)

func init() {
	slog.SetDefault(slog.New(internal.NewPrettyJSONSlogHandler(os.Stdout, nil)))
}

func main() {
	slog.Info("start")

	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("failed to run", slog.Any("err", err), internal.NewStackTraceSlogAttr(err))
		os.Exit(1)
	}

	slog.Info("end")
}

func run(ctx context.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cli := internal.NewCLI(os.Stdout, os.Stderr, os.Stdin, cwd)
	opts := internal.NewCLIOptions("world")

	if err := cli.Run(ctx, opts); err != nil {
		return err
	}

	return nil
}
