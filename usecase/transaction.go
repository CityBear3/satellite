package usecase

import (
	"context"
	"github.com/CityBear3/satellite/usecase/dto"
)

type Operation func(ctxWithTx context.Context) (dto.IResult, error)

type ITxManager interface {
	DoInTx(ctx context.Context, operation Operation) (dto.IResult, error)
}
