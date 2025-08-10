package server

import (
	"encoding/json"
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

		type resp struct {
			Status string `json:"status"`
		}
		b, err := json.Marshal(resp{Status: "ok"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	return mux
}
