package shcema

import (
	"time"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type ArchiveEventSchema struct {
	ID          string    `db:"id"`
	DeviceID    string    `db:"device_id"`
	ClientID    string    `db:"client_id"`
	RequestedAt time.Time `db:"requested_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (s ArchiveEventSchema) MapToArchiveEvent() (entity.ArchiveEvent, error) {
	id, err := primitive.ParseID(s.ID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	deviceID, err := primitive.ParseID(s.DeviceID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	clientID, err := primitive.ParseID(s.ClientID)
	if err != nil {
		return entity.ArchiveEvent{}, err
	}

	return entity.ArchiveEvent{
		ID:          id,
		DeviceID:    deviceID,
		ClientID:    clientID,
		RequestedAt: s.RequestedAt,
	}, nil
}
