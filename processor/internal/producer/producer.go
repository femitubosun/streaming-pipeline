package producer

import (
	"context"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
	topic  string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.RequestRetries(3),
		kgo.RetryTimeout(30*time.Second),
		kgo.ProducerBatchMaxBytes(1_000_000),
		kgo.BrokerMaxWriteBytes(10_000_000),
	)

	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return &Producer{
		client: cl,
		topic:  topic,
	}, nil
}

func (p *Producer) Close() {
	p.client.Close()
}

func (p *Producer) SendMessage(ctx context.Context, key []byte, value []byte) {
	p.client.Produce(ctx, &kgo.Record{
		Topic: p.topic,
		Key:   key,
		Value: value,
	}, func(_ *kgo.Record, _ error) {})
}
