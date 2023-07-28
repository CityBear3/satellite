package mysql

import (
	"context"
	"github.com/CityBear3/satellite/adaptor/repository/mysql/shcema"
	"github.com/CityBear3/satellite/domain/model"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ArchiveRepository struct{}

func NewArchiveRepository() *ArchiveRepository {
	return &ArchiveRepository{}
}

func (i *ArchiveRepository) Save(ctx context.Context, archive model.Archive) error {
	tx, err := GetTransaction(ctx)
	if err != nil {
		return err
	}

	archiveSchema := schema.Archive{
		ID:       archive.Id.String(),
		Size:     archive.Size,
		Ext:      archive.Ext,
		DeviceID: archive.DeviceId.String(),
	}
	if err := archiveSchema.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (i *ArchiveRepository) GetArchive(ctx context.Context, imageId ulid.ULID) (*model.Archive, error) {
	//TODO implement me
	panic("implement me")
}
