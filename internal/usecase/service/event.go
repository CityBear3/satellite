package service

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type (
	IEventService interface {
		PublishArchiveEvent(ctx context.Context, event entity.ArchiveEvent) error
		ReceiveArchiveEvent(ctx context.Context, deviceID primitive.ID) (<-chan ArchiveEventMessage, error)
	}

	ArchiveEventMessage struct {
		ID       string `json:"id"`
		ClientID string `json:"client_id"`
	}
)
