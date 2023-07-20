package logic

import (
	"context"
	"github.com/CityBear3/satellite/logic/dto"
)

type ITxManager interface {
	DoInTxWithResult(ctx context.Context, operation func(context.Context) (dto.Result, error)) (dto.Result, error)
	DoInTx(ctx context.Context, operation func(context.Context) error) error
}
