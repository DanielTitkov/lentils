package domain

import (
	"fmt"
	"time"
)

func (t *Take) Begin() error {
	if t.Status != TestStepIntro {
		return fmt.Errorf("begin is only possible from '%s' status but got %s", TestStepIntro, t.Status)
	}
	t.Status = TestStepQuestions
	t.Progress = 0
	t.Page = 1
	now := time.Now()
	t.StartTime = &now

	return nil
}

func (t *Take) End() error {
	if t.Status != TestStepQuestions {
		return fmt.Errorf("end is only possible from '%s' status but got %s", TestStepQuestions, t.Status)
	}

	t.Status = TestStepFinish
	t.Progress = 100
	now := time.Now()
	t.EndTime = &now
	t.calculateSuspicion()

	return nil
}

func (t *Take) calculateSuspicion() {
	if t.StartTime == nil || t.EndTime == nil {
		return
	}
	if t.EndTime.Sub(*t.StartTime) < TakeMinTime {
		t.Suspicious = true
	}
	if t.EndTime.Sub(*t.StartTime) > TakeMaxTime {
		t.Suspicious = true
	}
}
