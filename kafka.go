package tools

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct {
	BrokerAddress []string
}
type KafkaMgrType struct {
	Writer *kafka.Writer
	Config *KafkaConfig
}

var KafkaMgr KafkaMgrType

func (k *KafkaMgrType) Init() {
	KafkaMgr.Writer = &kafka.Writer{
		Addr:     kafka.TCP(k.Config.BrokerAddress...),
		Balancer: &kafka.LeastBytes{},
	}
}

func (k *KafkaMgrType) Write(c context.Context, input interface{}) error {
	b, err := jsoniter.Marshal(input)
	if err != nil {
		return err
	}
	err = k.Writer.WriteMessages(c,
		kafka.Message{Topic: k.Config.Topic, Value: b},
	)
	return err
}
