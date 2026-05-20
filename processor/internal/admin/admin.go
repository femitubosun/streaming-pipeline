package admin

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Admin struct {
	client *kadm.Client
}

func NewAdmin(brokers []string) (*Admin, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)

	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return &Admin{
		client: kadm.NewClient(cl),
	}, nil
}

func (a *Admin) TopicExists(topic string) (bool, error) {
	topics, err := a.client.ListTopics(context.Background())
	if err != nil {
		return false, fmt.Errorf("List topics: %w", err)
	}

	_, exists := topics[topic]
	return exists, nil
}

func (a *Admin) CreateTopic(topic string) error {
	resp, err := a.client.CreateTopics(context.Background(), 1, 1, nil, topic)
	if err != nil {
		return fmt.Errorf("create topic: %w", err)
	}

	for _, t := range resp {
		if t.Err != nil {
			return fmt.Errorf("topic %s: %w", t.Topic, t.Err)
		}
	}

	return nil
}

func (a *Admin) Close() {
	a.client.Close()
}
