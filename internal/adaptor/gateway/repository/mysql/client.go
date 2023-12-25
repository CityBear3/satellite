package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/domain/primitive/client"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type ClientRepository struct {
	db Executor
}

func NewClientRepository(db Executor) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

const getClientQuery = `
SELECT 
	c.id,
	c.name,
	c.secrets,
	c.created_at,
	c.updated_at,
	
	d.id AS "device.id",
	d.name AS "device.name",
	d.secrets AS "device.secrets",
	d.client_id AS "device.client_id",
	d.created_at AS "device.created_at",
	d.updated_at AS "device.updated_at"
FROM satellite.client AS c
LEFT JOIN satellite.device AS d ON d.client_id = c.id
WHERE c.id = ?;
`

func (i *ClientRepository) GetClient(ctx context.Context, clientID primitive.ID) (entity.Client, error) {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	var records []shcema.ClientWithDevicesSchema
	err := exec.SelectContext(ctx, &records, getClientQuery, clientID.Value().String())

	if errors.Is(err, sql.ErrNoRows) || len(records) == 0 {
		return entity.Client{}, apperrs.NotFoundClientError
	}

	if err != nil {
		return entity.Client{}, err
	}

	clientName, err := client.NewName(records[0].Name)
	if err != nil {
		return entity.Client{}, err
	}

	secret, err := authentication.NewHashedSecrets(records[0].Secrets)
	if err != nil {
		return entity.Client{}, err
	}

	var deviceEntities []entity.Device
	for _, r := range records {
		deviceEntity, err := r.Device.MapToDevice()
		if err != nil {
			return entity.Client{}, err
		}

		deviceEntities = append(deviceEntities, deviceEntity)
	}

	return entity.Client{
		ID:      clientID,
		Name:    clientName,
		Secrets: secret,
		Devices: deviceEntities,
	}, nil
}
