package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Shauryan10/GoFlow/config"
	"github.com/Shauryan10/GoFlow/internal/api"
	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/Shauryan10/GoFlow/internal/metrics"
	"github.com/Shauryan10/GoFlow/internal/queue"
	"github.com/Shauryan10/GoFlow/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	cfg := config.LoadConfig()

	fmt.Println("Database Host :", cfg.DBHost)
	fmt.Println("Database Port :", cfg.DBPort)
	fmt.Println("Application Port :", cfg.AppPort)

	// Connect to PostgreSQL
	database.ConnectDatabase(cfg)

	// Register Prometheus metrics
	metrics.Register()

	// Initialize task queue
	queue.InitQueue(100)

	// Create shutdown context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// WaitGroup keeps track of workers
	var wg sync.WaitGroup

	// Start worker pool
	worker.StartWorkerPool(ctx, 3, &wg)

	router := gin.Default()

	// Prometheus metrics endpoint
	router.GET(
		"/metrics",
		gin.WrapH(promhttp.Handler()),
	)

	// Register API routes
	api.RegisterRoutes(router)

	fmt.Println("🚀 GoFlow API is running on port", cfg.AppPort)

	// Start server in a goroutine
	go func() {

		err := router.Run(":" + cfg.AppPort)

		if err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// Listen for Ctrl+C / SIGTERM
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(
		signalChannel,
		os.Interrupt,
		syscall.SIGTERM,
	)

	// Wait until shutdown signal arrives
	<-signalChannel

	fmt.Println("\nShutdown signal received...")

	// Tell workers to stop
	cancel()

	// Wait for workers to finish
	wg.Wait()

	fmt.Println("All workers stopped")
	fmt.Println("GoFlow shutdown complete")
}
