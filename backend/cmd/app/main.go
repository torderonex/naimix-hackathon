package main

import (
	"context"
	"fmt"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/config"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/handler"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/logger"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/service"
	"github.com/fullstakilla/naimix-hackathon/backend/pkg/server"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	// init cfg
	cfg := config.MustLoad()
	// init logger
	slog.SetDefault(logger.MustCreate(cfg.Env))
	slog.Info(fmt.Sprintf("The server starts on port %s", cfg.HTTPServer.Port))
	fmt.Println(cfg)
	//init repo
	store := repo.New(cfg)

	//init services
	services := service.New(store)

	// handler init
	h := handler.New(services)
	// init server
	srv := server.New(cfg.HTTPServer.Port, h.InitRoutes(), cfg.HTTPServer.Timeout)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// run server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	slog.Info("Server exiting")
}
