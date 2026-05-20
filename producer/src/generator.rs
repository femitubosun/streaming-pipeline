use chrono::Utc;
use rand::RngExt;
use uuid::{NoContext, Timestamp, Uuid};

use crate::events::{Country, Currency, Money, TransactionEvent, TransactionEventType};

pub fn generate_event() -> TransactionEvent {
    let money = get_random_money();
    let now = Utc::now();

    TransactionEvent {
        event_id: Uuid::new_v7(Timestamp::now(NoContext)).to_string(),
        idempotency_key: Uuid::new_v4().to_string(),
        event_time: now.to_rfc3339(),
        event_type: get_random_event_type(),
        customer_id: Uuid::new_v4().to_string(),
        amount: money.amount,
        currency: money.currency,
        merchant_id: Uuid::new_v4().to_string(),
        country: get_random_country(),
        instrument_fingerprint: format!("fp_{}", Uuid::new_v4()),
    }
}

fn get_random_money() -> Money {
    let mut rng = rand::rng();

    let amount = rng.random_range(1_00..200_000_00);
    let currency = match rng.random_range(0..5) {
        0 => Currency::USD,
        1 => Currency::NGN,
        2 => Currency::EUR,
        3 => Currency::GBP,
        _ => Currency::CAD,
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

fn get_random_country() -> Country {
    let mut rng = rand::rng();

    match rng.random_range(0..6) {
        0 => Country::US,
        1 => Country::NG,
        2 => Country::GB,
        3 => Country::CA,
        4 => Country::FR,
        _ => Country::ES,
    }
}
