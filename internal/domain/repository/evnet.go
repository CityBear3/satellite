//go:generate mockgen -source=$GOFILE -package=mock_repository -destination=../../adaptor/repository/mock/$GOFILE
package repository

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type IEventRepository interface {
	SaveArchiveEvent(ctx context.Context, tx ITx, archiveEvent entity.ArchiveEvent) error
	GetArchiveEvent(ctx context.Context, archiveEventID primitive.ID) (entity.ArchiveEvent, error)
}
