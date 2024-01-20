package service

import (
	"context"
	"fmt"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/infrastructure/minio"
	"go.uber.org/zap"
)

type FileService struct {
	client     *minio.FileClient
	logger     *zap.Logger
	bucketName string
}

func NewFileService(client *minio.FileClient, bucketName string, logger *zap.Logger) *FileService {
	return &FileService{
		client:     client,
		bucketName: bucketName,
		logger:     logger,
	}
}

func (f FileService) Save(
	ctx context.Context,
	archiveID primitive.ID,
	contentType archive.ContentType,
	data archive.Data,
) error {
	if err := f.client.CreateObject(
		ctx,
		f.bucketName,
		fmt.Sprintf("%s.%s", archiveID.String(), contentType.GetExt()),
		data.Chunks,
	); err != nil {
		f.logger.Error("error while saving file", zap.Error(err))
		return err
	}

	return nil
}

func (f FileService) GetFile(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType) (archive.Data, error) {
	buf, err := f.client.GetObject(
		ctx,
		f.bucketName,
		fmt.Sprintf("%s.%s", archiveID.String(), contentType.GetExt()),
	)

	if err != nil {
		f.logger.Error("error while getting file", zap.Error(err))
		return archive.Data{}, err
	}

	data, err := archive.NewData(buf)
	if err != nil {
		return archive.Data{}, err
	}

	return data, nil
}
