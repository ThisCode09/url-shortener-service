package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"url-shortener/internal/handlers"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"
)

func main() {
	store, err := storage.NewJSONStore("urls.json")
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	svc := shortener.NewService(store)

	mux := http.NewServeMux()
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("GET /shorten/{url...}", handlers.HandleShorten(svc))
	mux.HandleFunc("GET /{code}", handlers.HandleRedirect(svc))

	go func() {
		slog.Info("starting server.....")
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "err", err)
			os.Exit(1)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

		<-sigChan
		slog.Info("shutting down server")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error("error shutting down HTTP server", "err", err)
		}
	}()

	wg.Wait()
	slog.Info("server shutdown complete")
}
