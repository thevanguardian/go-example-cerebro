package httpapi

import "net/http"

// simple function that links routes to their handlers
// request that comes in for /healthz will be served by healthHandler for example
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", healthHandler)
}
