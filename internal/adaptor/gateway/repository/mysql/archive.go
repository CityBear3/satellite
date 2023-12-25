package mysql

import (
	"context"
	"database/sql"
	"errors"

	schema "github.com/CityBear3/satellite/internal/adaptor/gateway/repository/mysql/shcema"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/gateway/transfer"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type ArchiveRepository struct {
	db           Executor
	fileTransfer transfer.IFileTransfer
}

func NewArchiveRepository(db Executor, fileTransfer transfer.IFileTransfer) *ArchiveRepository {
	return &ArchiveRepository{
		db:           db,
		fileTransfer: fileTransfer,
	}
}

const saveArchiveCmd = `
INSERT INTO archive (id, archive_event_id, content_type, device_id) VALUE (?, ?, ?, ?) AS new
	ON DUPLICATE KEY UPDATE
		id = new.id,
		archive_event_id = new.archive_event_id,
		content_type = new.content_type,
		device_id = new.device_id;
`

func (i *ArchiveRepository) Save(
	ctx context.Context,
	archive entity.Archive,
) error {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	if _, err := exec.ExecContext(ctx, saveArchiveCmd, archive.ID.Value().String(), archive.ArchiveEventID.Value().String(),
		archive.ContentType.Value(), archive.DeviceID.Value().String()); err != nil {
		return err
	}

	if err := i.fileTransfer.Save(ctx, archive.ID, archive.ContentType, archive.Data); err != nil {
		return err
	}

	return nil
}

const getArchiveQuery = `
SELECT id, archive_event_id, content_type, device_id FROM archive WHERE id = ?;
`

func (i *ArchiveRepository) GetArchive(
	ctx context.Context,
	archiveID primitive.ID,
) (entity.Archive, error) {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	var record schema.ArchiveScheme
	err := exec.QueryRowxContext(ctx, getArchiveQuery, archiveID.Value().String()).StructScan(&record)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Archive{}, apperrs.NotFoundArchiveError
	}

	if err != nil {
		return entity.Archive{}, err
	}

	var data *archive.Data
	archiveEntity, err := record.MapToArchive(data)
	if err != nil {
		return entity.Archive{}, err
	}

	data, err = i.fileTransfer.GetFile(ctx, archiveEntity.ID, archiveEntity.ContentType)
	if err != nil {
		return entity.Archive{}, err
	}

	return archiveEntity, nil
}

const getArchiveByArchiveEventIDQuery = `
SELECT id, archive_event_id, content_type, device_id FROM archive WHERE archive_event_id = ?;
`

func (i *ArchiveRepository) GetArchiveByArchiveEventID(
	ctx context.Context,
	archiveEventID primitive.ID,
) (entity.Archive, error) {
	var exec Executor
	exec, ok := getTxFromCtx(ctx)
	if !ok {
		exec = i.db
	}

	var record schema.ArchiveScheme
	err := exec.QueryRowxContext(ctx, getArchiveByArchiveEventIDQuery, archiveEventID.Value().String()).StructScan(&record)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Archive{}, apperrs.NotFoundArchiveError
	}

	if err != nil {
		return entity.Archive{}, err
	}

	var data *archive.Data
	archiveEntity, err := record.MapToArchive(data)
	if err != nil {
		return entity.Archive{}, err
	}

	data, err = i.fileTransfer.GetFile(ctx, archiveEntity.ID, archiveEntity.ContentType)
	if err != nil {
		return entity.Archive{}, err
	}

	return archiveEntity, nil
}
