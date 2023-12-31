package table

import (
	"context"
	"database/sql"
)

type ClientTable struct {
	Id        string
	Name      string
	Secret    string
	CreatedAt any
	UpdatedAt any
}

func (t ClientTable) Insert(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO `client` (id, name, secret) VALUE (?, ?, ?)", t.Id, t.Name, t.Secret)
	if err != nil {
		return err
	}
	return nil
}
