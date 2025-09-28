package log

import (
	"log/slog"
	"os"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
)

// GetLogger 環境に応じた [slog.Logger] を取得する。
func GetLogger(e env.Env) *slog.Logger {
	switch e {
	case env.LCL:
		opts := &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		}
		handler := NewPrettyJSONSlogHandler(os.Stdout, opts)
		return slog.New(handler)
	case env.DEV:
		opts := &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		}
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler)
	case env.STG, env.PRD:
		fallthrough
	default:
		opts := &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelInfo,
			ReplaceAttr: nil,
		}
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler)
	}
}
