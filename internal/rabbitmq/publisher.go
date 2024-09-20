package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"github.com/kxrxh/queue-master/pgk/utils"
)

type Publisher struct {
	channl *amqp.Channel
	queue  amqp.Queue
}

func NewPublisher(ch *amqp.Channel) *Publisher {
	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")

	return &Publisher{channl: ch, queue: q}
}

func (p *Publisher) Publish(message string) error {
	err := p.channl.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		logger.Sugar().Errorf("Failed to publish a message: %s", err)
		return err
	}

	logger.Sugar().Infof("Published a message: %s", message)
	return nil
}
