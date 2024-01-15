package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const exchangeName = "event"
const exchangeKind = "direct"

type EventHandler struct {
	logger *zap.Logger
	conn   *amqp.Connection
}

func NewEventHandler(logger *zap.Logger, conn *amqp.Connection) *EventHandler {
	return &EventHandler{
		logger: logger,
		conn:   conn,
	}
}

func (h EventHandler) PublishArchiveEvent(ctx context.Context, event entity.ArchiveEvent) error {
	ch, err := h.conn.Channel()
	if err != nil {
		return err
	}

	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			h.logger.Error(err.Error())
		}
	}(ch)

	if err = ch.ExchangeDeclare(exchangeName, exchangeKind, true, false, false, false, nil); err != nil {
		return err
	}

	body, err := json.Marshal(usecase.ArchiveEventMessage{
		ID:       event.ID.Value().String(),
		ClientID: event.ClientID.Value().String(),
	})

	if err != nil {
		log.Fatalln(err)
	}

	key := event.DeviceID.Value().String()
	if err = ch.PublishWithContext(ctx, exchangeName, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}); err != nil {
		return err
	}
	return nil
}

func (h EventHandler) ReceiveArchiveEvent(ctx context.Context, deviceID primitive.ID) (<-chan usecase.ArchiveEventMessage, error) {
	ch, err := h.conn.Channel()
	if err != nil {
		return nil, err
	}

	if err = ch.ExchangeDeclare(
		exchangeName,
		exchangeKind,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	queueName := deviceID.Value().String()
	queue, err := ch.QueueDeclare(
		queueName,
		true,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	key := queueName
	if err = ch.QueueBind(queueName, key, exchangeName, false, nil); err != nil {
		return nil, err
	}

	messages, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	subscribe := make(chan usecase.ArchiveEventMessage, 10000)
	go func() {
		var message usecase.ArchiveEventMessage
		for m := range messages {
			select {
			case <-ctx.Done():
				err := ch.Close()
				if err != nil {
					h.logger.Error(err.Error())
				}
				close(subscribe)
			default:
				if err := json.Unmarshal(m.Body, &message); err != nil {
					h.logger.Error(err.Error())
				}

				subscribe <- message
			}
		}
	}()

	return subscribe, nil
}
