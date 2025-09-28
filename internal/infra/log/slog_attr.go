package log

import (
	"fmt"
	"log/slog"

	"github.com/go-errors/errors"
)

// NewStackTraceSlogAttr [github.com/go-errors/errors.Error] のエラーからスタックトレースを取得し、
// [slog.Attr] として返す。
func NewStackTraceSlogAttr(err error) slog.Attr {
	if err == nil {
		return slog.Any("stacktrace", []any{})
	}

	// go-errors/errors の Error かどうか
	var goerror *errors.Error
	if errors.As(err, &goerror) {
		return slog.Any("stacktrace", goerror.StackFrames())
	}

	return slog.String("details", fmt.Sprintf("%+v", err))
}
