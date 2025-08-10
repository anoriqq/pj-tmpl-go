/*
Package server provides the HTTP server implementation.
*/
package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/port"
	"github.com/anoriqq/pj-tmpl-go/internal/infra/cli"
	"github.com/go-errors/errors"
)

// Serve HTTPサーバーを起動する
func Serve(ctx context.Context, p port.Port) error {
	s := &http.Server{
		Addr:    ":" + p.String(),
		Handler: newHandler(),
	}

	errCh := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errCh <- cli.ParsePanic(r)
			}
		}()

		slog.Info("starting HTTP server", slog.String("addr", s.Addr))

		err := s.ListenAndServe()
		if err != nil {
			errCh <- errors.Wrap(err, 0)
		}
	}()

	select {
	case <-ctx.Done():
	case err := <-errCh:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	if err := gracefulShutdown(s); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

// nolint:contextcheck
func gracefulShutdown(srv *http.Server) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	slog.Info("shutting down HTTP server", slog.String("addr", srv.Addr))

	err := srv.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
