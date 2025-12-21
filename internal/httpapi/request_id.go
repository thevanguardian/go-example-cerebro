package httpapi

import "context"

type ctxKeyRequestID struct{}

// we want the following to be exported for use in logging middleware
// so the functions will be capitalized appropriately

// WithRequestID returns a new context with the given request ID.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ctxKeyRequestID{}, requestID)
}

func RequestID(ctx context.Context) string {
	v := ctx.Value(ctxKeyRequestID{})
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
