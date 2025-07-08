package server

import (
	"io"
	"log/slog"
	"net/http"
)

func newHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "received request",
			slog.String("user-agent", r.UserAgent()),
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()))

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello, World!\n")
	})

	return mux
}
