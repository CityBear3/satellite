package rpc

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

func AuthenticatedDevice(ctx context.Context) (entity.Device, error) {
	device, ok := ctx.Value("device").(entity.Device)
	if !ok {
		return entity.Device{}, apperrs.ForbiddenError
	}

	return device, nil
}

func AuthenticatedClient(ctx context.Context) (entity.Client, error) {
	client, ok := ctx.Value("client").(entity.Client)
	if !ok {
		return entity.Client{}, apperrs.ForbiddenError
	}

	return client, nil
}
