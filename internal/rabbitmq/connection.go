package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/kxrxh/queue-master/pgk/utils"
)

func NewConnection(username, password string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	return conn, ch
}
