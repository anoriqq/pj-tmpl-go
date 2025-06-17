package internal

import (
	"context"
	"io"
	"log/slog"

	"github.com/go-errors/errors"
)

type cli struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
	cwd    string
}

func (c *cli) Run(ctx context.Context, opts cliOptions) error {
	if c == nil {
		return errors.New("cli is nil")
	}

	slog.Info("working directory", slog.String("cwd", c.cwd))
	slog.Info("running CLI command", slog.Any("options", opts))

	// Check if the context is done before proceeding
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// TODO: Implement the actual logic of the CLI command

	return nil
}

func NewCLI(stdout, stderr io.Writer, stdin io.Reader, cwd string) *cli {
	return &cli{
		stdout: stdout,
		stderr: stderr,
		stdin:  stdin,
		cwd:    cwd,
	}
}
