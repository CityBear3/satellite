//go:generate mockgen -source=$GOFILE -package=mock_usecase -destination=./mock/$GOFILE
package usecase

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/usecase/dto"
)

type ArchiveUseCase interface {
	CreateArchive(ctx context.Context, request dto.UploadArchiveRequest, device entity.Device) error
	GetArchive(ctx context.Context, request dto.GetArchiveRequest, client entity.Client) (dto.GetArchiveResult, error)
}
