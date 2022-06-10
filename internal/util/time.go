package util

import (
	"fmt"
	"time"

	"github.com/DanielTitkov/lentils/internal/logger"
)

func LogExecutionTime(start time.Time, name string, logger *logger.Logger) {
	elapsed := time.Since(start)
	logger.Info(fmt.Sprintf("%s exited", name), fmt.Sprintf("%s took %s", name, elapsed))
}
