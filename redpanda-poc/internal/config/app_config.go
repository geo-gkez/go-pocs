package config

import (
	"fmt"
	"os"

	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/config/models"
	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/routes"
	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Settings struct {
	KafkaClient *kgo.Client
}

func SetupApp() {
	config, err := loadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	kafkaClient := setUpKafka(config, service.ProcessKafkaMessage)

	kafkaService := service.NewKafkaService(kafkaClient)

	router := gin.Default()

	router = routes.SetupRoutesAndRegister(router, kafkaService)

	router.Run(":8085")
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
