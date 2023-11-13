package convertors

import (
	"net/http"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/archive/v1"
)

func CreateArchiveRequestToInput(requests []*archivePb.CreateArchiveRequest) (usecase.CreateArchiveInput, error) {
	var meta *archivePb.CreateArchiveMetaInfo
	var chunks []byte

	for _, request := range requests {
		if m := request.GetMeta(); m != nil {
			meta = m
		}

		chunks = append(chunks, request.GetChunk()...)
	}

	if meta == nil {
		return usecase.CreateArchiveInput{}, apperrs.UnexpectedError
	}

	id, err := primitive.ParseID(meta.ArchiveEventId)
	if err != nil {
		return usecase.CreateArchiveInput{}, err
	}

	contentType, err := archive.NewContentType(http.DetectContentType(chunks))
	if err != nil {
		return usecase.CreateArchiveInput{}, err
	}

	data, err := archive.NewData(chunks)
	if err != nil {
		return usecase.CreateArchiveInput{}, err
	}

	return usecase.CreateArchiveInput{
		ArchiveEventID: id,
		ContentType:    contentType,
		Data:           data,
	}, nil
}

func GetArchiveRequestToInput(request *archivePb.GetArchiveRequest) (usecase.GetArchiveInput, error) {
	archiveEventID, err := primitive.ParseID(request.GetArchiveEventId())
	if err != nil {
		return usecase.GetArchiveInput{}, err
	}

	return usecase.GetArchiveInput{
		ArchiveEventID: archiveEventID,
	}, nil
}

func GetArchiveResultToResponse(result usecase.GetArchiveResult) []*archivePb.GetArchiveResponse {
	size := result.Data.Size.Value()
	responses := []*archivePb.GetArchiveResponse{
		{
			Value: &archivePb.GetArchiveResponse_Meta{
				Meta: &archivePb.GetArchiveMetaInfo{
					ArchiveId:   result.ID.String(),
					ContentType: result.ContentType.Value(),
					Size:        int64(size),
				},
			},
		},
	}

	step := 2000000
	switch {
	case size <= step:
		responses = append(responses, &archivePb.GetArchiveResponse{
			Value: &archivePb.GetArchiveResponse_Chunk{
				Chunk: result.Data.Chunks,
			},
		})
	default:
		for i := 0; i < size; i += step {
			end := i + step
			if end > result.Data.Size.Value() {
				end = result.Data.Size.Value()
			}

			responses = append(responses, &archivePb.GetArchiveResponse{
				Value: &archivePb.GetArchiveResponse_Chunk{
					Chunk: result.Data.Chunks[i:end],
				},
			})
		}
	}

	return responses
}
