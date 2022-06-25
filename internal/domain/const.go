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
	NormMinBase  = 10
	NormOKBase   = 50
	NormGoodBase = 100
	// take
	TakeMinTime = 8 * time.Second
	TakeMaxTime = 2 * time.Hour
)

func Locales() []string {
	return []string{
		LocaleEn,
		LocaleRu,
	}
}

func DefaultLocale() string {
	return LocaleEn
}

func IsValidLocale(locale string) bool {
	for _, l := range Locales() {
		if l == locale {
			return true
		}
	}

	return false
}

func AreValidLocales(locales []string) bool {
	for _, l := range locales {
		if ok := IsValidLocale(l); !ok {
			return false
		}
	}

	return true
}
