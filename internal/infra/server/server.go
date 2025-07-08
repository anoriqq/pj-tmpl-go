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
		if err := s.ListenAndServe(); err != nil {
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

	// Shutdown the server gracefully
	{
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}
