use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Copy, Serialize, Deserialize)]
pub enum Country {
    #[serde(rename = "US")]
    US,

    #[serde(rename = "NG")]
    NG,

    #[serde(rename = "GB")]
    GB,

    #[serde(rename = "CA")]
    CA,

    #[serde(rename = "FR")]
    FR,

    #[serde(rename = "ES")]
    ES,
}

#[derive(Clone, Debug, Copy, Serialize, Deserialize)]
pub enum Currency {
    #[serde(rename = "USD")]
    USD,

    #[serde(rename = "NGN")]
    NGN,

    #[serde(rename = "EUR")]
    EUR,

    #[serde(rename = "GBP")]
    GBP,

    #[serde(rename = "CAD")]
    CAD,
}

#[derive(Clone, Copy, Debug, Serialize, Deserialize)]
pub enum TransactionEventType {
    #[serde(rename = "transaction.initiated")]
    Initiated,

    #[serde(rename = "transaction.authorized")]
    Authorized,

    #[serde(rename = "transaction.completed")]
    Completed,

    #[serde(rename = "transaction.declined")]
    Declined,

    #[serde(rename = "transaction.flagged")]
    Flagged,

    #[serde(rename = "transaction.reversed")]
    Reversed,

    #[serde(rename = "transaction.expired")]
    Expired,
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
    pub event_type: TransactionEventType,

    pub customer_id: String,
    pub amount: u64,
    pub currency: Currency,
    pub merchant_id: String,
    pub country: Country,

    pub instrument_fingerprint: String,
}
