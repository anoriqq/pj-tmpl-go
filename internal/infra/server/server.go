package server

import (
	"context"
	"net/http"
	"time"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/port"
	"github.com/go-errors/errors"
)

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

func gracefulShutdown(s *http.Server) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
