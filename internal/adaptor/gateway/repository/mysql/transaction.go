package mysql

import (
	"context"
	"database/sql"

	"github.com/CityBear3/satellite/internal/usecase"
)

type TxManager struct {
	db *sql.DB
}

func NewTxManger(db *sql.DB) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (t *TxManager) DoInTx(ctx context.Context, operation usecase.Operation) (context.Context, error) {
	tx, err := t.db.BeginTx(ctx, nil)
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

func getTxFromCtx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	return tx, ok
}
