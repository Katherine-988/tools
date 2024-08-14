package tools

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

type KafkaMgrType struct {
	Writer        *kafka.Writer
	BrokerAddress []string
	Topic         string
}

var KafkaMgr KafkaMgrType

func (k *KafkaMgrType) Init() {
	KafkaMgr.Writer = &kafka.Writer{
		Addr:     kafka.TCP(k.BrokerAddress...),
		Topic:    k.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (k *KafkaMgrType) Write(c context.Context, input interface{}) error {
	b, err := jsoniter.Marshal(input)
	if err != nil {
		return err
	}
	err = k.Writer.WriteMessages(c,
		kafka.Message{Topic: k.Topic, Value: b},
	)
	return err
}
