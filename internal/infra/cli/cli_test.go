package cli_test

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"testing"
	"time"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/cli"
	"github.com/tenntenn/golden"
)

var flagUpdate bool

func init() {
	flag.BoolVar(&flagUpdate, "update", false, "update golden files")
}

func TestCLI_Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	stdin := new(bytes.Buffer)
	tempDir := t.TempDir()

	sut := cli.NewCLI(stdout, stderr, stdin, tempDir)
	args := []string{"-env", "lcl"}
	opts, err := cli.NewOptions(args)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	err = sut.Run(ctx, opts)

	//' Assert
	if !errors.Is(err, nil) {
		t.Fatalf("Run failed: %v", err)
	}

	c := golden.New(t, flagUpdate, "testdata", "TestCLI_Run")
	if diff := c.Check("_stdout", stdout); diff != "" {
		t.Error("stdout\n", diff)
	}
	if diff := c.Check("_stderr", stderr); diff != "" {
		t.Error("stderr\n", diff)
	}
}
