//go:generate mockgen -source=$GOFILE -package=mock_repository -destination=../../../adaptor/gateway/repository/mock/$GOFILE
package repository

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type IArchiveRepository interface {
	Save(ctx context.Context, archive entity.Archive) error
	GetArchive(ctx context.Context, archiveId primitive.ID) (entity.Archive, error)
	GetArchiveByArchiveEventID(ctx context.Context, archiveEventID primitive.ID) (entity.Archive, error)
}
