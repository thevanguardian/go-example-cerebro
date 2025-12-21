package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDGenerated(t *testing.T) {
	handler := requestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id := RequestID(r.Context()); id == "" {
			t.Fatal("expected request id to be set")
		}
	}))

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Header().Get(requestIDHeader) == "" {
		t.Fatal("expected X-Request-Id header to be set")
	}
}
