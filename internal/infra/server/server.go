package server

import (
	"context"
	"net/http"
	"time"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/port"
	"github.com/go-errors/errors"
)

// Serve HTTPサーバーを起動する
func Serve(ctx context.Context, port port.Port) error {
	s := &http.Server{
		Addr:    ":" + port.String(),
		Handler: newHandler(),
	}

	errCh := make(chan error, 1)

	go func() {
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

	gracefulShutdown(s)

	return nil
}

func gracefulShutdown(srv *http.Server) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
