package service

import (
	"context"
	"fmt"
	"time"

	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/model"
	"github.com/twmb/franz-go/pkg/kgo"
)

type IKafkaService interface {
	ProduceMessage(message model.ProduceMessageRequest) error
}

type kafkaService struct {
	client *kgo.Client
}

func NewKafkaService(client *kgo.Client) IKafkaService {
	return &kafkaService{client: client}
}

func (s *kafkaService) ProduceMessage(message model.ProduceMessageRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// fire and forget approach
	record := &kgo.Record{Key: []byte(message.Key), Value: []byte(message.Message)}
	s.client.Produce(ctx, record, func(r *kgo.Record, err error) {
		defer cancel()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		} else {
			fmt.Printf("Successfully produced record to %s [%d] at offset %d\n",
				r.Topic, r.Partition, r.Offset)
		}
	})

	return nil
}
