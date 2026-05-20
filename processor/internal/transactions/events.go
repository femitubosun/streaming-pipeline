package transactions

type TransactionEventType string

const (
	TransactionInitiated  TransactionEventType = "transaction.initiated"
	TransactionAuthorized TransactionEventType = "transaction.authorized"
	TransactionCompleted  TransactionEventType = "transaction.completed"
	TransactionDeclined   TransactionEventType = "transaction.declined"
	TransactionFlagged    TransactionEventType = "transaction.flagged"
	TransactionReversed   TransactionEventType = "transaction.reversed"
	TransactionExpired    TransactionEventType = "transaction.expired"
)

type Currency string

const (
	USD Currency = "USD"
	NGN Currency = "NGN"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

type RawTransactionEvent struct {
	EventID        string               `json:"event_id"`
	IdempotencyKey string               `json:"idempotency_key"`
	EventTime      string               `json:"event_time"`
	EventType      TransactionEventType `json:"event_type"`

	CustomerID            string   `json:"customer_id"`
	Amount                uint64   `json:"amount"`
	Currency              Currency `json:"currency"`
	MerchantID            string   `json:"merchant_id"`
	Country               string   `json:"country"`
	InstrumentFingerprint string   `json:"instrument_fingerprint"`
}

type ProcessedTransactionEvent struct {
	EventID        string               `json:"event_id"`
	IdempotencyKey string               `json:"idempotency_key"`
	EventTime      string               `json:"event_time"`
	EventType      TransactionEventType `json:"event_type"`

	CustomerID            string   `json:"customer_id"`
	Amount                uint64   `json:"amount"`
	Currency              Currency `json:"currency"`
	MerchantID            string   `json:"merchant_id"`
	Country               string   `json:"country"`
	InstrumentFingerprint string   `json:"instrument_fingerprint"`
	IsFraud               bool     `json:"is_fraud"`
}
