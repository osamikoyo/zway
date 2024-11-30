package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"zway/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.App(ctx); err != nil {
		slog.Error(err.Error())
	}
}
