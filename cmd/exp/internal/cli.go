package internal

import (
	"fmt"
	"io"

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
	stdout    io.Writer
	stderr    io.Writer
	stdin     io.Reader
	cwd       string
	outputDir string
}

func (c *cli) Run(opts cliOptions) error {
	if c == nil {
		return errors.New("cli is nil")
	}

	fmt.Fprintln(c.stdout, "Hello ", opts.Name)
	return nil
}

func NewCLI(stdout, stderr io.Writer, stdin io.Reader, cwd, outputDir string) *cli {
	return &cli{
		stdout:    stdout,
		stderr:    stderr,
		stdin:     stdin,
		cwd:       cwd,
		outputDir: outputDir,
	}
}
