package main

import (
	"log/slog"
	"os"

	"github.com/anoriqq/pj-tmpl-go/cmd/exp/internal"
)

func init() {
	slog.SetDefault(slog.New(internal.NewPrettyJSONSlogHandler(os.Stdout, nil)))
}

func main() {
	slog.Info("start")

	if err := run(); err != nil {
		slog.Error("failed to run", slog.Any("err", err), internal.NewStackTraceSlogAttr(err))
		os.Exit(1)
	}

	slog.Info("end")
}

func run() error {
	cli := internal.NewCLI(os.Stdout, os.Stderr, os.Stdin, os.TempDir())
	opts := internal.CLIOptions{
		Name: "world",
	}

	if err := cli.Run(opts); err != nil {
		return err
	}

	return nil
}
