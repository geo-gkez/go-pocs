package config

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/config/models"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

type pconsumer struct {
	quit chan struct{}
	recs chan []*kgo.Record
}

func (pc *pconsumer) consume(topic string, partition int32, handler MessageHandler) {
	fmt.Printf("starting, t %s p %d\n", topic, partition)
	// Log when the function exits (stops consuming from this partition)
	defer fmt.Printf("killing, t %s p %d\n", topic, partition)

	// Main processing loop
	for {
		select {
		// Channel to signal this consumer to quit
		case <-pc.quit:
			fmt.Printf("quitting, t %s p %d\n", topic, partition)
			return

		// Channel to receive batches of records for this partition
		case recs := <-pc.recs:

			// Process each record in the batch
			for _, rec := range recs {

				// handle the record
				if handler != nil {
					if err := handler(rec.Key, rec.Value); err != nil {
						fmt.Printf("Error handling message: %v\n", err)
					}
				}

			}
		}
	}
}

type MessageHandler func(key, value []byte) error

type splitConsume struct {
	mu        sync.Mutex // gaurds assigning / losing vs. polling
	consumers map[string]map[int32]pconsumer
	handler   MessageHandler
}

func (s *splitConsume) assigned(_ context.Context, cl *kgo.Client, assigned map[string][]int32) {
	// Lock the mutex to prevent concurrent access to the consumers map
	s.mu.Lock()
	defer s.mu.Unlock()

	// Iterate through each topic and its assigned partitions
	for topic, partitions := range assigned {
		// If this is the first partition for this topic, initialize the map
		if s.consumers[topic] == nil {
			s.consumers[topic] = make(map[int32]pconsumer)
		}

		// For each partition assigned to this consumer...
		for _, partition := range partitions {
			// Create a new partition consumer with communication channels
			pc := pconsumer{
				quit: make(chan struct{}),          // Channel to signal shutdown
				recs: make(chan []*kgo.Record, 10), // Buffered channel for records
			}

			// Store the partition consumer in the map for later access
			s.consumers[topic][partition] = pc

			// Launch a dedicated goroutine to process this partition
			go pc.consume(topic, partition, s.handler)
		}
	}
}

func (s *splitConsume) lost(_ context.Context, cl *kgo.Client, lost map[string][]int32) {
	// Lock the mutex to prevent concurrent access to the consumers map
	s.mu.Lock()
	defer s.mu.Unlock()

	// Iterate through each topic and its lost partitions
	for topic, partitions := range lost {
		// Get the map of partition consumers for this topic
		ptopics := s.consumers[topic]

		// For each partition that was lost...
		for _, partition := range partitions {
			// Get the partition consumer object
			pc := ptopics[partition]

			// Remove this partition from the map
			delete(ptopics, partition)

			// If this was the last partition for this topic, clean up the topic entry
			if len(ptopics) == 0 {
				delete(s.consumers, topic)
			}

			// Signal the partition consumer goroutine to stop
			close(pc.quit)
		}
	}
}

func (s *splitConsume) poll(cl *kgo.Client) {
	// Infinite loop to continuously poll for messages
	for {
		// Poll for new messages from Kafka
		fetches := cl.PollFetches(context.Background())

		// Check if the client has been closed
		if fetches.IsClientClosed() {
			fmt.Println("Client is closed, stopping consumption")
			return
		}

		// Handle any errors in the fetched data
		fetches.EachError(func(_ string, _ int32, err error) {
			//TODO: Handle errors appropriately, e.g., log them
			fmt.Println("Error in fetches:", err)
		})

		// Process each topic in the fetched data
		fetches.EachTopic(func(t kgo.FetchTopic) {
			// Safely get the consumers for this topic
			s.mu.Lock()
			tconsumers := s.consumers[t.Topic]
			s.mu.Unlock()

			// If there are no consumers for this topic, skip it
			if tconsumers == nil {
				return
			}

			// Process each partition in this topic
			t.EachPartition(func(p kgo.FetchPartition) {
				// Get the consumer for this partition
				pc, ok := tconsumers[p.Partition]
				if !ok {
					return
				}

				// Try to send records to the partition consumer or check if it's quitting
				select {
				case pc.recs <- p.Records: // Send records to the partition consumer
				case <-pc.quit: // Check if the consumer is shutting down
				}
			})
		})
	}
}

func setUpKafka(appConfig *config_models.AppConfiguration, handler MessageHandler) *kgo.Client {
	s := &splitConsume{
		consumers: make(map[string]map[int32]pconsumer),
		handler:   handler,
	}

	topics := appConfig.Kafka.Topics

	client, err := kgo.NewClient(
		kgo.SeedBrokers(appConfig.Kafka.Connection.Brokers...),
		kgo.DefaultProduceTopic(topics.DefaultProducer),
		kgo.ConsumerGroup(topics.DefaultConsumerGroup),
		kgo.ConsumeTopics(topics.DefaultConsumer),
		kgo.OnPartitionsAssigned(s.assigned),
		kgo.OnPartitionsRevoked(s.lost),
		kgo.OnPartitionsLost(s.lost),
	)

	if err != nil {
		panic(fmt.Sprintf("failed to create Kafka client: %v", err))
	}

	createTopics(client, topics)

	// Start the polling in a separate goroutine
	go func() {
		s.poll(client)
	}()

	return client
}

func createTopics(client *kgo.Client, topics config_models.KafkaTopics) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := kadm.NewClient(client).CreateTopics(ctx, 3, -1, nil, topics.DefaultProducer, topics.DefaultConsumer)

	if err != nil && !strings.Contains(err.Error(), "TOPIC_ALREADY_EXISTS") {
		panic(fmt.Sprintf("failed to create topics %s and/or %s: %v", topics.DefaultProducer, topics.DefaultConsumer, err))
	}
}
