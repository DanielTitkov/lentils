package domain

import (
	"html/template"
	"time"

	"github.com/google/uuid"
)

// test logic types
type (
	Test struct {
		ID                uuid.UUID
		Code              string        // unique code for url
		Title             string        // translatable
		Description       string        // translatable
		Details           template.HTML // translatable
		Instruction       template.HTML // translatable
		ResultPreambule   template.HTML // translatable
		Locale            string
		AvailableLocales  []string
		GenerateQuestions string
		Published         bool
		Questions         []*Question
		Scales            []*Scale
		Tags              []*Tag
		Image             string
		Display           TestDisplay
		Mark              float64
		Duration          time.Duration
		QuestionCount     int     // for display (and less joins in requests)
		Take              *Take   // for use in handler
		Takes             []*Take // for calculations
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
		Steps      int    // number of steps in response scale
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
		Abbreviation    string // translatable
		Global          bool   // if scale can be used by more than one test
		Items           []*Item
		Interpretations []*Interpretation
		Result          *ScaleResult // not save in db as yet
		Norm            *Norm        // to use in calculation
	}

	Norm struct {
		ID       uuid.UUID
		SampleID uuid.UUID
		ScaleID  uuid.UUID
		Name     string
		Base     int
		Mean     float64
		Sigma    float64
		Rank     int
		Meta     map[string]interface{}
	}

	Sample struct {
		ID       uuid.UUID
		Code     string
		Criteria SampleCriteria
	}

	SampleCriteria struct {
		NotSuspicious bool   `json:"notSuspicious,omitempty"`
		Locale        string `json:"locale,omitempty"`
	}

	ScaleResult struct {
		RawScore       float64
		Score          float64
		Min            float64
		Max            float64
		Interpretation *Interpretation
		Formula        string
		Elapsed        time.Duration
		Unit           string
		Meta           map[string]interface{}
	}

	// Result is saved to db and used for norms calculation
	// there's no need as yet to save calculatable data like min, max, etc.
	Result struct {
		ID         uuid.UUID
		TakeID     uuid.UUID
		ScaleID    uuid.UUID
		RawScore   float64
		FinalScore float64
		CreateTime time.Time
		UpdateTime time.Time
		Meta       map[string]interface{}
	}

	Interpretation struct {
		ID      uuid.UUID
		Content template.HTML // translatable
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
		StartTime  *time.Time
		EndTime    *time.Time
		Suspicious bool
		InLocale   string
		CreateTime time.Time
		UpdateTime time.Time
		Page       int
		Status     string
		Progress   int
		Mark       *int
		Meta       map[string]interface{}
	}

	Tag struct {
		ID      uuid.UUID
		Code    string
		Type    string
		Content string
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
		UseDarkTheme bool
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
	SystemSummary struct {
		ID            int
		Users         int
		Tests         int
		FinishedTakes int
		Responses     int
		CreateTime    time.Time
	}
	Event struct {
		Name      string
		StartTime time.Time
		EndTime   time.Time
		Elapsed   time.Duration
	}
)
