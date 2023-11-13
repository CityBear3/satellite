package entity

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/client"
)

type Client struct {
	ID      primitive.ID
	Name    client.ClientName
	Secret  authentication.Secret
	Devices []Device
}

func NewClient(id primitive.ID, name client.ClientName, secret authentication.Secret, devices []Device) Client {
	return Client{
		ID:      id,
		Name:    name,
		Secret:  secret,
		Devices: devices,
	}
}
