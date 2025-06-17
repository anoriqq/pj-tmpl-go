package internal

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/go-errors/errors"
)

type cliOptions struct {
	Name string
}

func NewCLIOptions(name string) cliOptions {
	return cliOptions{
		Name: name,
	}
}

type cli struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
	cwd    string
}

func (c *cli) Run(opts cliOptions) error {
	if c == nil {
		return errors.New("cli is nil")
	}

	slog.Info("working directory", slog.String("cwd", c.cwd))

	fmt.Fprintln(c.stdout, "Hello ", opts.Name)
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
