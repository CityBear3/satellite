package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/gateway/repository"
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

func (r *EventRepository) SaveArchiveEvent(ctx context.Context, rtx repository.ITx, archiveEvent entity.ArchiveEvent) error {
	tx, err := ConvertToSqlTx(rtx)
	if err != nil {
		return err
	}

	archiveEventSchema := schema.ArchiveEvent{
		ID:          archiveEvent.ID.Value().String(),
		DeviceID:    archiveEvent.DeviceID.Value().String(),
		ClientID:    archiveEvent.ClientID.Value().String(),
		RequestedAt: archiveEvent.RequestedAt,
	}

	if err := archiveEventSchema.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) GetArchiveEvent(ctx context.Context, archiveEventID primitive.ID) (entity.ArchiveEvent, error) {
	event, err := schema.ArchiveEvents(schema.ArchiveEventWhere.ID.EQ(archiveEventID.Value().String())).One(ctx, r.db)
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
