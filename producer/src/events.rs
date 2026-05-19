use serde::{Deserialize, Serialize};

#[derive(Clone)]
pub enum Currency {
    USD,
    NGN,
    EUR,
}

impl From<Currency> for String {
    fn from(c: Currency) -> Self {
        match c {
            Currency::USD => "USD".to_string(),
            Currency::NGN => "NGN".to_string(),
            Currency::EUR => "EUR".to_string(),
        }
    }
}

#[derive(Clone)]
pub enum TransactionEventType {
    Initiated,
    Authorized,
    Completed,
    Declined,
    Flagged,
    Reversed,
    Expired,
}

impl From<TransactionEventType> for String {
    fn from(t: TransactionEventType) -> Self {
        match t {
            TransactionEventType::Initiated => "transaction.initiated".to_string(),
            TransactionEventType::Authorized => "transaction.authorized".to_string(),
            TransactionEventType::Completed => "transaction.completed".to_string(),
            TransactionEventType::Declined => "transaction.declined".to_string(),
            TransactionEventType::Flagged => "transaction.flagged".to_string(),
            TransactionEventType::Reversed => "transaction.reversed".to_string(),
            TransactionEventType::Expired => "transaction.expired".to_string(),
        }
    }
}

#[derive(Clone)]
pub struct Money {
    pub amount: u64,
    pub currency: Currency,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransactionEvent {
    pub event_id: String,
    pub idempotency_key: String,
    pub event_time: String,
    pub event_type: String,

    pub customer_id: String,
    pub amount: u64,
    pub currency: String,
    pub merchant_id: String,
    pub country: String,

    pub instrument_fingerprint: String,
}
