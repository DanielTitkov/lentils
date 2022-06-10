package logger

import (
	"time"

	"go.uber.org/zap"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(env string) *Logger {
	var logger *zap.Logger
	switch env {
	case "dev":
		logger, _ = zap.NewDevelopment()
	default:
		logger, _ = zap.NewProduction()
	}
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(msg, info string) {
	l.logger.Debug(msg,
		zap.String("info", info),
	)
}

func (l *Logger) Info(msg, info string) {
	l.logger.Info(msg,
		zap.String("time", getTime()),
		zap.String("info", info),
	)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal(msg,
		zap.String("time", getTime()),
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Error(msg,
		zap.String("time", getTime()),
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Sync() {
	_ = l.logger.Sync()
}

func getTime() string {
	return time.Now().Format(time.RFC3339Nano)
}
