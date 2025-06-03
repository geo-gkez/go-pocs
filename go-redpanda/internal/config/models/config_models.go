package config_models

// AppConfig represents the root configuration structure
type AppConfiguration struct {
	Kafka KafkaProperties
}

// KafkaProperties holds all Kafka-related configuration
type KafkaProperties struct {
	Connection KafkaConnection
	Topics     KafkaTopics
}

// KafkaConnection holds Kafka connection details
type KafkaConnection struct {
	Brokers []string
}

// KafkaTopics holds Kafka topic configurations
type KafkaTopics struct {
	DefaultProducer      string `mapstructure:"default-producer"`
	DefaultConsumer      string `mapstructure:"default-consumer" `
	DefaultConsumerGroup string `mapstructure:"default-consumer-group"`
}
