package domain

import (
	"time"
)

const (
	// env
	EnvDev = "dev"
	// locales
	LocaleRu = "ru"
	LocaleEn = "en"
	// generate questions methods
	GenerateQuestionsNone     = "none"
	GenerateQuestionsEachItem = "each-item"
	// scale types
	ScaleTypeSten   = "sten"
	ScaleTypeSum    = "sum"
	ScaleTypePerc   = "perc"
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
	TakeMinMark = 1
	TakeMaxMark = 5
	// tags
	TagTypeTheme   = "theme"
	TagTypeLen     = "len"
	TagTypeFeature = "feature"
	TagCodeShort   = "short"
	TagCodeMedium  = "medium"
	TagCodeLong    = "long"
	TagLenShort    = 10
	TagLenLong     = 50
	// time
	DefaultDisplayTime = time.RFC822
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

func LocaleIcon(locale string) string {
	switch locale {
	case LocaleEn:
		return "🇬🇧"
	case LocaleRu:
		return "🇷🇺"
	default:
		return locale
	}
}
