package job

import (
	"context"
	"errors"
	"time"
)

func (j *Job) UpdateSystemSummary() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(j.cfg.App.SystemSummaryTimeout)*time.Millisecond)
	processDone := make(chan bool)
	go func() {
		err := j.app.UpdateSystemSummary(ctx)
		if err != nil {
			j.logger.Error("failed to update system summary", err)
			j.app.AddError(err)
		}
		processDone <- true
	}()

	select {
	case <-ctx.Done():
		err := errors.New("timeout reached while updating system summary")
		j.logger.Error("failed to update system summary", err)
		j.app.AddError(err)
	case <-processDone:
	}

	cancel()
	j.app.AddEvent("job.UpdateSystemSummary", start)
}
