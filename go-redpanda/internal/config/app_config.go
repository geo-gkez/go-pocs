package config

import (
	"fmt"
	"os"

	config_models "github.com/geo-gkez/go-pocs/redpanda-poc/internal/config/models"
	"github.com/spf13/viper"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Settings struct {
	KafkaClient *kgo.Client
}

func SetupApp() *Settings {
	config, err := loadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	kafkaClient := setUpKafka(config)

	return &Settings{
		KafkaClient: kafkaClient,
	}
}

func loadConfig() (*config_models.AppConfiguration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Get the project root directory
	projectRoot, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	// Try multiple possible config locations
	viper.AddConfigPath(fmt.Sprintf("%s/configs", projectRoot)) // From project root
	viper.AddConfigPath("configs")                              // Direct subfolder
	viper.AddConfigPath("../../configs")                        // Two levels up

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config config_models.AppConfiguration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func setUpKafka(appConfig *config_models.AppConfiguration) *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(appConfig.Kafka.Connection.Brokers...),
		kgo.DefaultProduceTopic(appConfig.Kafka.Topics.DefaultProducer),
		kgo.ConsumerGroup(appConfig.Kafka.Topics.DefaultConsumerGroup),
		kgo.ConsumeTopics(appConfig.Kafka.Topics.DefaultConsumer),
	)

	if err != nil {
		panic(fmt.Sprintf("failed to create Kafka client: %v", err))
	}

	return client
}
