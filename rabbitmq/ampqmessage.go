package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"moneyx.golang.framework/logger"
)

type rabbitMessage struct {
	msg    *amqp.Delivery
	logger *logger.GoFrLogger
}

func newRabbitMessage(msg *amqp.Delivery, logger *logger.GoFrLogger) *rabbitMessage {
	return &rabbitMessage{
		msg:    msg,
		logger: logger,
	}
}

func (r *rabbitMessage) Commit() {
	err := r.msg.Ack(false)
	if err != nil {
		r.logger.Errorf("failed to ack message from rabbit topic %s: %v", string(r.msg.Body), err)
	}
}
