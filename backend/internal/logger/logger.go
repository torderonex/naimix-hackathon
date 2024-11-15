package logger

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/config"
	"log"
	"log/slog"
	"os"
)

func MustCreate(environment string) *slog.Logger {
	var logger *slog.Logger
	switch environment {
	case config.EnvLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log.Fatalf("logger error: unknown environment '%s'", environment)

	}

	return logger
}
