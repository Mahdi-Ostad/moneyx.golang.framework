package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/trace"
	"gofr.dev/pkg/gofr/datasource"
	"gofr.dev/pkg/gofr/datasource/pubsub"
	"moneyx.golang.framework/logger"
)

const (
	defaultRetryTimeout   = 10 * time.Second
	maxRetryTimeout       = 1 * time.Minute
	Err_Subscriber_Closed = "subscriber channel closed"
)

type Metrics interface {
	IncrementCounter(ctx context.Context, name string, labels ...string)
	RecordHistogram(ctx context.Context, name string, value float64, labels ...string)
}

type RabbitMQ struct {
	connectionString string
	exchangeName     string
	conn             *amqp.Connection
	ch               *amqp.Channel
	logger           logger.GoFrLogger
	metrics          Metrics
	tracer           trace.Tracer
	bindedQueues     map[string]bool
	declaredQueues   map[string]bool
	subscribers      map[string]<-chan amqp.Delivery
}

func NewRabbitMQPubSub(connectionString, exchangeName string) *RabbitMQ {
	return &RabbitMQ{
		connectionString: connectionString,
		exchangeName:     exchangeName,
		bindedQueues:     make(map[string]bool),
		declaredQueues:   make(map[string]bool),
		subscribers:      make(map[string]<-chan amqp.Delivery),
	}
}

func (r *RabbitMQ) Publish(ctx context.Context, topic string, message []byte) (err error) {
	queueName := queueNameMaker(topic)
	err = r.bindQueue(topic)
	if err != nil {
		return
	}
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         message,
		Headers: amqp.Table{
			"message-type": topic,
		},
	}
	err = r.ch.Publish(r.exchangeName, queueName, true, false, msg)
	return err
}

func (r *RabbitMQ) Subscribe(ctx context.Context, topic string) (message *pubsub.Message, err error) {
	err = r.declareQueue(topic)
	if err != nil {
		return
	}
	err = r.bindQueue(topic)
	if err != nil {
		return
	}
	if _, ok := r.subscribers[topic]; !ok {
		queueName := queueNameMaker(topic)
		msgs, err := r.ch.Consume(queueName, "", false, false, false, false, nil)
		if err != nil {
			return nil, err
		}
		r.subscribers[topic] = msgs
	}
	subscriber := r.subscribers[topic]
	select {
	case msg, ok := <-subscriber:
		if !ok {
			return nil, errors.New(Err_Subscriber_Closed)
		}
		message = pubsub.NewMessage(ctx)
		message.Value = msg.Body
		message.Topic = topic
		message.Committer = newRabbitMessage(&msg, &r.logger)
		return
	default:
		return nil, nil
	}
}

func (r *RabbitMQ) Health() datasource.Health {
	res := datasource.Health{
		Status: "DOWN",
		Details: map[string]any{
			"backend": "RABBITMQ",
			"host":    r.connectionString,
		},
	}

	if r.conn == nil {
		r.logger.Errorf("%v", "connection not initialized")

		return res
	}
	if !r.conn.ConnectionState().HandshakeComplete {
		r.logger.Errorf("%v", "handshake not complete")

		return res
	}

	if r.ch == nil {
		r.logger.Errorf("%v", "channel not initialized")

		return res
	}
	if r.ch.IsClosed() {
		r.logger.Errorf("%v", "channel is closed")

		return res
	}

	res.Status = "UP"

	return res
}

func (r *RabbitMQ) CreateTopic(context context.Context, name string) error {
	return r.ch.ExchangeDeclare(name, "topic", false, false, false, false, nil)
}

func (r *RabbitMQ) DeleteTopic(context context.Context, name string) error {
	return errors.New("deleting exchange not allowed on rabbitmq")
}

func (r *RabbitMQ) Close() error {
	err := r.ch.Close()
	if err != nil {
		return err
	}
	err = r.conn.Close()
	return err
}

func (r *RabbitMQ) UseLogger(customLogger any) {
	if l, ok := customLogger.(logger.GoFrLogger); ok {
		r.logger = l
	}
}

func (r *RabbitMQ) UseMetrics(metrics any) {
	if l, ok := metrics.(Metrics); ok {
		r.metrics = l
	}
}

func (c *RabbitMQ) UseTracer(tracer any) {
	if t, ok := tracer.(trace.Tracer); ok {
		c.tracer = t
	}
}

func (r *RabbitMQ) Connect() {
	conn, err := amqp.Dial(r.connectionString)
	if err != nil {
		r.logger.Errorf("Could not connect rabbitmq. Starting retry GoRoutine: %v", err)
		go r.retryConnect()
	} else {
		r.conn = conn
		r.makeChannel()
	}
}

func (r *RabbitMQ) makeChannel() {
	ch, err := r.conn.Channel()
	if err != nil {
		r.logger.Errorf("Could not make rabbitmq channel. Starting retry GoRoutine: %v", err)
		go r.retryChannel()
	} else {
		r.ch = ch
		r.CreateTopic(context.Background(), r.exchangeName)
	}
}

func (r *RabbitMQ) retryConnect() {
	for {
		conn, err := amqp.Dial(r.connectionString)
		if err == nil {
			r.conn = conn
			r.makeChannel()
			break
		}
		r.logger.Errorf("Could not connect rabbitmq. Retrying in 10 seconds: %v", err)
		time.Sleep(defaultRetryTimeout)
	}
}

func (r *RabbitMQ) retryChannel() {
	for {
		ch, err := r.conn.Channel()
		if err == nil {
			r.ch = ch
			r.CreateTopic(context.Background(), r.exchangeName)
			break
		}
		r.logger.Errorf("Could not make rabbitmq channel. Retrying in 10 seconds: %v", err)
		time.Sleep(defaultRetryTimeout)
	}
}

func (r *RabbitMQ) bindQueue(topic string) error {
	queueName := queueNameMaker(topic)
	if _, ok := r.bindedQueues[topic]; !ok {
		err := r.ch.QueueBind(queueName, queueName, r.exchangeName, false, nil)
		if err != nil {
			return err
		}
		r.bindedQueues[topic] = true
	}
	return nil
}

func (r *RabbitMQ) declareQueue(topic string) error {
	queueName := queueNameMaker(topic)
	if _, ok := r.declaredQueues[topic]; !ok {
		_, err := r.ch.QueueDeclare(queueName, false, false, false, false, nil)
		if err != nil {
			return err
		}
		r.declaredQueues[topic] = true
	}
	return nil
}

func queueNameMaker(cmdName string) string {
	return fmt.Sprintf("%s.workers", strings.ToLower(cmdName))
}
