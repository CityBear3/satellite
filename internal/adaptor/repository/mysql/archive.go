package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CityBear3/satellite/internal/adaptor/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/internal/usecase/service"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ArchiveRepository struct {
	db           boil.ContextExecutor
	fileTransfer service.IFileService
}

func NewArchiveRepository(db boil.ContextExecutor, fileTransfer service.IFileService) *ArchiveRepository {
	return &ArchiveRepository{
		db:           db,
		fileTransfer: fileTransfer,
	}
}

func (i *ArchiveRepository) Save(
	ctx context.Context,
	archive entity.Archive,
) error {
	var exec boil.ContextExecutor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	archiveSchema := schema.Archive{
		ID:             archive.ID.Value().String(),
		DeviceID:       archive.DeviceID.Value().String(),
		ArchiveEventID: archive.ArchiveEventID.Value().String(),
		ContentType:    archive.ContentType.Value(),
	}

	if err := archiveSchema.Upsert(ctx, exec, boil.Infer(), boil.Infer()); err != nil {
		return err
	}

	if err := i.fileTransfer.Save(ctx, archive.ID, archive.ContentType, archive.Data); err != nil {
		return err
	}

	return nil
}

func (i *ArchiveRepository) GetArchive(
	ctx context.Context,
	archiveID primitive.ID,
) (entity.Archive, error) {
	var exec boil.ContextExecutor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	archiveSchema, err := schema.Archives(schema.ArchiveWhere.ID.EQ(archiveID.Value().String())).One(ctx, exec)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Archive{}, apperrs.NotFoundArchiveError
	}
	if err != nil {
		return entity.Archive{}, err
	}

	id, err := primitive.ParseID(archiveSchema.ID)
	if err != nil {
		return entity.Archive{}, err
	}

	archiveEventID, err := primitive.ParseID(archiveSchema.ArchiveEventID)
	if err != nil {
		return entity.Archive{}, err
	}

	deviceID, err := primitive.ParseID(archiveSchema.DeviceID)
	if err != nil {
		return entity.Archive{}, err
	}

	contentType, err := archive.NewContentType(archiveSchema.ContentType)
	if err != nil {
		return entity.Archive{}, err
	}

	data, err := i.fileTransfer.GetFile(ctx, archiveID, contentType)
	if err != nil {
		return entity.Archive{}, err
	}

	return entity.NewArchive(id, archiveEventID, contentType, deviceID, data), nil
}

func (i *ArchiveRepository) GetArchiveByArchiveEventID(
	ctx context.Context,
	archiveEventID primitive.ID,
) (entity.Archive, error) {
	var exec boil.ContextExecutor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	archiveSchema, err := schema.Archives(schema.ArchiveWhere.ArchiveEventID.EQ(archiveEventID.Value().String())).One(ctx, exec)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Archive{}, apperrs.NotFoundArchiveError
	}

	if err != nil {
		return entity.Archive{}, err
	}

	id, err := primitive.ParseID(archiveSchema.ID)
	if err != nil {
		return entity.Archive{}, err
	}

	archiveEventID, err = primitive.ParseID(archiveSchema.ArchiveEventID)
	if err != nil {
		return entity.Archive{}, err
	}

	deviceID, err := primitive.ParseID(archiveSchema.DeviceID)
	if err != nil {
		return entity.Archive{}, err
	}

	contentType, err := archive.NewContentType(archiveSchema.ContentType)
	if err != nil {
		return entity.Archive{}, err
	}

	data, err := i.fileTransfer.GetFile(ctx, id, contentType)
	if err != nil {
		return entity.Archive{}, err
	}

	return entity.NewArchive(id, archiveEventID, contentType, deviceID, data), nil
}
