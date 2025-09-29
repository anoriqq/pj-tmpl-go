package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"sync"

	"github.com/go-errors/errors"
)

const (
	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

// Handler ログレコードをきれいなJSONとしてフォーマットするカスタムslogハンドラー。
type Handler struct {
	w io.Writer
	h slog.Handler
	b *bytes.Buffer
	m *sync.Mutex
}

var _ slog.Handler = (*Handler)(nil)

// NewPrettyJSONSlogHandler きれいなJSON形式でログを出力するslogハンドラーを作成する。
func NewPrettyJSONSlogHandler(w io.Writer, opts *slog.HandlerOptions) *Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{
			AddSource:   false,
			Level:       nil,
			ReplaceAttr: nil,
		}
	}

	buf := &bytes.Buffer{}

	return &Handler{
		w: w,
		h: slog.NewJSONHandler(buf, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: suppressDefaults(opts.ReplaceAttr),
		}),
		b: buf,
		m: &sync.Mutex{},
	}
}

// Enabled implements [slog.Handler] interface.
func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

const (
	timeFormat = "[15:04:05.000]"
)

// Handle implements [slog.Handler] interface.
func (h *Handler) Handle(ctx context.Context, rec slog.Record) error {
	level := rec.Level.String() + ":"

	switch rec.Level {
	case slog.LevelDebug:
		level = colorize(darkGray, level)
	case slog.LevelInfo:
		level = colorize(cyan, level)
	case slog.LevelWarn:
		level = colorize(lightYellow, level)
	case slog.LevelError:
		level = colorize(lightRed, level)
	}

	attrs, err := h.computeAttrs(ctx, rec)
	if err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("error when marshaling attrs: %w", err)
	}

	if _, err := fmt.Fprintln(
		h.w,
		colorize(lightGray, rec.Time.Format(timeFormat)),
		level,
		colorize(white, rec.Message),
		colorize(darkGray, string(jsonBytes)),
	); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

// WithAttrs implements [slog.Handler] interface.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		w: h.w,
		h: h.h.WithAttrs(attrs),
		b: h.b,
		m: h.m,
	}
}

// WithGroup implements [slog.Handler] interface.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		w: h.w,
		h: h.h.WithGroup(name),
		b: h.b,
		m: h.m,
	}
}

func (h *Handler) computeAttrs(
	ctx context.Context,
	rec slog.Record,
) (map[string]any, error) {
	h.m.Lock()
	defer h.m.Unlock()

	defer h.b.Reset()

	if err := h.h.Handle(ctx, rec); err != nil {
		return nil, fmt.Errorf("error when calling inner handler's Handle: %w", err)
	}

	var attrs map[string]any

	err := json.Unmarshal(h.b.Bytes(), &attrs)
	if err != nil {
		return nil, fmt.Errorf("error when unmarshaling inner handler's Handle result: %w", err)
	}

	return attrs, nil
}

func suppressDefaults(
	next func([]string, slog.Attr) slog.Attr,
) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, attr slog.Attr) slog.Attr {
		if attr.Key == slog.TimeKey ||
			attr.Key == slog.LevelKey ||
			attr.Key == slog.MessageKey {
			var zero slog.Attr

			return zero
		}

		if next == nil {
			return attr
		}

		return next(groups, attr)
	}
}
