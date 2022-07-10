package job

import (
	"context"
	"errors"
	"time"
)

func (j *Job) UpdateNorms() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(j.cfg.App.UpdateNormsTimeout))
	processDone := make(chan bool)
	go func() {
		err := j.app.UpdateNorms(ctx)
		if err != nil {
			j.logger.Error("failed to update norms", err)
			j.app.AddError(err)
		}
		processDone <- true
	}()

	select {
	case <-ctx.Done():
		err := errors.New("timeout reached while updating norms")
		j.logger.Error("failed to update norms", err)
		j.app.AddError(err)
	case <-processDone:
	}

	cancel()
	j.app.AddEvent("job.UpdateNorms", start)
}
