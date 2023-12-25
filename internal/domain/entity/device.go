package entity

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/device"
)

type Device struct {
	ID        primitive.ID
	Name      device.DeviceName
	Secrets   authentication.Secrets
	ClientID  primitive.ID
	IsDeleted bool
}

func NewDevice(
	id primitive.ID,
	name device.DeviceName,
	secrets authentication.Secrets,
	clientID primitive.ID,
) Device {
	return Device{
		ID:       id,
		Name:     name,
		Secrets:  secrets,
		ClientID: clientID,
	}
}
