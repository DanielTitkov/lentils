package job

import (
	"context"
	"errors"
	"time"
)

func (j *Job) UpdateMarks() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(j.cfg.App.UpdateMarksTimeout))
	processDone := make(chan bool)
	go func() {
		err := j.app.UpdateTestMarks(ctx)
		if err != nil {
			j.logger.Error("failed to update marks", err)
			j.app.AddError(err)
		}
		processDone <- true
	}()

	select {
	case <-ctx.Done():
		err := errors.New("timeout reached while updating marks")
		j.logger.Error("failed to update marks", err)
		j.app.AddError(err)
	case <-processDone:
	}

	cancel()
	j.app.AddEvent("job.UpdateMarks", start)
}

func (j *Job) UpdateDurations() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(j.cfg.App.UpdateDurationsTimeout))
	processDone := make(chan bool)
	go func() {
		err := j.app.UpdateTestDurations(ctx)
		if err != nil {
			j.logger.Error("failed to update durations", err)
			j.app.AddError(err)
		}
		processDone <- true
	}()

	select {
	case <-ctx.Done():
		err := errors.New("timeout reached while updating durations")
		j.logger.Error("failed to update durations", err)
		j.app.AddError(err)
	case <-processDone:
	}

	cancel()
	j.app.AddEvent("job.UpdateDurations", start)
}
