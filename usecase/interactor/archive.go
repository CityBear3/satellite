package interactor

import (
	"context"
	"fmt"
	"github.com/CityBear3/satellite/domain/model"
	"github.com/CityBear3/satellite/pkg/apperrs"
	"github.com/CityBear3/satellite/usecase"
	"github.com/CityBear3/satellite/usecase/dto"
	"strings"
)

const (
	FileSizeLimitation = 8e6
)

type ArchiveInteractor struct {
	archiveRepository model.IArchiveRepository
	txManager         usecase.ITxManager
}

func NewArchiveInteractor(archiveRepository model.IArchiveRepository, txManager usecase.ITxManager) *ArchiveInteractor {
	return &ArchiveInteractor{
		archiveRepository: archiveRepository,
		txManager:         txManager,
	}
}

// UploadArchive do operation for upload image
func (i *ArchiveInteractor) UploadArchive(ctx context.Context, request dto.UploadArchiveRequest) error {
	if _, err := i.txManager.DoInTx(ctx, func(ctxWithTx context.Context) (dto.IResult, error) {
		archive := model.Archive{
			Id:       request.Id,
			Size:     len(request.Data),
			Ext:      getExt(request.ContentType),
			DeviceId: request.DeviceId,
		}
		if !archive.CheckDataSize(FileSizeLimitation) {
			return nil, apperrs.NewError(apperrs.BadRequest, fmt.Sprintf(apperrs.InvalidFileSizeMSG, "8MB"))
		}
		if err := i.archiveRepository.Save(ctxWithTx, archive); err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func getExt(contentType string) string {
	return strings.TrimPrefix(contentType, "image/")
}
