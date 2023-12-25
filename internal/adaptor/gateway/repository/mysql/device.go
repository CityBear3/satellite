package mysql

import (
	"context"
	"database/sql"
	"errors"

	schema "github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type DeviceRepository struct {
	db Executor
}

func NewDeviceRepository(db Executor) *DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}

const getDeviceQuery = `SELECT 
    	id AS id,
    	name AS name,
    	secrets AS secrets,
    	client_id AS client_id,
    	created_at AS created_at,
    	updated_at AS updated_at
    FROM satellite.device 
    WHERE id = ?;`

func (d *DeviceRepository) GetDevice(ctx context.Context, deviceID primitive.ID) (entity.Device, error) {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = d.db
	}

	var record schema.DeviceSchema
	err := exec.QueryRowxContext(ctx, getDeviceQuery, deviceID.Value().String()).StructScan(&record)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Device{}, apperrs.NotFoundDeviceError
	}

	if err != nil {
		return entity.Device{}, err
	}

	return record.MapToDevice()
}
