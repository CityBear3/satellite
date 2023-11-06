package entity

import (
	"slices"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type Archive struct {
	ID             primitive.ID
	ArchiveEventID primitive.ID
	ContentType    archive.ContentType
	DeviceID       primitive.ID
}

func NewArchive(id primitive.ID, archiveEventID primitive.ID, contentType archive.ContentType, deviceId primitive.ID) Archive {
	return Archive{
		ID:             id,
		ContentType:    contentType,
		ArchiveEventID: archiveEventID,
		DeviceID:       deviceId,
	}
}

func (a Archive) CheckCorrectCall(client Client) error {
	var deviceIDs []primitive.ID
	for _, device := range client.Devices {
		deviceIDs = append(deviceIDs, device.ID)
	}

	if !slices.Contains(deviceIDs, a.DeviceID) {
		return apperrs.InvalidClientCallingArchiveError
	}

	return nil
}
