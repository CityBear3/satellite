package table

import (
	"context"
	"database/sql"
)

type DeviceTable struct {
	Id          string
	Name        string
	Description string
	Secret      string
	ClientId    string
}

func (t DeviceTable) Insert(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO `device` (id, name, description, secret, client_id) VALUE (?, ?, ?, ?, ?)",
		t.Id, t.Name, t.Description, t.Secret, t.ClientId)
	if err != nil {
		return err
	}
	return nil
}