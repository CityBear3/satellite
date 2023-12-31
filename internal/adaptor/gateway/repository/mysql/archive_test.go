package mysql

import (
	"context"
	"database/sql"
	"testing"

	mock_transfer "github.com/CityBear3/satellite/internal/adaptor/gateway/transfer/mock"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/testutils/helper"
	"github.com/CityBear3/satellite/testutils/table"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestArchiveRepository_Save(t *testing.T) {
	type args struct {
		ctx     context.Context
		tx      *sql.Tx
		archive entity.Archive
	}
	type test struct {
		name    string
		args    args
		tables  []helper.TableOperator
		queries []string
		mocks   func(mockFileTransfer *mock_transfer.MockIFileTransfer)
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

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockFileTransfer := mock_transfer.NewMockIFileTransfer(mockController)

	sut := NewArchiveRepository(tx, mockFileTransfer)

	clientID := ulid.Make()

	archiveEventID := primitive.NewID()

	deviceID := primitive.NewID()

	archiveID := primitive.NewID()

	contentType, err := archive.NewContentType("image/jpeg")
	if err != nil {
		t.Fatal(err)
	}

	tests := []test{
		{
			name: "insert image",
			args: args{
				ctx: context.Background(),
				tx:  tx,
				archive: entity.Archive{
					ID:             archiveID,
					ArchiveEventID: archiveEventID,
					ContentType:    contentType,
					DeviceID:       deviceID,
				},
			},
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:     clientID.String(),
					Name:   "test",
					Secret: "",
				},
				table.DeviceTable{
					Id:       deviceID.Value().String(),
					Name:     "test",
					Secret:   "",
					ClientId: clientID.String(),
				},
			},
			queries: []string{
				"SELECT * FROM `archive` WHERE `id`=?",
			},
			mocks: func(mockFileTransfer *mock_transfer.MockIFileTransfer) {
				mockFileTransfer.EXPECT().Save(ctx, archiveID, contentType, archive.Data{}).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		for _, operator := range tt.tables {
			if err := operator.Insert(ctx, tx); err != nil {
				t.Error(err)
				return
			}
		}

		savedResult := table.ArchiveTable{}
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(mockFileTransfer)
			if err = sut.Save(tt.args.ctx, tt.args.archive); err != nil {
				t.Error(err)
				return
			}

			if err := tx.QueryRowContext(
				tt.args.ctx,
				tt.queries[0],
				tt.args.archive.ID.Value().String()).Scan(
				&savedResult.ID,
				&savedResult.DeviceId,
				&savedResult.ArchiveEventID,
				&savedResult.ContentType,
				&savedResult.CreatedAt,
				&savedResult.UpdatedAt,
			); err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, tt.args.archive.ID.Value().String(), savedResult.ID)
			assert.Equal(t, tt.args.archive.ArchiveEventID.Value().String(), savedResult.ArchiveEventID)
			assert.Equal(t, tt.args.archive.ContentType.Value(), savedResult.ContentType)
			assert.Equal(t, tt.args.archive.DeviceID.Value().String(), savedResult.DeviceId)

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestArchiveRepository_GetArchive(t *testing.T) {
	type args struct {
		ctx context.Context
		id  primitive.ID
	}
	type test struct {
		name        string
		args        args
		want        entity.Archive
		expectedErr error
		tables      []helper.TableOperator
		mocks       func(mockFileTransfer *mock_transfer.MockIFileTransfer)
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

	ctx := context.Background()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockFileTransfer := mock_transfer.NewMockIFileTransfer(mockController)

	clientID := primitive.NewID()
	deviceID := primitive.NewID()
	archiveID := primitive.NewID()
	archiveEventID := primitive.NewID()
	contentType, err := archive.NewContentType("image/jpeg")
	if err != nil {
		t.Fatal(err)
	}

	tests := []test{
		{
			name: "find archive by id",
			args: args{
				ctx,
				archiveID,
			},
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:     clientID.Value().String(),
					Name:   "test",
					Secret: "",
				},
				table.DeviceTable{
					Id:       deviceID.Value().String(),
					Name:     "test",
					Secret:   "",
					ClientId: clientID.Value().String(),
				},
				table.ArchiveTable{
					ID:             archiveID.Value().String(),
					ArchiveEventID: archiveEventID.Value().String(),
					ContentType:    "image/jpeg",
					DeviceId:       deviceID.Value().String(),
				},
			},
			want: entity.Archive{
				ID:             archiveID,
				DeviceID:       deviceID,
				ArchiveEventID: archiveEventID,
				ContentType:    contentType,
			},
			mocks: func(mockFileTransfer *mock_transfer.MockIFileTransfer) {
				mockFileTransfer.EXPECT().GetFile(ctx, archiveID, contentType).Return(archive.Data{}, nil)
			},
		},
		{
			name: "not found error when archive is not exists",
			args: args{
				ctx,
				archiveID,
			},
			expectedErr: apperrs.NotFoundArchiveError,
			mocks: func(mockFileTransfer *mock_transfer.MockIFileTransfer) {
				mockFileTransfer.EXPECT().GetFile(ctx, archiveID, contentType).Return(archive.Data{}, nil).Times(0)
			},
		},
	}

	for _, tt := range tests {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		sut := NewArchiveRepository(tx, mockFileTransfer)

		t.Run(tt.name, func(t *testing.T) {
			for _, operator := range tt.tables {
				if err := operator.Insert(tt.args.ctx, tx); err != nil {
					t.Error(err)
					return
				}
			}

			tt.mocks(mockFileTransfer)
			got, err := sut.GetArchive(tt.args.ctx, tt.args.id)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestArchiveRepository_GetArchiveByArchiveEventID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  primitive.ID
	}
	type test struct {
		name        string
		args        args
		want        entity.Archive
		expectedErr error
		tables      []helper.TableOperator
		mocks       func(mockFileTransfer *mock_transfer.MockIFileTransfer)
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

	ctx := context.Background()

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockFileTransfer := mock_transfer.NewMockIFileTransfer(mockController)

	clientID := primitive.NewID()
	deviceID := primitive.NewID()
	archiveID := primitive.NewID()
	archiveEventID := primitive.NewID()
	contentType, err := archive.NewContentType("image/jpeg")
	if err != nil {
		t.Fatal(err)
	}

	tests := []test{
		{
			name: "find archive by archive event id",
			args: args{
				ctx,
				archiveEventID,
			},
			tables: []helper.TableOperator{
				table.ClientTable{
					Id:     clientID.Value().String(),
					Name:   "test",
					Secret: "",
				},
				table.DeviceTable{
					Id:       deviceID.Value().String(),
					Name:     "test",
					Secret:   "",
					ClientId: clientID.Value().String(),
				},
				table.ArchiveTable{
					ID:             archiveID.Value().String(),
					ArchiveEventID: archiveEventID.Value().String(),
					ContentType:    "image/jpeg",
					DeviceId:       deviceID.Value().String(),
				},
			},
			want: entity.Archive{
				ID:             archiveID,
				DeviceID:       deviceID,
				ArchiveEventID: archiveEventID,
				ContentType:    contentType,
			},
			mocks: func(mockFileTransfer *mock_transfer.MockIFileTransfer) {
				mockFileTransfer.EXPECT().GetFile(ctx, archiveID, contentType).Return(archive.Data{}, nil)
			},
		},
		{
			name: "not found error when archive is not exists",
			args: args{
				ctx,
				primitive.NewID(),
			},
			expectedErr: apperrs.NotFoundArchiveError,
			mocks: func(mockFileTransfer *mock_transfer.MockIFileTransfer) {
				mockFileTransfer.EXPECT().GetFile(ctx, archiveID, contentType).Return(archive.Data{}, nil).Times(0)
			},
		},
	}

	for _, tt := range tests {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		tt.mocks(mockFileTransfer)
		sut := NewArchiveRepository(tx, mockFileTransfer)

		for _, operator := range tt.tables {
			if err := operator.Insert(tt.args.ctx, tx); err != nil {
				t.Error(err)
				return
			}
		}

		t.Run(tt.name, func(t *testing.T) {

			got, err := sut.GetArchiveByArchiveEventID(tt.args.ctx, tt.args.id)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)

			t.Cleanup(func() {
				if err = tx.Rollback(); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
