package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/thevanguardian/go-example-cerebro/internal/httpapi"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg := httpapi.NewServer(httpapi.Config{
		Addr: ":8080",
	})

	go func() {
		slog.Info("Cerebro Sentinel starting", "addr", cfg.Addr)
		if err := cfg.ListenAndServe(); err != nil {
			slog.Error("server error", "err", err)

		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	slog.Info("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := cfg.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown error", "err", err)
	}
}
