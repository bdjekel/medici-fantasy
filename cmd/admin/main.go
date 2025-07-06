package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	// "go.uber.org/zap"
)

func main() {
// TODO: Add logger
	// Initialize logger
	// var logger *zap.Logger
	// var err error
	// if *verbose {
	// 	logger, err = zap.NewDevelopment()
	// } else {
	// 	logger, err = zap.NewProduction()
	// }
	// if err != nil {
	// 	log.Fatalf("Failed to initialize logger: %v", err)
	// }
	// defer logger.Sync()

	// logger.Info("Starting admin CLI application")
	fmt.Println("Starting admin CLI application")

	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		// logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
		fmt.Printf("Received shutdown signal: %s\n", sig.String())
		cancel()
	}()

	// TODO: Add configuration loading
	// TODO: Add database connection
	// TODO: Add command handling logic

	// Wait for context cancellation
	<-ctx.Done()
	// logger.Info("Shutting down admin CLI application")
	fmt.Println("Shutting down admin CLI application")
}

