package transactions

import (
	"context"
	"os"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Enrich(ctx context.Context, event RawTransactionEvent) (ProcessedTransactionEvent, error) {
	now := time.Now().UTC()

	score, factors := calculateRisk(event)

	status := ValidationStatusApproved

	if score > 80 {
		status = ValidationStatusRejected
	} else if score > 50 {
		status = ValidationStatusFlagged
	}

	processorID, err := os.Hostname()
	if err != nil {
		return ProcessedTransactionEvent{}, err
	}

	eventTime, err := time.Parse(time.RFC3339Nano, event.EventTime)
	if err != nil {
		return ProcessedTransactionEvent{}, err
	}

	return ProcessedTransactionEvent{
		EventID:               event.EventID,
		IdempotencyKey:        event.IdempotencyKey,
		EventTime:             event.EventTime,
		EventType:             event.EventType,
		CustomerID:            event.CustomerID,
		Amount:                event.Amount,
		Currency:              event.Currency,
		MerchantID:            event.MerchantID,
		Country:               event.Country,
		InstrumentFingerprint: event.InstrumentFingerprint,
		IsFraud:               status == ValidationStatusRejected,

		RiskScore:           score,
		RiskFactors:         factors,
		ValidationStatus:    status,
		ProcessedAt:         now.Format(time.RFC3339),
		ProcessorID:         processorID,
		ProcessingLatencyMs: now.Sub(eventTime).Milliseconds(),
	}, nil
}

func calculateRisk(event RawTransactionEvent) (float64, []string) {
	score := 0.0
	var factors []string

	if event.Amount > 90_000_00 {
		score += 60
		factors = append(factors, "high_amount")
	} else if event.Amount > 50_000_00 {
		score += 40
		factors = append(factors, "elevated_amount")
	} else if event.Amount > 10_000_00 {
		score += 20
		factors = append(factors, "moderate_amount")
	}

	if event.Country != "US" && event.Country != "GB" && event.Country != "CA" {
		score += 25
		factors = append(factors, "foreign_country")
	}

	if event.Currency != USD && event.Currency != GBP && event.Currency != CAD {
		score += 25
		factors = append(factors, "foreign_currency")
	}

	switch event.EventType {
	case TransactionFlagged:
		score += 50
		factors = append(factors, "transaction_flagged")
	case TransactionDeclined:
		score += 30
		factors = append(factors, "transaction_declined")
	}

	return score, factors
}
