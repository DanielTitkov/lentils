package entgo

import (
	"context"
	"errors"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/response"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"
)

func (r *EntgoRepository) AddOrUpdateResponse(ctx context.Context, takeID uuid.UUID, itm *domain.Item) (*domain.Response, error) {
	if takeID == uuid.Nil {
		return nil, errors.New("take id mustn't be nil")
	}

	if itm.Response == nil {
		return nil, errors.New("item must have response")
	}

	respQuery := r.client.Response.Query()
	if itm.Response.ID != uuid.Nil {
		respQuery.Where(response.IDEQ(itm.Response.ID))
	} else {
		respQuery.Where(
			response.HasTakeWith(take.IDEQ(takeID)),
			response.HasItemWith(item.IDEQ(itm.ID)),
		)
	}

	resp, err := respQuery.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		// response not found, create response
		resp, err = r.client.Response.Create().
			SetItemID(itm.ID).
			SetTakeID(takeID).
			SetMeta(itm.Response.Meta).
			SetValue(itm.Response.Value).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		// response successfully created
		return entToDomainResponse(resp, takeID, itm.ID), nil
	}

	resp, err = resp.Update().
		SetMeta(itm.Response.Meta).
		SetValue(itm.Response.Value).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainResponse(resp, takeID, itm.ID), nil
}

func entToDomainResponse(r *ent.Response, takeID, itemID uuid.UUID) *domain.Response {
	if r.Edges.Take != nil {
		takeID = r.Edges.Take.ID
	}

	if r.Edges.Item != nil {
		itemID = r.Edges.Item.ID
	}

	return &domain.Response{
		ID:     r.ID,
		Value:  r.Value,
		Meta:   r.Meta,
		TakeID: takeID,
		ItemID: itemID,
	}
}
