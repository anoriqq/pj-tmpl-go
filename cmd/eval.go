package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/log"
	"github.com/anoriqq/pj-tmpl-go/internal/infra/pnc"
	"github.com/go-errors/errors"
)

// eval アプリケーションのライフサイクルを管理し、エラーハンドリングとロギングを行う。
// 基本的に [main] 関数から1度だけ呼び出されるものであり、複数回の呼び出しを想定していない。
func eval(ctx context.Context, run func(context.Context) error) (err error) {
	defer func() {
		if err != nil {
			slog.Error(
				"failed to run",
				slog.Any("err", err),
				log.NewStackTraceSlogAttr(err),
			)
		}
		slog.Info("exiting")
	}()

	defer func() {
		if r := recover(); r != nil {
			err = pnc.Parse(r)
		}
	}()

	cfg := loadConfig()
	logger := log.GetLogger(cfg.env)
	slog.SetDefault(logger)
	slog.Info("starting")
	slog.Info("load config", slog.Any("config", cfg))

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	if err = run(ctx); err != nil {
		return errors.Wrap(err, 0)
	}

	slog.Info("shutting down")

	return nil
}
