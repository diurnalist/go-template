package main

import (
	"log/slog"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
	))

	log.Info("Hello, World!")
}
