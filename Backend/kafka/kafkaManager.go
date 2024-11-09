package kafka

import (
	"fmt"

	"github.com/adityjoshi/avinya/Backend/kafka"
)

// KafkaManager is responsible for managing Kafka producers and sending messages to topics
type KafkaManager struct {
	northProducer *kafka.NorthProducer
	southProducer *kafka.SouthProducer
	// Add other region producers if needed
}

// NewKafkaManager initializes and returns a KafkaManager instance
func NewKafkaManager(northBrokers, southBrokers []string) (*KafkaManager, error) {
	northProducer, err := kafka.NewNorthProducer(northBrokers)
	if err != nil {
		return nil, fmt.Errorf("Error initializing North producer: %w", err)
	}

	southProducer, err := kafka.NewSouthProducer(southBrokers)
	if err != nil {
		return nil, fmt.Errorf("Error initializing South producer: %w", err)
	}

	return &KafkaManager{
		northProducer: northProducer,
		southProducer: southProducer,
	}, nil
}

// SendUserRegistrationMessage sends the user registration data to the appropriate Kafka topic based on the region
func (km *KafkaManager) SendUserRegistrationMessage(region, topic, message string) error {
	var producer kafka.Producer

	// Determine which producer to use based on the region
	switch region {
	case "north":
		producer = km.northProducer
	case "south":
		producer = km.southProducer
	default:
		return fmt.Errorf("invalid region: %s", region)
	}

	// Send the message to the specified Kafka topic
	return producer.SendMessage(topic, message)
}
