package minio

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type FileClient struct {
	client *minio.Client
}

func NewFileClient(client *minio.Client) *FileClient {
	return &FileClient{
		client: client,
	}
}

func (f FileClient) CreateObject(ctx context.Context, bucketName string, objectName string, data []byte) error {
	if _, err := f.client.PutObject(
		ctx, bucketName, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{},
	); err != nil {
		return err
	}

	return nil
}

func (f FileClient) GetObject(ctx context.Context, bucketName string, objectName string) ([]byte, error) {
	object, err := f.client.GetObject(
		ctx,
		bucketName,
		objectName,
		minio.GetObjectOptions{},
	)

	if err != nil {
		return nil, err
	}

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

	return buf, nil
}
