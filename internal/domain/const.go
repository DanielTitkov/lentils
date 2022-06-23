package domain

import (
	"time"
)

const (
	// locales
	LocaleRu = "ru"
	LocaleEn = "en"
	// generate questions methods
	GenerateQuestionsNone     = "none"
	GenerateQuestionsEachItem = "each-item"
	// scale types
	ScaleTypeSten   = "sten"
	ScaleTypeSum    = "sum"
	ScaleTypeZScore = "zscore"
	ScaleTypeMean   = "mean"
	// question types
	QuestionTypeSimple = "simple"
	// test steps
	TestStepIntro     = "intro"
	TestStepQuestions = "questions"
	TestStepFinish    = "finish"
	TestStepResult    = "result"
	// samples
	SampleAllCode              = "all"
	SampleAllNonSuspiciousCode = "all-not-suspicious"
	// norms
	NormMinBase = 5 // FIXME
	// take
	TakeMinTime = 9 * time.Second // FIXME
	TakeMaxTime = 2 * time.Hour
)
