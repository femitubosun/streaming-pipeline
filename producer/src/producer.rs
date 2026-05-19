use std::time::Duration;

use rdkafka::{
    ClientConfig,
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
            .create()
            .expect("Producer creation failed");

        AppProducer {
            producer,
            topic: topic.to_string(),
        }
    }

    pub async fn send_message(&self, message: TransactionEvent) {
        let payload = serde_json::to_string(&message).expect("Failed to serialize message");

        match self
            .producer
            .send(
                FutureRecord::to(&self.topic)
                    .payload(&payload)
                    .key(&message.customer_id),
                Timeout::After(Duration::from_secs(10)),
            )
            .await
        {
            Ok(_) => println!("Message sent"),
            Err((err, _)) => eprintln!("Failed to send: {err:?}"),
        }
    }
}
