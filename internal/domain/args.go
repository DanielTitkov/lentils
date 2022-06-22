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
		// TODO:
		// for each user
		// force generate
		// randomize items
	}

	TestTranslation struct {
		Locale      string
		Title       string
		Description string
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
		Type         string
		Code         string
		Translations []ScaleTranslation
		Items        []CreateItemArgs
	}

	ScaleTranslation struct {
		Locale      string
		Title       string
		Description string
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
)

// norm calculation args
type (
	NormCalculationData struct {
		ScaleID uuid.UUID
		Takes   []struct { // takes must be complete
			// values will be summed up to get raw value
			// them get mean and sigma from them
			Responses []struct {
				Value   int
				Reverse bool
			}
		}
	}
)

// function args
type (
	PrepareTestArgs struct {
		UserID  uuid.UUID
		Session string
	}
)
