package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventClient struct {
	conn *amqp.Connection
}

func NewEventClient(conn *amqp.Connection) *EventClient {
	return &EventClient{
		conn: conn,
	}
}

func (h EventClient) GetChannel() (*amqp.Channel, error) {
	return h.conn.Channel()
}

func (h EventClient) PublishMessage(ctx context.Context, ch *amqp.Channel, exchangeName, exchangeKind, key string, body []byte) error {
	if err := ch.ExchangeDeclare(exchangeName, exchangeKind, true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.PublishWithContext(ctx, exchangeName, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}); err != nil {
		return err
	}

	return nil
}

func (h EventClient) ReceiveMessage(ctx context.Context, ch *amqp.Channel, exchangeName, exchangeKind, key string) (<-chan amqp.Delivery, error) {
	if err := ch.ExchangeDeclare(exchangeName, exchangeKind, true, false, false, false, nil); err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		key,
		true,
		true,
		true,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	if err := ch.QueueBind(queue.Name, key, exchangeName, false, nil); err != nil {
		return nil, err
	}

	messages, err := ch.ConsumeWithContext(ctx, queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
