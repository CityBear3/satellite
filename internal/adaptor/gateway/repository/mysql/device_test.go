package mysql

import (
	"context"
	"testing"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/device"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/testutils/helper"
	"github.com/CityBear3/satellite/testutils/table"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestDeviceRepository_GetDevice(t *testing.T) {
	ctx := context.Background()

	db, err := helper.GetTestDB()
	if err != nil {
		t.Fatal(err)
	}

	defer func(db *sqlx.DB) {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}(db)

	deviceName, err := device.NewDeviceName("test")
	if err != nil {
		t.Fatal(err)
	}
	secret, err := helper.GenerateSecret()
	if err != nil {
		t.Fatal(err)
	}

	device := entity.NewDevice(primitive.NewID(), deviceName, secret, primitive.NewID())

	secretValue, err := secret.Value()
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx      context.Context
		deviceID primitive.ID
	}

	tests := []struct {
		name        string
		args        args
		want        entity.Device
		expectedErr error
		tables      []helper.TableOperator
	}{
		{
			name: "normal case",
			args: args{
				ctx:      ctx,
				deviceID: device.ID,
			},
			want: device,
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:     device.ClientID.Value().String(),
					Name:   "test",
					Secret: "test",
				},
				table.DeviceTable{
					Id:       device.ID.Value().String(),
					Name:     device.Name.String(),
					Secret:   secretValue,
					ClientId: device.ClientID.Value().String(),
				},
			},
		},
		{
			name: "not found device",
			args: args{
				ctx:      ctx,
				deviceID: primitive.NewID(),
			},
			expectedErr: apperrs.NotFoundDeviceError,
		},
	}

	for _, tt := range tests {
		tx, err := db.BeginTxx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		sut := NewDeviceRepository(tx)

		for _, operator := range tt.tables {
			if err := operator.Insert(ctx, tx.Tx); err != nil {
				t.Error(err)
				return
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.GetDevice(tt.args.ctx, tt.args.deviceID)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.ClientID, got.ClientID)
			assert.Equal(t, tt.want.IsDeleted, got.IsDeleted)

			insertedSecret, err := got.Secrets.Value()
			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, secretValue, insertedSecret)

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
