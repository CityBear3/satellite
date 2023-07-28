package archive

import (
	"context"
	"fmt"
	"github.com/CityBear3/satellite/domain/model"
	"github.com/CityBear3/satellite/logic"
	"github.com/CityBear3/satellite/logic/dto"
	"github.com/CityBear3/satellite/pkg/apperrs"
	"strings"
)

const (
	FileSizeLimitation = 8e6
)

type UploadArchiveInteractor struct {
	archiveRepository model.IArchiveRepository
	txManager         logic.ITxManager
}

func NewArchiveImageInteractor(archiveRepository model.IArchiveRepository, txManager logic.ITxManager) *UploadArchiveInteractor {
	return &UploadArchiveInteractor{
		archiveRepository: archiveRepository,
		txManager:         txManager,
	}
}

// Handle do operation for upload image
func (i *UploadArchiveInteractor) Handle(ctx context.Context, request dto.UploadArchiveRequest) error {
	if err := i.txManager.DoInTx(ctx, func(ctxWithTx context.Context) error {
		archive := model.Archive{
			Id:       request.Id,
			Size:     len(request.Data),
			Ext:      getExt(request.ContentType),
			DeviceId: request.DeviceId,
		}
		if !archive.CheckDataSize(FileSizeLimitation) {
			return apperrs.NewError(apperrs.BadRequest, fmt.Sprintf(apperrs.InvalidFileSizeMSG, "8MB"))
		}
		if err := i.archiveRepository.Save(ctxWithTx, archive); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func getExt(contentType string) string {
	return strings.TrimPrefix(contentType, "image/")
}
