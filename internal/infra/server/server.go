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
	"github.com/anoriqq/pj-tmpl-go/internal/infra/pnc"
	"github.com/go-errors/errors"
)

// Serve HTTPサーバーを起動する。
func Serve(ctx context.Context, p port.Port) error {
	srv := &http.Server{
		Addr:                         ":" + p.String(),
		Handler:                      newHandler(),
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}

	errCh := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errCh <- pnc.Parse(r)
			}
		}()

		slog.Info("starting HTTP server", slog.String("addr", srv.Addr))

		err := srv.ListenAndServe()
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

	//nolint:contextcheck // 親ctxはすでにcancel済みなので、新しいctxを作成して使う
	err := gracefulShutdown(srv)
	if err != nil {
		return err
	}

	return nil
}

const gracefulShutdownTimeout = 10 * time.Second

func gracefulShutdown(srv *http.Server) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, gracefulShutdownTimeout)
	defer cancel()

	slog.Info("shutting down HTTP server", slog.String("addr", srv.Addr))

	err := srv.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
