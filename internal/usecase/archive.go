//go:generate mockgen -source=$GOFILE -package=mock_usecase -destination=./mock/$GOFILE
package usecase

import (
	"context"
	"time"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/domain/repository"
)

type (
	ArchiveUseCase interface {
		CreateArchive(ctx context.Context, request CreateArchiveInput, device entity.Device) error
		GetArchive(ctx context.Context, request GetArchiveInput, client entity.Client) (GetArchiveResult, error)
	}

	CreateArchiveInput struct {
		ArchiveEventID primitive.ID
		ContentType    archive.ContentType
		Data           archive.Data
	}

	GetArchiveInput struct {
		ArchiveEventID primitive.ID
	}

	GetArchiveResult struct {
		ID             primitive.ID
		ArchiveEventID primitive.ID
		ContentType    archive.ContentType
		DeviceID       primitive.ID
		Data           archive.Data
	}
)

type ArchiveInteractor struct {
	archiveRepository repository.IArchiveRepository
	eventRepository   repository.IEventRepository
	txManager         ITxManager
}

func NewArchiveInteractor(
	archiveRepository repository.IArchiveRepository,
	eventRepository repository.IEventRepository,
	txManager ITxManager,
) *ArchiveInteractor {
	return &ArchiveInteractor{
		archiveRepository: archiveRepository,
		eventRepository:   eventRepository,
		txManager:         txManager,
	}
}

// CreateArchive do operation for upload archive
func (i *ArchiveInteractor) CreateArchive(
	ctx context.Context,
	request CreateArchiveInput,
	device entity.Device,
) error {
	if err := i.txManager.DoInTx(ctx, func(rtx repository.ITx) error {
		archiveID := primitive.NewID()

		event, err := i.eventRepository.GetArchiveEvent(ctx, request.ArchiveEventID)
		if err != nil {
			return err
		}

		if err := event.CheckCorrectCall(device.ID, device.ClientID, time.Now()); err != nil {
			return err
		}

		archiveEntity := entity.NewArchive(archiveID, request.ArchiveEventID, request.ContentType, device.ID, request.Data)
		if err := i.archiveRepository.Save(ctx, rtx, archiveEntity); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GetArchive do operation for get archive
func (i *ArchiveInteractor) GetArchive(
	ctx context.Context,
	request GetArchiveInput,
	client entity.Client,
) (GetArchiveResult, error) {
	archiveEntity, err := i.archiveRepository.GetArchiveByArchiveEventID(ctx, request.ArchiveEventID)
	if err != nil {
		return GetArchiveResult{}, err
	}

	if err = archiveEntity.CheckCorrectCall(client); err != nil {
		return GetArchiveResult{}, err
	}

	return GetArchiveResult{
		ID:             archiveEntity.ID,
		ArchiveEventID: archiveEntity.ArchiveEventID,
		ContentType:    archiveEntity.ContentType,
		DeviceID:       archiveEntity.DeviceID,
		Data:           archiveEntity.Data,
	}, nil
}
