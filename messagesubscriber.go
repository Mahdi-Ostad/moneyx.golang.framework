package moneyx

import "gofr.dev/pkg/gofr"

type IntegratedMessageSubscriber struct {
	topic   string
	handler gofr.SubscribeFunc
}

func NewIntegratedMessageSubscriber(topic string, handler gofr.SubscribeFunc) *IntegratedMessageSubscriber {
	return &IntegratedMessageSubscriber{
		topic:   topic,
		handler: handler,
	}
}
