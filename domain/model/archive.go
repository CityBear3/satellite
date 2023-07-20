package model

import (
	"context"
	"github.com/oklog/ulid/v2"
)

type Archive struct {
	Id       ulid.ULID
	Ext      string
	Size     int
	DeviceId ulid.ULID
	Data     []byte
}

type IArchiveRepository interface {
	Save(ctx context.Context, archive Archive) error
	GetArchive(ctx context.Context, imageId ulid.ULID) (*Archive, error)
}

func NewImageModel(id ulid.ULID, ext string, size int, deviceId ulid.ULID, data []byte) *Archive {
	return &Archive{
		Id:       id,
		Ext:      ext,
		Size:     size,
		DeviceId: deviceId,
		Data:     data,
	}
}

func (i *Archive) CheckDataSize(limit int) bool {
	return i.Size <= limit
}
