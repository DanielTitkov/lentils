package domain

import (
	"time"

	"github.com/google/uuid"
)

// test logic types
type (
	Test struct {
		ID          uuid.UUID
		Code        string
		Title       string
		Description string
		Instruction string
		Locale      string
		Questions   []*Question
		Scales      []*Scale
	}

	Item struct {
		ID         uuid.UUID
		TestID     uuid.UUID
		ScaleID    uuid.UUID
		QuestionID uuid.UUID
		Content    string
		Type       string
		Steps      int // number of stepes in response scale
	}

	Question struct {
		ID            uuid.UUID
		TestID        uuid.UUID
		Content       string
		HeaderContent string
		FooterConent  string
		Items         []Item
	}

	Scale struct {
		ID              uuid.UUID
		Code            string
		Type            string
		Title           string
		Description     string
		Items           []Item
		Interpretations []Interpretation
	}

	Interpretation struct {
		ID      uuid.UUID
		Content string
		Range   [2]float64
	}

	Response struct {
		ID     uuid.UUID
		ItemID uuid.UUID
		UserID uuid.UUID
		Value  int
		Meta   map[string]interface{}
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
		Meta         map[string]interface{}
	}

	UserSummary struct {
		UserID               uuid.UUID
		CorrectPredictions   int
		IncorrectPredictions int
		UnknownPredictions   int
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

// legacy types // FIXME
type (
	Badge struct {
		ID     int // probably not needed
		UserID uuid.UUID
		Type   string
		Active bool
		Meta   map[string]interface{}
	}

	Challenge struct {
		ID             uuid.UUID
		AuthorID       uuid.UUID
		Type           string
		Content        string
		Description    string
		Outcome        *bool
		Published      bool
		StartTime      time.Time
		EndTime        time.Time
		Predictions    []*Prediction
		Proofs         []*Proof
		UserPrediction *Prediction
	}

	Prediction struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		ChallengeID uuid.UUID
		Prognosis   bool
		Meta        map[string]interface{}
	}

	Proof struct {
		ID          uuid.UUID
		ChallengeID uuid.UUID
		Content     string
		Link        string
		Meta        map[string]interface{}
	}

	SystemSymmary struct {
		ID                 int
		Users              int
		Challenges         int
		OngoingChallenges  int
		FinishedChallenges int
		Predictions        int
		CreateTime         time.Time
	}
)
