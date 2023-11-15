package mysql

import (
	"context"
	"database/sql"

	"github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/client"
	"github.com/CityBear3/satellite/internal/domain/primitive/device"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ClientRepository struct {
	db boil.ContextExecutor
}

func NewClientRepository(db boil.ContextExecutor) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (i *ClientRepository) GetClient(ctx context.Context, clientID primitive.ID) (entity.Client, error) {
	clientSchema, err := schema.Clients(
		schema.ClientWhere.ID.EQ(clientID.Value().String()),
		qm.Load("Devices"),
	).One(ctx, i.db)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Client{}, apperrs.NotFoundClientError
	}

	if err != nil {
		return entity.Client{}, err
	}

	clientName, err := client.NewClientName(clientSchema.Name)
	if err != nil {
		return entity.Client{}, err
	}

	secret, err := authentication.NewHashedSecret(clientSchema.Secret)
	if err != nil {
		return entity.Client{}, err
	}

	var deviceEntities []entity.Device
	for _, d := range clientSchema.R.Devices {
		deviceID, err := primitive.ParseID(d.ID)
		if err != nil {
			return entity.Client{}, err
		}

		deviceName, err := device.NewDeviceName(d.Name)
		if err != nil {
			return entity.Client{}, err
		}

		deviceSecret, err := authentication.NewHashedSecret(d.Secret)
		if err != nil {
			return entity.Client{}, err
		}

		deviceEntities = append(deviceEntities, entity.NewDevice(
			deviceID,
			deviceName,
			deviceSecret,
			clientID,
		))
	}

	return entity.Client{
		ID:      clientID,
		Name:    clientName,
		Secret:  secret,
		Devices: deviceEntities,
	}, nil
}
