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

	// Shutdown the server gracefully
	{
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.Shutdown(ctx)
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}
