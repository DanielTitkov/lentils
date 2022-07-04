package job

import (
	"github.com/DanielTitkov/orrery/internal/app"
	"github.com/DanielTitkov/orrery/internal/configs"
	"github.com/DanielTitkov/orrery/internal/logger"
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
