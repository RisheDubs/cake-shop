// Command server runs the Cake Shop API.
//
// Locally: go run ./cmd/server, then open http://localhost:8080/v1/cakes
//
// On AWS the same binary runs inside Lambda behind API Gateway, using the
// AWS Lambda Web Adapter layer (see infra/template.yaml). No code changes.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cakeshop/api/internal/cakes"
	"cakeshop/api/internal/httpapi"
	"cakeshop/api/internal/platform/config"
	"cakeshop/api/internal/store"
)

func main() {
	cfg := config.Load()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("env", cfg.Env))

	// Storage: in-memory today; DynamoDB arrives in Sprint 2 behind the same interface.
	repo := store.NewMemory(store.SeedCakes())

	handler := httpapi.NewRouter(httpapi.Deps{
		Cakes:          cakes.NewService(repo),
		AllowedOrigins: cfg.AllowedOrigins,
	})

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		slog.Info("api listening", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server stopped", "err", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown on Ctrl+C (local) or SIGTERM (Lambda/containers).
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	slog.Info("api stopped")
}
