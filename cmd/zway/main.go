package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"zway/internal/app"
)

func main() {
	file,_ := os.Open("logs/log.txt")
	loger := slog.New(slog.NewJSONHandler(file, nil))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.App(ctx); err != nil {
		loger.Error(err.Error())
		slog.Error(err.Error())
	}
}
