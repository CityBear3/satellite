package usecase

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/repository"
)

type (
	EventUseCase interface {
		PublishArchiveEvent(ctx context.Context, client entity.Client) (primitive.ID, error)
		ReceiveArchiveEvent(ctx context.Context, device entity.Device) (<-chan ArchiveEventMessage, error)
	}

	IEventHandler interface {
		PublishArchiveEvent(ctx context.Context, event entity.ArchiveEvent) error
		ReceiveArchiveEvent(ctx context.Context, deviceID primitive.ID) (<-chan ArchiveEventMessage, error)
	}

	ArchiveEventMessage struct {
		ID       string `json:"id"`
		ClientID string `json:"client_id"`
	}
)

type EventInteractor struct {
	eventRepository repository.IEventRepository
	eventHandler    IEventHandler
	txManager       ITxManager
}

func NewEventInteractor(
	eventRepository repository.IEventRepository,
	eventHandler IEventHandler,
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
	if err := i.txManager.DoInTx(ctx, func(rtx repository.ITx) error {
		if err := i.eventRepository.SaveArchiveEvent(ctx, rtx, archiveEvent); err != nil {
			return err
		}

		if err := i.eventHandler.PublishArchiveEvent(ctx, archiveEvent); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return primitive.ID{}, err
	}

	return archiveEvent.ID, nil
}

func (i EventInteractor) ReceiveArchiveEvent(ctx context.Context, device entity.Device) (<-chan ArchiveEventMessage, error) {
	archiveEvents, err := i.eventHandler.ReceiveArchiveEvent(ctx, device.ID)
	if err != nil {
		return nil, err
	}

	return archiveEvents, nil
}
