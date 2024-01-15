//go:generate mockgen -source=$GOFILE -package=mock_transfer -destination=../../../adaptor/handler/mock/$GOFILE
package transfer

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
)

type IFileTransfer interface {
	Save(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType, data archive.Data) error
	GetFile(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType) (archive.Data, error)
}
