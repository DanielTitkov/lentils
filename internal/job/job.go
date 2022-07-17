package job

import (
	"time"

	"go.uber.org/multierr"

	"github.com/tinygodsdev/orrery/internal/app"
	"github.com/tinygodsdev/orrery/internal/configs"
	"github.com/tinygodsdev/orrery/internal/logger"
	"github.com/go-co-op/gocron"
)

type Job struct {
	cfg       configs.Config
	logger    *logger.Logger
	app       *app.App
	scheduler *gocron.Scheduler
}

func New(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) (*Job, error) {
	s := gocron.NewScheduler(time.UTC)
	s.SetMaxConcurrentJobs(1, gocron.WaitMode)
	s.SingletonModeAll()
	// s.WaitForScheduleAll()

	job := &Job{
		cfg:       cfg,
		logger:    logger,
		app:       app,
		scheduler: s,
	}

	err := job.scheduleTasks()
	if err != nil {
		logger.Error("errors while scheduling tasks", err)
		app.AddError(err)
		if app.IsDev() {
			return nil, err
		}
	}

	return job, nil
}

func (j *Job) Run() {
	j.scheduler.StartAsync()
	j.scheduler.RunAll()
}

func (j *Job) scheduleTasks() (errs error) {
	// update norms
	_, err := j.scheduler.Cron(j.cfg.App.UpdateNormsSchedule).Tag("norms").Do(j.UpdateNorms)
	if err != nil {
		errs = multierr.Append(errs, err)
	}
	// update test marks
	_, err = j.scheduler.Cron(j.cfg.App.UpdateMarksSchedule).Tag("marks").Do(j.UpdateMarks)
	if err != nil {
		errs = multierr.Append(errs, err)
	}
	// update system summary
	_, err = j.scheduler.Cron(j.cfg.App.SystemSummarySchedule).Tag("summary").Do(j.UpdateSystemSummary)
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	return errs
}
