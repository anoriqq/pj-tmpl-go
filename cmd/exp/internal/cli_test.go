package internal_test

import (
	"bytes"
	"context"
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
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	cli := internal.NewCLI(stdout, stderr, stdin, cwd)
	ctx := context.Background()
	opts := internal.NewCLIOptions("world")

	{
		// Act
		err := cli.Run(ctx, opts)
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
