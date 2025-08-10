package log

import (
	"log/slog"
	"os"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
)

func init() {
	var e env.Env
	slog.SetDefault(GetLogger(e))
}

// GetLogger 環境に応じた[slog.Logger]を取得する
func GetLogger(e env.Env) *slog.Logger {
	switch e {
	case env.LCL:
		handler := NewPrettyJSONSlogHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		return slog.New(handler)
	case env.DEV:
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		return slog.New(handler)
	case env.STG, env.PRD:
		fallthrough
	default:
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		return slog.New(handler)
	}
}
