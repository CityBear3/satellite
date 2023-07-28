package mysql

import (
	"context"
	"database/sql"
	"github.com/CityBear3/satellite/logic/dto"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TxManager struct{}

func NewTxManger() *TxManager {
	return &TxManager{}
}

func (t *TxManager) DoInTxWithResult(ctx context.Context, operation func(ctx context.Context) (dto.Result, error)) (dto.Result, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "tx", tx)

	result, err := operation(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TxManager) DoInTx(ctx context.Context, operation func(ctx2 context.Context) error) error {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, "tx", tx)

	if err = operation(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func GetTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok || tx == nil {
		return nil, errors.New("Database transaction was not found.")
	}
	return tx, nil
}
