package usecase

import (
	"context"
)

type Operation func(ctx context.Context) error

type ITxManager interface {
	DoInTx(ctx context.Context, operation Operation) error
}
