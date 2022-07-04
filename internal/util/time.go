package util

import (
	"fmt"
	"time"

	"github.com/DanielTitkov/orrery/internal/logger"
)

func InfoExecutionTime(start time.Time, name string, logger *logger.Logger) {
	elapsed := time.Since(start)
	logger.Info(fmt.Sprintf("%s exited", name), fmt.Sprintf("%s took %s", name, elapsed))
}

func DebugExecutionTime(start time.Time, name string, logger *logger.Logger) {
	elapsed := time.Since(start)
	logger.Debug(fmt.Sprintf("%s exited", name), fmt.Sprintf("%s took %s", name, elapsed))
}
