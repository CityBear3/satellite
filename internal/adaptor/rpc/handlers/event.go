package handlers

import (
	"context"
	"encoding/json"

	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type EventHandler struct {
	logger *zap.Logger
	conn   *amqp091.Connection
}

func (h EventHandler) PublishArchiveEvent(ctx context.Context, event entity.ArchiveEvent) error {
	//TODO implement me
	panic("implement me")
}

func (h EventHandler) ReceiveArchiveEvent(ctx context.Context, deviceID primitive.ID) (<-chan entity.ArchiveEvent, error) {
	ch, err := h.conn.Channel()
	if err != nil {
		return nil, err
	}

	if err = ch.ExchangeDeclare(
		"event",
		"fanout",
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

	if err = ch.QueueBind(queueName, "", "event", false, nil); err != nil {
		return nil, err
	}

	messages, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	subscribe := make(chan entity.ArchiveEvent)
	go func() {
		var message entity.ArchiveEvent
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
