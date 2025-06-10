package service

import "fmt"

func ProcessKafkaMessage(key, value []byte) error {

	fmt.Printf("Processing Kafka message with key: %s, value: %s\n", key, value)
	//TODO: Add processing logic

	return nil
}
