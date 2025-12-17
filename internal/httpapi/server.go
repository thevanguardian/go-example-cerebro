package httpapi

import (
	"net/http"
	"time"
)

// We want Config to have a fixed record
// Since struct are static and ordered, they're less memory and compute intensive than the unordered and dynamic setup of a map.
type Config struct {
	Addr string
}

func NewServer(cfg Config) *http.Server {
	// Setup
	mux := http.NewServeMux()
	registerRoutes(mux)

	handler := middlewareChain(
		mux,
		requestLogger,
		recoverPanic,
	)

	return &http.Server{
		Addr:              cfg.Addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
