package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"zway/internal/handler"
)

func App(ctx context.Context) error {
	file,_ := os.Open("logs/log.txt")
	loger := slog.New(slog.NewJSONHandler(file, nil))
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	handler.RegisterRoute(r)

	go func() {
		<-ctx.Done()
		slog.Info("server shutdown!Goodbye :3")
		loger.Info("server shutdown!Goodbye :3")
		s.Shutdown(ctx)
	}()

	slog.Info("Server starting!:3")
	loger.Info("server shutdown!Goodbye :3")

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
