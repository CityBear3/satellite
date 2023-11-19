package rpc

import (
	"io"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
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

func (s ArchiveRPCService) handleError(err error) error {
	return convertors.ConvertError(s.logger, err)
}

func (s ArchiveRPCService) CreateArchive(server archivePb.ArchiveService_CreateArchiveServer) error {
	ctx := server.Context()

	device, err := auth.AuthenticatedDevice(ctx)
	if err != nil {
		return s.handleError(err)
	}

	var requests []*archivePb.CreateArchiveRequest
	for {
		request, err := server.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return s.handleError(err)
		}

		requests = append(requests, request)
	}

	request, err := convertors.CreateArchiveRequestToInput(requests)
	if err != nil {
		return s.handleError(err)
	}

	if err := s.uploadArchiveInteractor.CreateArchive(ctx, request, device); err != nil {
		return s.handleError(err)
	}

	if err := server.SendAndClose(&emptypb.Empty{}); err != nil {
		return s.handleError(err)
	}

	return nil
}

func (s ArchiveRPCService) GetArchive(request *archivePb.GetArchiveRequest, server archivePb.ArchiveService_GetArchiveServer) error {
	ctx := server.Context()

	client, err := auth.AuthenticatedClient(ctx)
	if err != nil {
		return s.handleError(err)
	}

	getArchiveRequest, err := convertors.GetArchiveRequestToInput(request)
	if err != nil {
		return s.handleError(err)
	}

	result, err := s.uploadArchiveInteractor.GetArchive(ctx, getArchiveRequest, client)
	if err != nil {
		return s.handleError(err)
	}

	for _, response := range convertors.GetArchiveResultToResponse(result) {
		if err = server.Send(response); err != nil {
			return s.handleError(err)
		}
	}

	return nil
}
