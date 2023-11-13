//go:generate mockgen -source=$GOFILE -package=mock_repository -destination=../../adaptor/repository/mock/$GOFILE
package repository

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
)

type IArchiveRepository interface {
	Save(ctx context.Context, tx ITx, archive entity.Archive) error
	GetArchive(ctx context.Context, archiveId primitive.ID) (entity.Archive, error)
	GetArchiveByArchiveEventID(ctx context.Context, archiveEventID primitive.ID) (entity.Archive, error)
}

type IFileTransfer interface {
	Save(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType, data archive.Data) error
	GetFile(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType) (archive.Data, error)
}
