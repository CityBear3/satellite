package usecase

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/gateway/repository"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/usecase/service"
)

type (
	EventUseCase interface {
		PublishArchiveEvent(ctx context.Context, client entity.Client) (primitive.ID, error)
		ReceiveArchiveEvent(ctx context.Context, device entity.Device) (<-chan service.ArchiveEventMessage, error)
	}
)

type EventInteractor struct {
	eventRepository repository.IEventRepository
	eventHandler    service.IEventService
	txManager       ITxManager
}

func NewEventInteractor(
	eventRepository repository.IEventRepository,
	eventHandler service.IEventService,
	txManager ITxManager,
) *EventInteractor {
	return &EventInteractor{
		eventRepository: eventRepository,
		eventHandler:    eventHandler,
		txManager:       txManager,
	}
}

func (i EventInteractor) PublishArchiveEvent(ctx context.Context, client entity.Client) (primitive.ID, error) {
	archiveEvent := entity.NewArchiveEvent(primitive.NewID(), client.Devices[0].ID, client.ID)
	if _, err := i.txManager.DoInTx(ctx, func(ctx2 context.Context) error {
		if err := i.eventRepository.SaveArchiveEvent(ctx2, archiveEvent); err != nil {
			return err
		}

		if err := i.eventHandler.PublishArchiveEvent(ctx2, archiveEvent); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return primitive.ID{}, err
	}

	return archiveEvent.ID, nil
}

func (i EventInteractor) ReceiveArchiveEvent(ctx context.Context, device entity.Device) (<-chan service.ArchiveEventMessage, error) {
	archiveEvents, err := i.eventHandler.ReceiveArchiveEvent(ctx, device.ID)
	if err != nil {
		return nil, err
	}

	return archiveEvents, nil
}
