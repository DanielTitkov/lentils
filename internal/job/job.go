package job

import (
	"github.com/DanielTitkov/lentils/internal/app"
	"github.com/DanielTitkov/lentils/internal/configs"
	"github.com/DanielTitkov/lentils/internal/logger"
)

type Job struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func New(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Job {
	return &Job{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}
