use std::time::Duration;

use rdkafka::{
    ClientConfig,
    producer::{FutureProducer, FutureRecord},
    util::Timeout,
};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AppMessage {
    pub id: String,
    pub message: String,
}

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

    pub async fn send_message(&self, message: AppMessage) {
        let payload = serde_json::to_string(&message).expect("Failed to serialize message");

        match self
            .producer
            .send(
                FutureRecord::to(&self.topic)
                    .payload(&payload)
                    .key(&message.id),
                Timeout::After(Duration::from_secs(10)),
            )
            .await
        {
            Ok(_) => println!("Message sent"),
            Err((err, _)) => eprintln!("Filed to send: {err:?}"),
        }

        println!("Message sent")
    }
}
