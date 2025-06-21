package internal

import (
	"fmt"
	"log/slog"

	"github.com/go-errors/errors"
)


func NewStackTraceSlogAttr(err error) slog.Attr {
	if err == nil {
		return slog.Any("stacktrace", []any{})
	}

	// go-errors/errors の Error かどうか
	if goerror, ok := err.(*errors.Error); ok {
		return slog.Any("stacktrace", goerror.StackFrames())
	}

	return slog.String("details", fmt.Sprintf("%+v", err))
}
