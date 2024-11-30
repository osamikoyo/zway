package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"zway/internal/handler"
)

func App(ctx context.Context) error {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	handler.RegisterRoute(r)

	go func() {
		<-ctx.Done()
		slog.Info("server shutdown!")
		s.Shutdown(ctx)
	}()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
