package dto

import (
	"github.com/oklog/ulid/v2"
)

type Result interface{}

type UploadArchiveRequest struct {
	Id          ulid.ULID
	ContentType string
	Data        []byte
	Result
}

type GetArchiveRequest struct {
	Id ulid.ULID
}

type GetArchiveResult struct {
	Id   ulid.ULID
	Ext  string
	Data []byte
	Result
}

func NewUploadArchiveRequest(id ulid.ULID, contentType string, data []byte) UploadArchiveRequest {
	return UploadArchiveRequest{
		Id:          id,
		ContentType: contentType,
		Data:        data,
	}
}