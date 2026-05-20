package transactions

import "context"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Enrich(ctx context.Context, event RawTransactionEvent) (ProcessedTransactionEvent, error) {
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
		IsFraud:               false,
	}, nil
}
