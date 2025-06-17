package internal_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/anoriqq/pj-tmpl-go/cmd/exp/internal"
	"github.com/tenntenn/golden"
)

func TestCLI_Run(t *testing.T) {
	t.Parallel()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	stdin := new(bytes.Buffer)
	outputDir := t.TempDir()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	cli := internal.NewCLI(stdout, stderr, stdin, cwd, outputDir)
	opts := internal.NewCLIOptions("world")

	{
		// Act
		err := cli.Run(opts)
		// Assert
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		c := golden.New(t, true, "testdata", "TestCLI_Run")
		if diff := c.Check("_stdout", stdout); diff != "" {
			t.Error("stdout\n", diff)
		}
		if diff := c.Check("_stderr", stderr); diff != "" {
			t.Error("stderr\n", diff)
		}
	}
}
