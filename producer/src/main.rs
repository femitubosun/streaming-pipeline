use std::time::Duration;

use tokio::time::interval;

use crate::utils::get_env;

mod admin;
mod events;
mod generator;
mod producer;
mod utils;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("Initializing Redpanda Admin Client");
    let brokers = get_env("KAFKA_BROKERS", None);
    let topic = get_env("TOPIC", Some("raw-events"));

    let admin = admin::Admin::new(&brokers)?;

    let exists = admin.topic_exists(&topic).await?;

    if exists {
        println!("Topic {topic} already exists")
    } else if let Err(err) = admin.create_topic(&topic).await {
        eprintln!("Failed to create topic {topic}: {err:?}")
    }

    let producer = producer::AppProducer::new(&brokers, &topic);

    let mut ticker = interval(Duration::from_millis(10)); // 100 events/sec

    loop {
        ticker.tick().await;
        let msg = generator::generate_event();
        producer.send_message(msg).await;
    }
}
