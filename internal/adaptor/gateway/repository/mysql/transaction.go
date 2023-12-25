package mysql

import (
	"context"

	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/jmoiron/sqlx"
)

type Executor interface {
	sqlx.ExtContext
	sqlx.Ext
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type TxManager struct {
	db *sqlx.DB
}

func NewTxManger(db *sqlx.DB) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (t *TxManager) DoInTx(ctx context.Context, operation usecase.Operation) (context.Context, error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, "tx", tx)

	if err := operation(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return ctx, err
		}
		return ctx, err
	}

	if err = tx.Commit(); err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, "tx", nil)
	return ctx, nil
}

func getTxFromCtx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value("tx").(*sqlx.Tx)
	return tx, ok
}
