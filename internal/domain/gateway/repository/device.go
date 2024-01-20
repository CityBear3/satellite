//go:generate mockgen -source=$GOFILE -package=mock_repository -destination=../../../adaptor/repository/mock/$GOFILE
package repository

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type IDeviceRepository interface {
	GetDevice(ctx context.Context, deviceID primitive.ID) (entity.Device, error)
}
