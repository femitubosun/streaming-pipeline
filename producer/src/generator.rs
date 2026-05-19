use chrono::Utc;
use rand::RngExt;
use uuid::{NoContext, Timestamp, Uuid};

use crate::events::{Currency, Money, TransactionEvent, TransactionEventType};

pub fn generate_event() -> TransactionEvent {
    let money = get_random_money();
    let now = Utc::now();

    TransactionEvent {
        event_id: Uuid::new_v7(Timestamp::now(NoContext)).to_string(),
        idempotency_key: Uuid::new_v4().to_string(),
        event_time: now.to_rfc3339(),
        event_type: get_random_event_type().into(),
        customer_id: Uuid::new_v4().to_string(),
        amount: money.amount,
        currency: money.currency.into(),
        merchant_id: Uuid::new_v4().to_string(),
        country: String::from("US"),
        instrument_fingerprint: format!("fp_{}", Uuid::new_v4()),
    }
}

fn get_random_money() -> Money {
    let mut rng = rand::rng();

    let amount = rng.random_range(100..100000);
    let currency = match rng.random_range(0..3) {
        0 => Currency::USD,
        1 => Currency::NGN,
        _ => Currency::EUR,
    };

    Money { amount, currency }
}

fn get_random_event_type() -> TransactionEventType {
    let mut rng = rand::rng();

    match rng.random_range(0..7) {
        0 => TransactionEventType::Initiated,
        1 => TransactionEventType::Authorized,
        2 => TransactionEventType::Completed,
        3 => TransactionEventType::Declined,
        4 => TransactionEventType::Flagged,
        5 => TransactionEventType::Reversed,
        _ => TransactionEventType::Expired,
    }
}
