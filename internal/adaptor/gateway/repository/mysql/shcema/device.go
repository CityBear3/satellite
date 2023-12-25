package shcema

import (
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/device"
)

type DeviceSchema struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Secrets   string `db:"secrets"`
	ClientID  string `db:"client_id"`
	IsDeleted bool   `db:"is_deleted"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (s DeviceSchema) MapToDevice() (entity.Device, error) {
	id, err := primitive.ParseID(s.ID)
	if err != nil {
		return entity.Device{}, err
	}

	name, err := device.NewDeviceName(s.Name)
	if err != nil {
		return entity.Device{}, err
	}

	secrets, err := authentication.NewHashedSecrets(s.Secrets)
	if err != nil {
		return entity.Device{}, err
	}

	clientID, err := primitive.ParseID(s.ClientID)
	if err != nil {
		return entity.Device{}, err
	}

	return entity.NewDevice(id, name, secrets, clientID), nil
}
