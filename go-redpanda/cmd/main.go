package main

import (
	"fmt"

	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/config"
	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/service"
)

func main() {
	settings := config.SetupApp()

	kafkaService := service.NewKafkaService(settings.KafkaClient)

	message := "Hello, Redpanda!"
	if err := kafkaService.ProduceMessage(message); err != nil {
		fmt.Printf("Failed to produce message: %v\n", err)
	} else {
		fmt.Println("Message produced successfully!")
	}
	fmt.Println("Kafka client setup complete. Ready to produce messages.")
}
