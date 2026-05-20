package consumer

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	client *kgo.Client
}

func NewConsumer(brokers []string, topics []string) (*Consumer, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup("stream-processor"),
		kgo.ConsumeTopics(topics...),
		kgo.AutoCommitMarks(),
	)

	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return &Consumer{
		client: cl,
	}, nil
}

func (c *Consumer) Close() {
	c.client.Close()
}

func (c *Consumer) Poll(ctx context.Context) ([]*kgo.Record, error) {
	fetches := c.client.PollFetches(ctx)

	if errs := fetches.Errors(); len(errs) > 0 {
		return nil, fmt.Errorf("fetch errors: %v", errs)
	}

	var records []*kgo.Record
	iter := fetches.RecordIter()
	for !iter.Done() {
		records = append(records, iter.Next())
	}

	return records, nil
}

func (c *Consumer) MarkCommitted(record *kgo.Record) {
	c.client.MarkCommitRecords(record)
}
