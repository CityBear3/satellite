package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EventRepository struct {
	db boil.ContextExecutor
}

func NewEventRepository(db boil.ContextExecutor) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) SaveArchiveEvent(ctx context.Context, archiveEvent entity.ArchiveEvent) error {
	var exec boil.ContextExecutor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = r.db
	}

	archiveEventSchema := schema.ArchiveEvent{
		ID:          archiveEvent.ID.Value().String(),
		DeviceID:    archiveEvent.DeviceID.Value().String(),
		ClientID:    archiveEvent.ClientID.Value().String(),
		RequestedAt: archiveEvent.RequestedAt,
	}

	if err := archiveEventSchema.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) GetArchiveEvent(ctx context.Context, archiveEventID primitive.ID) (entity.ArchiveEvent, error) {
	var exec boil.ContextExecutor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = r.db
	}

	event, err := schema.ArchiveEvents(schema.ArchiveEventWhere.ID.EQ(archiveEventID.Value().String())).One(ctx, exec)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.ArchiveEvent{}, apperrs.NotFoundArchiveEventError
	}

	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	id, err := primitive.ParseID(event.ID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	deviceID, err := primitive.ParseID(event.DeviceID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	clientID, err := primitive.ParseID(event.ClientID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	return entity.NewArchiveEvent(id, deviceID, clientID), nil
}
