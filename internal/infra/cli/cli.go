package cli

import (
	"context"
	"io"
	"log/slog"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/server"
	"github.com/go-errors/errors"
)

type cli struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
	cwd    string
}

func (c *cli) Main(ctx context.Context, opts options) error {
	if c == nil {
		return errors.New("cli is nil")
	}

	slog.Info("working directory", slog.String("cwd", c.cwd))
	slog.Info("running CLI", slog.Any("options", opts))

	// Check if the context is done before proceeding
	select {
	case <-ctx.Done():
		return errors.Wrap(context.Cause(ctx), 0)
	default:
	}

	if err := server.Serve(ctx, opts.Port()); err != nil {
		return err
	}

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
