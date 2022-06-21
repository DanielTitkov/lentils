package domain

import (
	"time"

	"github.com/google/uuid"
)

// test logic types
type (
	Test struct {
		ID                uuid.UUID
		Code              string // unique code for url
		Title             string // translatable
		Description       string // translatable
		Instruction       string // translatable
		Locale            string
		AvailableLocales  []string
		GenerateQuestions string
		Published         bool
		Questions         []*Question
		Scales            []*Scale
		Display           TestDisplay
	}

	TestDisplay struct {
		RandomizeOrder   bool
		QuestionsPerPage int
	}

	Item struct {
		ID         uuid.UUID
		TestID     uuid.UUID
		ScaleID    uuid.UUID
		QuestionID uuid.UUID
		Code       string
		Content    string // translatable
		Steps      int    // number of stepes in response scale
		Reverse    bool
		// Type       string
		Response *Response // for use in handler, not for saving
	}

	Question struct {
		ID            uuid.UUID
		TestID        uuid.UUID
		Code          string
		Order         int
		Content       string // translatable
		HeaderContent string // translatable
		FooterContent string // translatable
		Items         []*Item
		// Type          string // not needed as yet
	}

	Scale struct {
		ID              uuid.UUID
		Code            string
		Type            string
		Title           string // translatable
		Description     string // translatable
		Global          bool   // if scale can be used by more than one test
		Items           []*Item
		Interpretations []*Interpretation
		Result          *ScaleResult // not save in db as yet
	}

	Norm struct {
		ID    uuid.UUID
		Mean  float64
		Sigma float64
	}

	ScaleResult struct {
		Value          float64
		Min            float64
		Max            float64
		Interpretation *Interpretation
		Formula        string
		Elaplsed       time.Duration
		Meta           map[string]interface{}
	}

	Interpretation struct {
		ID      uuid.UUID
		Content string // translatable
		Range   [2]float64
	}

	Response struct {
		ID         uuid.UUID
		ItemID     uuid.UUID
		TakeID     uuid.UUID
		Value      int
		CreateTime time.Time
		UpdateTime time.Time
		Meta       map[string]interface{}
	}

	// Take is one instance of user taking a test
	Take struct {
		ID         uuid.UUID
		UserID     uuid.UUID
		TestID     uuid.UUID
		Seed       int64
		CreateTime time.Time
		UpdateTime time.Time
		Page       int
		Status     string
		Progress   int
		Meta       map[string]interface{}
	}
)

// user types
type (
	User struct {
		ID           uuid.UUID
		Name         string
		Email        string
		Admin        bool
		Picture      string
		Password     string
		PasswordHash string
		Locale       string
		Anonymous    bool
		AnonymousID  []uuid.UUID
		Meta         map[string]interface{}
	}

	UserSummary struct {
		UserID uuid.UUID
	}

	UserSession struct {
		ID           int    // probably uuid not needed here, sessions are temporary anyways
		SID          string // code to identify the session
		UserID       uuid.UUID
		IP           string
		UserAgent    string
		CreateTime   time.Time
		UpdateTime   time.Time
		LastActivity time.Time
		Meta         map[string]interface{}
		Active       bool
	}
)

// system types
type (
	SystemSymmary struct {
		ID         int
		Users      int
		CreateTime time.Time
	}
)
