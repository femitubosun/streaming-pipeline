package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/femitubosun/streaming-pipeline/processor/internal/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Could not load env")
		os.Exit(1)
	}
	fmt.Println("KAFKA_BROKERS: ", cfg.KafkaBrokers)

	topic := "raw-events"

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.KafkaBrokers),
		kgo.ConsumerGroup("stream-processor"),
		kgo.ConsumeTopics(topic),
	)

	if err != nil {
		panic(err)
	}

	defer cl.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	for {
		fetches := cl.PollFetches(ctx)
		if ctx.Err() != nil {
			break
		}

		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			fmt.Printf("fetch errors: %v\n", errs)
			continue
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}
	}

}
