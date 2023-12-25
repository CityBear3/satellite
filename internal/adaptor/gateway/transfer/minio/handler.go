package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/minio/minio-go/v7"
)

type FileTransfer struct {
	client     *minio.Client
	bucketName string
}

func NewFileTransfer(client *minio.Client, bucketName string) *FileTransfer {
	return &FileTransfer{
		client:     client,
		bucketName: bucketName,
	}
}

func (f FileTransfer) Save(
	ctx context.Context,
	archiveID primitive.ID,
	contentType archive.ContentType,
	data *archive.Data,
) error {
	if data == nil {
		return nil
	}

	if _, err := f.client.PutObject(
		ctx, f.bucketName,
		fmt.Sprintf("%s.%s", archiveID.String(), contentType.GetExt()),
		bytes.NewReader(data.Chunks),
		int64(data.Size.Value()), minio.PutObjectOptions{},
	); err != nil {
		return err
	}

	return nil
}

func (f FileTransfer) GetFile(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType) (*archive.Data, error) {
	object, err := f.client.GetObject(
		ctx,
		f.bucketName,
		fmt.Sprintf("%s.%s", archiveID.String(), contentType.GetExt()),
		minio.GetObjectOptions{},
	)

	var buf []byte
	for {
		tmp := make([]byte, 1e6)
		_, err := object.Read(tmp)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		buf = append(buf, tmp...)
	}

	data, err := archive.NewData(buf)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
