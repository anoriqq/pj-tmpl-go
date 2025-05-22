package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/go-errors/errors"
)

func main() {
	slog.Info("start")

	if err := run(); err != nil {
		slog.Error("failed to run", slog.Any("err", err), slogStackTrace(err))
		os.Exit(1)
	}

	slog.Info("end")
}

func run() error {
	fmt.Println("Hello World")
	return nil
}

func slogStackTrace(err error) slog.Attr {
	if err == nil {
		return slog.Any("stacktrace", []any{})
	}

	// go-errors/errors の Error かどうか
	if goerror, ok := err.(*errors.Error); ok {
		return slog.Any("stacktrace", goerror.StackFrames())
	}

	return slog.String("details", fmt.Sprintf("%+v", err))
}
