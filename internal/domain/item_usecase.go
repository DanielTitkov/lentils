package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (i *Item) CurrentResponseValue() int {
	if i.Response == nil {
		return -1
	}

	return i.Response.Value
}

func (i *Item) AddResponse(takeID uuid.UUID, value int, meta map[string]interface{}) error {
	if takeID == uuid.Nil {
		return errors.New("got nil take uuid")
	}

	if value < 0 {
		return fmt.Errorf("value cannot be less than 0, got %d", value)
	}

	if value > i.Steps-1 {
		return fmt.Errorf("value cannot be greater than %d, got %d", i.Steps, value)
	}

	if meta == nil {
		meta = make(map[string]interface{})
	}

	meta["timestamp"] = time.Now().UnixNano()

	i.Response = &Response{
		ItemID: i.ID,
		TakeID: takeID,
		Value:  value,
		Meta:   meta,
	}

	return nil
}
