//go:generate mockgen -source=$GOFILE -package=mock_repository -destination=../../../adaptor/gateway/repository/mock/$GOFILE
package repository

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type IClientRepository interface {
	GetClient(ctx context.Context, clientID primitive.ID) (entity.Client, error)
}
