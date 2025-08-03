package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/diurnalist/go-template/internal/server"
)

const defaultPort = 8080

func main() {
	// Set up structured logging.
	log := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:       slog.LevelInfo,
			AddSource:   true,
			ReplaceAttr: nil,
		},
	))

	// Create context that can be cancelled.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for graceful shutdown.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create and start the gRPC server.
	grpcServer := server.NewGRPCServer(defaultPort)

	// Start the server in a goroutine.
	go func() {
		log.InfoContext(ctx, "Starting gRPC server", "port", defaultPort)
		if err := grpcServer.Start(ctx); err != nil {
			log.ErrorContext(ctx, "Failed to start gRPC server", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for shutdown signal.
	<-sigChan
	log.InfoContext(ctx, "Received shutdown signal, stopping server")

	// Gracefully stop the server.
	grpcServer.Stop()
	log.InfoContext(ctx, "Server stopped gracefully")
}
