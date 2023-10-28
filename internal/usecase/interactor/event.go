package interactor

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/internal/usecase/dto"
)

type EventInteractor struct {
	eventRepository repository.IEventRepository
	eventHandler    usecase.IEventHandler
	txManager       usecase.ITxManager
}

func NewEventInteractor(
	eventRepository repository.IEventRepository,
	eventHandler usecase.IEventHandler,
	txManager usecase.ITxManager,
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

func (i EventInteractor) ReceiveArchiveEvent(ctx context.Context, device entity.Device) (<-chan dto.ArchiveEventMessage, error) {
	archiveEvents, err := i.eventHandler.ReceiveArchiveEvent(ctx, device.ID)
	if err != nil {
		return nil, err
	}

	return archiveEvents, nil
}
