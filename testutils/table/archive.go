package table

import (
	"context"
	"database/sql"
)

type ArchiveTable struct {
	ID             string
	ArchiveEventID string
	ContentType    string
	DeviceId       string
	CreatedAt      any
	UpdatedAt      any
}

func (t ArchiveTable) Insert(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO `archive` (satellite.archive.id, satellite.archive.device_id, satellite.archive.archive_event_id, satellite.archive.content_type) VALUE (?, ?, ?, ?)",
		t.ID, t.DeviceId, t.ArchiveEventID, t.ContentType)
	if err != nil {
		return err
	}

	return nil
}
