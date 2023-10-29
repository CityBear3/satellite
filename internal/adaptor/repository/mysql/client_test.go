package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/testutils/helper"
	"github.com/CityBear3/satellite/testutils/table"
	"github.com/stretchr/testify/assert"
)

func TestClientRepository_GetClient(t *testing.T) {
	ctx := context.Background()

	db, err := helper.GetTestDB()
	if err != nil {
		t.Fatal(err)
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}(db)

	clientID := primitive.NewID()

	generatedSecret, err := helper.GenerateSecret()
	if err != nil {
		t.Fatal(err)
	}
	secretValue, err := generatedSecret.Value()
	if err != nil {
		t.Fatal(err)
	}

	deviceName, err := primitive.NewDeviceName("test")
	if err != nil {
		t.Fatal(err)
	}

	device := entity.NewDevice(primitive.NewID(), deviceName, generatedSecret, clientID)

	clientName, err := primitive.NewClientName("test")
	if err != nil {
		t.Fatal(err)
	}

	client := entity.NewClient(clientID, clientName, generatedSecret, []entity.Device{device})

	type args struct {
		ctx      context.Context
		clientID primitive.ID
	}

	tests := []struct {
		name        string
		args        args
		want        entity.Client
		expectedErr error
		tables      []helper.TableOperator
	}{
		{
			name: "normal case",
			args: args{
				ctx:      ctx,
				clientID: clientID,
			},
			want: client,
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:     clientID.Value().String(),
					Name:   client.Name.String(),
					Secret: secretValue,
				},
				table.DeviceTable{
					Id:       device.ID.Value().String(),
					Name:     device.Name.String(),
					Secret:   secretValue,
					ClientId: clientID.Value().String(),
				},
			},
		},
		{
			name: "not found client",
			args: args{
				ctx:      ctx,
				clientID: primitive.NewID(),
			},
			expectedErr: apperrs.NotFoundClientError,
		},
	}

	for _, tt := range tests {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		sut := NewClientRepository(tx)

		for _, operator := range tt.tables {
			if err := operator.Insert(ctx, tx); err != nil {
				t.Error(err)
				return
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.GetClient(tt.args.ctx, tt.args.clientID)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, len(tt.want.Devices), len(got.Devices))

			value, err := got.Secret.Value()
			assert.NoError(t, err)

			assert.Equal(t, secretValue, value)

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
