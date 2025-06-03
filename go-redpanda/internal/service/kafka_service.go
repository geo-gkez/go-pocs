package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type IKafkaService interface {
	ProduceMessage(message string) error
}

type kafkaService struct {
	client *kgo.Client
}

func NewKafkaService(client *kgo.Client) IKafkaService {
	return &kafkaService{client: client}
}

func (s *kafkaService) ProduceMessage(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{Value: []byte(message)}
	s.client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}
	})
	wg.Wait()

	return nil
}
