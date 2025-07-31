package cli

import (
	"context"
	"io"
	"log/slog"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/server"
	"github.com/go-errors/errors"
)

// CLI Command Line Interface
type CLI struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
	cwd    string
}

// Main メインの処理を実行する
func (c *CLI) Main(ctx context.Context, opts Options) error {
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

	err := server.Serve(ctx, opts.Port())
	if err != nil {
		return err
	}

	return nil
}

// NewCLI cliを作成する
func NewCLI(stdout, stderr io.Writer, stdin io.Reader, cwd string) *CLI {
	return &CLI{
		stdout: stdout,
		stderr: stderr,
		stdin:  stdin,
		cwd:    cwd,
	}
}
