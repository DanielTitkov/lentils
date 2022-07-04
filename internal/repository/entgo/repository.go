package entgo

import (
	"github.com/DanielTitkov/orrery/internal/logger"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
)

type EntgoRepository struct {
	client *ent.Client
	logger *logger.Logger
}

func NewEntgoRepository(
	client *ent.Client,
	logger *logger.Logger,
) *EntgoRepository {
	return &EntgoRepository{
		client: client,
		logger: logger,
	}
}
