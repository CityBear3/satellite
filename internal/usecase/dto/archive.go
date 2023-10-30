package dto

import (
	"github.com/oklog/ulid/v2"
)

type UploadArchiveRequest struct {
	ArchiveEventID string
	ContentType    string
	Data           []byte
}

type GetArchiveRequest struct {
	ArchiveEventID string
}

type GetArchiveResult struct {
	ID             ulid.ULID
	ArchiveEventID ulid.ULID
	ContentType    string
	DeviceID       ulid.ULID
	Data           []byte
	Size           int
}
