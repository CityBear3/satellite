package rpc

import (
	"io"
	"net/http"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/adaptor/rpc/validations"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/internal/usecase/dto"
	"github.com/CityBear3/satellite/pb/archive/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArchiveRPCService struct {
	logger                  *zap.Logger
	uploadArchiveInteractor usecase.ArchiveUseCase
	archive.UnimplementedArchiveServiceServer
}

func NewArchiveRPCService(logger *zap.Logger, archiveInteractor usecase.ArchiveUseCase) *ArchiveRPCService {
	return &ArchiveRPCService{
		logger:                  logger,
		uploadArchiveInteractor: archiveInteractor,
	}
}

func (s ArchiveRPCService) CreateArchive(server archive.ArchiveService_CreateArchiveServer) error {
	ctx := server.Context()

	device, err := AuthenticatedDevice(ctx)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	var meta *archive.Meta
	var data []byte
	for {
		request, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return convertors.ConvertError(s.logger, err)
		}

		if m := request.GetMeta(); m != nil {
			meta = m
		}
		if c := request.GetChunk(); c != nil {
			data = append(data, c...)
		}
	}

	contentType := http.DetectContentType(data)

	if err := validations.ValidateCreateArchive(meta, data); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	request := dto.UploadArchiveRequest{
		ArchiveEventID: meta.ArchiveEventId,
		ContentType:    contentType,
		Data:           data,
	}
	if err := s.uploadArchiveInteractor.CreateArchive(ctx, request, device); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	if err := server.SendAndClose(&emptypb.Empty{}); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	return nil
}

func (s ArchiveRPCService) GetArchive(request *archive.GetArchiveRequest, server archive.ArchiveService_GetArchiveServer) error {
	ctx := server.Context()

	client, err := AuthenticatedClient(ctx)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	getArchiveRequest := dto.GetArchiveRequest{ArchiveEventID: request.ArchiveEventId}

	result, err := s.uploadArchiveInteractor.GetArchive(ctx, getArchiveRequest, client)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	if err = server.Send(&archive.GetArchiveResponse{Value: &archive.GetArchiveResponse_Meta{
		Meta: &archive.Meta{
			ArchiveEventId: result.ID.String(),
		},
	}}); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	return nil
}
