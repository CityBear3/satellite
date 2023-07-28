package mysql

import (
	"context"
	"database/sql"
	"github.com/CityBear3/satellite/domain/model"
	"github.com/CityBear3/satellite/tests/helper"
	"github.com/CityBear3/satellite/tests/table"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"testing"
)

func setTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, "tx", tx)
}

func TestArchiveRepository_Save(t *testing.T) {
	type args struct {
		ctx     context.Context
		archive model.Archive
	}
	type test struct {
		name        string
		args        args
		wantErr     bool
		expectedErr string
		tables      []helper.TableOperator
		queries     []string
	}

	db, err := helper.GetTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	boil.SetDB(db)

	sut := NewArchiveRepository()

	clientId := ulid.Make()
	deviceId := ulid.Make()
	archiveId := ulid.Make()

	tests := []test{
		{
			name: "insert image",
			args: args{
				ctx: context.Background(),
				archive: model.Archive{
					Id:       archiveId,
					Size:     1024,
					Ext:      "jpg",
					DeviceId: deviceId,
				},
			},
			wantErr: false,
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:          clientId.String(),
					Name:        "test",
					Description: "test",
					Secret:      "",
				},
				table.DeviceTable{
					Id:          deviceId.String(),
					Name:        "test",
					Description: "test",
					Secret:      "",
					ClientId:    clientId.String(),
				},
			},
			queries: []string{
				"SELECT * FROM `archive` WHERE `id`=?",
			},
		},
	}

	for _, tt := range tests {
		tx, err := boil.BeginTx(tt.args.ctx, nil)
		if err != nil {
			t.Error(err)
		}
		ctx := setTx(tt.args.ctx, tx)
		for _, operator := range tt.tables {
			if err := operator.Insert(ctx, tx); err != nil {
				t.Error(err)
				return
			}
		}

		savedResult := table.ArchiveTable{}
		t.Run(tt.name, func(t *testing.T) {
			if err = sut.Save(ctx, tt.args.archive); err != nil {
				t.Error(err)
				return
			}
			err := tx.QueryRowContext(ctx, tt.queries[0], tt.args.archive.Id.String()).Scan(&savedResult.Id, &savedResult.DeviceId, &savedResult.Size, &savedResult.Ext, &savedResult.CreatedAt, &savedResult.UpdatedAt)
			if err != nil {
				t.Error(err)
			} else {
				assert.Equal(t, tt.args.archive.Id.String(), savedResult.Id)
				assert.Equal(t, tt.args.archive.Size, savedResult.Size)
				assert.Equal(t, tt.args.archive.Ext, savedResult.Ext)
			}

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
