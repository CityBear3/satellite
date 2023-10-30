package usecase

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/usecase/dto"
)

type EventUseCase interface {
	PublishArchiveEvent(ctx context.Context, client entity.Client) (primitive.ID, error)
	ReceiveArchiveEvent(ctx context.Context, device entity.Device) (<-chan dto.ArchiveEventMessage, error)
}

type IEventHandler interface {
	PublishArchiveEvent(ctx context.Context, event entity.ArchiveEvent) error
	ReceiveArchiveEvent(ctx context.Context, deviceID primitive.ID) (<-chan dto.ArchiveEventMessage, error)
}
