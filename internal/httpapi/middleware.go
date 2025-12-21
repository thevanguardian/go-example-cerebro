package httpapi

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"strings"
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
			"request_id", RequestID(r.Context()),
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

// assign a unique request ID to each incoming HTTP request
// useful for tracing and debugging, especially in distributed systems
const requestIDHeader = "X-Request-Id"

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(r.Header.Get(requestIDHeader))
		if id == "" {
			id = newRequestID()
		}

		w.Header().Set(requestIDHeader, id)
		ctx := WithRequestID(r.Context(), id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func newRequestID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		slog.Error("failed to generate request id", "err", err)
		return hex.EncodeToString([]byte(time.Now().Format(time.RFC3339Nano)))
	}
	return hex.EncodeToString(b[:])
}
