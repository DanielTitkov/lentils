package entgo

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/badge"
)

func (r *EntgoRepository) CreateOrUpdateBadgeByType(ctx context.Context, b *domain.Badge) (*domain.Badge, error) {
	badge, err := r.client.Badge.
		Query().
		Where(badge.TypeEQ(b.Type)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create badge
		badge, err := r.client.Badge.
			Create().
			SetActive(b.Active).
			SetType(b.Type).
			SetMeta(b.Meta).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainBadge(badge), nil
	}

	// update challenge
	badge, err = badge.Update().
		SetActive(b.Active).
		SetMeta(b.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainBadge(badge), nil
}

func entToDomainBadge(b *ent.Badge) *domain.Badge {
	return &domain.Badge{
		Active: b.Active,
		Type:   b.Type,
		Meta:   b.Meta,
	}
}
