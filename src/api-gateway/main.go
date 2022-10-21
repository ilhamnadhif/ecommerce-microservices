package main

import (
	"api-gateway/app"
	"api-gateway/config"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	service = "api-gateway"
	version = "latest"
)

func main() {
	// log
	file, _ := os.OpenFile(fmt.Sprintf("logs/app_%s.log", time.Now().Format("2006_01_02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// setup config
	config.InitConfig()

	// route
	route := app.Route()

	// Start server
	go func() {
		if err := route.Start(config.Config.Server.HostPort); err != nil && err != http.ErrServerClosed {
			route.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := route.Shutdown(ctx); err != nil {
		route.Logger.Fatal(err)
	}
}
