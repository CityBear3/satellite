package entity

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/client"
)

type Client struct {
	ID      primitive.ID
	Name    client.Name
	Secrets authentication.Secrets
	Devices []Device
}

func NewClient(id primitive.ID, name client.Name, secrets authentication.Secrets, devices []Device) Client {
	return Client{
		ID:      id,
		Name:    name,
		Secrets: secrets,
		Devices: devices,
	}
}
