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

// Fairly standard multiplexer for setting up a basic http server
// allowing routing of requests to specific application pieces.
// Will also be the basis for the basic healthz check which is standard
// across many orchestration tools as healthcheck endpoints
func NewServer(cfg Config) *http.Server {
	mux := http.NewServeMux()
	// grabs routes and links them to their handlers
	registerRoutes(mux)
	// add new middleware functions here
	handler := middlewareChain(
		mux,
		requestID,
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
