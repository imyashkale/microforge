package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imyashkale/microforge/pkg/configs"
	"github.com/imyashkale/microforge/pkg/dynamodb"
	"github.com/imyashkale/microforge/pkg/log"
)

type Application struct {
	router    *gin.Engine
	logger    log.Logger
	dyanamodb *dynamodb.DyanamoDB
	config    *configs.Config
}

func (app Application) Run() {

	addr := fmt.Sprintf(":%d", app.config.Application.Port)

	// Create the server
	server := &http.Server{
		Addr:    addr,
		Handler: app.router,
	}

	fmt.Printf("Starting application...\n")
	fmt.Printf("Version			: %s\n", app.config.Application.Version)
	fmt.Printf("Server Port		: %d\n", app.config.Application.Port)
	fmt.Printf("Config File Path: %s\n", app.config.AppConfigPath)

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Infof("listen: %s\n", err)
		}
	}()

	// Create a channel to listen for signals
	quit := make(chan os.Signal, 1)

	// Listen for SIGINT and SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	app.logger.Info("shuting down server ...")

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		app.logger.Info("Server forced to shutdown: %v", err)
	}
}
