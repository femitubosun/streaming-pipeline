use std::time::Duration;

use rdkafka::{
    ClientConfig,
    error::KafkaError,
    message::OwnedMessage,
    producer::{FutureProducer, FutureRecord},
    util::Timeout,
};

use crate::events::TransactionEvent;

pub struct AppProducer {
    producer: FutureProducer,
    topic: String,
}

impl AppProducer {
    pub fn new(brokers: &str, topic: &str) -> Self {
        let producer: FutureProducer = ClientConfig::new()
            .set("bootstrap.servers", brokers)
            .set("batch.size", "32768")
            .set("linger.ms", "10")
            .set("compression.type", "lz4")
            .set("acks", "1")
            .set("retries", "3")
            .create()
            .expect("Producer creation failed");

        AppProducer {
            producer,
            topic: topic.to_string(),
        }
    }

    pub async fn send_message(
        &self,
        message: TransactionEvent,
    ) -> Result<(), (KafkaError, OwnedMessage)> {
        let payload = serde_json::to_string(&message).expect("Failed to serialize message");

        self.producer
            .send(
                FutureRecord::to(&self.topic)
                    .payload(&payload)
                    .key(&message.customer_id),
                Timeout::After(Duration::from_secs(10)),
            )
            .await
            .map(|_| ())
    }
}
