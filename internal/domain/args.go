package domain

import (
	"github.com/google/uuid"
)

// test import args
type (
	CreateTestArgs struct {
		Code             string
		Published        bool
		AvailableLocales []string `yaml:"availableLocales"`
		ForceUpdate      bool     `yaml:"forceUpdate"`
		Generate         GenerateQuestionsArgs
		Tags             []string
		Translations     []TestTranslation
		Questions        []CreateQuestionArgs
		Scales           []CreateScaleArgs
		Display          CreateTestDisplayArgs
	}

	CreateTestDisplayArgs struct {
		QuestionsPerPage int  `yaml:"questionsPerPage"`
		RandomizeOrder   bool `yaml:"randomizeOrder"`
	}

	GenerateQuestionsArgs struct {
		Method   string
		Template CreateQuestionArgs
	}

	TestTranslation struct {
		Locale      string
		Title       string
		Description string
		Details     string
		Instruction string
	}

	CreateQuestionArgs struct {
		Code         string
		Type         string
		Order        int
		Translations []QuestionTranslation
		Items        []CreateItemArgs
	}

	QuestionTranslation struct {
		Locale        string
		Content       string
		HeaderContent string
		FooterConent  string
	}

	CreateScaleArgs struct {
		Type            string
		Code            string
		Translations    []ScaleTranslation
		Items           []CreateItemArgs
		Interpretations []CreateInterpretationArgs
	}

	ScaleTranslation struct {
		Locale       string
		Title        string
		Description  string
		Abbreviation string
	}

	CreateInterpretationArgs struct {
		Range        [2]float64
		Translations []InterpretationTranslation
	}

	InterpretationTranslation struct {
		Locale  string
		Content string
	}

	CreateItemArgs struct {
		Code         string
		Steps        int
		Reverse      bool
		Translations []ItemTranslation
	}

	ItemTranslation struct {
		Locale  string
		Content string
	}

	CreateTagArgs struct {
		Code         string
		Type         string
		Translations []TagTranslation
	}

	TagTranslation struct {
		Locale  string
		Content string
	}
)

// query args
type (
	QueryTestsArgs struct {
		Locale        string
		Tags          []*Tag
		TagIDs        []uuid.UUID
		FilterModeAny bool
	}
)

// norm calculation args
type (
	NormCalculationData struct {
		ScaleID   uuid.UUID
		ScaleCode string
		Results   []float64
	}
)

// function args
type (
	PrepareTestArgs struct {
		UserID  uuid.UUID
		Session string
	}
)
