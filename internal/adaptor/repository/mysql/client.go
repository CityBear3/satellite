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
	client, err := schema.Clients(
		schema.ClientWhere.ID.EQ(clientID.Value().String()),
		qm.Load("Devices"),
	).One(ctx, i.db)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Client{}, apperrs.NotFoundClientError
	}

	if err != nil {
		return entity.Client{}, err
	}

	clientName, err := primitive.NewClientName(client.Name)
	if err != nil {
		return entity.Client{}, err
	}

	secret, err := primitive.NewHashedSecret(client.Secret)
	if err != nil {
		return entity.Client{}, err
	}

	var deviceEntitys []entity.Device
	for _, device := range client.R.Devices {
		deviceID, err := primitive.ParseID(device.ID)
		if err != nil {
			return entity.Client{}, err
		}

		deviceName, err := primitive.NewDeviceName(device.Name)
		if err != nil {
			return entity.Client{}, err
		}

		deviceSecret, err := primitive.NewHashedSecret(device.Secret)
		if err != nil {
			return entity.Client{}, err
		}

		deviceEntitys = append(deviceEntitys, entity.NewDevice(
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
		Devices: deviceEntitys,
	}, nil
}
