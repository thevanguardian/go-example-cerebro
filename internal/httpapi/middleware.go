package httpapi

import (
	"log/slog"
	"net/http"
	"time"
)

// common behavior for wrapping additional behavior into handlers while
// maintaining the http.Handler interface for compatibility
// accepts a variable number of middleware functions in reverse order (known as variadic functions)
type middleware func(http.Handler) http.Handler

// Using the ellipsis to accept a variable number of middleware functions
// we iterate over them in reverse order to ensure the first middleware
// wraps the entire chain, preserving the intended order of execution
func middlewareChain(h http.Handler, m ...middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		// wrap the current handler with the middleware function
		h = m[i](h)
	}
	return h
}

// Log request details such as method, path, duration, etc. using standard slog package
// get's wrapped up in middleware chain to be applied to all requests, a simple and straight-forward
// unified logging approach
func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // call the next handler in the chain
		slog.Info(
			"request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}

// handle panic recovery to prevent server crashes
// if a panic occurs in any handler, this middleware will catch it
// log the error and return a 500 Internal Server Error response
func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// deferred anonymous function; executes just after the surrounding function exits
		// doesn't really need a name since it's inline, so we'll keep it anonymous
		defer func() {
			if v := recover(); v != nil {
				slog.Error("panic recovered", "value", v)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
