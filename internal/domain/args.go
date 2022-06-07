package domain

import (
	"github.com/google/uuid"
)

type (
	CreateChallengeArgs struct {
		Type        string    `json:"type"`
		Content     string    `json:"content"`
		Description string    `json:"description"`
		StartTime   string    `json:"startTime"`
		EndTime     string    `json:"endTime"`
		Published   bool      `json:"published"`
		TimeLayout  string    `json:"timeLayout"`
		Outcome     *bool     `json:"outcome"`
		AuthorID    uuid.UUID `json:"authorID"`
	}
	CreatePredictionFromPresetArgs struct {
		UserEmail        string `json:"userEmail"`
		ChallengeContent string `json:"challengeContent"`
		Prognosis        bool   `json:"prognosis"`
	}
	FilterChallengesArgs struct {
		UserID      uuid.UUID
		AuthorID    uuid.UUID
		Unvoted     bool
		Ongoing     bool
		Pending     bool
		Unpublished bool
		Finished    bool
		Limit       int
		Offset      int
	}
)
