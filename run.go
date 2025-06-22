package pjtmplgo

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-errors/errors"
)

type options struct{}

func NewOptions() (options, error) {
	return options{}, nil
}

func Run(ctx context.Context, opts options) error {
	// API server example

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "Received request",
			slog.String("user-agent", r.UserAgent()),
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()))
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello, World!\n")
	})

	s := &http.Server{
		Addr:    ":8888",
		Handler: mux,
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

	{
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}
