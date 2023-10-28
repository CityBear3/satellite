package rpc

import (
	"context"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/event/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventRPCService struct {
	logger          *zap.Logger
	eventInteractor usecase.IEventUseCase
	event.UnimplementedArchiveEventServiceServer
}

func NewEventRPCService(logger *zap.Logger, eventInteractor usecase.IEventUseCase) *EventRPCService {
	return &EventRPCService{
		logger:          logger,
		eventInteractor: eventInteractor,
	}
}

func (s EventRPCService) PublishEvent(ctx context.Context, req *emptypb.Empty) (*event.PublishEventResponse, error) {
	client, err := AuthenticatedClient(ctx)
	if err != nil {
		return nil, convertors.ConvertError(s.logger, err)
	}

	archiveEventID, err := s.eventInteractor.PublishArchiveEvent(ctx, client)
	if err != nil {
		return nil, convertors.ConvertError(s.logger, err)
	}

	return &event.PublishEventResponse{
		ArchiveEventId: archiveEventID.Value().String(),
	}, nil
}

func (s EventRPCService) ReceiveEvent(req *emptypb.Empty, server event.ArchiveEventService_ReceiveEventServer) error {
	ctx := server.Context()

	device, err := AuthenticatedDevice(ctx)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	archiveEvents, err := s.eventInteractor.ReceiveArchiveEvent(ctx, device)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	for archiveEvent := range archiveEvents {
		archiveEventID := archiveEvent.ID
		s.logger.Info("event received", zap.String("archive_event_id", archiveEventID))
		archiveEventResponse := &event.ArchiveEvent{
			ArchiveEventId: archiveEventID,
		}

		if err := server.Send(archiveEventResponse); err != nil {
			return convertors.ConvertError(s.logger, err)
		}
	}

	return nil
}
