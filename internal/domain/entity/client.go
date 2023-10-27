package entity

import "github.com/CityBear3/satellite/internal/domain/primitive"

type Client struct {
	ID      primitive.ID
	Name    primitive.ClientName
	Secret  primitive.Secret
	Devices []Device
}

func NewClient(id primitive.ID, name primitive.ClientName, secret primitive.Secret, devices []Device) Client {
	return Client{
		ID:      id,
		Name:    name,
		Secret:  secret,
		Devices: devices,
	}
}
