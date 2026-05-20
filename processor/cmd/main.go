package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/femitubosun/streaming-pipeline/processor/internal/admin"
	"github.com/femitubosun/streaming-pipeline/processor/internal/config"
	"github.com/femitubosun/streaming-pipeline/processor/internal/consumer"
	"github.com/femitubosun/streaming-pipeline/processor/internal/producer"
	"github.com/femitubosun/streaming-pipeline/processor/internal/transactions"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Could not load env")
		os.Exit(1)
	}
	fmt.Println("KAFKA_BROKERS: ", cfg.KafkaBrokers)

	brokers := []string{cfg.KafkaBrokers}

	adm, err := admin.NewAdmin(brokers)

	if err != nil {
		fmt.Println("Could not create admin:", err)
		os.Exit(1)
	}
	defer adm.Close()

	topic := "raw-events"
	processedTopic := "processed-events"

	inputExists, err := adm.TopicExists(topic)
	if err != nil {
		fmt.Println("Could not check input topic exists", err)
		os.Exit(1)
	}

	if !inputExists {
		fmt.Printf("Input Topic %s does not exist\n", topic)
		os.Exit(1)
	}

	processedExists, err := adm.TopicExists(processedTopic)
	if err != nil {
		fmt.Printf("Could not check output topic: %v\n", err)
		os.Exit(1)
	}

	if !processedExists {
		fmt.Printf("Creating topic %s\n", processedTopic)
		if err := adm.CreateTopic(processedTopic); err != nil {
			fmt.Printf("Could not create topic: %v\n", err)
			os.Exit(1)
		}
	}

	prd, err := producer.NewProducer(brokers, processedTopic)

	if err != nil {
		fmt.Println("Could not create producer", err)
		os.Exit(1)
	}

	defer prd.Close()

	fmt.Printf("Topic %s exists\n", topic)

	topics := []string{"raw-events"}

	cs, err := consumer.NewConsumer(brokers, topics)

	if err != nil {
		panic(err)
	}

	defer cs.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	s := transactions.NewService()

	for {
		records, err := cs.Poll(ctx)

		if ctx.Err() != nil {
			break
		}

		if err != nil {
			fmt.Printf("poll error: %v\n", err)
			continue
		}

		for _, record := range records {
			var event transactions.RawTransactionEvent
			if err := json.Unmarshal(record.Value, &event); err != nil {
				slog.Error("unmarshal failed", "error", err)
				continue
			}

			processed, err := s.Enrich(context.Background(), event)
			if err != nil {
				slog.Error("enrich failed", "error", err)
				continue
			}

			payload, err := json.Marshal(processed)
			if err != nil {
				slog.Error("marshal failed", "error", err)
				continue
			}

			if err := prd.SendMessage(ctx, []byte(event.EventID), payload); err != nil {
				slog.Error("produce failed", "error", err)
				continue
			}

			slog.Info("processed",
				"event_id", processed.EventID,
				"risk_score", processed.RiskScore,
				"status", processed.ValidationStatus,
			)
			cs.MarkCommitted(record)
		}

	}

}
