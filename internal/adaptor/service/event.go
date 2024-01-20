package service

import (
	"context"
	"encoding/json"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/infrastructure/rabbitmq"
	"github.com/CityBear3/satellite/internal/usecase/service"
	"go.uber.org/zap"
)

const exchangeName = "event"
const exchangeKind = "direct"

type EventService struct {
	logger *zap.Logger
	client *rabbitmq.EventClient
}

func NewEventService(logger *zap.Logger, client *rabbitmq.EventClient) *EventService {
	return &EventService{
		logger: logger,
		client: client,
	}
}

func (h EventService) PublishArchiveEvent(ctx context.Context, event entity.ArchiveEvent) error {
	ch, err := h.client.GetChannel()
	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	defer func() {
		if err := ch.Close(); err != nil {
			h.logger.Error(err.Error())
		}
	}()

	body, err := json.Marshal(service.ArchiveEventMessage{
		ID:       event.ID.Value().String(),
		ClientID: event.ClientID.Value().String(),
	})

	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return h.client.PublishMessage(ctx, ch, exchangeName, exchangeKind, event.DeviceID.Value().String(), body)
}

func (h EventService) ReceiveArchiveEvent(ctx context.Context, deviceID primitive.ID) (<-chan service.ArchiveEventMessage, error) {
	ch, err := h.client.GetChannel()
	if err != nil {
		h.logger.Error("failed to get rabbitmq channel", zap.Error(err))
		return nil, err
	}

	messages, err := h.client.ReceiveMessage(ctx, ch, exchangeName, exchangeKind, deviceID.Value().String())
	if err != nil {
		h.logger.Error("failed to receive event message", zap.Error(err))
		return nil, err
	}

	subscribe := make(chan service.ArchiveEventMessage, 10000)
	go func() {
		defer func() {
			if err := ch.Close(); err != nil {
				h.logger.Error(err.Error())
			}
		}()

		var message service.ArchiveEventMessage
		for m := range messages {
			select {
			case <-ctx.Done():
				close(subscribe)
				return
			default:
				if err := json.Unmarshal(m.Body, &message); err != nil {
					h.logger.Error("failed to unmarshal event message", zap.Error(err))
					continue
				}

				subscribe <- message
			}
		}
	}()

	return subscribe, nil
}
