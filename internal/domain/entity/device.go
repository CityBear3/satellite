package entity

import "github.com/CityBear3/satellite/internal/domain/primitive"

type Device struct {
	ID        primitive.ID
	Name      primitive.DeviceName
	Secret    primitive.Secret
	ClientID  primitive.ID
	IsDeleted bool
}

func NewDevice(
	id primitive.ID,
	name primitive.DeviceName,
	secret primitive.Secret,
	clientID primitive.ID,
) Device {
	return Device{
		ID:       id,
		Name:     name,
		Secret:   secret,
		ClientID: clientID,
	}
}
