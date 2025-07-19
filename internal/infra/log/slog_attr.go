package log

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
	goerror := &errors.Error{}
	if errors.As(err, &goerror) {
		return slog.Any("stacktrace", goerror.StackFrames())
	}

	return slog.String("details", fmt.Sprintf("%+v", err))
}
