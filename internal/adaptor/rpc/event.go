package rpc

import (
	"context"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/pkg/auth"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/event/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventRPCService struct {
	logger          *zap.Logger
	eventInteractor usecase.EventUseCase
	eventPb.UnimplementedArchiveEventServiceServer
}

func NewEventRPCService(logger *zap.Logger, eventInteractor usecase.EventUseCase) *EventRPCService {
	return &EventRPCService{
		logger:          logger,
		eventInteractor: eventInteractor,
	}
}

func (s EventRPCService) handleError(err error) error {
	return convertors.ConvertError(s.logger, err)
}

func (s EventRPCService) PublishEvent(ctx context.Context, _ *emptypb.Empty) (*eventPb.PublishEventResponse, error) {
	client, err := auth.AuthenticatedClient(ctx)
	if err != nil {
		return nil, s.handleError(err)
	}

	archiveEventID, err := s.eventInteractor.PublishArchiveEvent(ctx, client)
	if err != nil {
		return nil, s.handleError(err)
	}

	return &eventPb.PublishEventResponse{
		ArchiveEventId: archiveEventID.Value().String(),
	}, nil
}

func (s EventRPCService) ReceiveEvent(_ *emptypb.Empty, server eventPb.ArchiveEventService_ReceiveEventServer) error {
	ctx := server.Context()

	device, err := auth.AuthenticatedDevice(ctx)
	if err != nil {
		return s.handleError(err)
	}

	archiveEvents, err := s.eventInteractor.ReceiveArchiveEvent(ctx, device)
	if err != nil {
		return s.handleError(err)
	}

	for archiveEvent := range archiveEvents {
		archiveEventID := archiveEvent.ID
		s.logger.Info("event received", zap.String("archive_event_id", archiveEventID))
		archiveEventResponse := &eventPb.ArchiveEvent{
			ArchiveEventId: archiveEventID,
		}

		if err := server.Send(archiveEventResponse); err != nil {
			return s.handleError(err)
		}
	}

	return nil
}
