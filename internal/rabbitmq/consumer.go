package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/kxrxh/queue-master/pgk/utils"
)

type Consumer struct {
	channl *amqp.Channel
	queue  amqp.Queue
}

func NewConsumer(ch *amqp.Channel) *Consumer {
	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")

	return &Consumer{channl: ch, queue: q}
}

// Consume returns a channel of messages from the consumer's queue.
//
// The returned channel will receive messages as they are sent to the queue. The
// messages are automatically acknowledged, so the consumer does not need to
// explicitly ack them.
//
// If an error occurs while consuming, the channel will be closed.
func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := c.channl.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}
