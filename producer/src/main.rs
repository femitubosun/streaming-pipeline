use crate::utils::get_env;
use futures::future::join_all;

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

    const BATCH_SIZE: usize = 5000;

    loop {
        let futures: Vec<_> = (0..BATCH_SIZE)
            .map(|_| producer.send_message(generator::generate_event()))
            .collect();

        let results = join_all(futures).await;
        let sent = results.iter().filter(|r| r.is_ok()).count();

        if sent < BATCH_SIZE {
            eprintln!("Dropped {}/{BATCH_SIZE}", BATCH_SIZE - sent)
        }
    }
}
