package entity

import (
	"time"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type ArchiveEvent struct {
	ID          primitive.ID
	DeviceID    primitive.ID
	ClientID    primitive.ID
	RequestedAt time.Time
}

func NewArchiveEvent(id, deviceID, clientID primitive.ID) ArchiveEvent {
	return ArchiveEvent{
		ID:          id,
		DeviceID:    deviceID,
		ClientID:    clientID,
		RequestedAt: time.Now(),
	}
}

func (a ArchiveEvent) CheckCorrectCall(deviceID, clientID primitive.ID, now time.Time) error {
	if a.DeviceID != deviceID {
		return apperrs.InvalidEventDeviceIDError
	}

	if a.ClientID != clientID {
		return apperrs.InvalidEventClientIDError
	}

	if now.After(a.RequestedAt.Add(30 * time.Second)) {
		return apperrs.EventTimeOutError
	}

	return nil
}
