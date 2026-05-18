use std::time::Duration;

use tokio::time::interval;

mod admin;
mod producer;

#[tokio::main]
async fn main() {
    println!("Initializing Redpanda Admin Client");
    let brokers = std::env::var("KAFKA_BROKERS").expect("KAFKA_BROKERS env variable must be set");
    let topic = std::env::var("TOPIC").unwrap_or_else(|_| "raw-events".to_string());

    let admin = admin::Admin::new(&brokers);

    if let Ok(exists) = admin.topic_exists(&topic).await {
        if exists {
            println!("Topic {topic} already exists")
        } else if let Err(err) = admin.create_topic(&topic).await {
            eprintln!("Failed to create topic {topic}: {err:?}")
        }
    } else {
        eprintln!("Could not check if topic {topic} exists")
    }

    let producer = producer::AppProducer::new(&brokers, &topic);

    let mut ticker = interval(Duration::from_millis(10)); // 100 events/sec

    loop {
        ticker.tick().await;
        let msg = producer::AppMessage {
            id: uuid::Uuid::new_v4().to_string(),
            message: "synthetic_tx".to_string(),
        };
        producer.send_message(msg).await;
    }
}
