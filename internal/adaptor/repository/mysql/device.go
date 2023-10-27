package mysql

import (
	"context"
	"database/sql"

	schema "github.com/CityBear3/satellite/internal/adaptor/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type DeviceRepository struct {
	db boil.ContextExecutor
}

func NewDeviceRepository(db boil.ContextExecutor) *DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}

func (d *DeviceRepository) GetDevice(ctx context.Context, deviceID primitive.ID) (entity.Device, error) {
	device, err := schema.Devices(schema.DeviceWhere.ID.EQ(deviceID.Value().String())).One(ctx, d.db)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Device{}, apperrs.NotFoundDeviceError
	}

	if err != nil {
		return entity.Device{}, err
	}

	id, err := primitive.ParseID(device.ID)
	if err != nil {
		return entity.Device{}, err
	}

	name, err := primitive.NewDeviceName(device.Name)
	if err != nil {
		return entity.Device{}, err
	}

	secret, err := primitive.NewHashedSecret(device.Secret)
	if err != nil {
		return entity.Device{}, err
	}

	clientID, err := primitive.ParseID(device.ClientID)
	if err != nil {
		return entity.Device{}, err
	}

	return entity.Device{
		ID:       id,
		Name:     name,
		Secret:   secret,
		ClientID: clientID,
	}, nil
}
