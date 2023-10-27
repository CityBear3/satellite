package mysql

import (
	"context"
	"database/sql"

	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
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

func (t *TxManager) DoInTx(ctx context.Context, operation usecase.Operation) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := operation(tx); err != nil {
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

func ConvertToSqlTx(rtx repository.ITx) (*sql.Tx, error) {
	tx, ok := rtx.(*sql.Tx)
	if !ok {
		return nil, apperrs.UnexpectedError
	}

	return tx, nil
}
