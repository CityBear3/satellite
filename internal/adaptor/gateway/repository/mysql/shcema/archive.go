package shcema

import (
	"time"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
)

type ArchiveScheme struct {
	ID             string    `db:"id"`
	ArchiveEventID string    `db:"archive_event_id"`
	ContentType    string    `db:"content_type"`
	DeviceID       string    `db:"device_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (s ArchiveScheme) MapToArchive(data *archive.Data) (entity.Archive, error) {
	id, err := primitive.ParseID(s.ID)
	if err != nil {
		return entity.Archive{}, err
	}

	archiveEventID, err := primitive.ParseID(s.ArchiveEventID)
	if err != nil {
		return entity.Archive{}, err
	}

	deviceID, err := primitive.ParseID(s.DeviceID)
	if err != nil {
		return entity.Archive{}, err
	}

	contentType, err := archive.NewContentType(s.ContentType)
	if err != nil {
		return entity.Archive{}, err
	}

	return entity.NewArchive(id, archiveEventID, contentType, deviceID, data), nil
}
