package interactor

import (
	"context"
	"time"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/internal/usecase/dto"
)

type ArchiveInteractor struct {
	archiveRepository repository.IArchiveRepository
	eventRepository   repository.IEventRepository
	txManager         usecase.ITxManager
}

func NewArchiveInteractor(
	archiveRepository repository.IArchiveRepository,
	eventRepository repository.IEventRepository,
	txManager usecase.ITxManager,
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
	request dto.UploadArchiveRequest,
	device entity.Device,
) error {
	if err := i.txManager.DoInTx(ctx, func(rtx repository.ITx) error {
		archiveID := primitive.NewID()

		contentType, err := primitive.NewContentType(request.ContentType)
		if err != nil {
			return err
		}

		archiveEventID, err := primitive.ParseID(request.ArchiveEventID)
		if err != nil {
			return err
		}

		event, err := i.eventRepository.GetArchiveEvent(ctx, archiveEventID)
		if err != nil {
			return err
		}

		if err := event.CheckCorrectCall(device.ID, device.ClientID, time.Now()); err != nil {
			return err
		}

		archive := entity.NewArchive(archiveID, archiveEventID, contentType, device.ID)
		if err := i.archiveRepository.Save(ctx, rtx, archive); err != nil {
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
	request dto.GetArchiveRequest,
	client entity.Client,
) (dto.GetArchiveResult, error) {
	archiveEventID, err := primitive.ParseID(request.ArchiveEventID)
	if err != nil {
		return dto.GetArchiveResult{}, err
	}

	archive, err := i.archiveRepository.GetArchiveByArchiveEventID(ctx, archiveEventID)
	if err != nil {
		return dto.GetArchiveResult{}, err
	}

	if err = archive.CheckCorrectCall(client); err != nil {
		return dto.GetArchiveResult{}, err
	}

	return dto.GetArchiveResult{
		ID:             archive.ID.Value(),
		ArchiveEventID: archive.ArchiveEventID.Value(),
		ContentType:    archive.ContentType.GetExt(),
		DeviceID:       archive.DeviceID.Value(),
	}, nil
}
