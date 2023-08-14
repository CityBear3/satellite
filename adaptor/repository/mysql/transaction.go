package mysql

import (
	"context"
	"database/sql"
	"github.com/CityBear3/satellite/usecase"
	"github.com/CityBear3/satellite/usecase/dto"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TxManager struct{}

func NewTxManger() *TxManager {
	return &TxManager{}
}

func (t *TxManager) DoInTx(ctx context.Context, operation usecase.Operation) (dto.IResult, error) {
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

func GetTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok || tx == nil {
		return nil, errors.New("Database transaction was not found.")
	}
	return tx, nil
}
