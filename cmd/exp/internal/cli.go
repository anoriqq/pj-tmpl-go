package internal

import (
	"fmt"
	"io"
)

type cli struct {
	stdout    io.Writer
	stderr    io.Writer
	stdin     io.Reader
	cwd       string
	outputDir string
}

type CLIOptions struct {
	Name string
}

func (c *cli) Run(opts CLIOptions) error {
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
