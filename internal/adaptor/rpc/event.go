package rpc

import (
	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/pb/event/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventRPCService struct {
	logger       *zap.Logger
	eventHandler entity.IArchiveEventHandler
	event.UnimplementedArchiveEventServiceServer
}

func (s EventRPCService) ReceiveEvent(_ *emptypb.Empty, server event.ArchiveEventService_ReceiveEventServer) error {
	ctx := server.Context()

	device, err := AuthenticatedDevice(ctx)
	if err != nil {
		return convertors.ConvertError(s.logger, apperrs.ForbiddenError)
	}

	archiveEvents, err := s.eventHandler.ReceiveArchiveEvent(ctx, device.ID)
	if err != nil {
		return convertors.ConvertError(s.logger, err)
	}

	for archiveEvent := range archiveEvents {
		archiveEventResponse := &event.ArchiveEvent{
			ArchiveEventId: archiveEvent.ID.Value().String(),
		}

		if err := server.Send(archiveEventResponse); err != nil {
			return convertors.ConvertError(s.logger, err)
		}
	}

	return nil
}
