package mysql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/testutils/helper"
	"github.com/CityBear3/satellite/testutils/table"
	"github.com/stretchr/testify/assert"
)

func TestEventRepository_SaveArchiveEvent(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx          context.Context
		archiveEvent entity.ArchiveEvent
	}

	db, err := helper.GetTestDB()
	if err != nil {
		t.Fatal(err)
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}(db)

	eventId := primitive.NewID()
	clientID := primitive.NewID()
	deviceID := primitive.NewID()

	tests := []struct {
		name        string
		args        args
		expectedErr error
		tables      []helper.TableOperator
		queries     []string
	}{
		{
			name: "normal case",
			args: args{
				ctx:          ctx,
				archiveEvent: entity.NewArchiveEvent(eventId, deviceID, clientID),
			},
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:     clientID.Value().String(),
					Name:   "test",
					Secret: "test",
				},
				table.DeviceTable{
					Id:       deviceID.Value().String(),
					Name:     "test",
					Secret:   "test",
					ClientId: clientID.Value().String(),
				},
			},
			queries: []string{
				"SELECT * FROM `archive_event` WHERE `id`=?",
			},
		},
	}

	for _, tt := range tests {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		sut := NewEventRepository(tx)

		for _, operator := range tt.tables {
			if err := operator.Insert(tt.args.ctx, tx); err != nil {
				t.Error(err)
				return
			}
		}

		saveResult := table.ArchiveEventTable{}
		t.Run(tt.name, func(t *testing.T) {
			if err := sut.SaveArchiveEvent(tt.args.ctx, tx, tt.args.archiveEvent); err != nil {
				t.Error(err)
				return
			}

			if err := tx.QueryRowContext(
				tt.args.ctx,
				tt.queries[0],
				tt.args.archiveEvent.ID.Value().String(),
			).Scan(
				&saveResult.ID,
				&saveResult.DeviceID,
				&saveResult.ClientID,
				&saveResult.RequestedAt,
				&saveResult.CreatedAt,
				&saveResult.UpdatedAt,
			); err != nil {
				t.Error(err)
				return
			}

			want := tt.args.archiveEvent
			assert.Equal(t, want.ID.Value().String(), saveResult.ID)
			assert.Equal(t, want.DeviceID.Value().String(), saveResult.DeviceID)
			assert.Equal(t, want.ClientID.Value().String(), saveResult.ClientID)
			assert.Equal(t, want.RequestedAt.UTC().Truncate(time.Minute), saveResult.RequestedAt.Truncate(time.Minute))

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestEventRepository_GetArchiveEvent(t *testing.T) {
	type args struct {
		ctx            context.Context
		archiveEventID primitive.ID
	}

	db, err := helper.GetTestDB()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	sut := NewEventRepository(tx)

	archiveEventID := primitive.NewID()
	deviceID := primitive.NewID()
	clientID := primitive.NewID()

	tests := []struct {
		name        string
		args        args
		want        entity.ArchiveEvent
		expectedErr error
		tables      []helper.TableOperator
	}{
		{
			name: "normal case",
			args: args{
				ctx:            ctx,
				archiveEventID: archiveEventID,
			},
			want: entity.NewArchiveEvent(
				archiveEventID,
				deviceID,
				clientID,
			),
			tables: []helper.TableOperator{
				table.ArchiveEventTable{
					ID:          archiveEventID.Value().String(),
					DeviceID:    deviceID.Value().String(),
					ClientID:    clientID.Value().String(),
					RequestedAt: time.Now(),
				},
			},
		},
	}

	for _, tt := range tests {
		for _, operator := range tt.tables {
			if err := operator.Insert(tt.args.ctx, tx); err != nil {
				t.Error(err)
				return
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.GetArchiveEvent(tt.args.ctx, tt.args.archiveEventID)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.DeviceID, got.DeviceID)
			assert.Equal(t, tt.want.ClientID, got.ClientID)
			assert.Equal(t, tt.want.RequestedAt.Truncate(time.Second), got.RequestedAt.Truncate(time.Second))

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
