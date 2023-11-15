package usecase

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/gateway/repository"
)

type Operation func(rtx repository.ITx) error

type ITxManager interface {
	DoInTx(ctx context.Context, operation Operation) error
}
