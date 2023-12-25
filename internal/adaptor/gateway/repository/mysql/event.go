package mysql

import (
	"context"
	"database/sql"
	"errors"

	schema "github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type EventRepository struct {
	db Executor
}

func NewEventRepository(db Executor) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

const saveArchiveEventQuery = `
INSERT INTO archive_event (id, device_id, client_id, requested_at) VALUE (?, ?, ?, ?) AS new
	ON DUPLICATE KEY UPDATE
	                    id = new.id,
	                 	device_id = new.device_id,
	                    client_id = new.client_id,
	                	requested_at = new.requested_at;
`

func (r *EventRepository) SaveArchiveEvent(ctx context.Context, archiveEvent entity.ArchiveEvent) error {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = r.db
	}

	if _, err := exec.ExecContext(ctx, saveArchiveEventQuery, archiveEvent.ID.Value().String(), archiveEvent.DeviceID.Value().String(),
		archiveEvent.ClientID.Value().String(), archiveEvent.RequestedAt); err != nil {
		return err
	}

	return nil
}

const getArchiveEventQuery = `
SELECT * FROM archive_event WHERE id = ?;
`

func (r *EventRepository) GetArchiveEvent(ctx context.Context, archiveEventID primitive.ID) (entity.ArchiveEvent, error) {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = r.db
	}

	var record schema.ArchiveEventSchema
	err := exec.QueryRowxContext(ctx, getArchiveEventQuery, archiveEventID.Value().String()).StructScan(&record)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.ArchiveEvent{}, apperrs.NotFoundArchiveEventError
	}

	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	id, err := primitive.ParseID(record.ID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	deviceID, err := primitive.ParseID(record.DeviceID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	clientID, err := primitive.ParseID(record.ClientID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	return entity.NewArchiveEvent(id, deviceID, clientID), nil
}
