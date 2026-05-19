use std::time::Duration;

use rdkafka::{
    ClientConfig,
    admin::{AdminClient, AdminOptions, NewTopic, TopicReplication},
    client::DefaultClientContext,
    error::{KafkaError, KafkaResult},
    util::Timeout,
};

pub struct Admin {
    client: AdminClient<DefaultClientContext>,
}
impl Admin {
    pub fn new(brokers: &str) -> Result<Self, KafkaError> {
        let client: AdminClient<DefaultClientContext> = ClientConfig::new()
            .set("bootstrap.servers", brokers)
            .create()?;

        Ok(Admin { client })
    }

    pub async fn topic_exists(&self, topic: &str) -> KafkaResult<bool> {
        let metadata = self
            .client
            .inner()
            .fetch_metadata(None, Timeout::After(Duration::from_secs(5)))?;

        Ok(metadata.topics().iter().any(|t| t.name() == topic))
    }

    pub async fn create_topic(&self, topic: &str) -> KafkaResult<()> {
        let new_topic = NewTopic::new(topic, 1, TopicReplication::Fixed(1));
        let res = self
            .client
            .create_topics(
                &[new_topic],
                &AdminOptions::new()
                    .operation_timeout(Some(Timeout::After(Duration::from_secs(10)))),
            )
            .await?;

        for result in res {
            match result {
                Ok(_) => println!("Topic {topic} was created successfully"),
                Err((err, _)) => eprintln!("Failed to create topic {topic}: {err:?}"),
            }
        }

        Ok(())
    }
}
