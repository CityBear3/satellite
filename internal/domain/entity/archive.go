package entity

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
)

type Archive struct {
	ID             primitive.ID
	ArchiveEventID primitive.ID
	ContentType    primitive.ContentType
	DeviceID       primitive.ID
}

func NewArchive(id primitive.ID, archiveEventID primitive.ID, contentType primitive.ContentType, deviceId primitive.ID) Archive {
	return Archive{
		ID:             id,
		ContentType:    contentType,
		ArchiveEventID: archiveEventID,
		DeviceID:       deviceId,
	}
}
