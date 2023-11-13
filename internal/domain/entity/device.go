package entity

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/device"
)

type Device struct {
	ID        primitive.ID
	Name      device.DeviceName
	Secret    authentication.Secret
	ClientID  primitive.ID
	IsDeleted bool
}

func NewDevice(
	id primitive.ID,
	name device.DeviceName,
	secret authentication.Secret,
	clientID primitive.ID,
) Device {
	return Device{
		ID:       id,
		Name:     name,
		Secret:   secret,
		ClientID: clientID,
	}
}
