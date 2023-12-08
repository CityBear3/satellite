//go:generate mockgen -source=$GOFILE -package=mock_usecase -destination=./mock/$GOFILE
package usecase

import (
	"context"
)

type Operation func(ctx context.Context) error

type ITxManager interface {
	DoInTx(ctx context.Context, operation Operation) (context.Context, error)
}
