package auth

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"golang.org/x/crypto/bcrypt"
)

func CompareSecret(hashed authentication.Secret, row authentication.Secret) error {
	h, err := hashed.Value()
	if err != nil {
		return err
	}

	r, err := row.Value()
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(h), []byte(r)); err != nil {
		return apperrs.UnauthenticatedError
	}

	return nil
}

func AuthenticatedDevice(ctx context.Context) (entity.Device, error) {
	device, ok := ctx.Value("device").(entity.Device)
	if !ok {
		return entity.Device{}, apperrs.UnauthenticatedError
	}

	return device, nil
}

func AuthenticatedClient(ctx context.Context) (entity.Client, error) {
	client, ok := ctx.Value("client").(entity.Client)
	if !ok {
		return entity.Client{}, apperrs.UnauthenticatedError
	}

	return client, nil
}
