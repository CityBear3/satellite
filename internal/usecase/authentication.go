package usecase

import (
	"context"

	"github.com/CityBear3/satellite/internal/usecase/dto"
)

type AuthenticationUseCase interface {
	AuthenticateClient(ctx context.Context, request dto.AuthenticateRequest) error
	AuthenticateDevice(ctx context.Context, request dto.AuthenticateRequest) error
}
