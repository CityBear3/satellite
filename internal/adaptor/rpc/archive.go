package rpc

import (
	"io"
	"net/http"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/adaptor/rpc/validations"
	"github.com/CityBear3/satellite/internal/pkg/auth"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/archive/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArchiveRPCService struct {
	logger                  *zap.Logger
	uploadArchiveInteractor usecase.ArchiveUseCase
	archivePb.UnimplementedArchiveServiceServer
}

func NewArchiveRPCService(logger *zap.Logger, archiveInteractor usecase.ArchiveUseCase) *ArchiveRPCService {
	return &ArchiveRPCService{
		logger:                  logger,
		uploadArchiveInteractor: archiveInteractor,
	}
}

func (s ArchiveRPCService) CreateArchive(server archivePb.ArchiveService_CreateArchiveServer) error {
	ctx := server.Context()

	device, err := auth.AuthenticatedDevice(ctx)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	var meta *archivePb.CreateArchiveMetaInfo
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

	request, err := convertors.ConvertToCreateArchiveRequest(meta.ArchiveEventId, contentType, data)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	if err := s.uploadArchiveInteractor.CreateArchive(ctx, request, device); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	if err := server.SendAndClose(&emptypb.Empty{}); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	return nil
}

func (s ArchiveRPCService) GetArchive(request *archivePb.GetArchiveRequest, server archivePb.ArchiveService_GetArchiveServer) error {
	ctx := server.Context()

	client, err := auth.AuthenticatedClient(ctx)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	getArchiveRequest, err := convertors.ConvertToGetArchiveRequest(request)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	result, err := s.uploadArchiveInteractor.GetArchive(ctx, getArchiveRequest, client)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	if err = server.Send(&archivePb.GetArchiveResponse{Value: &archivePb.GetArchiveResponse_Meta{
		Meta: &archivePb.GetArchiveMetaInfo{
			ArchiveId:   result.ID.String(),
			ContentType: result.ContentType.Value(),
			Size:        int64(result.Size),
		},
	}}); err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	return nil
}
