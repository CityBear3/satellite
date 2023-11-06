package convertors

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/archive/v1"
)

func ConvertToCreateArchiveRequest(archiveEventID string, contentType string, data []byte) (usecase.CreateArchiveRequest, error) {
	id, err := primitive.ParseID(archiveEventID)
	if err != nil {
		return usecase.CreateArchiveRequest{}, err
	}

	ct, err := archive.NewContentType(contentType)
	if err != nil {
		return usecase.CreateArchiveRequest{}, err
	}

	return usecase.CreateArchiveRequest{
		ArchiveEventID: id,
		ContentType:    ct,
		Data:           data,
	}, nil
}

func ConvertToGetArchiveRequest(request *archivePb.GetArchiveRequest) (usecase.GetArchiveRequest, error) {
	archiveEventID, err := primitive.ParseID(request.GetArchiveEventId())
	if err != nil {
		return usecase.GetArchiveRequest{}, err
	}

	return usecase.GetArchiveRequest{
		ArchiveEventID: archiveEventID,
	}, nil
}
