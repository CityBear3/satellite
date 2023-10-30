package interactor

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/internal/pkg/auth"
	"github.com/CityBear3/satellite/internal/usecase/dto"
)

type AuthenticationInteractor struct {
	clientRepository repository.IClientRepository
	deviceRepository repository.IDeviceRepository
}

func NewAuthenticationInteractor(
	clientRepository repository.IClientRepository,
	deviceRepository repository.IDeviceRepository,
) *AuthenticationInteractor {
	return &AuthenticationInteractor{
		clientRepository: clientRepository,
		deviceRepository: deviceRepository,
	}
}

func (i AuthenticationInteractor) AuthenticateClient(ctx context.Context, request dto.AuthenticateRequest) error {
	if request.Secret == nil {
		return apperrs.UnexpectedError
	}

	client, err := i.clientRepository.GetClient(ctx, request.ID)
	if err != nil {
		return err
	}

	return auth.CompareSecret(client.Secret, request.Secret)
}

func (i AuthenticationInteractor) AuthenticateDevice(ctx context.Context, request dto.AuthenticateRequest) error {
	if request.Secret == nil {
		return apperrs.UnexpectedError
	}

	device, err := i.deviceRepository.GetDevice(ctx, request.ID)
	if err != nil {
		return err
	}

	return auth.CompareSecret(device.Secret, request.Secret)
}
