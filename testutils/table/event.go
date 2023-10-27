package table

import (
	"context"
	"database/sql"
	"time"
)

type ArchiveEventTable struct {
	ID          string
	DeviceID    string
	ClientID    string
	RequestedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t ArchiveEventTable) Insert(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO `archive_event` (id, device_id, client_id, requested_at) VALUE (?, ?, ?, ?)",
		t.ID, t.DeviceID, t.ClientID, t.RequestedAt)
	if err != nil {
		return err
	}

	return nil
}
