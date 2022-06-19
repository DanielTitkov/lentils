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
	t.Meta["begin"] = time.Now()

	return nil
}

func (t *Take) End() error {
	if t.Status != TestStepQuestions {
		return fmt.Errorf("end is only possible from '%s' status but got %s", TestStepQuestions, t.Status)
	}

	t.Status = TestStepFinish
	t.Progress = 100
	t.Meta["end"] = time.Now()

	return nil
}
