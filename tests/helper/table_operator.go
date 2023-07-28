package helper

import (
	"context"
	"database/sql"
)

type TableOperator interface {
	Insert(ctx context.Context, tx *sql.Tx) error
}
